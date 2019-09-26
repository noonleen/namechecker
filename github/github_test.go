package github_test

import (
	"testing"

	"github.com/noonleen/goprjt/github"
)

var gtb github.Github

func TestValidateTooShort(t *testing.T) {
	if gtb.Validate("") != false {
		t.Error("Erreur too short")
	}
}

func TestValidateTooLong(t *testing.T) {
	if gtb.Validate("kkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkk") != false {
		t.Error("Erreur too long")
	}
}

func TestValidateDoubleMinus(t *testing.T) {
	if gtb.Validate("hel--lo") != false {
		t.Error("Erreur double minus")
	}
}

func TestValidateSpecialChars(t *testing.T) {
	if gtb.Validate("hello_") != false {
		t.Error("Erreur special chars")
	}
}

func TestValidatePrefSuff(t *testing.T) {
	if gtb.Validate("-hello-") != false {
		t.Error("Erreur pref suff")
	}
}

func TestValidateAllRight(t *testing.T) {
	if gtb.Validate("hello") != true {
		t.Error("Erreur all right")
	}
}
