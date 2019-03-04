# chuck-norris-go

Golang wrapper for "The Internet Chuck Norris Database" http://www.icndb.com/api/


#### Installation:

```bash
$ go get github.com/Keda87/chucknorrisgo
```


#### Example usage:
```go
import (
	"fmt"
	"github.com/Keda87/chucknorrisgo"	
) 


cn := &chucknorrisgo.ChuckNorris{}

// Build the payload.
joke := cn.Build()
joke := cn.FirstName("John").Build()                                          // Change main name character.
joke := cn.FirstName("John").LastName("Doe").Build()                          // Change main name character.
category := []string{"nerdy", "explicit"}
joke := cn.FirstName("John").LastName("Doe").Categories(category...).Build()  // Filter jokes by categories.

// Fetch the jokes.
result := joke.Random()  // Get the jokes randomly.
result := joke.Get(128)  // Get the jokes by specific ID.

fmt.Println(result.JokeID)
fmt.Println(result.JokeText)
```

