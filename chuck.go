package chucknorrisgo

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const BaseURL string = "http://api.icndb.com/jokes/"

type Joke struct {
	FirstName     string
	LastName      string
	Categories    []string
	RequestParams string
}

func (joke *Joke) Random() JokeResponse {
	requestURL := BaseURL + "random/" + joke.RequestParams
	result, err := MakeRequest(requestURL)

	if err != nil {
		return JokeResponse{}
	}

	return WrapResponse(result)
}

func (joke *Joke) Get(JokeID int) JokeResponse {
	requestURL := BaseURL + strconv.Itoa(JokeID) + "/" + joke.RequestParams
	result, err := MakeRequest(requestURL)

	if err != nil {
		return JokeResponse{}
	}

	return WrapResponse(result)
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

type JokeResponse struct {
	JokeID   float64
	JokeText string
}

func WrapResponse(jsonResp string) JokeResponse {
	var result map[string]interface{}
	err := json.Unmarshal([]byte(jsonResp), &result)
	if err != nil {
		log.Fatalln(err)
	}
	result = result["value"].(map[string]interface{})
	return JokeResponse{JokeID: result["id"].(float64), JokeText: result["joke"].(string)}
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
