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
	"time"
)

// BillingServer struct for BillingServer
type BillingServer struct {
	RefType *string `json:"refType,omitempty"`
	RefUUID *string `json:"refUUID,omitempty"`
	RefId NullableString `json:"refId,omitempty"`
	StartDate *time.Time `json:"startDate,omitempty"`
	EndDate *time.Time `json:"endDate,omitempty"`
	Cost *float32 `json:"cost,omitempty"`
	Price *float32 `json:"price,omitempty"`
	NumUnits *float32 `json:"numUnits,omitempty"`
	Unit *string `json:"unit,omitempty"`
	Currency *string `json:"currency,omitempty"`
	Usages *[]BillingServerUsages `json:"usages,omitempty"`
	NumUsages *int64 `json:"numUsages,omitempty"`
	TotalUsages *int64 `json:"totalUsages,omitempty"`
	HasMoreUsages *bool `json:"hasMoreUsages,omitempty"`
	FoundPricing *bool `json:"foundPricing,omitempty"`
	Name *string `json:"name,omitempty"`
	ServerUUID *string `json:"serverUUID,omitempty"`
	ServerUniqueId NullableString `json:"serverUniqueId,omitempty"`
	ServerExternalId *string `json:"serverExternalId,omitempty"`
	ServerInternalId NullableString `json:"serverInternalId,omitempty"`
	ResourcePoolId *int64 `json:"resourcePoolId,omitempty"`
	ResourcePoolName *string `json:"resourcePoolName,omitempty"`
}

// NewBillingServer instantiates a new BillingServer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingServer() *BillingServer {
	this := BillingServer{}
	return &this
}

// NewBillingServerWithDefaults instantiates a new BillingServer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingServerWithDefaults() *BillingServer {
	this := BillingServer{}
	return &this
}

// GetRefType returns the RefType field value if set, zero value otherwise.
func (o *BillingServer) GetRefType() string {
	if o == nil || o.RefType == nil {
		var ret string
		return ret
	}
	return *o.RefType
}

// GetRefTypeOk returns a tuple with the RefType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingServer) GetRefTypeOk() (*string, bool) {
	if o == nil || o.RefType == nil {
		return nil, false
	}
	return o.RefType, true
}

// HasRefType returns a boolean if a field has been set.
func (o *BillingServer) HasRefType() bool {
	if o != nil && o.RefType != nil {
		return true
	}

	return false
}

// SetRefType gets a reference to the given string and assigns it to the RefType field.
func (o *BillingServer) SetRefType(v string) {
	o.RefType = &v
}

// GetRefUUID returns the RefUUID field value if set, zero value otherwise.
func (o *BillingServer) GetRefUUID() string {
	if o == nil || o.RefUUID == nil {
		var ret string
		return ret
	}
	return *o.RefUUID
}

// GetRefUUIDOk returns a tuple with the RefUUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingServer) GetRefUUIDOk() (*string, bool) {
	if o == nil || o.RefUUID == nil {
		return nil, false
	}
	return o.RefUUID, true
}

// HasRefUUID returns a boolean if a field has been set.
func (o *BillingServer) HasRefUUID() bool {
	if o != nil && o.RefUUID != nil {
		return true
	}

	return false
}

// SetRefUUID gets a reference to the given string and assigns it to the RefUUID field.
func (o *BillingServer) SetRefUUID(v string) {
	o.RefUUID = &v
}

// GetRefId returns the RefId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingServer) GetRefId() string {
	if o == nil || o.RefId.Get() == nil {
		var ret string
		return ret
	}
	return *o.RefId.Get()
}

// GetRefIdOk returns a tuple with the RefId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingServer) GetRefIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RefId.Get(), o.RefId.IsSet()
}

// HasRefId returns a boolean if a field has been set.
func (o *BillingServer) HasRefId() bool {
	if o != nil && o.RefId.IsSet() {
		return true
	}

	return false
}

// SetRefId gets a reference to the given NullableString and assigns it to the RefId field.
func (o *BillingServer) SetRefId(v string) {
	o.RefId.Set(&v)
}
// SetRefIdNil sets the value for RefId to be an explicit nil
func (o *BillingServer) SetRefIdNil() {
	o.RefId.Set(nil)
}

