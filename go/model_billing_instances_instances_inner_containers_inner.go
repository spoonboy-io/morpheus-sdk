/*
Morpheus API

Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.

API version: 6.1.1
Contact: dev@morpheusdata.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the BillingInstancesInstancesInnerContainersInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillingInstancesInstancesInnerContainersInner{}

// BillingInstancesInstancesInnerContainersInner struct for BillingInstancesInstancesInnerContainersInner
type BillingInstancesInstancesInnerContainersInner struct {
	RefType *string `json:"refType,omitempty"`
	RefUUID *string `json:"refUUID,omitempty"`
	RefId *int64 `json:"refId,omitempty"`
	StartDate *time.Time `json:"startDate,omitempty"`
	EndDate *time.Time `json:"endDate,omitempty"`
	Cost *float32 `json:"cost,omitempty"`
	Price *float32 `json:"price,omitempty"`
	NumUnits *float32 `json:"numUnits,omitempty"`
	Unit *string `json:"unit,omitempty"`
	Currency *string `json:"currency,omitempty"`
	Usages []BillingInstancesInstancesInnerContainersInnerUsagesInner `json:"usages,omitempty"`
	NumUsages *int64 `json:"numUsages,omitempty"`
	TotalUsages *int64 `json:"totalUsages,omitempty"`
	HasMoreUsages *bool `json:"hasMoreUsages,omitempty"`
	FoundPricing *bool `json:"foundPricing,omitempty"`
	Name *string `json:"name,omitempty"`
	ServerId *int64 `json:"serverId,omitempty"`
	ServerUUID *string `json:"serverUUID,omitempty"`
	ServerUniqueId *string `json:"serverUniqueId,omitempty"`
	ServerExternalId *string `json:"serverExternalId,omitempty"`
	ServerInternalId *string `json:"serverInternalId,omitempty"`
	ResourcePoolId *int64 `json:"resourcePoolId,omitempty"`
	ResourcePoolName *string `json:"resourcePoolName,omitempty"`
}

// NewBillingInstancesInstancesInnerContainersInner instantiates a new BillingInstancesInstancesInnerContainersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingInstancesInstancesInnerContainersInner() *BillingInstancesInstancesInnerContainersInner {
	this := BillingInstancesInstancesInnerContainersInner{}
	return &this
}

// NewBillingInstancesInstancesInnerContainersInnerWithDefaults instantiates a new BillingInstancesInstancesInnerContainersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingInstancesInstancesInnerContainersInnerWithDefaults() *BillingInstancesInstancesInnerContainersInner {
	this := BillingInstancesInstancesInnerContainersInner{}
	return &this
}

// GetRefType returns the RefType field value if set, zero value otherwise.
func (o *BillingInstancesInstancesInnerContainersInner) GetRefType() string {
	if o == nil || IsNil(o.RefType) {
		var ret string
		return ret
	}
	return *o.RefType
}

// GetRefTypeOk returns a tuple with the RefType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInstancesInstancesInnerContainersInner) GetRefTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RefType) {
		return nil, false
	}
	return o.RefType, true
}

// HasRefType returns a boolean if a field has been set.
func (o *BillingInstancesInstancesInnerContainersInner) HasRefType() bool {
	if o != nil && !IsNil(o.RefType) {
		return true
	}

	return false
}

// SetRefType gets a reference to the given string and assigns it to the RefType field.
func (o *BillingInstancesInstancesInnerContainersInner) SetRefType(v string) {
	o.RefType = &v
}

// GetRefUUID returns the RefUUID field value if set, zero value otherwise.
func (o *BillingInstancesInstancesInnerContainersInner) GetRefUUID() string {
	if o == nil || IsNil(o.RefUUID) {
		var ret string
		return ret
	}
	return *o.RefUUID
}

// GetRefUUIDOk returns a tuple with the RefUUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInstancesInstancesInnerContainersInner) GetRefUUIDOk() (*string, bool) {
	if o == nil || IsNil(o.RefUUID) {
		return nil, false
	}
	return o.RefUUID, true
}

// HasRefUUID returns a boolean if a field has been set.
func (o *BillingInstancesInstancesInnerContainersInner) HasRefUUID() bool {
	if o != nil && !IsNil(o.RefUUID) {
		return true
	}

	return false
}

// SetRefUUID gets a reference to the given string and assigns it to the RefUUID field.
func (o *BillingInstancesInstancesInnerContainersInner) SetRefUUID(v string) {
	o.RefUUID = &v
}

// GetRefId returns the RefId field value if set, zero value otherwise.
func (o *BillingInstancesInstancesInnerContainersInner) GetRefId() int64 {
	if o == nil || IsNil(o.RefId) {
		var ret int64
		return ret
	}
	return *o.RefId
}

// GetRefIdOk returns a tuple with the RefId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInstancesInstancesInnerContainersInner) GetRefIdOk() (*int64, bool) {
	if o == nil || IsNil(o.RefId) {
		return nil, false
	}
	return o.RefId, true
}

// HasRefId returns a boolean if a field has been set.
func (o *BillingInstancesInstancesInnerContainersInner) HasRefId() bool {
	if o != nil && !IsNil(o.RefId) {
		return true
	}

	return false
}

// SetRefId gets a reference to the given int64 and assigns it to the RefId field.
func (o *BillingInstancesInstancesInnerContainersInner) SetRefId(v int64) {
	o.RefId = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *BillingInstancesInstancesInnerContainersInner) GetStartDate() time.Time {
	if o == nil || IsNil(o.StartDate) {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInstancesInstancesInnerContainersInner) GetStartDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *BillingInstancesInstancesInnerContainersInner) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *BillingInstancesInstancesInnerContainersInner) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *BillingInstancesInstancesInnerContainersInner) GetEndDate() time.Time {
	if o == nil || IsNil(o.EndDate) {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInstancesInstancesInnerContainersInner) GetEndDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *BillingInstancesInstancesInnerContainersInner) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *BillingInstancesInstancesInnerContainersInner) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetCost returns the Cost field value if set, zero value otherwise.
func (o *BillingInstancesInstancesInnerContainersInner) GetCost() float32 {
	if o == nil || IsNil(o.Cost) {
		var ret float32
		return ret
	}
	return *o.Cost
}

// GetCostOk returns a tuple with the Cost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInstancesInstancesInnerContainersInner) GetCostOk() (*float32, bool) {
	if o == nil || IsNil(o.Cost) {
		return nil, false
	}
	return o.Cost, true
}

// HasCost returns a boolean if a field has been set.
func (o *BillingInstancesInstancesInnerContainersInner) HasCost() bool {
	if o != nil && !IsNil(o.Cost) {
		return true
	}

	return false
}

// SetCost gets a reference to the given float32 and assigns it to the Cost field.
func (o *BillingInstancesInstancesInnerContainersInner) SetCost(v float32) {
	o.Cost = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *BillingInstancesInstancesInnerContainersInner) GetPrice() float32 {
	if o == nil || IsNil(o.Price) {
		var ret float32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInstancesInstancesInnerContainersInner) GetPriceOk() (*float32, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *BillingInstancesInstancesInnerContainersInner) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float32 and assigns it to the Price field.
func (o *BillingInstancesInstancesInnerContainersInner) SetPrice(v float32) {
	o.Price = &v
}

// GetNumUnits returns the NumUnits field value if set, zero value otherwise.
func (o *BillingInstancesInstancesInnerContainersInner) GetNumUnits() float32 {
	if o == nil || IsNil(o.NumUnits) {
		var ret float32
		return ret
	}
	return *o.NumUnits
}

// GetNumUnitsOk returns a tuple with the NumUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInstancesInstancesInnerContainersInner) GetNumUnitsOk() (*float32, bool) {
	if o == nil || IsNil(o.NumUnits) {
		return nil, false
	}
	return o.NumUnits, true
}

// HasNumUnits returns a boolean if a field has been set.
func (o *BillingInstancesInstancesInnerContainersInner) HasNumUnits() bool {
	if o != nil && !IsNil(o.NumUnits) {
		return true
	}

	return false
}

// SetNumUnits gets a reference to the given float32 and assigns it to the NumUnits field.
func (o *BillingInstancesInstancesInnerContainersInner) SetNumUnits(v float32) {
	o.NumUnits = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *BillingInstancesInstancesInnerContainersInner) GetUnit() string {
	if o == nil || IsNil(o.Unit) {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInstancesInstancesInnerContainersInner) GetUnitOk() (*string, bool) {
	if o == nil || IsNil(o.Unit) {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *BillingInstancesInstancesInnerContainersInner) HasUnit() bool {
	if o != nil && !IsNil(o.Unit) {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *BillingInstancesInstancesInnerContainersInner) SetUnit(v string) {
	o.Unit = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *BillingInstancesInstancesInnerContainersInner) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInstancesInstancesInnerContainersInner) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *BillingInstancesInstancesInnerContainersInner) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *BillingInstancesInstancesInnerContainersInner) SetCurrency(v string) {
	o.Currency = &v
}

// GetUsages returns the Usages field value if set, zero value otherwise.
func (o *BillingInstancesInstancesInnerContainersInner) GetUsages() []BillingInstancesInstancesInnerContainersInnerUsagesInner {
	if o == nil || IsNil(o.Usages) {
		var ret []BillingInstancesInstancesInnerContainersInnerUsagesInner
		return ret
	}
	return o.Usages
}

// GetUsagesOk returns a tuple with the Usages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInstancesInstancesInnerContainersInner) GetUsagesOk() ([]BillingInstancesInstancesInnerContainersInnerUsagesInner, bool) {
	if o == nil || IsNil(o.Usages) {
		return nil, false
	}
	return o.Usages, true
}

// HasUsages returns a boolean if a field has been set.
func (o *BillingInstancesInstancesInnerContainersInner) HasUsages() bool {
	if o != nil && !IsNil(o.Usages) {
		return true
	}

	return false
}

// SetUsages gets a reference to the given []BillingInstancesInstancesInnerContainersInnerUsagesInner and assigns it to the Usages field.
func (o *BillingInstancesInstancesInnerContainersInner) SetUsages(v []BillingInstancesInstancesInnerContainersInnerUsagesInner) {
	o.Usages = v
}

// GetNumUsages returns the NumUsages field value if set, zero value otherwise.
func (o *BillingInstancesInstancesInnerContainersInner) GetNumUsages() int64 {
	if o == nil || IsNil(o.NumUsages) {
		var ret int64
		return ret
	}
	return *o.NumUsages
}

// GetNumUsagesOk returns a tuple with the NumUsages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInstancesInstancesInnerContainersInner) GetNumUsagesOk() (*int64, bool) {
	if o == nil || IsNil(o.NumUsages) {
		return nil, false
	}
	return o.NumUsages, true
}

// HasNumUsages returns a boolean if a field has been set.
func (o *BillingInstancesInstancesInnerContainersInner) HasNumUsages() bool {
	if o != nil && !IsNil(o.NumUsages) {
		return true
	}

	return false
}

// SetNumUsages gets a reference to the given int64 and assigns it to the NumUsages field.
func (o *BillingInstancesInstancesInnerContainersInner) SetNumUsages(v int64) {
	o.NumUsages = &v
}

// GetTotalUsages returns the TotalUsages field value if set, zero value otherwise.
func (o *BillingInstancesInstancesInnerContainersInner) GetTotalUsages() int64 {
	if o == nil || IsNil(o.TotalUsages) {
		var ret int64
		return ret
	}
	return *o.TotalUsages
}

// GetTotalUsagesOk returns a tuple with the TotalUsages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInstancesInstancesInnerContainersInner) GetTotalUsagesOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalUsages) {
		return nil, false
	}
	return o.TotalUsages, true
}

// HasTotalUsages returns a boolean if a field has been set.
func (o *BillingInstancesInstancesInnerContainersInner) HasTotalUsages() bool {
	if o != nil && !IsNil(o.TotalUsages) {
		return true
	}

	return false
}

// SetTotalUsages gets a reference to the given int64 and assigns it to the TotalUsages field.
func (o *BillingInstancesInstancesInnerContainersInner) SetTotalUsages(v int64) {
	o.TotalUsages = &v
}

// GetHasMoreUsages returns the HasMoreUsages field value if set, zero value otherwise.
func (o *BillingInstancesInstancesInnerContainersInner) GetHasMoreUsages() bool {
	if o == nil || IsNil(o.HasMoreUsages) {
		var ret bool
		return ret
	}
	return *o.HasMoreUsages
}

// GetHasMoreUsagesOk returns a tuple with the HasMoreUsages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInstancesInstancesInnerContainersInner) GetHasMoreUsagesOk() (*bool, bool) {
	if o == nil || IsNil(o.HasMoreUsages) {
		return nil, false
	}
	return o.HasMoreUsages, true
}

// HasHasMoreUsages returns a boolean if a field has been set.
func (o *BillingInstancesInstancesInnerContainersInner) HasHasMoreUsages() bool {
	if o != nil && !IsNil(o.HasMoreUsages) {
		return true
	}

	return false
}

// SetHasMoreUsages gets a reference to the given bool and assigns it to the HasMoreUsages field.
func (o *BillingInstancesInstancesInnerContainersInner) SetHasMoreUsages(v bool) {
	o.HasMoreUsages = &v
}

// GetFoundPricing returns the FoundPricing field value if set, zero value otherwise.
func (o *BillingInstancesInstancesInnerContainersInner) GetFoundPricing() bool {
	if o == nil || IsNil(o.FoundPricing) {
		var ret bool
		return ret
	}
	return *o.FoundPricing
}

// GetFoundPricingOk returns a tuple with the FoundPricing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInstancesInstancesInnerContainersInner) GetFoundPricingOk() (*bool, bool) {
	if o == nil || IsNil(o.FoundPricing) {
		return nil, false
	}
	return o.FoundPricing, true
}

// HasFoundPricing returns a boolean if a field has been set.
func (o *BillingInstancesInstancesInnerContainersInner) HasFoundPricing() bool {
	if o != nil && !IsNil(o.FoundPricing) {
		return true
	}

	return false
}

// SetFoundPricing gets a reference to the given bool and assigns it to the FoundPricing field.
func (o *BillingInstancesInstancesInnerContainersInner) SetFoundPricing(v bool) {
	o.FoundPricing = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BillingInstancesInstancesInnerContainersInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInstancesInstancesInnerContainersInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BillingInstancesInstancesInnerContainersInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BillingInstancesInstancesInnerContainersInner) SetName(v string) {
	o.Name = &v
}

// GetServerId returns the ServerId field value if set, zero value otherwise.
func (o *BillingInstancesInstancesInnerContainersInner) GetServerId() int64 {
	if o == nil || IsNil(o.ServerId) {
		var ret int64
		return ret
	}
	return *o.ServerId
}

// GetServerIdOk returns a tuple with the ServerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInstancesInstancesInnerContainersInner) GetServerIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ServerId) {
		return nil, false
	}
	return o.ServerId, true
}

// HasServerId returns a boolean if a field has been set.
func (o *BillingInstancesInstancesInnerContainersInner) HasServerId() bool {
	if o != nil && !IsNil(o.ServerId) {
		return true
	}

	return false
}

// SetServerId gets a reference to the given int64 and assigns it to the ServerId field.
func (o *BillingInstancesInstancesInnerContainersInner) SetServerId(v int64) {
	o.ServerId = &v
}

// GetServerUUID returns the ServerUUID field value if set, zero value otherwise.
func (o *BillingInstancesInstancesInnerContainersInner) GetServerUUID() string {
	if o == nil || IsNil(o.ServerUUID) {
		var ret string
		return ret
	}
	return *o.ServerUUID
}

// GetServerUUIDOk returns a tuple with the ServerUUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInstancesInstancesInnerContainersInner) GetServerUUIDOk() (*string, bool) {
	if o == nil || IsNil(o.ServerUUID) {
		return nil, false
	}
	return o.ServerUUID, true
}

// HasServerUUID returns a boolean if a field has been set.
func (o *BillingInstancesInstancesInnerContainersInner) HasServerUUID() bool {
	if o != nil && !IsNil(o.ServerUUID) {
		return true
	}

	return false
}

// SetServerUUID gets a reference to the given string and assigns it to the ServerUUID field.
func (o *BillingInstancesInstancesInnerContainersInner) SetServerUUID(v string) {
	o.ServerUUID = &v
}

// GetServerUniqueId returns the ServerUniqueId field value if set, zero value otherwise.
func (o *BillingInstancesInstancesInnerContainersInner) GetServerUniqueId() string {
	if o == nil || IsNil(o.ServerUniqueId) {
		var ret string
		return ret
	}
	return *o.ServerUniqueId
}

// GetServerUniqueIdOk returns a tuple with the ServerUniqueId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInstancesInstancesInnerContainersInner) GetServerUniqueIdOk() (*string, bool) {
	if o == nil || IsNil(o.ServerUniqueId) {
		return nil, false
	}
	return o.ServerUniqueId, true
}

// HasServerUniqueId returns a boolean if a field has been set.
func (o *BillingInstancesInstancesInnerContainersInner) HasServerUniqueId() bool {
	if o != nil && !IsNil(o.ServerUniqueId) {
		return true
	}

	return false
}

// SetServerUniqueId gets a reference to the given string and assigns it to the ServerUniqueId field.
func (o *BillingInstancesInstancesInnerContainersInner) SetServerUniqueId(v string) {
	o.ServerUniqueId = &v
}

// GetServerExternalId returns the ServerExternalId field value if set, zero value otherwise.
func (o *BillingInstancesInstancesInnerContainersInner) GetServerExternalId() string {
	if o == nil || IsNil(o.ServerExternalId) {
		var ret string
		return ret
	}
	return *o.ServerExternalId
}

// GetServerExternalIdOk returns a tuple with the ServerExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInstancesInstancesInnerContainersInner) GetServerExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ServerExternalId) {
		return nil, false
	}
	return o.ServerExternalId, true
}

// HasServerExternalId returns a boolean if a field has been set.
func (o *BillingInstancesInstancesInnerContainersInner) HasServerExternalId() bool {
	if o != nil && !IsNil(o.ServerExternalId) {
		return true
	}

	return false
}

// SetServerExternalId gets a reference to the given string and assigns it to the ServerExternalId field.
func (o *BillingInstancesInstancesInnerContainersInner) SetServerExternalId(v string) {
	o.ServerExternalId = &v
}

// GetServerInternalId returns the ServerInternalId field value if set, zero value otherwise.
func (o *BillingInstancesInstancesInnerContainersInner) GetServerInternalId() string {
	if o == nil || IsNil(o.ServerInternalId) {
		var ret string
		return ret
	}
	return *o.ServerInternalId
}

// GetServerInternalIdOk returns a tuple with the ServerInternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInstancesInstancesInnerContainersInner) GetServerInternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ServerInternalId) {
		return nil, false
	}
	return o.ServerInternalId, true
}

// HasServerInternalId returns a boolean if a field has been set.
func (o *BillingInstancesInstancesInnerContainersInner) HasServerInternalId() bool {
	if o != nil && !IsNil(o.ServerInternalId) {
		return true
	}

	return false
}

// SetServerInternalId gets a reference to the given string and assigns it to the ServerInternalId field.
func (o *BillingInstancesInstancesInnerContainersInner) SetServerInternalId(v string) {
	o.ServerInternalId = &v
}

// GetResourcePoolId returns the ResourcePoolId field value if set, zero value otherwise.
func (o *BillingInstancesInstancesInnerContainersInner) GetResourcePoolId() int64 {
	if o == nil || IsNil(o.ResourcePoolId) {
		var ret int64
		return ret
	}
	return *o.ResourcePoolId
}

// GetResourcePoolIdOk returns a tuple with the ResourcePoolId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInstancesInstancesInnerContainersInner) GetResourcePoolIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ResourcePoolId) {
		return nil, false
	}
	return o.ResourcePoolId, true
}

// HasResourcePoolId returns a boolean if a field has been set.
func (o *BillingInstancesInstancesInnerContainersInner) HasResourcePoolId() bool {
	if o != nil && !IsNil(o.ResourcePoolId) {
		return true
	}

	return false
}

// SetResourcePoolId gets a reference to the given int64 and assigns it to the ResourcePoolId field.
func (o *BillingInstancesInstancesInnerContainersInner) SetResourcePoolId(v int64) {
	o.ResourcePoolId = &v
}

// GetResourcePoolName returns the ResourcePoolName field value if set, zero value otherwise.
func (o *BillingInstancesInstancesInnerContainersInner) GetResourcePoolName() string {
	if o == nil || IsNil(o.ResourcePoolName) {
		var ret string
		return ret
	}
	return *o.ResourcePoolName
}

// GetResourcePoolNameOk returns a tuple with the ResourcePoolName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInstancesInstancesInnerContainersInner) GetResourcePoolNameOk() (*string, bool) {
	if o == nil || IsNil(o.ResourcePoolName) {
		return nil, false
	}
	return o.ResourcePoolName, true
}

// HasResourcePoolName returns a boolean if a field has been set.
func (o *BillingInstancesInstancesInnerContainersInner) HasResourcePoolName() bool {
	if o != nil && !IsNil(o.ResourcePoolName) {
		return true
	}

	return false
}

// SetResourcePoolName gets a reference to the given string and assigns it to the ResourcePoolName field.
func (o *BillingInstancesInstancesInnerContainersInner) SetResourcePoolName(v string) {
	o.ResourcePoolName = &v
}

func (o BillingInstancesInstancesInnerContainersInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillingInstancesInstancesInnerContainersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RefType) {
		toSerialize["refType"] = o.RefType
	}
	if !IsNil(o.RefUUID) {
		toSerialize["refUUID"] = o.RefUUID
	}
	if !IsNil(o.RefId) {
		toSerialize["refId"] = o.RefId
	}
	if !IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	if !IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	if !IsNil(o.Cost) {
		toSerialize["cost"] = o.Cost
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.NumUnits) {
		toSerialize["numUnits"] = o.NumUnits
	}
	if !IsNil(o.Unit) {
		toSerialize["unit"] = o.Unit
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.Usages) {
		toSerialize["usages"] = o.Usages
	}
	if !IsNil(o.NumUsages) {
		toSerialize["numUsages"] = o.NumUsages
	}
	if !IsNil(o.TotalUsages) {
		toSerialize["totalUsages"] = o.TotalUsages
	}
	if !IsNil(o.HasMoreUsages) {
		toSerialize["hasMoreUsages"] = o.HasMoreUsages
	}
	if !IsNil(o.FoundPricing) {
		toSerialize["foundPricing"] = o.FoundPricing
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ServerId) {
		toSerialize["serverId"] = o.ServerId
	}
	if !IsNil(o.ServerUUID) {
		toSerialize["serverUUID"] = o.ServerUUID
	}
	if !IsNil(o.ServerUniqueId) {
		toSerialize["serverUniqueId"] = o.ServerUniqueId
	}
	if !IsNil(o.ServerExternalId) {
		toSerialize["serverExternalId"] = o.ServerExternalId
	}
	if !IsNil(o.ServerInternalId) {
		toSerialize["serverInternalId"] = o.ServerInternalId
	}
	if !IsNil(o.ResourcePoolId) {
		toSerialize["resourcePoolId"] = o.ResourcePoolId
	}
	if !IsNil(o.ResourcePoolName) {
		toSerialize["resourcePoolName"] = o.ResourcePoolName
	}
	return toSerialize, nil
}

type NullableBillingInstancesInstancesInnerContainersInner struct {
	value *BillingInstancesInstancesInnerContainersInner
	isSet bool
}

func (v NullableBillingInstancesInstancesInnerContainersInner) Get() *BillingInstancesInstancesInnerContainersInner {
	return v.value
}

func (v *NullableBillingInstancesInstancesInnerContainersInner) Set(val *BillingInstancesInstancesInnerContainersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingInstancesInstancesInnerContainersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingInstancesInstancesInnerContainersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingInstancesInstancesInnerContainersInner(val *BillingInstancesInstancesInnerContainersInner) *NullableBillingInstancesInstancesInnerContainersInner {
	return &NullableBillingInstancesInstancesInnerContainersInner{value: val, isSet: true}
}

func (v NullableBillingInstancesInstancesInnerContainersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingInstancesInstancesInnerContainersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


