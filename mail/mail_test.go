package mail

import (
	"fmt"
	"testing"
)

// TestServerName tests with a port greater than 0.
func TestServerName(t *testing.T) {
	port := 2000
	host := "localhost"

	got := serverName(host, port)
	wanted := "localhost:2000"

	if got != wanted {
		t.Errorf("serverName with port - got: %+v, want: %+v.", got, wanted)
	}
}

// TestServerNameWithoutPort tests with a port equal to 0.
func TestServerNameWithoutPort(t *testing.T) {
	port := 0
	host := "localhost"

	got := serverName(host, port)
	wanted := "localhost"

	if got != wanted {
		t.Errorf("serverName without port - got: %+v, want: %+v.", got, wanted)
	}
}

func TestBuildMessageWithoutFrom(t *testing.T) {
	mail := Mail{
		To: make([]string, 1),
	}

	_, err := mail.buildMessage()
	if err == nil {
		t.Errorf("buildMessage without From must return an error.")
	}
}

func TestBuildMessageWithoutTo(t *testing.T) {
	mail := Mail{
		From: "bob@alice.test",
	}

	_, err := mail.buildMessage()
	if err == nil {
		t.Errorf("buildMessage without To must return an error.")
	}
}

func TestBuildMessage(t *testing.T) {
	to := make([]string, 2)
	to[0] = "alice@bob.test"
	to[1] = "john@doe.test"
	mail := Mail{
		From:    "bob@alice.test",
		To:      to,
		Subject: "Test",
		Body:    "Body",
	}

	got, _ := mail.buildMessage()
	wanted := fmt.Sprint("From: bob@alice.test\r\n")
	wanted += fmt.Sprint("To: alice@bob.test,john@doe.test\r\n")
	wanted += fmt.Sprint("Subject: Test\r\n")
	wanted += fmt.Sprint("MIME-version: 1.0\r\n")
	wanted += fmt.Sprint("Content-Type: text/html; charset: UTF-8\r\n")
	wanted += fmt.Sprint("\r\nBody")

	if got != wanted {
		t.Errorf("Mail buildMessage - got: %+v, want: %+v.", got, wanted)
	}
}
