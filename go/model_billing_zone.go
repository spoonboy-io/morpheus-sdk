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

// BillingZone struct for BillingZone
type BillingZone struct {
	ZoneName *string `json:"zoneName,omitempty"`
	ZoneId *int64 `json:"zoneId,omitempty"`
	ZoneUUID *string `json:"zoneUUID,omitempty"`
	ZoneCode NullableString `json:"zoneCode,omitempty"`
	StartDate *time.Time `json:"startDate,omitempty"`
	EndDate *time.Time `json:"endDate,omitempty"`
	PriceUnit *string `json:"priceUnit,omitempty"`
	ComputeServers *BillingComputeServers `json:"computeServers,omitempty"`
	Instances *BillingInstances `json:"instances,omitempty"`
	DiscoveredServers *BillingComputeServers `json:"discoveredServers,omitempty"`
	LoadBalancers *BillingLoadBalancers `json:"loadBalancers,omitempty"`
	VirtualImages *BillingVirtualImages `json:"virtualImages,omitempty"`
	Snapshots *BillingSnapshots `json:"snapshots,omitempty"`
	Price *float32 `json:"price,omitempty"`
	Cost *float32 `json:"cost,omitempty"`
}

// NewBillingZone instantiates a new BillingZone object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingZone() *BillingZone {
	this := BillingZone{}
	return &this
}

// NewBillingZoneWithDefaults instantiates a new BillingZone object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingZoneWithDefaults() *BillingZone {
	this := BillingZone{}
	return &this
}

// GetZoneName returns the ZoneName field value if set, zero value otherwise.
func (o *BillingZone) GetZoneName() string {
	if o == nil || o.ZoneName == nil {
		var ret string
		return ret
	}
	return *o.ZoneName
}

// GetZoneNameOk returns a tuple with the ZoneName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingZone) GetZoneNameOk() (*string, bool) {
	if o == nil || o.ZoneName == nil {
		return nil, false
	}
	return o.ZoneName, true
}

// HasZoneName returns a boolean if a field has been set.
func (o *BillingZone) HasZoneName() bool {
	if o != nil && o.ZoneName != nil {
		return true
	}

	return false
}

// SetZoneName gets a reference to the given string and assigns it to the ZoneName field.
func (o *BillingZone) SetZoneName(v string) {
	o.ZoneName = &v
}

// GetZoneId returns the ZoneId field value if set, zero value otherwise.
func (o *BillingZone) GetZoneId() int64 {
	if o == nil || o.ZoneId == nil {
		var ret int64
		return ret
	}
	return *o.ZoneId
}

// GetZoneIdOk returns a tuple with the ZoneId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingZone) GetZoneIdOk() (*int64, bool) {
	if o == nil || o.ZoneId == nil {
		return nil, false
	}
	return o.ZoneId, true
}

// HasZoneId returns a boolean if a field has been set.
func (o *BillingZone) HasZoneId() bool {
	if o != nil && o.ZoneId != nil {
		return true
	}

	return false
}

// SetZoneId gets a reference to the given int64 and assigns it to the ZoneId field.
func (o *BillingZone) SetZoneId(v int64) {
	o.ZoneId = &v
}

// GetZoneUUID returns the ZoneUUID field value if set, zero value otherwise.
func (o *BillingZone) GetZoneUUID() string {
	if o == nil || o.ZoneUUID == nil {
		var ret string
		return ret
	}
	return *o.ZoneUUID
}

// GetZoneUUIDOk returns a tuple with the ZoneUUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingZone) GetZoneUUIDOk() (*string, bool) {
	if o == nil || o.ZoneUUID == nil {
		return nil, false
	}
	return o.ZoneUUID, true
}

// HasZoneUUID returns a boolean if a field has been set.
func (o *BillingZone) HasZoneUUID() bool {
	if o != nil && o.ZoneUUID != nil {
		return true
	}

	return false
}

// SetZoneUUID gets a reference to the given string and assigns it to the ZoneUUID field.
func (o *BillingZone) SetZoneUUID(v string) {
	o.ZoneUUID = &v
}

// GetZoneCode returns the ZoneCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingZone) GetZoneCode() string {
	if o == nil || o.ZoneCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.ZoneCode.Get()
}

// GetZoneCodeOk returns a tuple with the ZoneCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingZone) GetZoneCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ZoneCode.Get(), o.ZoneCode.IsSet()
}

// HasZoneCode returns a boolean if a field has been set.
func (o *BillingZone) HasZoneCode() bool {
	if o != nil && o.ZoneCode.IsSet() {
		return true
	}

	return false
}

// SetZoneCode gets a reference to the given NullableString and assigns it to the ZoneCode field.
func (o *BillingZone) SetZoneCode(v string) {
	o.ZoneCode.Set(&v)
}
// SetZoneCodeNil sets the value for ZoneCode to be an explicit nil
func (o *BillingZone) SetZoneCodeNil() {
	o.ZoneCode.Set(nil)
}

