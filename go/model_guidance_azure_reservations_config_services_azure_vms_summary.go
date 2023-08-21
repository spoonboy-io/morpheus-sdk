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
)

// GuidanceAzureReservationsConfigServicesAzureVmsSummary struct for GuidanceAzureReservationsConfigServicesAzureVmsSummary
type GuidanceAzureReservationsConfigServicesAzureVmsSummary struct {
	TotalSavings *float32 `json:"totalSavings,omitempty"`
	CurrencyCode *string `json:"currencyCode,omitempty"`
	TotalSavingsPercent *float32 `json:"totalSavingsPercent,omitempty"`
	Term *string `json:"term,omitempty"`
	PaymentOption *string `json:"paymentOption,omitempty"`
	Service *string `json:"service,omitempty"`
	OnDemandCount *int64 `json:"onDemandCount,omitempty"`
	OnDemandCost *float32 `json:"onDemandCost,omitempty"`
	ReservedCount *int64 `json:"reservedCount,omitempty"`
	ReservedCost *int64 `json:"reservedCost,omitempty"`
	RecommendedCount *int64 `json:"recommendedCount,omitempty"`
	RecommendedCost *float32 `json:"recommendedCost,omitempty"`
}

// NewGuidanceAzureReservationsConfigServicesAzureVmsSummary instantiates a new GuidanceAzureReservationsConfigServicesAzureVmsSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGuidanceAzureReservationsConfigServicesAzureVmsSummary() *GuidanceAzureReservationsConfigServicesAzureVmsSummary {
	this := GuidanceAzureReservationsConfigServicesAzureVmsSummary{}
	return &this
}

// NewGuidanceAzureReservationsConfigServicesAzureVmsSummaryWithDefaults instantiates a new GuidanceAzureReservationsConfigServicesAzureVmsSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGuidanceAzureReservationsConfigServicesAzureVmsSummaryWithDefaults() *GuidanceAzureReservationsConfigServicesAzureVmsSummary {
	this := GuidanceAzureReservationsConfigServicesAzureVmsSummary{}
	return &this
}

// GetTotalSavings returns the TotalSavings field value if set, zero value otherwise.
func (o *GuidanceAzureReservationsConfigServicesAzureVmsSummary) GetTotalSavings() float32 {
	if o == nil || o.TotalSavings == nil {
		var ret float32
		return ret
	}
	return *o.TotalSavings
}

// GetTotalSavingsOk returns a tuple with the TotalSavings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceAzureReservationsConfigServicesAzureVmsSummary) GetTotalSavingsOk() (*float32, bool) {
	if o == nil || o.TotalSavings == nil {
		return nil, false
	}
	return o.TotalSavings, true
}

// HasTotalSavings returns a boolean if a field has been set.
func (o *GuidanceAzureReservationsConfigServicesAzureVmsSummary) HasTotalSavings() bool {
	if o != nil && o.TotalSavings != nil {
		return true
	}

	return false
}

// SetTotalSavings gets a reference to the given float32 and assigns it to the TotalSavings field.
func (o *GuidanceAzureReservationsConfigServicesAzureVmsSummary) SetTotalSavings(v float32) {
	o.TotalSavings = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *GuidanceAzureReservationsConfigServicesAzureVmsSummary) GetCurrencyCode() string {
	if o == nil || o.CurrencyCode == nil {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceAzureReservationsConfigServicesAzureVmsSummary) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || o.CurrencyCode == nil {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *GuidanceAzureReservationsConfigServicesAzureVmsSummary) HasCurrencyCode() bool {
	if o != nil && o.CurrencyCode != nil {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *GuidanceAzureReservationsConfigServicesAzureVmsSummary) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetTotalSavingsPercent returns the TotalSavingsPercent field value if set, zero value otherwise.
func (o *GuidanceAzureReservationsConfigServicesAzureVmsSummary) GetTotalSavingsPercent() float32 {
	if o == nil || o.TotalSavingsPercent == nil {
		var ret float32
		return ret
	}
	return *o.TotalSavingsPercent
}

// GetTotalSavingsPercentOk returns a tuple with the TotalSavingsPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceAzureReservationsConfigServicesAzureVmsSummary) GetTotalSavingsPercentOk() (*float32, bool) {
	if o == nil || o.TotalSavingsPercent == nil {
		return nil, false
	}
	return o.TotalSavingsPercent, true
}

// HasTotalSavingsPercent returns a boolean if a field has been set.
func (o *GuidanceAzureReservationsConfigServicesAzureVmsSummary) HasTotalSavingsPercent() bool {
	if o != nil && o.TotalSavingsPercent != nil {
		return true
	}

	return false
}

// SetTotalSavingsPercent gets a reference to the given float32 and assigns it to the TotalSavingsPercent field.
func (o *GuidanceAzureReservationsConfigServicesAzureVmsSummary) SetTotalSavingsPercent(v float32) {
	o.TotalSavingsPercent = &v
}

// GetTerm returns the Term field value if set, zero value otherwise.
func (o *GuidanceAzureReservationsConfigServicesAzureVmsSummary) GetTerm() string {
	if o == nil || o.Term == nil {
		var ret string
		return ret
	}
	return *o.Term
}

// GetTermOk returns a tuple with the Term field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceAzureReservationsConfigServicesAzureVmsSummary) GetTermOk() (*string, bool) {
	if o == nil || o.Term == nil {
		return nil, false
	}
	return o.Term, true
}

