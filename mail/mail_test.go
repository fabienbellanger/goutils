package mail

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestServerName tests with a port greater than 0.
func TestServerName(t *testing.T) {
	port := 2000
	host := "localhost"

	got := serverName(host, port)
	wanted := "localhost:2000"

	assert.Equal(t, got, wanted)
}

// TestServerNameWithoutPort tests with a port equal to 0.
func TestServerNameWithoutPort(t *testing.T) {
	port := 0
	host := "localhost"

	got := serverName(host, port)
	wanted := "localhost"

	assert.Equal(t, got, wanted)
}

// TestBuildMessageWithoutFrom tests BuildMessage without mail from.
func TestBuildMessageWithoutFrom(t *testing.T) {
	mail := Mail{
		To: make([]string, 1),
	}

	err := mail.buildMessage()

	assert.NotNil(t, err)
}

// TestBuildMessageWithouTo tests BuildMessage without mail to.
func TestBuildMessageWithoutTo(t *testing.T) {
	mail := Mail{
		From: "bob@alice.test",
	}

	err := mail.buildMessage()

	assert.NotNil(t, err)
}

// TestBuildMessage tests BuildMessage with good parameters.
func TestBuildMessage(t *testing.T) {
	mail := Mail{}
	mail.From = "john.doe@test.com"
	mail.To = []string{"john.doe@test.com"}
	mail.Cc = []string{"copy1@gmail.com"}
	mail.Subject = "test"
	mail.Body = "ghbzuhgzughzeughzuighzeughzughzugzhughzughzguzhguzih"
	mail.Attachments = []string{"../resources/tests/image.jpg"}
	got := mail.buildMessage()

	// This is the separator used for the various parts of the MIME message structure.
	bPlaceholder := "our-custom-separator"

	// Create a buffer for the MIME message.
	mime := bytes.NewBuffer(nil)

	// Construct the main MIME headers.
	mime.WriteString("From: john.doe@test.com\r\n")
	mime.WriteString("To: john.doe@test.com\r\n")

	if len(mail.Cc) > 0 {
		mime.WriteString("Cc: copy1@gmail.com\r\n")
	}

	mime.WriteString("Subject: test\r\n")
	mime.WriteString("MIME-Version: 1.0\r\n")
	mime.WriteString(fmt.Sprintf("Content-Type: multipart/mixed; boundary=%s\r\n\r\n", bPlaceholder))

	// Add the message body.
	mime.WriteString(fmt.Sprintf("--%s\r\n", bPlaceholder))
	mime.WriteString("Content-Type: text/plain; charset=utf-8\r\n\r\n")
	mime.WriteString("ghbzuhgzughzeughzuighzeughzughzugzhughzughzguzhguzih")
	mime.WriteString("\r\n")

	// Attach files from the filesystem.
	for _, file := range mail.Attachments {
		mime.WriteString(fmt.Sprintf("--%s\r\n", bPlaceholder))
		mime.WriteString("Content-Type: application/octet-stream\r\n")
		mime.WriteString("Content-Description: image.jpg\r\n")
		mime.WriteString("Content-Transfer-Encoding: base64\r\n")
		mime.WriteString("Content-Disposition: attachment; filename=\"image.jpg\"\r\n\r\n")

		fileContent, _ := os.ReadFile(file)

		b := make([]byte, base64.StdEncoding.EncodedLen(len(fileContent)))
		base64.StdEncoding.Encode(b, fileContent)
		mime.Write(b)
		mime.WriteString("\r\n")
	}

	// End of the message.
	mime.WriteString(fmt.Sprintf("--%s--\r\n", bPlaceholder))

	wanted := mime.Bytes()

	assert.Equal(t, got, wanted)
}

// TestSend tests mail send construction.
func TestSend(t *testing.T) {
	err := Send("from", []string{"test@example.com"}, []string{"test2@example.com"}, []string{"image.jpg"}, "subject", "body", "smptUser", "smtpPassword", "smtpHost", 1234)

	assert.NotNil(t, err)
}

// TestSendMail tests mail send.
func TestSendMail(t *testing.T) {
	mail := Mail{}
	err := sendMail(mail, mail.buildMessage(), "", 1234)

	assert.NotNil(t, err)
}
