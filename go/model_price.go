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

// Price struct for Price
type Price struct {
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Code *string `json:"code,omitempty"`
	Active *bool `json:"active,omitempty"`
	PriceType *string `json:"priceType,omitempty"`
	PriceUnit *string `json:"priceUnit,omitempty"`
	AdditionalPriceUnit NullableString `json:"additionalPriceUnit,omitempty"`
	Price NullableFloat32 `json:"price,omitempty"`
	CustomPrice NullableFloat32 `json:"customPrice,omitempty"`
	MarkupType NullableString `json:"markupType,omitempty"`
	Markup NullableFloat32 `json:"markup,omitempty"`
	MarkupPercent NullableFloat32 `json:"markupPercent,omitempty"`
	Cost NullableFloat32 `json:"cost,omitempty"`
	Currency *string `json:"currency,omitempty"`
	IncurCharges *string `json:"incurCharges,omitempty"`
	Platform NullableString `json:"platform,omitempty"`
	Software NullableString `json:"software,omitempty"`
	VolumeType NullablePriceSetVolumeType `json:"volumeType,omitempty"`
	Datastore NullableInlineResponse20082LoadBalancerInstanceSslCert `json:"datastore,omitempty"`
	CrossCloudApply NullableBool `json:"crossCloudApply,omitempty"`
	Account NullableString `json:"account,omitempty"`
}

// NewPrice instantiates a new Price object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrice() *Price {
	this := Price{}
	return &this
}

// NewPriceWithDefaults instantiates a new Price object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceWithDefaults() *Price {
	this := Price{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Price) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Price) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Price) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *Price) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Price) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Price) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Price) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Price) SetName(v string) {
	o.Name = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *Price) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Price) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *Price) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *Price) SetCode(v string) {
	o.Code = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *Price) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Price) GetActiveOk() (*bool, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *Price) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *Price) SetActive(v bool) {
	o.Active = &v
}

// GetPriceType returns the PriceType field value if set, zero value otherwise.
func (o *Price) GetPriceType() string {
	if o == nil || o.PriceType == nil {
		var ret string
		return ret
	}
	return *o.PriceType
}

// GetPriceTypeOk returns a tuple with the PriceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Price) GetPriceTypeOk() (*string, bool) {
	if o == nil || o.PriceType == nil {
		return nil, false
	}
	return o.PriceType, true
}

// HasPriceType returns a boolean if a field has been set.
func (o *Price) HasPriceType() bool {
	if o != nil && o.PriceType != nil {
		return true
	}

	return false
}

// SetPriceType gets a reference to the given string and assigns it to the PriceType field.
func (o *Price) SetPriceType(v string) {
	o.PriceType = &v
}

// GetPriceUnit returns the PriceUnit field value if set, zero value otherwise.
func (o *Price) GetPriceUnit() string {
	if o == nil || o.PriceUnit == nil {
		var ret string
		return ret
	}
	return *o.PriceUnit
}

// GetPriceUnitOk returns a tuple with the PriceUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Price) GetPriceUnitOk() (*string, bool) {
	if o == nil || o.PriceUnit == nil {
		return nil, false
	}
	return o.PriceUnit, true
}

// HasPriceUnit returns a boolean if a field has been set.
func (o *Price) HasPriceUnit() bool {
	if o != nil && o.PriceUnit != nil {
		return true
	}

	return false
}

// SetPriceUnit gets a reference to the given string and assigns it to the PriceUnit field.
func (o *Price) SetPriceUnit(v string) {
	o.PriceUnit = &v
}

// GetAdditionalPriceUnit returns the AdditionalPriceUnit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Price) GetAdditionalPriceUnit() string {
	if o == nil || o.AdditionalPriceUnit.Get() == nil {
		var ret string
		return ret
	}
	return *o.AdditionalPriceUnit.Get()
}

// GetAdditionalPriceUnitOk returns a tuple with the AdditionalPriceUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Price) GetAdditionalPriceUnitOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AdditionalPriceUnit.Get(), o.AdditionalPriceUnit.IsSet()
}

// HasAdditionalPriceUnit returns a boolean if a field has been set.
func (o *Price) HasAdditionalPriceUnit() bool {
	if o != nil && o.AdditionalPriceUnit.IsSet() {
		return true
	}

	return false
}

