package resty2curl

import (
	"fmt"
	"testing"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func TestExampleGetCurFromRestyRequest(t *testing.T) {

	t.Run(`Test GET from github`, func(t *testing.T) {
		client := resty.New().
			SetHostURL("https://api.github.com").
			SetTimeout(1 * time.Minute).
			SetRetryCount(2).
			SetRetryMaxWaitTime(1 * time.Minute).
			SetRetryWaitTime(1 * time.Minute)

		response, err := client.R().Get("/users/gabriellmandelli")

		assert.NotNil(t, response)
		assert.Nil(t, err)

		curl, err := GetCurFromRestyRequest(*response.Request)

		assert.Nil(t, err)

		stringCurl := curl.String()
		fmt.Println(stringCurl)
		assert.NotEmpty(t, stringCurl)

	})

	t.Run(`Test Post from github`, func(t *testing.T) {
		client := resty.New().
			SetHostURL("https://api.github.com").
			SetTimeout(1 * time.Minute).
			SetRetryCount(2).
			SetRetryMaxWaitTime(1 * time.Minute).
			SetRetryWaitTime(1 * time.Minute)

		example := exampleStruct{
			Name:  "Gabriel Longarete Mandelli",
			Email: "gabriel-mandelli@hotmail.com",
		}

		response, err := client.R().SetBody(example).Post("/users/gabriellmandelli")
		assert.NotNil(t, response)
		assert.Nil(t, err)

		curl, err := GetCurFromRestyRequest(*response.Request)

		assert.Nil(t, err)

		stringCurl := curl.String()
		fmt.Println(stringCurl)
		assert.NotEmpty(t, stringCurl)

	})
}

type exampleStruct struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