// UnsetZoneCode ensures that no value is present for ZoneCode, not even an explicit nil
func (o *BillingZone) UnsetZoneCode() {
	o.ZoneCode.Unset()
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *BillingZone) GetStartDate() time.Time {
	if o == nil || o.StartDate == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingZone) GetStartDateOk() (*time.Time, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *BillingZone) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *BillingZone) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *BillingZone) GetEndDate() time.Time {
	if o == nil || o.EndDate == nil {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingZone) GetEndDateOk() (*time.Time, bool) {
	if o == nil || o.EndDate == nil {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *BillingZone) HasEndDate() bool {
	if o != nil && o.EndDate != nil {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *BillingZone) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetPriceUnit returns the PriceUnit field value if set, zero value otherwise.
func (o *BillingZone) GetPriceUnit() string {
	if o == nil || o.PriceUnit == nil {
		var ret string
		return ret
	}
	return *o.PriceUnit
}

// GetPriceUnitOk returns a tuple with the PriceUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingZone) GetPriceUnitOk() (*string, bool) {
	if o == nil || o.PriceUnit == nil {
		return nil, false
	}
	return o.PriceUnit, true
}

// HasPriceUnit returns a boolean if a field has been set.
func (o *BillingZone) HasPriceUnit() bool {
	if o != nil && o.PriceUnit != nil {
		return true
	}

	return false
}

// SetPriceUnit gets a reference to the given string and assigns it to the PriceUnit field.
func (o *BillingZone) SetPriceUnit(v string) {
	o.PriceUnit = &v
}

// GetComputeServers returns the ComputeServers field value if set, zero value otherwise.
func (o *BillingZone) GetComputeServers() BillingComputeServers {
	if o == nil || o.ComputeServers == nil {
		var ret BillingComputeServers
		return ret
	}
	return *o.ComputeServers
}

// GetComputeServersOk returns a tuple with the ComputeServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingZone) GetComputeServersOk() (*BillingComputeServers, bool) {
	if o == nil || o.ComputeServers == nil {
		return nil, false
	}
	return o.ComputeServers, true
}

// HasComputeServers returns a boolean if a field has been set.
func (o *BillingZone) HasComputeServers() bool {
	if o != nil && o.ComputeServers != nil {
		return true
	}

	return false
}

// SetComputeServers gets a reference to the given BillingComputeServers and assigns it to the ComputeServers field.
func (o *BillingZone) SetComputeServers(v BillingComputeServers) {
	o.ComputeServers = &v
}

// GetInstances returns the Instances field value if set, zero value otherwise.
func (o *BillingZone) GetInstances() BillingInstances {
	if o == nil || o.Instances == nil {
		var ret BillingInstances
		return ret
	}
	return *o.Instances
}

// GetInstancesOk returns a tuple with the Instances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingZone) GetInstancesOk() (*BillingInstances, bool) {
	if o == nil || o.Instances == nil {
		return nil, false
	}
	return o.Instances, true
}

// HasInstances returns a boolean if a field has been set.
func (o *BillingZone) HasInstances() bool {
	if o != nil && o.Instances != nil {
		return true
	}

	return false
}

// SetInstances gets a reference to the given BillingInstances and assigns it to the Instances field.
func (o *BillingZone) SetInstances(v BillingInstances) {
	o.Instances = &v
}

// GetDiscoveredServers returns the DiscoveredServers field value if set, zero value otherwise.
func (o *BillingZone) GetDiscoveredServers() BillingComputeServers {
	if o == nil || o.DiscoveredServers == nil {
		var ret BillingComputeServers
		return ret
	}
	return *o.DiscoveredServers
}

// GetDiscoveredServersOk returns a tuple with the DiscoveredServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingZone) GetDiscoveredServersOk() (*BillingComputeServers, bool) {
	if o == nil || o.DiscoveredServers == nil {
		return nil, false
	}
	return o.DiscoveredServers, true
}

// HasDiscoveredServers returns a boolean if a field has been set.
func (o *BillingZone) HasDiscoveredServers() bool {
	if o != nil && o.DiscoveredServers != nil {
		return true
	}

	return false
}

// SetDiscoveredServers gets a reference to the given BillingComputeServers and assigns it to the DiscoveredServers field.
func (o *BillingZone) SetDiscoveredServers(v BillingComputeServers) {
	o.DiscoveredServers = &v
}

// GetLoadBalancers returns the LoadBalancers field value if set, zero value otherwise.
func (o *BillingZone) GetLoadBalancers() BillingLoadBalancers {
	if o == nil || o.LoadBalancers == nil {
		var ret BillingLoadBalancers
		return ret
	}
	return *o.LoadBalancers
}

// GetLoadBalancersOk returns a tuple with the LoadBalancers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingZone) GetLoadBalancersOk() (*BillingLoadBalancers, bool) {
	if o == nil || o.LoadBalancers == nil {
		return nil, false
	}
	return o.LoadBalancers, true
}

// HasLoadBalancers returns a boolean if a field has been set.
func (o *BillingZone) HasLoadBalancers() bool {
	if o != nil && o.LoadBalancers != nil {
		return true
	}

	return false
}

// SetLoadBalancers gets a reference to the given BillingLoadBalancers and assigns it to the LoadBalancers field.
func (o *BillingZone) SetLoadBalancers(v BillingLoadBalancers) {
	o.LoadBalancers = &v
}

