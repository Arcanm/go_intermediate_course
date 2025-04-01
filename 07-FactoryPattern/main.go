package main

import "fmt"

type NotificacionFactory interface {
	SendNotification()
	GetSender() ISender
}

type ISender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}

// SMS IMPLEMENTATION
type SmsNotification struct{}

// ISENDER TYPE
type SmsNotificationSender struct{}

// PREPARING ISENDER INTERFACE
func (SmsNotificationSender) GetSenderMethod() string {
	return "SMS"
}

func (SmsNotificationSender) GetSenderChannel() string {
	return "Twilio"
}

// PREPARING NOTIFICATION FACTORY INTERFACE
func (SmsNotification) SendNotification() {
	fmt.Println("Sending SMS Notification")
}

func (SmsNotification) GetSender() ISender {
	return SmsNotificationSender{}
}

// EMAIL IMPLEMENTATION

type EmailNotification struct{}
type EmailNotificationSender struct{}

// PREPARING ISENDER INTERFACE
func (EmailNotificationSender) GetSenderMethod() string {
	return "EMAIL"
}

func (EmailNotificationSender) GetSenderChannel() string {
	return "SES"
}

// PREPARING NOTIFICATION FACTORY INTERFACE
func (EmailNotification) SendNotification() {
	fmt.Println("Sending EMAIL notification")
}

func (EmailNotification) GetSender() ISender {
	return EmailNotificationSender{}
}

func getNotificationFactory(notificationType string) (NotificacionFactory, error) {
	if notificationType == "SMS" {
		return &SmsNotification{}, nil
	}

	if notificationType == "EMAIL" {
		return &EmailNotification{}, nil
	}

	return nil, fmt.Errorf("Invalid Notification Type")
}

func sendNotification(nf NotificacionFactory) {
	nf.SendNotification()
}

func getMethod(f NotificacionFactory) {
	fmt.Println(f.GetSender().GetSenderMethod())
}

func main() {
	smsFactory, _ := getNotificationFactory("SMS")
	emailFactory, _ := getNotificationFactory("EMAIL")
	_, err := getNotificationFactory("LOQUESEA")

	sendNotification(smsFactory)
	sendNotification(emailFactory)
	getMethod(smsFactory)
	getMethod(emailFactory)

	fmt.Println(err)
}