// UnsetRefId ensures that no value is present for RefId, not even an explicit nil
func (o *BillingServer) UnsetRefId() {
	o.RefId.Unset()
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *BillingServer) GetStartDate() time.Time {
	if o == nil || o.StartDate == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingServer) GetStartDateOk() (*time.Time, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *BillingServer) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *BillingServer) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *BillingServer) GetEndDate() time.Time {
	if o == nil || o.EndDate == nil {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingServer) GetEndDateOk() (*time.Time, bool) {
	if o == nil || o.EndDate == nil {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *BillingServer) HasEndDate() bool {
	if o != nil && o.EndDate != nil {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *BillingServer) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetCost returns the Cost field value if set, zero value otherwise.
func (o *BillingServer) GetCost() float32 {
	if o == nil || o.Cost == nil {
		var ret float32
		return ret
	}
	return *o.Cost
}

// GetCostOk returns a tuple with the Cost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingServer) GetCostOk() (*float32, bool) {
	if o == nil || o.Cost == nil {
		return nil, false
	}
	return o.Cost, true
}

// HasCost returns a boolean if a field has been set.
func (o *BillingServer) HasCost() bool {
	if o != nil && o.Cost != nil {
		return true
	}

	return false
}

// SetCost gets a reference to the given float32 and assigns it to the Cost field.
func (o *BillingServer) SetCost(v float32) {
	o.Cost = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *BillingServer) GetPrice() float32 {
	if o == nil || o.Price == nil {
		var ret float32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingServer) GetPriceOk() (*float32, bool) {
	if o == nil || o.Price == nil {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *BillingServer) HasPrice() bool {
	if o != nil && o.Price != nil {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float32 and assigns it to the Price field.
func (o *BillingServer) SetPrice(v float32) {
	o.Price = &v
}

// GetNumUnits returns the NumUnits field value if set, zero value otherwise.
func (o *BillingServer) GetNumUnits() float32 {
	if o == nil || o.NumUnits == nil {
		var ret float32
		return ret
	}
	return *o.NumUnits
}

// GetNumUnitsOk returns a tuple with the NumUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingServer) GetNumUnitsOk() (*float32, bool) {
	if o == nil || o.NumUnits == nil {
		return nil, false
	}
	return o.NumUnits, true
}

// HasNumUnits returns a boolean if a field has been set.
func (o *BillingServer) HasNumUnits() bool {
	if o != nil && o.NumUnits != nil {
		return true
	}

	return false
}

// SetNumUnits gets a reference to the given float32 and assigns it to the NumUnits field.
func (o *BillingServer) SetNumUnits(v float32) {
	o.NumUnits = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *BillingServer) GetUnit() string {
	if o == nil || o.Unit == nil {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingServer) GetUnitOk() (*string, bool) {
	if o == nil || o.Unit == nil {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *BillingServer) HasUnit() bool {
	if o != nil && o.Unit != nil {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *BillingServer) SetUnit(v string) {
	o.Unit = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *BillingServer) GetCurrency() string {
	if o == nil || o.Currency == nil {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingServer) GetCurrencyOk() (*string, bool) {
	if o == nil || o.Currency == nil {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *BillingServer) HasCurrency() bool {
	if o != nil && o.Currency != nil {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *BillingServer) SetCurrency(v string) {
	o.Currency = &v
}

// GetUsages returns the Usages field value if set, zero value otherwise.
func (o *BillingServer) GetUsages() []BillingServerUsages {
	if o == nil || o.Usages == nil {
		var ret []BillingServerUsages
		return ret
	}
	return *o.Usages
}

// GetUsagesOk returns a tuple with the Usages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingServer) GetUsagesOk() (*[]BillingServerUsages, bool) {
	if o == nil || o.Usages == nil {
		return nil, false
	}
	return o.Usages, true
}

// HasUsages returns a boolean if a field has been set.
func (o *BillingServer) HasUsages() bool {
	if o != nil && o.Usages != nil {
		return true
	}

	return false
}

// SetUsages gets a reference to the given []BillingServerUsages and assigns it to the Usages field.
func (o *BillingServer) SetUsages(v []BillingServerUsages) {
	o.Usages = &v
}

// GetNumUsages returns the NumUsages field value if set, zero value otherwise.
func (o *BillingServer) GetNumUsages() int64 {
	if o == nil || o.NumUsages == nil {
		var ret int64
		return ret
	}
	return *o.NumUsages
}

// GetNumUsagesOk returns a tuple with the NumUsages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingServer) GetNumUsagesOk() (*int64, bool) {
	if o == nil || o.NumUsages == nil {
		return nil, false
	}
	return o.NumUsages, true
}

// HasNumUsages returns a boolean if a field has been set.
func (o *BillingServer) HasNumUsages() bool {
	if o != nil && o.NumUsages != nil {
		return true
	}

	return false
}

// SetNumUsages gets a reference to the given int64 and assigns it to the NumUsages field.
func (o *BillingServer) SetNumUsages(v int64) {
	o.NumUsages = &v
}

// GetTotalUsages returns the TotalUsages field value if set, zero value otherwise.
func (o *BillingServer) GetTotalUsages() int64 {
	if o == nil || o.TotalUsages == nil {
		var ret int64
		return ret
	}
	return *o.TotalUsages
}

// GetTotalUsagesOk returns a tuple with the TotalUsages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingServer) GetTotalUsagesOk() (*int64, bool) {
	if o == nil || o.TotalUsages == nil {
		return nil, false
	}
	return o.TotalUsages, true
}

// HasTotalUsages returns a boolean if a field has been set.
func (o *BillingServer) HasTotalUsages() bool {
	if o != nil && o.TotalUsages != nil {
		return true
	}

	return false
}

// SetTotalUsages gets a reference to the given int64 and assigns it to the TotalUsages field.
func (o *BillingServer) SetTotalUsages(v int64) {
	o.TotalUsages = &v
}

// GetHasMoreUsages returns the HasMoreUsages field value if set, zero value otherwise.
func (o *BillingServer) GetHasMoreUsages() bool {
	if o == nil || o.HasMoreUsages == nil {
		var ret bool
		return ret
	}
	return *o.HasMoreUsages
}

// GetHasMoreUsagesOk returns a tuple with the HasMoreUsages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingServer) GetHasMoreUsagesOk() (*bool, bool) {
	if o == nil || o.HasMoreUsages == nil {
		return nil, false
	}
	return o.HasMoreUsages, true
}

// HasHasMoreUsages returns a boolean if a field has been set.
func (o *BillingServer) HasHasMoreUsages() bool {
	if o != nil && o.HasMoreUsages != nil {
		return true
	}

	return false
}

// SetHasMoreUsages gets a reference to the given bool and assigns it to the HasMoreUsages field.
func (o *BillingServer) SetHasMoreUsages(v bool) {
	o.HasMoreUsages = &v
}

// GetFoundPricing returns the FoundPricing field value if set, zero value otherwise.
func (o *BillingServer) GetFoundPricing() bool {
	if o == nil || o.FoundPricing == nil {
		var ret bool
		return ret
	}
	return *o.FoundPricing
}

// GetFoundPricingOk returns a tuple with the FoundPricing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingServer) GetFoundPricingOk() (*bool, bool) {
	if o == nil || o.FoundPricing == nil {
		return nil, false
	}
	return o.FoundPricing, true
}

// HasFoundPricing returns a boolean if a field has been set.
func (o *BillingServer) HasFoundPricing() bool {
	if o != nil && o.FoundPricing != nil {
		return true
	}

	return false
}

// SetFoundPricing gets a reference to the given bool and assigns it to the FoundPricing field.
func (o *BillingServer) SetFoundPricing(v bool) {
	o.FoundPricing = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BillingServer) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingServer) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BillingServer) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BillingServer) SetName(v string) {
	o.Name = &v
}

// GetServerUUID returns the ServerUUID field value if set, zero value otherwise.
func (o *BillingServer) GetServerUUID() string {
	if o == nil || o.ServerUUID == nil {
		var ret string
		return ret
	}
	return *o.ServerUUID
}

// GetServerUUIDOk returns a tuple with the ServerUUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingServer) GetServerUUIDOk() (*string, bool) {
	if o == nil || o.ServerUUID == nil {
		return nil, false
	}
	return o.ServerUUID, true
}

// HasServerUUID returns a boolean if a field has been set.
func (o *BillingServer) HasServerUUID() bool {
	if o != nil && o.ServerUUID != nil {
		return true
	}

	return false
}

// SetServerUUID gets a reference to the given string and assigns it to the ServerUUID field.
func (o *BillingServer) SetServerUUID(v string) {
	o.ServerUUID = &v
}

// GetServerUniqueId returns the ServerUniqueId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingServer) GetServerUniqueId() string {
	if o == nil || o.ServerUniqueId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ServerUniqueId.Get()
}

// GetServerUniqueIdOk returns a tuple with the ServerUniqueId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingServer) GetServerUniqueIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ServerUniqueId.Get(), o.ServerUniqueId.IsSet()
}

