package namechecker

import "net/http"

type ErrNetworkProblem struct {
	Cause error
}

func (e *ErrNetworkProblem) Error() string {
	return "Network problem, could not check availability"
}

func (e *ErrNetworkProblem) Unwrap() error {
	return e.Cause
}

type Client interface {
	Get(string) (*http.Response, error)
}

type NameChecker interface {
	Validate(string) bool
	IsAvailable(string, Client)
}