// HasTerm returns a boolean if a field has been set.
func (o *GuidanceAzureReservationsConfigServicesAzureVmsSummary) HasTerm() bool {
	if o != nil && o.Term != nil {
		return true
	}

	return false
}

// SetTerm gets a reference to the given string and assigns it to the Term field.
func (o *GuidanceAzureReservationsConfigServicesAzureVmsSummary) SetTerm(v string) {
	o.Term = &v
}

// GetPaymentOption returns the PaymentOption field value if set, zero value otherwise.
func (o *GuidanceAzureReservationsConfigServicesAzureVmsSummary) GetPaymentOption() string {
	if o == nil || o.PaymentOption == nil {
		var ret string
		return ret
	}
	return *o.PaymentOption
}

// GetPaymentOptionOk returns a tuple with the PaymentOption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceAzureReservationsConfigServicesAzureVmsSummary) GetPaymentOptionOk() (*string, bool) {
	if o == nil || o.PaymentOption == nil {
		return nil, false
	}
	return o.PaymentOption, true
}

// HasPaymentOption returns a boolean if a field has been set.
func (o *GuidanceAzureReservationsConfigServicesAzureVmsSummary) HasPaymentOption() bool {
	if o != nil && o.PaymentOption != nil {
		return true
	}

	return false
}

// SetPaymentOption gets a reference to the given string and assigns it to the PaymentOption field.
func (o *GuidanceAzureReservationsConfigServicesAzureVmsSummary) SetPaymentOption(v string) {
	o.PaymentOption = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *GuidanceAzureReservationsConfigServicesAzureVmsSummary) GetService() string {
	if o == nil || o.Service == nil {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceAzureReservationsConfigServicesAzureVmsSummary) GetServiceOk() (*string, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *GuidanceAzureReservationsConfigServicesAzureVmsSummary) HasService() bool {
	if o != nil && o.Service != nil {
		return true
	}

	return false
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *GuidanceAzureReservationsConfigServicesAzureVmsSummary) SetService(v string) {
	o.Service = &v
}

// GetOnDemandCount returns the OnDemandCount field value if set, zero value otherwise.
func (o *GuidanceAzureReservationsConfigServicesAzureVmsSummary) GetOnDemandCount() int64 {
	if o == nil || o.OnDemandCount == nil {
		var ret int64
		return ret
	}
	return *o.OnDemandCount
}

// GetOnDemandCountOk returns a tuple with the OnDemandCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceAzureReservationsConfigServicesAzureVmsSummary) GetOnDemandCountOk() (*int64, bool) {
	if o == nil || o.OnDemandCount == nil {
		return nil, false
	}
	return o.OnDemandCount, true
}

// HasOnDemandCount returns a boolean if a field has been set.
func (o *GuidanceAzureReservationsConfigServicesAzureVmsSummary) HasOnDemandCount() bool {
	if o != nil && o.OnDemandCount != nil {
		return true
	}

	return false
}