// HasServerUniqueId returns a boolean if a field has been set.
func (o *BillingServer) HasServerUniqueId() bool {
	if o != nil && o.ServerUniqueId.IsSet() {
		return true
	}

	return false
}

// SetServerUniqueId gets a reference to the given NullableString and assigns it to the ServerUniqueId field.
func (o *BillingServer) SetServerUniqueId(v string) {
	o.ServerUniqueId.Set(&v)
}
// SetServerUniqueIdNil sets the value for ServerUniqueId to be an explicit nil
func (o *BillingServer) SetServerUniqueIdNil() {
	o.ServerUniqueId.Set(nil)
}

// UnsetServerUniqueId ensures that no value is present for ServerUniqueId, not even an explicit nil
func (o *BillingServer) UnsetServerUniqueId() {
	o.ServerUniqueId.Unset()
}

// GetServerExternalId returns the ServerExternalId field value if set, zero value otherwise.
func (o *BillingServer) GetServerExternalId() string {
	if o == nil || o.ServerExternalId == nil {
		var ret string
		return ret
	}
	return *o.ServerExternalId
}

// GetServerExternalIdOk returns a tuple with the ServerExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingServer) GetServerExternalIdOk() (*string, bool) {
	if o == nil || o.ServerExternalId == nil {
		return nil, false
	}
	return o.ServerExternalId, true
}

