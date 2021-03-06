package models

import (
	"encoding/json"
	errors2 "github.com/lusantisuper/api-rede-golang/apierrs"
	"github.com/lusantisuper/api-rede-golang/utils"
)

// Payment Struct of the request to de REDE's API
type Payment struct {
	// Optional: Define true if you want the request to be automatic captured
	Capture bool `json:"capture,omitempty"`
	// Optional: Payment method -> "credit" and "debit"
	Kind string `json:"kind,omitempty"`
	// Task code generated by the establishment
	Reference string `json:"reference"`
	// Price of the product R$10,00 = 1000
	Amount int `json:"amount"`
	// Optional: Number of payments in installments -> 2 until 12
	// Not setting this value make the one time pay
	Installments int `json:"installments,omitempty"`
	// Optional: Name for the card's owner
	CardHolderName string `json:"cardHolderName,omitempty"`
	// Card's number
	CardNumber int `json:"cardNumber"`
	// Expiration month -> 1 until 12
	ExpirationMonth int `json:"expirationMonth"`
	// Expiration year -> could be 2021 or 21
	ExpirationYear int `json:"expirationYear"`
	// Optional: Card's security code
	SecurityCode int `json:"securityCode,omitempty"`
	// Optional: This string will be printed on the bill
	SoftDescriptor string `json:"softDescriptor,omitempty"`
	// Optional: Is it a subscription? This is only a log, the payment need to be redone every time the store wants
	Subscription bool `json:"subscription,omitempty"`
	// TODO
	origin int
	// PV number
	DistributorAffiliation int `json:"distributorAffiliation"`
	// TODO
	brandTid string
}

// ToJSON Return a valid byte array of the Payment
func (r Payment) ToJSON() ([]byte, error) {
	result := Payment{
		Capture:                r.Capture,
		Kind:                   r.Kind,
		Reference:              r.Reference,
		Amount:                 r.Amount,
		Installments:           r.Installments,
		CardHolderName:         r.CardHolderName,
		CardNumber:             r.CardNumber,
		ExpirationMonth:        r.ExpirationMonth,
		ExpirationYear:         r.ExpirationYear,
		SecurityCode:           r.SecurityCode,
		SoftDescriptor:         r.SoftDescriptor,
		Subscription:           r.Subscription,
		origin:                 r.origin,
		DistributorAffiliation: r.DistributorAffiliation,
		brandTid:               r.brandTid,
	}

	// Adding all necessary parameters
	if utils.IsStringEmpty(r.Reference) {
		return nil, errors2.APIErr(errors2.INSUFFICIENTPARAMETERS)
	}
	if r.Amount < 0 || r.Amount > 1000000000 {
		return nil, errors2.APIErr(errors2.WRONGAMOUNT)
	}
	if r.CardNumber == 0 {
		return nil, errors2.APIErr(errors2.INSUFFICIENTPARAMETERS)
	}
	if r.ExpirationMonth < 0 || r.ExpirationMonth > 12 {
		return nil, errors2.APIErr(errors2.WRONGDATENUMBER)
	}
	if r.ExpirationYear < 20 || r.ExpirationYear > 60 {
		return nil, errors2.APIErr(errors2.WRONGDATENUMBER)
	}
	if r.DistributorAffiliation == 0 {
		return nil, errors2.APIErr(errors2.INSUFFICIENTPARAMETERS)
	}

	// Adding all optional parameters
	if !r.Capture {
		result.Capture = true
	}
	if !utils.IsStringEmpty(r.Kind) {
		result.Kind = r.Kind
	}
	if r.Installments >= 2 && r.Installments <= 12 {
		result.Installments = r.Installments
	} else {
		result.Installments = 1
	}
	if !utils.IsStringEmpty(r.CardHolderName) {
		result.CardHolderName = r.CardHolderName
	}
	if !utils.IsStringEmpty(r.SoftDescriptor) {
		result.SoftDescriptor = r.SoftDescriptor
	}

	jsonResult, err := json.Marshal(result)
	return jsonResult, err
}
