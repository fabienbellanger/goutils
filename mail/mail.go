package mail

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net/smtp"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// Mail struct
type Mail struct {
	From        string
	To          []string
	Cc          []string
	Subject     string
	Body        string
	Attachments []string
}

// Send sends mail.
func Send(from string, to, cc, attachments []string, subject, body, smptUser, smtpPassword, smtpHost string, smtpPort int) error {

	mail := Mail{}
	mail.From = from
	mail.To = to
	mail.Cc = cc
	mail.Subject = subject
	mail.Body = body
	mail.Attachments = attachments

	// Envoi avec authentification
	// ---------------------------
	if smptUser != "" && smtpPassword != "" {
		auth := smtp.PlainAuth("", smptUser, smtpPassword, smtpHost)
		err := smtp.SendMail(serverName(smtpHost, smtpPort),
			auth,
			mail.From,
			mail.To,
			mail.buildMessage())
		if err != nil {
			return err
		}

		return nil
	}

	// Envoi sans authentification
	// ---------------------------
	return sendMail(mail, mail.buildMessage(), smtpHost, smtpPort)
}

// serverName returns SMTP server name from host and port.
func serverName(host string, port int) (s string) {
	s = host

	if port > 0 {
		s += ":" + strconv.Itoa(port)
	}

	return s
}

// buildMessage constructs mail.
func (mail *Mail) buildMessage() []byte {
	// This is the separator used for the various parts of the MIME message structure.
	bPlaceholder := "our-custom-separator"

	// Create a buffer for the MIME message.
	mime := bytes.NewBuffer(nil)

	// Construct the main MIME headers.
	mime.WriteString(fmt.Sprintf("From: %s\r\n", mail.From))
	mime.WriteString(fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ", ")))

	if len(mail.Cc) > 0 {
		mime.WriteString(fmt.Sprintf("Cc: %s\r\n", strings.Join(mail.Cc, ", ")))
	}

	mime.WriteString(fmt.Sprintf("Subject: %s\r\n", mail.Subject))
	mime.WriteString(fmt.Sprintf("MIME-Version: 1.0\r\n"))
	mime.WriteString(fmt.Sprintf("Content-Type: multipart/mixed; boundary=%s\r\n\r\n", bPlaceholder))

	// Add the message body.
	mime.WriteString(fmt.Sprintf("--%s\r\n", bPlaceholder))
	mime.WriteString("Content-Type: text/plain; charset=utf-8\r\n\r\n")
	mime.WriteString(mail.Body)
	mime.WriteString("\r\n")

	// Attach files from the filesystem.
	for _, file := range mail.Attachments {
		_, filename := filepath.Split(file)
		mime.WriteString(fmt.Sprintf("--%s\r\n", bPlaceholder))
		mime.WriteString(fmt.Sprintf("Content-Type: application/octet-stream\r\n"))
		mime.WriteString(fmt.Sprintf("Content-Description: %s\r\n", filename))
		mime.WriteString(fmt.Sprintf("Content-Transfer-Encoding: base64\r\n"))
		mime.WriteString(fmt.Sprintf("Content-Disposition: attachment; filename=\"%s\"\r\n\r\n", filename))

		fileContent, err := os.ReadFile(file)
		if err != nil {
			return nil
		}
		b := make([]byte, base64.StdEncoding.EncodedLen(len(fileContent)))
		base64.StdEncoding.Encode(b, fileContent)
		mime.Write(b)
		mime.WriteString("\r\n")
	}

	// End of the message.
	mime.WriteString(fmt.Sprintf("--%s--\r\n", bPlaceholder))

	return mime.Bytes()
}

// sendMail constructs and sends mail.
func sendMail(mail Mail, msg []byte, smtpHost string, smtpPort int) error {
	c, err := smtp.Dial(serverName(smtpHost, smtpPort))
	if err != nil {
		return err
	}
	defer c.Close()

	if err = c.Mail(mail.From); err != nil {
		return err
	}

	for i := range mail.To {
		if err = c.Rcpt(mail.To[i]); err != nil {
			return err
		}
	}

	w, err := c.Data()
	if err != nil {
		return err
	}

	_, err = w.Write(msg)
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}

	return c.Quit()
}

/*func main() {
	from := "hugo.le-guen@apitic.com"
	to := []string{"hugo.le-guen@apitic.com"}
	cc := []string{"copy1@gmail.com", "copy2@gmail.com"}
	subject := "test"
	body := "ghbzuhgzughzeughzuighzeughzughzugzhughzughzguzhguzih"
	attachments := []string{"image.jpg"}

	Send(from, to, cc, attachments, subject, body, "", "", "localhost", 1025)
}*/