// SetAdditionalPriceUnit gets a reference to the given NullableString and assigns it to the AdditionalPriceUnit field.
func (o *Price) SetAdditionalPriceUnit(v string) {
	o.AdditionalPriceUnit.Set(&v)
}
// SetAdditionalPriceUnitNil sets the value for AdditionalPriceUnit to be an explicit nil
func (o *Price) SetAdditionalPriceUnitNil() {
	o.AdditionalPriceUnit.Set(nil)
}

// UnsetAdditionalPriceUnit ensures that no value is present for AdditionalPriceUnit, not even an explicit nil
func (o *Price) UnsetAdditionalPriceUnit() {
	o.AdditionalPriceUnit.Unset()
}

// GetPrice returns the Price field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Price) GetPrice() float32 {
	if o == nil || o.Price.Get() == nil {
		var ret float32
		return ret
	}
	return *o.Price.Get()
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Price) GetPriceOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Price.Get(), o.Price.IsSet()
}

// HasPrice returns a boolean if a field has been set.
func (o *Price) HasPrice() bool {
	if o != nil && o.Price.IsSet() {
		return true
	}

	return false
}

// SetPrice gets a reference to the given NullableFloat32 and assigns it to the Price field.
func (o *Price) SetPrice(v float32) {
	o.Price.Set(&v)
}
// SetPriceNil sets the value for Price to be an explicit nil
func (o *Price) SetPriceNil() {
	o.Price.Set(nil)
}

// UnsetPrice ensures that no value is present for Price, not even an explicit nil
func (o *Price) UnsetPrice() {
	o.Price.Unset()
}

// GetCustomPrice returns the CustomPrice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Price) GetCustomPrice() float32 {
	if o == nil || o.CustomPrice.Get() == nil {
		var ret float32
		return ret
	}
	return *o.CustomPrice.Get()
}

// GetCustomPriceOk returns a tuple with the CustomPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Price) GetCustomPriceOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CustomPrice.Get(), o.CustomPrice.IsSet()
}

// HasCustomPrice returns a boolean if a field has been set.
func (o *Price) HasCustomPrice() bool {
	if o != nil && o.CustomPrice.IsSet() {
		return true
	}

	return false
}

// SetCustomPrice gets a reference to the given NullableFloat32 and assigns it to the CustomPrice field.
func (o *Price) SetCustomPrice(v float32) {
	o.CustomPrice.Set(&v)
}
// SetCustomPriceNil sets the value for CustomPrice to be an explicit nil
func (o *Price) SetCustomPriceNil() {
	o.CustomPrice.Set(nil)
}

// UnsetCustomPrice ensures that no value is present for CustomPrice, not even an explicit nil
func (o *Price) UnsetCustomPrice() {
	o.CustomPrice.Unset()
}

// GetMarkupType returns the MarkupType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Price) GetMarkupType() string {
	if o == nil || o.MarkupType.Get() == nil {
		var ret string
		return ret
	}
	return *o.MarkupType.Get()
}

// GetMarkupTypeOk returns a tuple with the MarkupType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Price) GetMarkupTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MarkupType.Get(), o.MarkupType.IsSet()
}

// HasMarkupType returns a boolean if a field has been set.
func (o *Price) HasMarkupType() bool {
	if o != nil && o.MarkupType.IsSet() {
		return true
	}

	return false
}

// SetMarkupType gets a reference to the given NullableString and assigns it to the MarkupType field.
func (o *Price) SetMarkupType(v string) {
	o.MarkupType.Set(&v)
}
// SetMarkupTypeNil sets the value for MarkupType to be an explicit nil
func (o *Price) SetMarkupTypeNil() {
	o.MarkupType.Set(nil)
}

// UnsetMarkupType ensures that no value is present for MarkupType, not even an explicit nil
func (o *Price) UnsetMarkupType() {
	o.MarkupType.Unset()
}

// GetMarkup returns the Markup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Price) GetMarkup() float32 {
	if o == nil || o.Markup.Get() == nil {
		var ret float32
		return ret
	}
	return *o.Markup.Get()
}

// GetMarkupOk returns a tuple with the Markup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Price) GetMarkupOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Markup.Get(), o.Markup.IsSet()
}

// HasMarkup returns a boolean if a field has been set.
func (o *Price) HasMarkup() bool {
	if o != nil && o.Markup.IsSet() {
		return true
	}

	return false
}