// GetVirtualImages returns the VirtualImages field value if set, zero value otherwise.
func (o *BillingZone) GetVirtualImages() BillingVirtualImages {
	if o == nil || o.VirtualImages == nil {
		var ret BillingVirtualImages
		return ret
	}
	return *o.VirtualImages
}

// GetVirtualImagesOk returns a tuple with the VirtualImages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingZone) GetVirtualImagesOk() (*BillingVirtualImages, bool) {
	if o == nil || o.VirtualImages == nil {
		return nil, false
	}
	return o.VirtualImages, true
}

// HasVirtualImages returns a boolean if a field has been set.
func (o *BillingZone) HasVirtualImages() bool {
	if o != nil && o.VirtualImages != nil {
		return true
	}

	return false
}

// SetVirtualImages gets a reference to the given BillingVirtualImages and assigns it to the VirtualImages field.
func (o *BillingZone) SetVirtualImages(v BillingVirtualImages) {
	o.VirtualImages = &v
}

// GetSnapshots returns the Snapshots field value if set, zero value otherwise.
func (o *BillingZone) GetSnapshots() BillingSnapshots {
	if o == nil || o.Snapshots == nil {
		var ret BillingSnapshots
		return ret
	}
	return *o.Snapshots
}

// GetSnapshotsOk returns a tuple with the Snapshots field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingZone) GetSnapshotsOk() (*BillingSnapshots, bool) {
	if o == nil || o.Snapshots == nil {
		return nil, false
	}
	return o.Snapshots, true
}

// HasSnapshots returns a boolean if a field has been set.
func (o *BillingZone) HasSnapshots() bool {
	if o != nil && o.Snapshots != nil {
		return true
	}

	return false
}

// SetSnapshots gets a reference to the given BillingSnapshots and assigns it to the Snapshots field.
func (o *BillingZone) SetSnapshots(v BillingSnapshots) {
	o.Snapshots = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *BillingZone) GetPrice() float32 {
	if o == nil || o.Price == nil {
		var ret float32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingZone) GetPriceOk() (*float32, bool) {
	if o == nil || o.Price == nil {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *BillingZone) HasPrice() bool {
	if o != nil && o.Price != nil {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float32 and assigns it to the Price field.
func (o *BillingZone) SetPrice(v float32) {
	o.Price = &v
}

// GetCost returns the Cost field value if set, zero value otherwise.
func (o *BillingZone) GetCost() float32 {
	if o == nil || o.Cost == nil {
		var ret float32
		return ret
	}
	return *o.Cost
}

// GetCostOk returns a tuple with the Cost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingZone) GetCostOk() (*float32, bool) {
	if o == nil || o.Cost == nil {
		return nil, false
	}
	return o.Cost, true
}

// HasCost returns a boolean if a field has been set.
func (o *BillingZone) HasCost() bool {
	if o != nil && o.Cost != nil {
		return true
	}

	return false
}

// SetCost gets a reference to the given float32 and assigns it to the Cost field.
func (o *BillingZone) SetCost(v float32) {
	o.Cost = &v
}

func (o BillingZone) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ZoneName != nil {
		toSerialize["zoneName"] = o.ZoneName
	}
	if o.ZoneId != nil {
		toSerialize["zoneId"] = o.ZoneId
	}
	if o.ZoneUUID != nil {
		toSerialize["zoneUUID"] = o.ZoneUUID
	}
	if o.ZoneCode.IsSet() {
		toSerialize["zoneCode"] = o.ZoneCode.Get()
	}
	if o.StartDate != nil {
		toSerialize["startDate"] = o.StartDate
	}
	if o.EndDate != nil {
		toSerialize["endDate"] = o.EndDate
	}
	if o.PriceUnit != nil {
		toSerialize["priceUnit"] = o.PriceUnit
	}
	if o.ComputeServers != nil {
		toSerialize["computeServers"] = o.ComputeServers
	}
	if o.Instances != nil {
		toSerialize["instances"] = o.Instances
	}
	if o.DiscoveredServers != nil {
		toSerialize["discoveredServers"] = o.DiscoveredServers
	}
	if o.LoadBalancers != nil {
		toSerialize["loadBalancers"] = o.LoadBalancers
	}
	if o.VirtualImages != nil {
		toSerialize["virtualImages"] = o.VirtualImages
	}
	if o.Snapshots != nil {
		toSerialize["snapshots"] = o.Snapshots
	}
	if o.Price != nil {
		toSerialize["price"] = o.Price
	}
	if o.Cost != nil {
		toSerialize["cost"] = o.Cost
	}
	return json.Marshal(toSerialize)
}

type NullableBillingZone struct {
	value *BillingZone
	isSet bool
}

func (v NullableBillingZone) Get() *BillingZone {
	return v.value
}

func (v *NullableBillingZone) Set(val *BillingZone) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingZone) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingZone) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingZone(val *BillingZone) *NullableBillingZone {
	return &NullableBillingZone{value: val, isSet: true}
}

func (v NullableBillingZone) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingZone) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

