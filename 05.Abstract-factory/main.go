package main

import (
	"fmt"

	notification "github.com/dyden/go-medium/05.Abstract-factory/notification"
)

func main() {

	SMSFactory, err := notification.GetNotificationFactory("SMS")
	EMAILFactory, err := notification.GetNotificationFactory("EMAIL")
	if err != nil {
		fmt.Println(err)
	}

	notification.SendNotification(SMSFactory)
	notification.GetMethod(SMSFactory)
	notification.GetChannel(SMSFactory)

	notification.SendNotification(EMAILFactory)
	notification.GetMethod(EMAILFactory)
	notification.GetChannel(EMAILFactory)

}