// SetOnDemandCount gets a reference to the given int64 and assigns it to the OnDemandCount field.
func (o *GuidanceAzureReservationsConfigServicesAzureVmsSummary) SetOnDemandCount(v int64) {
	o.OnDemandCount = &v
}

// GetOnDemandCost returns the OnDemandCost field value if set, zero value otherwise.
func (o *GuidanceAzureReservationsConfigServicesAzureVmsSummary) GetOnDemandCost() float32 {
	if o == nil || o.OnDemandCost == nil {
		var ret float32
		return ret
	}
	return *o.OnDemandCost
}

// GetOnDemandCostOk returns a tuple with the OnDemandCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceAzureReservationsConfigServicesAzureVmsSummary) GetOnDemandCostOk() (*float32, bool) {
	if o == nil || o.OnDemandCost == nil {
		return nil, false
	}
	return o.OnDemandCost, true
}

// HasOnDemandCost returns a boolean if a field has been set.
func (o *GuidanceAzureReservationsConfigServicesAzureVmsSummary) HasOnDemandCost() bool {
	if o != nil && o.OnDemandCost != nil {
		return true
	}

	return false
}

// SetOnDemandCost gets a reference to the given float32 and assigns it to the OnDemandCost field.
func (o *GuidanceAzureReservationsConfigServicesAzureVmsSummary) SetOnDemandCost(v float32) {
	o.OnDemandCost = &v
}

// GetReservedCount returns the ReservedCount field value if set, zero value otherwise.
func (o *GuidanceAzureReservationsConfigServicesAzureVmsSummary) GetReservedCount() int64 {
	if o == nil || o.ReservedCount == nil {
		var ret int64
		return ret
	}
	return *o.ReservedCount
}

// GetReservedCountOk returns a tuple with the ReservedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceAzureReservationsConfigServicesAzureVmsSummary) GetReservedCountOk() (*int64, bool) {
	if o == nil || o.ReservedCount == nil {
		return nil, false
	}
	return o.ReservedCount, true
}

// HasReservedCount returns a boolean if a field has been set.
func (o *GuidanceAzureReservationsConfigServicesAzureVmsSummary) HasReservedCount() bool {
	if o != nil && o.ReservedCount != nil {
		return true
	}

	return false
}

// SetReservedCount gets a reference to the given int64 and assigns it to the ReservedCount field.
func (o *GuidanceAzureReservationsConfigServicesAzureVmsSummary) SetReservedCount(v int64) {
	o.ReservedCount = &v
}

// GetReservedCost returns the ReservedCost field value if set, zero value otherwise.
func (o *GuidanceAzureReservationsConfigServicesAzureVmsSummary) GetReservedCost() int64 {
	if o == nil || o.ReservedCost == nil {
		var ret int64
		return ret
	}
	return *o.ReservedCost
}

// GetReservedCostOk returns a tuple with the ReservedCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceAzureReservationsConfigServicesAzureVmsSummary) GetReservedCostOk() (*int64, bool) {
	if o == nil || o.ReservedCost == nil {
		return nil, false
	}
	return o.ReservedCost, true
}

// HasReservedCost returns a boolean if a field has been set.
func (o *GuidanceAzureReservationsConfigServicesAzureVmsSummary) HasReservedCost() bool {
	if o != nil && o.ReservedCost != nil {
		return true
	}

	return false
}

// SetReservedCost gets a reference to the given int64 and assigns it to the ReservedCost field.
func (o *GuidanceAzureReservationsConfigServicesAzureVmsSummary) SetReservedCost(v int64) {
	o.ReservedCost = &v
}

// GetRecommendedCount returns the RecommendedCount field value if set, zero value otherwise.
func (o *GuidanceAzureReservationsConfigServicesAzureVmsSummary) GetRecommendedCount() int64 {
	if o == nil || o.RecommendedCount == nil {
		var ret int64
		return ret
	}
	return *o.RecommendedCount
}

// GetRecommendedCountOk returns a tuple with the RecommendedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceAzureReservationsConfigServicesAzureVmsSummary) GetRecommendedCountOk() (*int64, bool) {
	if o == nil || o.RecommendedCount == nil {
		return nil, false
	}
	return o.RecommendedCount, true
}

