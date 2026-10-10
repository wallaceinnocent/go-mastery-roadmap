package interfaces

import "fmt"

type Notifier interface {
	SendMessage(text string)
}

type EmailService struct {
	Email string
}

// EmailService satisfies the Notifier interface because it implements SendNotification

func (e EmailService) SendMessage(msg string) {
	fmt.Printf("[Email] sending to  %s:  %s\n", e.Email, msg)
}

// Second implementation: SMSService struct
type SMSService struct {
	Phone string
}

// SMSService also satisfies the Notifier interface
func (s SMSService) SendMessage(message string) {
	fmt.Printf("[SMS] Sending to %s: %s\n", s.Phone, message)

}

// A function that accepts the *interface*, not a concrete struct
// This function doesn't care if it's an Email or an SMS, as long as it can SendNotification!
func AlertUser(n Notifier, msg string) {
	n.SendMessage(msg)
}
