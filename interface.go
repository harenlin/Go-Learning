package main
import "fmt"

// define an interface of message sender
type MessageSender interface {
	Send(content string)
}

// email sender part
type EmailSender struct {
	Address string
}

func (s *EmailSender) Send(content string) {
	fmt.Println("Send:", content, "to", s.Address)
}

// mobile sender part
type SmsSender struct {
	Mobile string
}

func (s *SmsSender) Send(content string) {
	fmt.Println("Send:", content, "to", s.Mobile)
}

func main() {
	var sender MessageSender // Interface Declaration
	sender = &EmailSender{Address: "test@gmail.com"}
	sender.Send("Hello, Email!") // Send: Hello, Email! to test@gmail.com
	sender = &SmsSender{Mobile: "0912345678"}
	sender.Send("Hello, Sms!") // Send: Hello, Sms! to 0912345678
}