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

// BillingInstancesUsages struct for BillingInstancesUsages
type BillingInstancesUsages struct {
	Name *string `json:"name,omitempty"`
	InstanceName *string `json:"instanceName,omitempty"`
	ZoneName *string `json:"zoneName,omitempty"`
	AccountName *string `json:"accountName,omitempty"`
	Volumes *[]BillingInstancesVolumes `json:"volumes,omitempty"`
	MaxMemory *int64 `json:"maxMemory,omitempty"`
	MaxCpu NullableString `json:"maxCpu,omitempty"`
	MaxCores *int64 `json:"maxCores,omitempty"`
	ServerExternalId *string `json:"serverExternalId,omitempty"`
	ServerInternalId *string `json:"serverInternalId,omitempty"`
	PlanName *string `json:"planName,omitempty"`
	HourlyPrice *float32 `json:"hourlyPrice,omitempty"`
	HourlyCost *float32 `json:"hourlyCost,omitempty"`
	Currency *string `json:"currency,omitempty"`
	PricesUsed *[]BillingInstancesPricesUsed `json:"pricesUsed,omitempty"`
	Cost *float32 `json:"cost,omitempty"`
	Price *float32 `json:"price,omitempty"`
	CreatedByUser *string `json:"createdByUser,omitempty"`
	CreatedByUserId *int64 `json:"createdByUserId,omitempty"`
	SiteId *int64 `json:"siteId,omitempty"`
	SiteName *string `json:"siteName,omitempty"`
	SiteUUID *string `json:"siteUUID,omitempty"`
	SiteCode NullableString `json:"siteCode,omitempty"`
	StartDate *time.Time `json:"startDate,omitempty"`
	EndDate *time.Time `json:"endDate,omitempty"`
	Status *string `json:"status,omitempty"`
	Tags []map[string]interface{} `json:"tags,omitempty"`
	ApplicablePrices *[]BillingInstancesApplicablePrices `json:"applicablePrices,omitempty"`
	ServicePlanId *int64 `json:"servicePlanId,omitempty"`
	ServicePlanName *string `json:"servicePlanName,omitempty"`
	ResourcePoolId *int64 `json:"resourcePoolId,omitempty"`
	ResourcePoolName *string `json:"resourcePoolName,omitempty"`
}

// NewBillingInstancesUsages instantiates a new BillingInstancesUsages object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingInstancesUsages() *BillingInstancesUsages {
	this := BillingInstancesUsages{}
	return &this
}

// NewBillingInstancesUsagesWithDefaults instantiates a new BillingInstancesUsages object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingInstancesUsagesWithDefaults() *BillingInstancesUsages {
	this := BillingInstancesUsages{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BillingInstancesUsages) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInstancesUsages) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BillingInstancesUsages) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BillingInstancesUsages) SetName(v string) {
	o.Name = &v
}

// GetInstanceName returns the InstanceName field value if set, zero value otherwise.
func (o *BillingInstancesUsages) GetInstanceName() string {
	if o == nil || o.InstanceName == nil {
		var ret string
		return ret
	}
	return *o.InstanceName
}

// GetInstanceNameOk returns a tuple with the InstanceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInstancesUsages) GetInstanceNameOk() (*string, bool) {
	if o == nil || o.InstanceName == nil {
		return nil, false
	}
	return o.InstanceName, true
}

// HasInstanceName returns a boolean if a field has been set.
func (o *BillingInstancesUsages) HasInstanceName() bool {
	if o != nil && o.InstanceName != nil {
		return true
	}

	return false
}

// SetInstanceName gets a reference to the given string and assigns it to the InstanceName field.
func (o *BillingInstancesUsages) SetInstanceName(v string) {
	o.InstanceName = &v
}

// GetZoneName returns the ZoneName field value if set, zero value otherwise.
func (o *BillingInstancesUsages) GetZoneName() string {
	if o == nil || o.ZoneName == nil {
		var ret string
		return ret
	}
	return *o.ZoneName
}

