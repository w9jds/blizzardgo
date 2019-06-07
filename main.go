package blizzard

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type LocaleTypes int

const (
	EN_US LocaleTypes = iota
)

func (locale LocaleTypes) String() string {
	return [...]string{
		"en_US",
	}[locale]
}

// API is an instance of a client to interact with the blizzard api
type API struct {
	client       *http.Client
	clientAuth   *AccessToken
	region       string
	baseURI      string
	authURI      string
	clientID     string
	rendersURI   string
	clientSecret string
	redirectURI  string
	locale       LocaleTypes
}

// Options is the configuration object used to build the blizzard api client
type Options struct {
	Region       string
	Locale       LocaleTypes
	ClientID     string
	ClientSecret string
	RedirectURI  string
}

// NewClient creates a new instance of the BlizzardAPI client
func NewClient(options *Options) *API {
	api := &API{
		region:       options.Region,
		clientID:     options.ClientID,
		clientSecret: options.ClientSecret,
		redirectURI:  options.RedirectURI,
		locale:       options.Locale,
		client:       &http.Client{Timeout: time.Second * 10},
		authURI:      fmt.Sprintf("https://%s.battle.net", options.Region),
		baseURI:      fmt.Sprintf("https://%s.api.blizzard.com", options.Region),
	}

	api.setClientToken()

	return api
}

func (api API) setClientToken() {
	clientToken, err := api.GetClientAccessToken()
	if err != nil {
		log.Fatal("Error getting client token for Blizzard API", err.Error())
	}

	api.clientAuth = clientToken
}

func (api API) newGet(path string) (*http.Request, error) {
	return http.NewRequest("GET", api.baseURI+path, nil)
}

func (api API) newPost(path string, body io.Reader) (*http.Request, error) {
	return http.NewRequest("POST", api.baseURI+path, body)
}

func (api API) setLocale(locale LocaleTypes) {
	api.locale = locale
}

func (api API) do(request *http.Request, out interface{}) error {
	if api.clientAuth != nil && request.Header.Get("Authorization") == "" {
		if time.Now().After(api.clientAuth.ExpiresAt) {
			api.setClientToken()
		}

		request.Header.Set("Authorization", api.clientAuth.AccessToken)
	}

	response, err := api.client.Do(request)
	if err != nil {
		return err
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if response.StatusCode > 199 && response.StatusCode < 300 {
		err := json.Unmarshal(body, out)
		if err != nil {
			return err
		}

		return nil
	}

	return errors.New(string(body))
}
