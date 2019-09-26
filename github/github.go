package github

import (
	"net/http"
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/noonleen/namechecker"
)

type Github struct {
}

func (t *Github) IsAvailable(username string, client namechecker.Client) (bool, error) {

	resp, err := client.Get("https://github.com/" + username)
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
func (g *Github) Validate(username string) bool {

	legalPattern := "^[0-9A-Za-z-]*$"
	legalRegexp := regexp.MustCompile(legalPattern)

	if utf8.RuneCountInString(username) < 1 || utf8.RuneCountInString(username) > 39 {
		return false
	}

	if strings.Contains(username, "--") {
		return false
	}

	if strings.HasPrefix(username, "-") || strings.HasSuffix(username, "-") {
		return false
	}

	if !legalRegexp.MatchString(username) {
		return false
	}

	return true
}