// GetZoneNameOk returns a tuple with the ZoneName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInstancesUsages) GetZoneNameOk() (*string, bool) {
	if o == nil || o.ZoneName == nil {
		return nil, false
	}
	return o.ZoneName, true
}

// HasZoneName returns a boolean if a field has been set.
func (o *BillingInstancesUsages) HasZoneName() bool {
	if o != nil && o.ZoneName != nil {
		return true
	}

	return false
}

// SetZoneName gets a reference to the given string and assigns it to the ZoneName field.
func (o *BillingInstancesUsages) SetZoneName(v string) {
	o.ZoneName = &v
}

// GetAccountName returns the AccountName field value if set, zero value otherwise.
func (o *BillingInstancesUsages) GetAccountName() string {
	if o == nil || o.AccountName == nil {
		var ret string
		return ret
	}
	return *o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInstancesUsages) GetAccountNameOk() (*string, bool) {
	if o == nil || o.AccountName == nil {
		return nil, false
	}
	return o.AccountName, true
}

// HasAccountName returns a boolean if a field has been set.
func (o *BillingInstancesUsages) HasAccountName() bool {
	if o != nil && o.AccountName != nil {
		return true
	}

	return false
}

// SetAccountName gets a reference to the given string and assigns it to the AccountName field.
func (o *BillingInstancesUsages) SetAccountName(v string) {
	o.AccountName = &v
}

// GetVolumes returns the Volumes field value if set, zero value otherwise.
func (o *BillingInstancesUsages) GetVolumes() []BillingInstancesVolumes {
	if o == nil || o.Volumes == nil {
		var ret []BillingInstancesVolumes
		return ret
	}
	return *o.Volumes
}

// GetVolumesOk returns a tuple with the Volumes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInstancesUsages) GetVolumesOk() (*[]BillingInstancesVolumes, bool) {
	if o == nil || o.Volumes == nil {
		return nil, false
	}
	return o.Volumes, true
}

// HasVolumes returns a boolean if a field has been set.
func (o *BillingInstancesUsages) HasVolumes() bool {
	if o != nil && o.Volumes != nil {
		return true
	}

	return false
}

// SetVolumes gets a reference to the given []BillingInstancesVolumes and assigns it to the Volumes field.
func (o *BillingInstancesUsages) SetVolumes(v []BillingInstancesVolumes) {
	o.Volumes = &v
}

// GetMaxMemory returns the MaxMemory field value if set, zero value otherwise.
func (o *BillingInstancesUsages) GetMaxMemory() int64 {
	if o == nil || o.MaxMemory == nil {
		var ret int64
		return ret
	}
	return *o.MaxMemory
}

// GetMaxMemoryOk returns a tuple with the MaxMemory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInstancesUsages) GetMaxMemoryOk() (*int64, bool) {
	if o == nil || o.MaxMemory == nil {
		return nil, false
	}
	return o.MaxMemory, true
}

// HasMaxMemory returns a boolean if a field has been set.
func (o *BillingInstancesUsages) HasMaxMemory() bool {
	if o != nil && o.MaxMemory != nil {
		return true
	}

	return false
}

// SetMaxMemory gets a reference to the given int64 and assigns it to the MaxMemory field.
func (o *BillingInstancesUsages) SetMaxMemory(v int64) {
	o.MaxMemory = &v
}

// GetMaxCpu returns the MaxCpu field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingInstancesUsages) GetMaxCpu() string {
	if o == nil || o.MaxCpu.Get() == nil {
		var ret string
		return ret
	}
	return *o.MaxCpu.Get()
}

// GetMaxCpuOk returns a tuple with the MaxCpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingInstancesUsages) GetMaxCpuOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MaxCpu.Get(), o.MaxCpu.IsSet()
}

// HasMaxCpu returns a boolean if a field has been set.
func (o *BillingInstancesUsages) HasMaxCpu() bool {
	if o != nil && o.MaxCpu.IsSet() {
		return true
	}

	return false
}

