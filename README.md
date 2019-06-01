# BlizzardGo

This library has built in functions to help you manage your sessions and requests with the Blizzard APIs. 

---

## Installation

In order to start using, download it using go commandline:
``` 
go get github.com/w9jds/blizzardgo
```
---

## Client Authentication

If you are using just client requests for data, the token is automatically pulled when you create the client. It also refreshs the token automatically for you to make sure that it never expires.

---

## User Authentication

You can also setup a full web api handling a user login like this:
``` golang

import (
	"log"
	"net/http"
	"os"

	blizzard "github.com/w9jds/blizzardgo"
)

var (
	clientID       string
	blizzardClient *blizzard.API
)

func main() {
	clientID = os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")

	blizzardClient = blizzard.NewClient(&blizzard.Options{
		Region:       "us",
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURI:  "http://localhost:8000/oauth",
	})

	http.HandleFunc("/oauth", blizzardAuth)

	http.ListenAndServe(":8000", nil)
}

func blizzardAuth(writer http.ResponseWriter, request *http.Request) {
	if code := request.URL.Query().Get("code"); code != "" {
		token, err := blizzardClient.GetUserAccessToken(code)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Println(token)
		return
	}

	http.Redirect(
		writer,
		request,
		blizzardClient.GetUserAuthorizationURI("wow.profile"),
		http.StatusTemporaryRedirect,
	)
}

```