// SetMarkup gets a reference to the given NullableFloat32 and assigns it to the Markup field.
func (o *Price) SetMarkup(v float32) {
	o.Markup.Set(&v)
}
// SetMarkupNil sets the value for Markup to be an explicit nil
func (o *Price) SetMarkupNil() {
	o.Markup.Set(nil)
}

// UnsetMarkup ensures that no value is present for Markup, not even an explicit nil
func (o *Price) UnsetMarkup() {
	o.Markup.Unset()
}

// GetMarkupPercent returns the MarkupPercent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Price) GetMarkupPercent() float32 {
	if o == nil || o.MarkupPercent.Get() == nil {
		var ret float32
		return ret
	}
	return *o.MarkupPercent.Get()
}

// GetMarkupPercentOk returns a tuple with the MarkupPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Price) GetMarkupPercentOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MarkupPercent.Get(), o.MarkupPercent.IsSet()
}

// HasMarkupPercent returns a boolean if a field has been set.
func (o *Price) HasMarkupPercent() bool {
	if o != nil && o.MarkupPercent.IsSet() {
		return true
	}

	return false
}

// SetMarkupPercent gets a reference to the given NullableFloat32 and assigns it to the MarkupPercent field.
func (o *Price) SetMarkupPercent(v float32) {
	o.MarkupPercent.Set(&v)
}
// SetMarkupPercentNil sets the value for MarkupPercent to be an explicit nil
func (o *Price) SetMarkupPercentNil() {
	o.MarkupPercent.Set(nil)
}

// UnsetMarkupPercent ensures that no value is present for MarkupPercent, not even an explicit nil
func (o *Price) UnsetMarkupPercent() {
	o.MarkupPercent.Unset()
}

// GetCost returns the Cost field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Price) GetCost() float32 {
	if o == nil || o.Cost.Get() == nil {
		var ret float32
		return ret
	}
	return *o.Cost.Get()
}

// GetCostOk returns a tuple with the Cost field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Price) GetCostOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Cost.Get(), o.Cost.IsSet()
}

// HasCost returns a boolean if a field has been set.
func (o *Price) HasCost() bool {
	if o != nil && o.Cost.IsSet() {
		return true
	}

	return false
}

// SetCost gets a reference to the given NullableFloat32 and assigns it to the Cost field.
func (o *Price) SetCost(v float32) {
	o.Cost.Set(&v)
}
// SetCostNil sets the value for Cost to be an explicit nil
func (o *Price) SetCostNil() {
	o.Cost.Set(nil)
}

// UnsetCost ensures that no value is present for Cost, not even an explicit nil
func (o *Price) UnsetCost() {
	o.Cost.Unset()
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *Price) GetCurrency() string {
	if o == nil || o.Currency == nil {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Price) GetCurrencyOk() (*string, bool) {
	if o == nil || o.Currency == nil {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *Price) HasCurrency() bool {
	if o != nil && o.Currency != nil {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *Price) SetCurrency(v string) {
	o.Currency = &v
}

// GetIncurCharges returns the IncurCharges field value if set, zero value otherwise.
func (o *Price) GetIncurCharges() string {
	if o == nil || o.IncurCharges == nil {
		var ret string
		return ret
	}
	return *o.IncurCharges
}

// GetIncurChargesOk returns a tuple with the IncurCharges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Price) GetIncurChargesOk() (*string, bool) {
	if o == nil || o.IncurCharges == nil {
		return nil, false
	}
	return o.IncurCharges, true
}

// HasIncurCharges returns a boolean if a field has been set.
func (o *Price) HasIncurCharges() bool {
	if o != nil && o.IncurCharges != nil {
		return true
	}

	return false
}

// SetIncurCharges gets a reference to the given string and assigns it to the IncurCharges field.
func (o *Price) SetIncurCharges(v string) {
	o.IncurCharges = &v
}

// GetPlatform returns the Platform field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Price) GetPlatform() string {
	if o == nil || o.Platform.Get() == nil {
		var ret string
		return ret
	}
	return *o.Platform.Get()
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Price) GetPlatformOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Platform.Get(), o.Platform.IsSet()
}

