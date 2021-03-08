package types

import "fmt"

type EMAILNotification struct {
}
type EMAILNotificationSender struct {
}

func (EMAILNotification) SendNotification() {
	fmt.Println("Sending notification via EMAIL")
}
func (EMAILNotification) GetSender() ISender {
	return EMAILNotificationSender{}
}
func (EMAILNotificationSender) GetSenderMethod() string {
	return "EMAIL"
}
func (EMAILNotificationSender) GetSenderChannel() string {
	return "SES"
}