// SetMaxCpu gets a reference to the given NullableString and assigns it to the MaxCpu field.
func (o *BillingInstancesUsages) SetMaxCpu(v string) {
	o.MaxCpu.Set(&v)
}
// SetMaxCpuNil sets the value for MaxCpu to be an explicit nil
func (o *BillingInstancesUsages) SetMaxCpuNil() {
	o.MaxCpu.Set(nil)
}

// UnsetMaxCpu ensures that no value is present for MaxCpu, not even an explicit nil
func (o *BillingInstancesUsages) UnsetMaxCpu() {
	o.MaxCpu.Unset()
}

// GetMaxCores returns the MaxCores field value if set, zero value otherwise.
func (o *BillingInstancesUsages) GetMaxCores() int64 {
	if o == nil || o.MaxCores == nil {
		var ret int64
		return ret
	}
	return *o.MaxCores
}

// GetMaxCoresOk returns a tuple with the MaxCores field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInstancesUsages) GetMaxCoresOk() (*int64, bool) {
	if o == nil || o.MaxCores == nil {
		return nil, false
	}
	return o.MaxCores, true
}

// HasMaxCores returns a boolean if a field has been set.
func (o *BillingInstancesUsages) HasMaxCores() bool {
	if o != nil && o.MaxCores != nil {
		return true
	}

	return false
}

// SetMaxCores gets a reference to the given int64 and assigns it to the MaxCores field.
func (o *BillingInstancesUsages) SetMaxCores(v int64) {
	o.MaxCores = &v
}

// GetServerExternalId returns the ServerExternalId field value if set, zero value otherwise.
func (o *BillingInstancesUsages) GetServerExternalId() string {
	if o == nil || o.ServerExternalId == nil {
		var ret string
		return ret
	}
	return *o.ServerExternalId
}

// GetServerExternalIdOk returns a tuple with the ServerExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInstancesUsages) GetServerExternalIdOk() (*string, bool) {
	if o == nil || o.ServerExternalId == nil {
		return nil, false
	}
	return o.ServerExternalId, true
}

// HasServerExternalId returns a boolean if a field has been set.
func (o *BillingInstancesUsages) HasServerExternalId() bool {
	if o != nil && o.ServerExternalId != nil {
		return true
	}

	return false
}

// SetServerExternalId gets a reference to the given string and assigns it to the ServerExternalId field.
func (o *BillingInstancesUsages) SetServerExternalId(v string) {
	o.ServerExternalId = &v
}

// GetServerInternalId returns the ServerInternalId field value if set, zero value otherwise.
func (o *BillingInstancesUsages) GetServerInternalId() string {
	if o == nil || o.ServerInternalId == nil {
		var ret string
		return ret
	}
	return *o.ServerInternalId
}

// GetServerInternalIdOk returns a tuple with the ServerInternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInstancesUsages) GetServerInternalIdOk() (*string, bool) {
	if o == nil || o.ServerInternalId == nil {
		return nil, false
	}
	return o.ServerInternalId, true
}

// HasServerInternalId returns a boolean if a field has been set.
func (o *BillingInstancesUsages) HasServerInternalId() bool {
	if o != nil && o.ServerInternalId != nil {
		return true
	}

	return false
}

// SetServerInternalId gets a reference to the given string and assigns it to the ServerInternalId field.
func (o *BillingInstancesUsages) SetServerInternalId(v string) {
	o.ServerInternalId = &v
}

// GetPlanName returns the PlanName field value if set, zero value otherwise.
func (o *BillingInstancesUsages) GetPlanName() string {
	if o == nil || o.PlanName == nil {
		var ret string
		return ret
	}
	return *o.PlanName
}

// GetPlanNameOk returns a tuple with the PlanName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInstancesUsages) GetPlanNameOk() (*string, bool) {
	if o == nil || o.PlanName == nil {
		return nil, false
	}
	return o.PlanName, true
}

// HasPlanName returns a boolean if a field has been set.
func (o *BillingInstancesUsages) HasPlanName() bool {
	if o != nil && o.PlanName != nil {
		return true
	}

	return false
}

// SetPlanName gets a reference to the given string and assigns it to the PlanName field.
func (o *BillingInstancesUsages) SetPlanName(v string) {
	o.PlanName = &v
}