// HasPlatform returns a boolean if a field has been set.
func (o *Price) HasPlatform() bool {
	if o != nil && o.Platform.IsSet() {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given NullableString and assigns it to the Platform field.
func (o *Price) SetPlatform(v string) {
	o.Platform.Set(&v)
}
// SetPlatformNil sets the value for Platform to be an explicit nil
func (o *Price) SetPlatformNil() {
	o.Platform.Set(nil)
}

// UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
func (o *Price) UnsetPlatform() {
	o.Platform.Unset()
}

// GetSoftware returns the Software field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Price) GetSoftware() string {
	if o == nil || o.Software.Get() == nil {
		var ret string
		return ret
	}
	return *o.Software.Get()
}

// GetSoftwareOk returns a tuple with the Software field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Price) GetSoftwareOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Software.Get(), o.Software.IsSet()
}

// HasSoftware returns a boolean if a field has been set.
func (o *Price) HasSoftware() bool {
	if o != nil && o.Software.IsSet() {
		return true
	}

	return false
}

// SetSoftware gets a reference to the given NullableString and assigns it to the Software field.
func (o *Price) SetSoftware(v string) {
	o.Software.Set(&v)
}
// SetSoftwareNil sets the value for Software to be an explicit nil
func (o *Price) SetSoftwareNil() {
	o.Software.Set(nil)
}

// UnsetSoftware ensures that no value is present for Software, not even an explicit nil
func (o *Price) UnsetSoftware() {
	o.Software.Unset()
}

// GetVolumeType returns the VolumeType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Price) GetVolumeType() PriceSetVolumeType {
	if o == nil || o.VolumeType.Get() == nil {
		var ret PriceSetVolumeType
		return ret
	}
	return *o.VolumeType.Get()
}

// GetVolumeTypeOk returns a tuple with the VolumeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Price) GetVolumeTypeOk() (*PriceSetVolumeType, bool) {
	if o == nil  {
		return nil, false
	}
	return o.VolumeType.Get(), o.VolumeType.IsSet()
}

// HasVolumeType returns a boolean if a field has been set.
func (o *Price) HasVolumeType() bool {
	if o != nil && o.VolumeType.IsSet() {
		return true
	}

	return false
}

// SetVolumeType gets a reference to the given NullablePriceSetVolumeType and assigns it to the VolumeType field.
func (o *Price) SetVolumeType(v PriceSetVolumeType) {
	o.VolumeType.Set(&v)
}
// SetVolumeTypeNil sets the value for VolumeType to be an explicit nil
func (o *Price) SetVolumeTypeNil() {
	o.VolumeType.Set(nil)
}

// UnsetVolumeType ensures that no value is present for VolumeType, not even an explicit nil
func (o *Price) UnsetVolumeType() {
	o.VolumeType.Unset()
}

// GetDatastore returns the Datastore field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Price) GetDatastore() InlineResponse20082LoadBalancerInstanceSslCert {
	if o == nil || o.Datastore.Get() == nil {
		var ret InlineResponse20082LoadBalancerInstanceSslCert
		return ret
	}
	return *o.Datastore.Get()
}

// GetDatastoreOk returns a tuple with the Datastore field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Price) GetDatastoreOk() (*InlineResponse20082LoadBalancerInstanceSslCert, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Datastore.Get(), o.Datastore.IsSet()
}

// HasDatastore returns a boolean if a field has been set.
func (o *Price) HasDatastore() bool {
	if o != nil && o.Datastore.IsSet() {
		return true
	}

	return false
}

// SetDatastore gets a reference to the given NullableInlineResponse20082LoadBalancerInstanceSslCert and assigns it to the Datastore field.
func (o *Price) SetDatastore(v InlineResponse20082LoadBalancerInstanceSslCert) {
	o.Datastore.Set(&v)
}
// SetDatastoreNil sets the value for Datastore to be an explicit nil
func (o *Price) SetDatastoreNil() {
	o.Datastore.Set(nil)
}

// UnsetDatastore ensures that no value is present for Datastore, not even an explicit nil
func (o *Price) UnsetDatastore() {
	o.Datastore.Unset()
}

// GetCrossCloudApply returns the CrossCloudApply field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Price) GetCrossCloudApply() bool {
	if o == nil || o.CrossCloudApply.Get() == nil {
		var ret bool
		return ret
	}
	return *o.CrossCloudApply.Get()
}

