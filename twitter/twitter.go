package twitter

import (
	"net/http"
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/noonleen/namechecker"
)

type Twitter struct {
}

func (t *Twitter) IsAvailable(username string, client namechecker.Client) (bool, error) {

	resp, err := client.Get("https://twitter.com/" + username)
	if err != nil {
		errnp := namechecker.ErrNetworkProblem{Cause: err}
		return false, &errnp
	}
	resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return true, nil
	}
	return false, nil
}

// min length = 1
// max length = 15
// twitter (min/maj combinées) n'apparaît pas
// chars : 0-9A-Z_a-z
func (t *Twitter) Validate(username string) bool {

	legalPattern := "^[0-9A-Z_a-z]*$"
	legalRegexp := regexp.MustCompile(legalPattern)

	if utf8.RuneCountInString(username) < 1 || utf8.RuneCountInString(username) > 15 {
		return false
	}

	if strings.Contains(strings.ToLower(username), "twitter") {
		return false
	}

	if !legalRegexp.MatchString(username) {
		return false
	}

	return true
}
