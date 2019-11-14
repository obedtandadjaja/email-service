package main

import (
	"fmt"
	"os"

	"github.com/mailgun/mailgun-go/v3"
	"github.com/obedtandadjaja/email-service/api"
	"github.com/sirupsen/logrus"
)

var Environment string
var AppHost, AppPort, AppUrl string
var MailgunDomain, MailgunApiKey string

func init() {
	Environment = os.Getenv("ENV")
	AppHost = os.Getenv("APP_HOST")
	AppPort = os.Getenv("APP_PORT")
	AppUrl = AppHost + ":" + AppPort
	MailgunDomain = os.Getenv("MAILGUN_DOMAIN")
	MailgunApiKey = os.Getenv("MAILGUN_API_KEY")
}

func main() {
	fmt.Println("hello")

	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)

	mailgunInstance := mailgun.NewMailgun(MailgunDomain, MailgunApiKey)

	api.Start(AppUrl, mailgunInstance)
}