// GetHourlyPrice returns the HourlyPrice field value if set, zero value otherwise.
func (o *BillingInstancesUsages) GetHourlyPrice() float32 {
	if o == nil || o.HourlyPrice == nil {
		var ret float32
		return ret
	}
	return *o.HourlyPrice
}

// GetHourlyPriceOk returns a tuple with the HourlyPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInstancesUsages) GetHourlyPriceOk() (*float32, bool) {
	if o == nil || o.HourlyPrice == nil {
		return nil, false
	}
	return o.HourlyPrice, true
}

// HasHourlyPrice returns a boolean if a field has been set.
func (o *BillingInstancesUsages) HasHourlyPrice() bool {
	if o != nil && o.HourlyPrice != nil {
		return true
	}

	return false
}

// SetHourlyPrice gets a reference to the given float32 and assigns it to the HourlyPrice field.
func (o *BillingInstancesUsages) SetHourlyPrice(v float32) {
	o.HourlyPrice = &v
}

// GetHourlyCost returns the HourlyCost field value if set, zero value otherwise.
func (o *BillingInstancesUsages) GetHourlyCost() float32 {
	if o == nil || o.HourlyCost == nil {
		var ret float32
		return ret
	}
	return *o.HourlyCost
}

// GetHourlyCostOk returns a tuple with the HourlyCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInstancesUsages) GetHourlyCostOk() (*float32, bool) {
	if o == nil || o.HourlyCost == nil {
		return nil, false
	}
	return o.HourlyCost, true
}

// HasHourlyCost returns a boolean if a field has been set.
func (o *BillingInstancesUsages) HasHourlyCost() bool {
	if o != nil && o.HourlyCost != nil {
		return true
	}

	return false
}

