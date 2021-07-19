package main

import "fmt"

//SMS Y Email

type INotificationFactory interface {
	SendNotification()
	GetSender() ISender
}

type ISender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}

type SMSNotification struct {
}

func (SMSNotification) SendNotification() {
	fmt.Println("Se envio la notificacion via SMS")
}

func (SMSNotification) GetSender() ISender {
	return SMSNotificationSender{}
}

type SMSNotificationSender struct {
}

func (SMSNotificationSender) GetSenderMethod() string {
	return "SMS"
}

func (SMSNotificationSender) GetSenderChannel() string {
	return "Channel SMS"
}

type EmailNotification struct {
}

func (EmailNotification) SendNotification() {
	fmt.Println("Se envio la notificacion via Email")
}

func (EmailNotification) GetSender() ISender {
	return EmailNotificationSender{}
}

type EmailNotificationSender struct {
}

func (EmailNotificationSender) GetSenderMethod() string {
	return "Email"
}

func (EmailNotificationSender) GetSenderChannel() string {
	return "Channel Email"
}

func getNotificationFactory(notificacion string) (INotificationFactory, error) {
	if notificacion == "SMS" {
		return SMSNotification{}, nil
	}
	if notificacion == "Email" {
		return EmailNotification{}, nil
	}
	return nil, fmt.Errorf("No Notification Type")
}

func sendNotification(f INotificationFactory) {
	f.SendNotification()
}

func getMethod(f INotificationFactory) {
	fmt.Println(f.GetSender().GetSenderMethod())
}

func main() {
	smsFactory, _ := getNotificationFactory("SMS")
	emailFactory, _ := getNotificationFactory("Email")

	sendNotification(smsFactory)
	sendNotification(emailFactory)

	getMethod(smsFactory)
	getMethod(emailFactory)
}
