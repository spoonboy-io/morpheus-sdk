/*
 * Morpheus API
 *
 * Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.
 *
 * API version: 6.2.1
 * Contact: dev@morpheusdata.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// CurrencyCode Currency Code (ISO 4217)
type CurrencyCode string

// List of CurrencyCode
const (
	AUD CurrencyCode = "AUD"
	BGN CurrencyCode = "BGN"
	BRL CurrencyCode = "BRL"
	CAD CurrencyCode = "CAD"
	CHF CurrencyCode = "CHF"
	CLF CurrencyCode = "CLF"
	CLP CurrencyCode = "CLP"
	CNY CurrencyCode = "CNY"
	CZK CurrencyCode = "CZK"
	DKK CurrencyCode = "DKK"
	EUR CurrencyCode = "EUR"
	GBP CurrencyCode = "GBP"
	HKD CurrencyCode = "HKD"
	HRK CurrencyCode = "HRK"
	HUF CurrencyCode = "HUF"
	IDR CurrencyCode = "IDR"
	ILS CurrencyCode = "ILS"
	INR CurrencyCode = "INR"
	JPY CurrencyCode = "JPY"
	KRW CurrencyCode = "KRW"
	MXN CurrencyCode = "MXN"
	MYR CurrencyCode = "MYR"
	NOK CurrencyCode = "NOK"
	NZD CurrencyCode = "NZD"
	PHP CurrencyCode = "PHP"
	PLN CurrencyCode = "PLN"
	RON CurrencyCode = "RON"
	RUB CurrencyCode = "RUB"
	SEK CurrencyCode = "SEK"
	SGD CurrencyCode = "SGD"
	THB CurrencyCode = "THB"
	TRY CurrencyCode = "TRY"
	USD CurrencyCode = "USD"
	ZAR CurrencyCode = "ZAR"
)

func (v *CurrencyCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CurrencyCode(value)
	for _, existing := range []CurrencyCode{ "AUD", "BGN", "BRL", "CAD", "CHF", "CLF", "CLP", "CNY", "CZK", "DKK", "EUR", "GBP", "HKD", "HRK", "HUF", "IDR", "ILS", "INR", "JPY", "KRW", "MXN", "MYR", "NOK", "NZD", "PHP", "PLN", "RON", "RUB", "SEK", "SGD", "THB", "TRY", "USD", "ZAR",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CurrencyCode", value)
}

// Ptr returns reference to CurrencyCode value
func (v CurrencyCode) Ptr() *CurrencyCode {
	return &v
}

type NullableCurrencyCode struct {
	value *CurrencyCode
	isSet bool
}

func (v NullableCurrencyCode) Get() *CurrencyCode {
	return v.value
}

func (v *NullableCurrencyCode) Set(val *CurrencyCode) {
	v.value = val
	v.isSet = true
}

func (v NullableCurrencyCode) IsSet() bool {
	return v.isSet
}

func (v *NullableCurrencyCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCurrencyCode(val *CurrencyCode) *NullableCurrencyCode {
	return &NullableCurrencyCode{value: val, isSet: true}
}

func (v NullableCurrencyCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCurrencyCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