// SetHourlyCost gets a reference to the given float32 and assigns it to the HourlyCost field.
func (o *BillingInstancesUsages) SetHourlyCost(v float32) {
	o.HourlyCost = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *BillingInstancesUsages) GetCurrency() string {
	if o == nil || o.Currency == nil {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInstancesUsages) GetCurrencyOk() (*string, bool) {
	if o == nil || o.Currency == nil {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *BillingInstancesUsages) HasCurrency() bool {
	if o != nil && o.Currency != nil {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *BillingInstancesUsages) SetCurrency(v string) {
	o.Currency = &v
}

// GetPricesUsed returns the PricesUsed field value if set, zero value otherwise.
func (o *BillingInstancesUsages) GetPricesUsed() []BillingInstancesPricesUsed {
	if o == nil || o.PricesUsed == nil {
		var ret []BillingInstancesPricesUsed
		return ret
	}
	return *o.PricesUsed
}

// GetPricesUsedOk returns a tuple with the PricesUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInstancesUsages) GetPricesUsedOk() (*[]BillingInstancesPricesUsed, bool) {
	if o == nil || o.PricesUsed == nil {
		return nil, false
	}
	return o.PricesUsed, true
}

// HasPricesUsed returns a boolean if a field has been set.
func (o *BillingInstancesUsages) HasPricesUsed() bool {
	if o != nil && o.PricesUsed != nil {
		return true
	}

	return false
}

// SetPricesUsed gets a reference to the given []BillingInstancesPricesUsed and assigns it to the PricesUsed field.
func (o *BillingInstancesUsages) SetPricesUsed(v []BillingInstancesPricesUsed) {
	o.PricesUsed = &v
}

// GetCost returns the Cost field value if set, zero value otherwise.
func (o *BillingInstancesUsages) GetCost() float32 {
	if o == nil || o.Cost == nil {
		var ret float32
		return ret
	}
	return *o.Cost
}

// GetCostOk returns a tuple with the Cost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInstancesUsages) GetCostOk() (*float32, bool) {
	if o == nil || o.Cost == nil {
		return nil, false
	}
	return o.Cost, true
}

// HasCost returns a boolean if a field has been set.
func (o *BillingInstancesUsages) HasCost() bool {
	if o != nil && o.Cost != nil {
		return true
	}

	return false
}

// SetCost gets a reference to the given float32 and assigns it to the Cost field.
func (o *BillingInstancesUsages) SetCost(v float32) {
	o.Cost = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *BillingInstancesUsages) GetPrice() float32 {
	if o == nil || o.Price == nil {
		var ret float32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInstancesUsages) GetPriceOk() (*float32, bool) {
	if o == nil || o.Price == nil {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *BillingInstancesUsages) HasPrice() bool {
	if o != nil && o.Price != nil {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float32 and assigns it to the Price field.
func (o *BillingInstancesUsages) SetPrice(v float32) {
	o.Price = &v
}

// GetCreatedByUser returns the CreatedByUser field value if set, zero value otherwise.
func (o *BillingInstancesUsages) GetCreatedByUser() string {
	if o == nil || o.CreatedByUser == nil {
		var ret string
		return ret
	}
	return *o.CreatedByUser
}

// GetCreatedByUserOk returns a tuple with the CreatedByUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInstancesUsages) GetCreatedByUserOk() (*string, bool) {
	if o == nil || o.CreatedByUser == nil {
		return nil, false
	}
	return o.CreatedByUser, true
}

// HasCreatedByUser returns a boolean if a field has been set.
func (o *BillingInstancesUsages) HasCreatedByUser() bool {
	if o != nil && o.CreatedByUser != nil {
		return true
	}

	return false
}

// SetCreatedByUser gets a reference to the given string and assigns it to the CreatedByUser field.
func (o *BillingInstancesUsages) SetCreatedByUser(v string) {
	o.CreatedByUser = &v
}

// GetCreatedByUserId returns the CreatedByUserId field value if set, zero value otherwise.
func (o *BillingInstancesUsages) GetCreatedByUserId() int64 {
	if o == nil || o.CreatedByUserId == nil {
		var ret int64
		return ret
	}
	return *o.CreatedByUserId
}

// GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInstancesUsages) GetCreatedByUserIdOk() (*int64, bool) {
	if o == nil || o.CreatedByUserId == nil {
		return nil, false
	}
	return o.CreatedByUserId, true
}

// HasCreatedByUserId returns a boolean if a field has been set.
func (o *BillingInstancesUsages) HasCreatedByUserId() bool {
	if o != nil && o.CreatedByUserId != nil {
		return true
	}

	return false
}

// SetCreatedByUserId gets a reference to the given int64 and assigns it to the CreatedByUserId field.
func (o *BillingInstancesUsages) SetCreatedByUserId(v int64) {
	o.CreatedByUserId = &v
}

// GetSiteId returns the SiteId field value if set, zero value otherwise.
func (o *BillingInstancesUsages) GetSiteId() int64 {
	if o == nil || o.SiteId == nil {
		var ret int64
		return ret
	}
	return *o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInstancesUsages) GetSiteIdOk() (*int64, bool) {
	if o == nil || o.SiteId == nil {
		return nil, false
	}
	return o.SiteId, true
}

// HasSiteId returns a boolean if a field has been set.
func (o *BillingInstancesUsages) HasSiteId() bool {
	if o != nil && o.SiteId != nil {
		return true
	}

	return false
}

// SetSiteId gets a reference to the given int64 and assigns it to the SiteId field.
func (o *BillingInstancesUsages) SetSiteId(v int64) {
	o.SiteId = &v
}

// GetSiteName returns the SiteName field value if set, zero value otherwise.
func (o *BillingInstancesUsages) GetSiteName() string {
	if o == nil || o.SiteName == nil {
		var ret string
		return ret
	}
	return *o.SiteName
}

// GetSiteNameOk returns a tuple with the SiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInstancesUsages) GetSiteNameOk() (*string, bool) {
	if o == nil || o.SiteName == nil {
		return nil, false
	}
	return o.SiteName, true
}

// HasSiteName returns a boolean if a field has been set.
func (o *BillingInstancesUsages) HasSiteName() bool {
	if o != nil && o.SiteName != nil {
		return true
	}

	return false
}

// SetSiteName gets a reference to the given string and assigns it to the SiteName field.
func (o *BillingInstancesUsages) SetSiteName(v string) {
	o.SiteName = &v
}

// GetSiteUUID returns the SiteUUID field value if set, zero value otherwise.
func (o *BillingInstancesUsages) GetSiteUUID() string {
	if o == nil || o.SiteUUID == nil {
		var ret string
		return ret
	}
	return *o.SiteUUID
}

// GetSiteUUIDOk returns a tuple with the SiteUUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInstancesUsages) GetSiteUUIDOk() (*string, bool) {
	if o == nil || o.SiteUUID == nil {
		return nil, false
	}
	return o.SiteUUID, true
}

// HasSiteUUID returns a boolean if a field has been set.
func (o *BillingInstancesUsages) HasSiteUUID() bool {
	if o != nil && o.SiteUUID != nil {
		return true
	}

	return false
}

// SetSiteUUID gets a reference to the given string and assigns it to the SiteUUID field.
func (o *BillingInstancesUsages) SetSiteUUID(v string) {
	o.SiteUUID = &v
}

// GetSiteCode returns the SiteCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingInstancesUsages) GetSiteCode() string {
	if o == nil || o.SiteCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.SiteCode.Get()
}

