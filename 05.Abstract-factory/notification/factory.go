package notification

import (
	"errors"
	"fmt"

	. "github.com/dyden/go-medium/05.Abstract-factory/notification/types"
)

type INotificationFactory interface {
	SendNotification()
	GetSender() ISender
}

func GetNotificationFactory(notificationType string) (INotificationFactory, error) {
	switch notificationType {
	case "SMS":
		return &SMSNotification{}, nil
	case "EMAIL":
		return &EMAILNotification{}, nil
	default:
		return nil, errors.New("Unsuported notification type")
	}
}

func SendNotification(factory INotificationFactory) {
	factory.SendNotification()
}

func GetMethod(factory INotificationFactory) {
	fmt.Println("METHOD -> ", factory.GetSender().GetSenderMethod())
}
func GetChannel(factory INotificationFactory) {
	fmt.Println("CHANNEL-> ", factory.GetSender().GetSenderChannel())
}
