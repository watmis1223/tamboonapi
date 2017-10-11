package service

import(
	"log"
	"strconv"
	"time"
	"github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
)

type DonationReq struct {
	Name   string `json:"name"`
	Pan  string `json:"pan"`
	ExpiryMonth int  `json:"expiryMonth"`
	ExpiryYear int  `json:"expiryYear"`
	Cvv2 string `json:"cvv2"`
	Amount string `json:"amount"`
	CharityName string `json:"charityName"`
}

func (model *DonationReq)DoDonate(client *omise.Client) (*omise.Charge, *operations.CreateCharge)  {

	// Creates a token from a test card.
	token, createToken := &omise.Token{}, &operations.CreateToken{
		Name:            model.Name,
		Number:          model.Pan,
		ExpirationMonth: time.Month(model.ExpiryMonth),
		ExpirationYear:  model.ExpiryYear,
	}
	if e := client.Do(token, createToken); e != nil {
		log.Fatal(e)
	}

	i, e := strconv.ParseInt(model.Amount, 10, 64)
	if e != nil {
		log.Printf("conver amount error")
		log.Fatal(e)
	}
	charge, createCharge := &omise.Charge{}, &operations.CreateCharge{
		Card:        token.ID,
		Amount:      i,
		Currency:    "THB",
		Description: model.CharityName,
	}
	if e := client.Do(charge, createCharge); e != nil {
		log.Fatal(e)
	}

	log.Printf("charge: %s  amount: %s %d\n", charge.ID, charge.Currency, charge.Amount)

	return charge, createCharge
}