// GetSiteCodeOk returns a tuple with the SiteCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingInstancesUsages) GetSiteCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SiteCode.Get(), o.SiteCode.IsSet()
}

// HasSiteCode returns a boolean if a field has been set.
func (o *BillingInstancesUsages) HasSiteCode() bool {
	if o != nil && o.SiteCode.IsSet() {
		return true
	}

	return false
}

// SetSiteCode gets a reference to the given NullableString and assigns it to the SiteCode field.
func (o *BillingInstancesUsages) SetSiteCode(v string) {
	o.SiteCode.Set(&v)
}
// SetSiteCodeNil sets the value for SiteCode to be an explicit nil
func (o *BillingInstancesUsages) SetSiteCodeNil() {
	o.SiteCode.Set(nil)
}

// UnsetSiteCode ensures that no value is present for SiteCode, not even an explicit nil
func (o *BillingInstancesUsages) UnsetSiteCode() {
	o.SiteCode.Unset()
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *BillingInstancesUsages) GetStartDate() time.Time {
	if o == nil || o.StartDate == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInstancesUsages) GetStartDateOk() (*time.Time, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *BillingInstancesUsages) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *BillingInstancesUsages) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *BillingInstancesUsages) GetEndDate() time.Time {
	if o == nil || o.EndDate == nil {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInstancesUsages) GetEndDateOk() (*time.Time, bool) {
	if o == nil || o.EndDate == nil {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *BillingInstancesUsages) HasEndDate() bool {
	if o != nil && o.EndDate != nil {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *BillingInstancesUsages) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BillingInstancesUsages) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInstancesUsages) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BillingInstancesUsages) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *BillingInstancesUsages) SetStatus(v string) {
	o.Status = &v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingInstancesUsages) GetTags() []map[string]interface{} {
	if o == nil  {
		var ret []map[string]interface{}
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingInstancesUsages) GetTagsOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *BillingInstancesUsages) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []map[string]interface{} and assigns it to the Tags field.
func (o *BillingInstancesUsages) SetTags(v []map[string]interface{}) {
	o.Tags = v
}

// GetApplicablePrices returns the ApplicablePrices field value if set, zero value otherwise.
func (o *BillingInstancesUsages) GetApplicablePrices() []BillingInstancesApplicablePrices {
	if o == nil || o.ApplicablePrices == nil {
		var ret []BillingInstancesApplicablePrices
		return ret
	}
	return *o.ApplicablePrices
}

// GetApplicablePricesOk returns a tuple with the ApplicablePrices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInstancesUsages) GetApplicablePricesOk() (*[]BillingInstancesApplicablePrices, bool) {
	if o == nil || o.ApplicablePrices == nil {
		return nil, false
	}
	return o.ApplicablePrices, true
}

