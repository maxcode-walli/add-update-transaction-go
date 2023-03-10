package p

import "encoding/json"

type TransactionAnomalyScoreCalculatedEventV1 struct {
	UserId            string
	ExternalAccountId string
	TransactionId     string
	Label             string
	Score             json.Number
}

type Money struct {
	Amount   int64  `json:"amount"`
	Currency string `json:"currency"`
}

type CounterPart struct {
	Name     string                 `json:"name,omitempty"`
	IBAN     string                 `json:"iban,omitempty"`
	MetaInfo map[string]interface{} `json:"metaInfo,omitempty"`
}

type TransactionUpdatedEventV1 struct {
	UserID                            string                 `json:"userID,omitempty"`
	IBAN                              string                 `json:"iban,omitempty"`
	ExternalAccountID                 string                 `json:"externalAccountID,omitempty"`
	TransactionStatus                 string                 `json:"transactionStatus,omitempty"`
	TransactionId                     string                 `json:"transactionId,omitempty"`
	Uuid                              string                 `json:"uuid,omitempty"`
	EndToEndId                        string                 `json:"endToEndId,omitempty"`
	CreditorAccount                   *CounterPart           `json:"creditorAccount,omitempty"`
	DebtorAccount                     *CounterPart           `json:"debtorAccount,omitempty"`
	TransactionAmount                 *Money                 `json:"transactionAmount,omitempty"`
	BookingDate                       string                 `json:"bookingDate,omitempty"`
	ValueDate                         string                 `json:"valueDate,omitempty"`
	BalanceAfterTransaction           string                 `json:"balanceAfterTransaction,omitempty"`
	RemittanceInformationUnstructured string                 `json:"remittanceInformationUnstructured,omitempty"`
	RemittanceInformationStructured   string                 `json:"remittanceInformationStructured,omitempty"`
	MetaInfo                          map[string]interface{} `json:"metaInfo,omitempty,omitempty"`
}
