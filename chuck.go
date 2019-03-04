package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const base_url string = "http://api.icndb.com/jokes/"

type Joke struct {
	FirstName     string
	LastName      string
	Categories    []string
	RequestParams string
}

func (joke *Joke) Random() string {
	request_url := base_url + "random/" + joke.RequestParams
	result, err := MakeRequest(request_url)

	if err != nil {
		return ""
	}

	return result
}

func (joke *Joke) Get(JokeID int) string {
	request_url := base_url + strconv.Itoa(JokeID) + "/" + joke.RequestParams
	result, err := MakeRequest(request_url)

	if err != nil {
		return ""
	}

	return result
}

type ChuckNorris struct {
	Joke
}

func (cn *ChuckNorris) Build() Joke {
	joke := cn.Joke
	params := ""

	if joke.FirstName == "" {
		params += "?firstName=Chuck"
	} else {
		params += "?firstName=" + joke.FirstName
	}

	if joke.LastName == "" {
		params += "&lastName=Norris"
	} else {
		params += "&lastName=" + joke.LastName
	}

	if len(joke.Categories) > 0 {
		if strings.Contains(params, "?") {
			params += "&limitTo=" + strings.Join(joke.Categories, ",")
		} else {
			params += "?limitTo=" + strings.Join(joke.Categories, ",")
		}
	}

	cn.Joke.RequestParams = params
	return cn.Joke
}

func (cn *ChuckNorris) FirstName(firstName string) *ChuckNorris {
	cn.Joke.FirstName = firstName
	return cn
}

func (cn *ChuckNorris) LastName(lastName string) *ChuckNorris {
	cn.Joke.LastName = lastName
	return cn
}

func (cn *ChuckNorris) Categories(cat ...string) *ChuckNorris {
	cn.Joke.Categories = cat
	return cn
}

func MakeRequest(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return "", err
	}

	return string(body), nil
}

func main() {
	cn := ChuckNorris{}

	cat := []string{"explicit"}
	joke := cn.FirstName("Adit").LastName("Keda").Categories(cat...).Build()

	fmt.Println(joke.Random())
	fmt.Println(joke.Get(528))
}