// HasApplicablePrices returns a boolean if a field has been set.
func (o *BillingInstancesUsages) HasApplicablePrices() bool {
	if o != nil && o.ApplicablePrices != nil {
		return true
	}

	return false
}

// SetApplicablePrices gets a reference to the given []BillingInstancesApplicablePrices and assigns it to the ApplicablePrices field.
func (o *BillingInstancesUsages) SetApplicablePrices(v []BillingInstancesApplicablePrices) {
	o.ApplicablePrices = &v
}

// GetServicePlanId returns the ServicePlanId field value if set, zero value otherwise.
func (o *BillingInstancesUsages) GetServicePlanId() int64 {
	if o == nil || o.ServicePlanId == nil {
		var ret int64
		return ret
	}
	return *o.ServicePlanId
}

// GetServicePlanIdOk returns a tuple with the ServicePlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInstancesUsages) GetServicePlanIdOk() (*int64, bool) {
	if o == nil || o.ServicePlanId == nil {
		return nil, false
	}
	return o.ServicePlanId, true
}

// HasServicePlanId returns a boolean if a field has been set.
func (o *BillingInstancesUsages) HasServicePlanId() bool {
	if o != nil && o.ServicePlanId != nil {
		return true
	}

	return false
}

// SetServicePlanId gets a reference to the given int64 and assigns it to the ServicePlanId field.
func (o *BillingInstancesUsages) SetServicePlanId(v int64) {
	o.ServicePlanId = &v
}

// GetServicePlanName returns the ServicePlanName field value if set, zero value otherwise.
func (o *BillingInstancesUsages) GetServicePlanName() string {
	if o == nil || o.ServicePlanName == nil {
		var ret string
		return ret
	}
	return *o.ServicePlanName
}

// GetServicePlanNameOk returns a tuple with the ServicePlanName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInstancesUsages) GetServicePlanNameOk() (*string, bool) {
	if o == nil || o.ServicePlanName == nil {
		return nil, false
	}
	return o.ServicePlanName, true
}

// HasServicePlanName returns a boolean if a field has been set.
func (o *BillingInstancesUsages) HasServicePlanName() bool {
	if o != nil && o.ServicePlanName != nil {
		return true
	}

	return false
}

// SetServicePlanName gets a reference to the given string and assigns it to the ServicePlanName field.
func (o *BillingInstancesUsages) SetServicePlanName(v string) {
	o.ServicePlanName = &v
}

// GetResourcePoolId returns the ResourcePoolId field value if set, zero value otherwise.
func (o *BillingInstancesUsages) GetResourcePoolId() int64 {
	if o == nil || o.ResourcePoolId == nil {
		var ret int64
		return ret
	}
	return *o.ResourcePoolId
}

// GetResourcePoolIdOk returns a tuple with the ResourcePoolId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInstancesUsages) GetResourcePoolIdOk() (*int64, bool) {
	if o == nil || o.ResourcePoolId == nil {
		return nil, false
	}
	return o.ResourcePoolId, true
}

// HasResourcePoolId returns a boolean if a field has been set.
func (o *BillingInstancesUsages) HasResourcePoolId() bool {
	if o != nil && o.ResourcePoolId != nil {
		return true
	}

	return false
}

// SetResourcePoolId gets a reference to the given int64 and assigns it to the ResourcePoolId field.
func (o *BillingInstancesUsages) SetResourcePoolId(v int64) {
	o.ResourcePoolId = &v
}

// GetResourcePoolName returns the ResourcePoolName field value if set, zero value otherwise.
func (o *BillingInstancesUsages) GetResourcePoolName() string {
	if o == nil || o.ResourcePoolName == nil {
		var ret string
		return ret
	}
	return *o.ResourcePoolName
}

// GetResourcePoolNameOk returns a tuple with the ResourcePoolName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInstancesUsages) GetResourcePoolNameOk() (*string, bool) {
	if o == nil || o.ResourcePoolName == nil {
		return nil, false
	}
	return o.ResourcePoolName, true
}

