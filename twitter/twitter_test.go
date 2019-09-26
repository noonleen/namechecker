package twitter_test

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/noonleen/goprjt/twitter"
)

var twt twitter.Twitter

type MockClient struct {
}

func (mc *MockClient) Get(string) (*http.Response, error) {
	resp := http.Response{StatusCode: http.StatusNotFound, Body: ioutil.NopCloser(nil)}
	return &resp, nil
}

func TestValidateTooShort(t *testing.T) {
	if twt.Validate("") != false {
		t.Error("Erreur too short")
	}
}

func TestValidateTooLong(t *testing.T) {
	if twt.Validate("kkkkkkkkkkkkkkkk") != false {
		t.Error("Erreur too long")
	}
}

func TestValidateTwitterWord(t *testing.T) {
	if twt.Validate("hello_twItter") != false {
		t.Error("Erreur twitter word")
	}
}

func TestValidateSpecialChars(t *testing.T) {
	if twt.Validate("hello$") != false {
		t.Error("Erreur special chars")
	}
}

func TestValidateAllRight(t *testing.T) {
	if twt.Validate("hello") != true {
		t.Error("Erreur all right")
	}
}

func TestIsAvailable(t *testing.T) {

	var client MockClient
	if isavail, _ := twt.IsAvailable("hello", &client); !isavail {
		t.Error("Erreur 404 non détecté")
	}
}
