package models

import (
	"testing"
)

func TestCheckPasswordCorrect(t *testing.T) {
	user := AppUser{"pawel", "pZHTjHUjKXfmyHjQzxL10w==:g2yMB7ukc6aU30Rhe4lB1FlXxCfIqcHvnXOciAe/IpE="}
	pass := "somepassword"
	res, err := user.CheckPassword(pass)
	if err != nil {
		t.Errorf("CheckPassword() should not return an error. Got %v", err)
	}
	if res != true {
		t.Errorf("CheckPassword() should accept this password")
	}
}

func TestCheckPasswordIncorrect(t *testing.T) {
	user := AppUser{"pawel", "pZHTjHUjKXfmyHjQzxL10w==:g2yMB7ukc6aU30Rhe4lB1FlXxCfIqcHvnXOciAe/IpE="}
	pass := "somepassword1"
	res, err := user.CheckPassword(pass)
	if err == nil {
		t.Errorf("CheckPassword() SHOULD return an error. Got nil")
	}
	if res != false {
		t.Errorf("CheckPassword() SHOULD NOT accept this password")
	}
}

func TestSetPassword(t *testing.T) {
	user := AppUser{"pawel", ""}
	password := "someotherpassword"
	err := user.SetPassword(password)
	if err != nil {
		t.Errorf("SetPassword() SHOULD NOT return an error. Got %v", err)
	}

	res, err := user.CheckPassword(password)
	if err != nil {
		t.Errorf("CheckPassword() should not return an error. Got %v", err)
	}
	if res != true {
		t.Errorf("CheckPassword() should accept this password")
	}
}