// HasResourcePoolName returns a boolean if a field has been set.
func (o *BillingInstancesUsages) HasResourcePoolName() bool {
	if o != nil && o.ResourcePoolName != nil {
		return true
	}

	return false
}

// SetResourcePoolName gets a reference to the given string and assigns it to the ResourcePoolName field.
func (o *BillingInstancesUsages) SetResourcePoolName(v string) {
	o.ResourcePoolName = &v
}

func (o BillingInstancesUsages) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.InstanceName != nil {
		toSerialize["instanceName"] = o.InstanceName
	}
	if o.ZoneName != nil {
		toSerialize["zoneName"] = o.ZoneName
	}
	if o.AccountName != nil {
		toSerialize["accountName"] = o.AccountName
	}
	if o.Volumes != nil {
		toSerialize["volumes"] = o.Volumes
	}
	if o.MaxMemory != nil {
		toSerialize["maxMemory"] = o.MaxMemory
	}
	if o.MaxCpu.IsSet() {
		toSerialize["maxCpu"] = o.MaxCpu.Get()
	}
	if o.MaxCores != nil {
		toSerialize["maxCores"] = o.MaxCores
	}
	if o.ServerExternalId != nil {
		toSerialize["serverExternalId"] = o.ServerExternalId
	}
	if o.ServerInternalId != nil {
		toSerialize["serverInternalId"] = o.ServerInternalId
	}
	if o.PlanName != nil {
		toSerialize["planName"] = o.PlanName
	}
	if o.HourlyPrice != nil {
		toSerialize["hourlyPrice"] = o.HourlyPrice
	}
	if o.HourlyCost != nil {
		toSerialize["hourlyCost"] = o.HourlyCost
	}
	if o.Currency != nil {
		toSerialize["currency"] = o.Currency
	}
	if o.PricesUsed != nil {
		toSerialize["pricesUsed"] = o.PricesUsed
	}
	if o.Cost != nil {
		toSerialize["cost"] = o.Cost
	}
	if o.Price != nil {
		toSerialize["price"] = o.Price
	}
	if o.CreatedByUser != nil {
		toSerialize["createdByUser"] = o.CreatedByUser
	}
	if o.CreatedByUserId != nil {
		toSerialize["createdByUserId"] = o.CreatedByUserId
	}
	if o.SiteId != nil {
		toSerialize["siteId"] = o.SiteId
	}
	if o.SiteName != nil {
		toSerialize["siteName"] = o.SiteName
	}
	if o.SiteUUID != nil {
		toSerialize["siteUUID"] = o.SiteUUID
	}
	if o.SiteCode.IsSet() {
		toSerialize["siteCode"] = o.SiteCode.Get()
	}
	if o.StartDate != nil {
		toSerialize["startDate"] = o.StartDate
	}
	if o.EndDate != nil {
		toSerialize["endDate"] = o.EndDate
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.ApplicablePrices != nil {
		toSerialize["applicablePrices"] = o.ApplicablePrices
	}
	if o.ServicePlanId != nil {
		toSerialize["servicePlanId"] = o.ServicePlanId
	}
	if o.ServicePlanName != nil {
		toSerialize["servicePlanName"] = o.ServicePlanName
	}
	if o.ResourcePoolId != nil {
		toSerialize["resourcePoolId"] = o.ResourcePoolId
	}
	if o.ResourcePoolName != nil {
		toSerialize["resourcePoolName"] = o.ResourcePoolName
	}
	return json.Marshal(toSerialize)
}

type NullableBillingInstancesUsages struct {
	value *BillingInstancesUsages
	isSet bool
}

func (v NullableBillingInstancesUsages) Get() *BillingInstancesUsages {
	return v.value
}

func (v *NullableBillingInstancesUsages) Set(val *BillingInstancesUsages) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingInstancesUsages) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingInstancesUsages) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingInstancesUsages(val *BillingInstancesUsages) *NullableBillingInstancesUsages {
	return &NullableBillingInstancesUsages{value: val, isSet: true}
}

func (v NullableBillingInstancesUsages) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingInstancesUsages) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


