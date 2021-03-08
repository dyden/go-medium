package types

import "fmt"

type SMSNotification struct {
}

type SMSNotificationSender struct {
}

func (sms SMSNotification) SendNotification() {
	fmt.Println("Sending notification via SMS")
}
func (SMSNotification) GetSender() ISender {
	return SMSNotificationSender{}
}
func (SMSNotificationSender) GetSenderMethod() string {
	return "SMS"
}
func (SMSNotificationSender) GetSenderChannel() string {
	return "Twilio"
}