// HasServerExternalId returns a boolean if a field has been set.
func (o *BillingServer) HasServerExternalId() bool {
	if o != nil && o.ServerExternalId != nil {
		return true
	}

	return false
}

// SetServerExternalId gets a reference to the given string and assigns it to the ServerExternalId field.
func (o *BillingServer) SetServerExternalId(v string) {
	o.ServerExternalId = &v
}

// GetServerInternalId returns the ServerInternalId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingServer) GetServerInternalId() string {
	if o == nil || o.ServerInternalId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ServerInternalId.Get()
}

// GetServerInternalIdOk returns a tuple with the ServerInternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingServer) GetServerInternalIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ServerInternalId.Get(), o.ServerInternalId.IsSet()
}

// HasServerInternalId returns a boolean if a field has been set.
func (o *BillingServer) HasServerInternalId() bool {
	if o != nil && o.ServerInternalId.IsSet() {
		return true
	}

	return false
}

// SetServerInternalId gets a reference to the given NullableString and assigns it to the ServerInternalId field.
func (o *BillingServer) SetServerInternalId(v string) {
	o.ServerInternalId.Set(&v)
}
// SetServerInternalIdNil sets the value for ServerInternalId to be an explicit nil
func (o *BillingServer) SetServerInternalIdNil() {
	o.ServerInternalId.Set(nil)
}

// UnsetServerInternalId ensures that no value is present for ServerInternalId, not even an explicit nil
func (o *BillingServer) UnsetServerInternalId() {
	o.ServerInternalId.Unset()
}

// GetResourcePoolId returns the ResourcePoolId field value if set, zero value otherwise.
func (o *BillingServer) GetResourcePoolId() int64 {
	if o == nil || o.ResourcePoolId == nil {
		var ret int64
		return ret
	}
	return *o.ResourcePoolId
}

