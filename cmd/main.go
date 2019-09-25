package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	messagebird "github.com/messagebird/go-rest-api"
	"github.com/messagebird/go-rest-api/sms"
)

// TODO:
// - Add my proper list of relatives
// - Host it somewhere (aws lambda on schedule 1pm deployed manually seems to be the best)

// SMSRequest is the type matching the body of a POST to /sms
type SMSRequest struct {
	Recipient  int    `json:"recipient"`
	Originator string `json:"originator"`
	Message    string `json:"message"`
}

// Call MessageBird service to send an SMS
func sendSMS(client *messagebird.Client, smsRequest SMSRequest) error {
	msg, err := sms.Create(
		client,
		smsRequest.Originator,
		[]string{strconv.Itoa(smsRequest.Recipient)},
		smsRequest.Message,
		nil,
	)
	if err != nil {
		return fmt.Errorf("Failed sending sms through messagebird : '%s'", err.Error())
	}

	log.Printf("SMS successfully sent to messagebird service '%v'", msg)
	return nil
}

// Singleton instance of a Message Bird client.
var MBClient *messagebird.Client

func init() {

	// Test API key
	// MBClient = messagebird.New("HLrjj0gVcebObnlTFQrt0E11U")

}

func main() {
	log.Println("Go HB!")

	birthdates := map[string]time.Time{
		"Arnaud": time.Date(1991, time.January, 22, 0, 0, 0, 0, time.UTC),
		"Igor":   time.Date(1991, time.August, 15, 0, 0, 0, 0, time.UTC),
		"Johny":  time.Date(1991, time.September, 1, 0, 0, 0, 0, time.UTC),
	}
	now := time.Now()
	now = time.Date(1991, time.September, 1, 0, 0, 0, 0, time.UTC)

	smsSent := false
	for name, birthdate := range birthdates {
		if now.Day() == birthdate.Day() && now.Month() == birthdate.Month() {
			age := now.Year() - birthdate.Year()
			log.Printf("It is %s's birthday! %d years of greatness :) \n", name, age)
			smsRequest := SMSRequest{
				Recipient:  316123456,
				Originator: "go-HB app",
				Message:    fmt.Sprintf("It is %s's birthday! %d years of greatness, wish it :) ", name, age),
			}
			err := sendSMS(MBClient, smsRequest)
			if err != nil {
				log.Fatal(err)
			}
			smsSent = true
		} else {
			log.Printf("It is not %s's birthday\n", name)
		}
	}

	if now.Day() == 1 && !smsSent {
		smsRequest := SMSRequest{
			Recipient:  316123456,
			Originator: "go-HB app",
			Message:    fmt.Sprintf("go-hb is still standing strong!"),
		}
		err := sendSMS(MBClient, smsRequest)
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Println("All the birthdays from the crew have been checked, bye!")
}