// GetCrossCloudApplyOk returns a tuple with the CrossCloudApply field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Price) GetCrossCloudApplyOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CrossCloudApply.Get(), o.CrossCloudApply.IsSet()
}

// HasCrossCloudApply returns a boolean if a field has been set.
func (o *Price) HasCrossCloudApply() bool {
	if o != nil && o.CrossCloudApply.IsSet() {
		return true
	}

	return false
}

// SetCrossCloudApply gets a reference to the given NullableBool and assigns it to the CrossCloudApply field.
func (o *Price) SetCrossCloudApply(v bool) {
	o.CrossCloudApply.Set(&v)
}
// SetCrossCloudApplyNil sets the value for CrossCloudApply to be an explicit nil
func (o *Price) SetCrossCloudApplyNil() {
	o.CrossCloudApply.Set(nil)
}

// UnsetCrossCloudApply ensures that no value is present for CrossCloudApply, not even an explicit nil
func (o *Price) UnsetCrossCloudApply() {
	o.CrossCloudApply.Unset()
}

// GetAccount returns the Account field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Price) GetAccount() string {
	if o == nil || o.Account.Get() == nil {
		var ret string
		return ret
	}
	return *o.Account.Get()
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Price) GetAccountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Account.Get(), o.Account.IsSet()
}

// HasAccount returns a boolean if a field has been set.
func (o *Price) HasAccount() bool {
	if o != nil && o.Account.IsSet() {
		return true
	}

	return false
}

// SetAccount gets a reference to the given NullableString and assigns it to the Account field.
func (o *Price) SetAccount(v string) {
	o.Account.Set(&v)
}
// SetAccountNil sets the value for Account to be an explicit nil
func (o *Price) SetAccountNil() {
	o.Account.Set(nil)
}

// UnsetAccount ensures that no value is present for Account, not even an explicit nil
func (o *Price) UnsetAccount() {
	o.Account.Unset()
}

func (o Price) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	if o.PriceType != nil {
		toSerialize["priceType"] = o.PriceType
	}
	if o.PriceUnit != nil {
		toSerialize["priceUnit"] = o.PriceUnit
	}
	if o.AdditionalPriceUnit.IsSet() {
		toSerialize["additionalPriceUnit"] = o.AdditionalPriceUnit.Get()
	}
	if o.Price.IsSet() {
		toSerialize["price"] = o.Price.Get()
	}
	if o.CustomPrice.IsSet() {
		toSerialize["customPrice"] = o.CustomPrice.Get()
	}
	if o.MarkupType.IsSet() {
		toSerialize["markupType"] = o.MarkupType.Get()
	}
	if o.Markup.IsSet() {
		toSerialize["markup"] = o.Markup.Get()
	}
	if o.MarkupPercent.IsSet() {
		toSerialize["markupPercent"] = o.MarkupPercent.Get()
	}
	if o.Cost.IsSet() {
		toSerialize["cost"] = o.Cost.Get()
	}
	if o.Currency != nil {
		toSerialize["currency"] = o.Currency
	}
	if o.IncurCharges != nil {
		toSerialize["incurCharges"] = o.IncurCharges
	}
	if o.Platform.IsSet() {
		toSerialize["platform"] = o.Platform.Get()
	}
	if o.Software.IsSet() {
		toSerialize["software"] = o.Software.Get()
	}
	if o.VolumeType.IsSet() {
		toSerialize["volumeType"] = o.VolumeType.Get()
	}
	if o.Datastore.IsSet() {
		toSerialize["datastore"] = o.Datastore.Get()
	}
	if o.CrossCloudApply.IsSet() {
		toSerialize["crossCloudApply"] = o.CrossCloudApply.Get()
	}
	if o.Account.IsSet() {
		toSerialize["account"] = o.Account.Get()
	}
	return json.Marshal(toSerialize)
}

type NullablePrice struct {
	value *Price
	isSet bool
}

func (v NullablePrice) Get() *Price {
	return v.value
}

func (v *NullablePrice) Set(val *Price) {
	v.value = val
	v.isSet = true
}

func (v NullablePrice) IsSet() bool {
	return v.isSet
}

func (v *NullablePrice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrice(val *Price) *NullablePrice {
	return &NullablePrice{value: val, isSet: true}
}

func (v NullablePrice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