// HasRecommendedCount returns a boolean if a field has been set.
func (o *GuidanceAzureReservationsConfigServicesAzureVmsSummary) HasRecommendedCount() bool {
	if o != nil && o.RecommendedCount != nil {
		return true
	}

	return false
}

// SetRecommendedCount gets a reference to the given int64 and assigns it to the RecommendedCount field.
func (o *GuidanceAzureReservationsConfigServicesAzureVmsSummary) SetRecommendedCount(v int64) {
	o.RecommendedCount = &v
}

// GetRecommendedCost returns the RecommendedCost field value if set, zero value otherwise.
func (o *GuidanceAzureReservationsConfigServicesAzureVmsSummary) GetRecommendedCost() float32 {
	if o == nil || o.RecommendedCost == nil {
		var ret float32
		return ret
	}
	return *o.RecommendedCost
}

// GetRecommendedCostOk returns a tuple with the RecommendedCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceAzureReservationsConfigServicesAzureVmsSummary) GetRecommendedCostOk() (*float32, bool) {
	if o == nil || o.RecommendedCost == nil {
		return nil, false
	}
	return o.RecommendedCost, true
}

// HasRecommendedCost returns a boolean if a field has been set.
func (o *GuidanceAzureReservationsConfigServicesAzureVmsSummary) HasRecommendedCost() bool {
	if o != nil && o.RecommendedCost != nil {
		return true
	}

	return false
}

// SetRecommendedCost gets a reference to the given float32 and assigns it to the RecommendedCost field.
func (o *GuidanceAzureReservationsConfigServicesAzureVmsSummary) SetRecommendedCost(v float32) {
	o.RecommendedCost = &v
}

func (o GuidanceAzureReservationsConfigServicesAzureVmsSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TotalSavings != nil {
		toSerialize["totalSavings"] = o.TotalSavings
	}
	if o.CurrencyCode != nil {
		toSerialize["currencyCode"] = o.CurrencyCode
	}
	if o.TotalSavingsPercent != nil {
		toSerialize["totalSavingsPercent"] = o.TotalSavingsPercent
	}
	if o.Term != nil {
		toSerialize["term"] = o.Term
	}
	if o.PaymentOption != nil {
		toSerialize["paymentOption"] = o.PaymentOption
	}
	if o.Service != nil {
		toSerialize["service"] = o.Service
	}
	if o.OnDemandCount != nil {
		toSerialize["onDemandCount"] = o.OnDemandCount
	}
	if o.OnDemandCost != nil {
		toSerialize["onDemandCost"] = o.OnDemandCost
	}
	if o.ReservedCount != nil {
		toSerialize["reservedCount"] = o.ReservedCount
	}
	if o.ReservedCost != nil {
		toSerialize["reservedCost"] = o.ReservedCost
	}
	if o.RecommendedCount != nil {
		toSerialize["recommendedCount"] = o.RecommendedCount
	}
	if o.RecommendedCost != nil {
		toSerialize["recommendedCost"] = o.RecommendedCost
	}
	return json.Marshal(toSerialize)
}

type NullableGuidanceAzureReservationsConfigServicesAzureVmsSummary struct {
	value *GuidanceAzureReservationsConfigServicesAzureVmsSummary
	isSet bool
}

func (v NullableGuidanceAzureReservationsConfigServicesAzureVmsSummary) Get() *GuidanceAzureReservationsConfigServicesAzureVmsSummary {
	return v.value
}

func (v *NullableGuidanceAzureReservationsConfigServicesAzureVmsSummary) Set(val *GuidanceAzureReservationsConfigServicesAzureVmsSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableGuidanceAzureReservationsConfigServicesAzureVmsSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableGuidanceAzureReservationsConfigServicesAzureVmsSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGuidanceAzureReservationsConfigServicesAzureVmsSummary(val *GuidanceAzureReservationsConfigServicesAzureVmsSummary) *NullableGuidanceAzureReservationsConfigServicesAzureVmsSummary {
	return &NullableGuidanceAzureReservationsConfigServicesAzureVmsSummary{value: val, isSet: true}
}

func (v NullableGuidanceAzureReservationsConfigServicesAzureVmsSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGuidanceAzureReservationsConfigServicesAzureVmsSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