// GetResourcePoolIdOk returns a tuple with the ResourcePoolId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingServer) GetResourcePoolIdOk() (*int64, bool) {
	if o == nil || o.ResourcePoolId == nil {
		return nil, false
	}
	return o.ResourcePoolId, true
}

// HasResourcePoolId returns a boolean if a field has been set.
func (o *BillingServer) HasResourcePoolId() bool {
	if o != nil && o.ResourcePoolId != nil {
		return true
	}

	return false
}

// SetResourcePoolId gets a reference to the given int64 and assigns it to the ResourcePoolId field.
func (o *BillingServer) SetResourcePoolId(v int64) {
	o.ResourcePoolId = &v
}

// GetResourcePoolName returns the ResourcePoolName field value if set, zero value otherwise.
func (o *BillingServer) GetResourcePoolName() string {
	if o == nil || o.ResourcePoolName == nil {
		var ret string
		return ret
	}
	return *o.ResourcePoolName
}

// GetResourcePoolNameOk returns a tuple with the ResourcePoolName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingServer) GetResourcePoolNameOk() (*string, bool) {
	if o == nil || o.ResourcePoolName == nil {
		return nil, false
	}
	return o.ResourcePoolName, true
}

// HasResourcePoolName returns a boolean if a field has been set.
func (o *BillingServer) HasResourcePoolName() bool {
	if o != nil && o.ResourcePoolName != nil {
		return true
	}

	return false
}

// SetResourcePoolName gets a reference to the given string and assigns it to the ResourcePoolName field.
func (o *BillingServer) SetResourcePoolName(v string) {
	o.ResourcePoolName = &v
}

func (o BillingServer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RefType != nil {
		toSerialize["refType"] = o.RefType
	}
	if o.RefUUID != nil {
		toSerialize["refUUID"] = o.RefUUID
	}
	if o.RefId.IsSet() {
		toSerialize["refId"] = o.RefId.Get()
	}
	if o.StartDate != nil {
		toSerialize["startDate"] = o.StartDate
	}
	if o.EndDate != nil {
		toSerialize["endDate"] = o.EndDate
	}
	if o.Cost != nil {
		toSerialize["cost"] = o.Cost
	}
	if o.Price != nil {
		toSerialize["price"] = o.Price
	}
	if o.NumUnits != nil {
		toSerialize["numUnits"] = o.NumUnits
	}
	if o.Unit != nil {
		toSerialize["unit"] = o.Unit
	}
	if o.Currency != nil {
		toSerialize["currency"] = o.Currency
	}
	if o.Usages != nil {
		toSerialize["usages"] = o.Usages
	}
	if o.NumUsages != nil {
		toSerialize["numUsages"] = o.NumUsages
	}
	if o.TotalUsages != nil {
		toSerialize["totalUsages"] = o.TotalUsages
	}
	if o.HasMoreUsages != nil {
		toSerialize["hasMoreUsages"] = o.HasMoreUsages
	}
	if o.FoundPricing != nil {
		toSerialize["foundPricing"] = o.FoundPricing
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ServerUUID != nil {
		toSerialize["serverUUID"] = o.ServerUUID
	}
	if o.ServerUniqueId.IsSet() {
		toSerialize["serverUniqueId"] = o.ServerUniqueId.Get()
	}
	if o.ServerExternalId != nil {
		toSerialize["serverExternalId"] = o.ServerExternalId
	}
	if o.ServerInternalId.IsSet() {
		toSerialize["serverInternalId"] = o.ServerInternalId.Get()
	}
	if o.ResourcePoolId != nil {
		toSerialize["resourcePoolId"] = o.ResourcePoolId
	}
	if o.ResourcePoolName != nil {
		toSerialize["resourcePoolName"] = o.ResourcePoolName
	}
	return json.Marshal(toSerialize)
}

type NullableBillingServer struct {
	value *BillingServer
	isSet bool
}

func (v NullableBillingServer) Get() *BillingServer {
	return v.value
}

func (v *NullableBillingServer) Set(val *BillingServer) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingServer) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingServer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingServer(val *BillingServer) *NullableBillingServer {
	return &NullableBillingServer{value: val, isSet: true}
}

func (v NullableBillingServer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingServer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

