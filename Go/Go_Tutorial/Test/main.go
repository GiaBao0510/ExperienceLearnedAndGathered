package main

import (
	"fmt"
	"log"


	"github.com/mailjet/mailjet-apiv3-go/v4"
)

func main() {
	mailjetClient := mailjet.NewMailjetClient("8bc4546bd33a45abb3717fb51bb19e5d", "767e328ce0dc56be36f81560d006d509")
	messagesInfo := []mailjet.InfoMessagesV31{
		{
			From: &mailjet.RecipientV31{
				Email: "pilot@mailjet.com",
				Name:  "Mailjet Pilot",
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: "baopg05102002@mailjet.com",
					Name:  "passenger 1",
				},
			},
			Subject:  "Your email flight plan!",
			TextPart: "Dear passenger 1, welcome to Mailjet! May the delivery force be with you!",
			HTMLPart: "<h3>Dear passenger 1, welcome to <a href=\"https://www.mailjet.com/\">Mailjet</a>!</h3><br />May the delivery force be with you!",
		},
	}
	messages := mailjet.MessagesV31{Info: messagesInfo}
	res, err := mailjetClient.SendMailV31(&messages)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Data: %+v\n", res)
}

// api key: 8bc4546bd33a45abb3717fb51bb19e5d
// secret key: 767e328ce0dc56be36f81560d006d509