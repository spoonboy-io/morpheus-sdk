# BillingZones

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ZoneName** | Pointer to **string** |  | [optional] 
**ZoneId** | Pointer to **int64** |  | [optional] 
**ZoneUUID** | Pointer to **string** |  | [optional] 
**ZoneCode** | Pointer to **NullableString** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**PriceUnit** | Pointer to **string** |  | [optional] 
**ComputeServers** | Pointer to [**BillingComputeServers**](billing_computeServers.md) |  | [optional] 
**Instances** | Pointer to [**BillingInstances**](billing_instances.md) |  | [optional] 
**DiscoveredServers** | Pointer to [**BillingComputeServers**](billing_computeServers.md) |  | [optional] 
**LoadBalancers** | Pointer to [**BillingLoadBalancers**](billing_loadBalancers.md) |  | [optional] 
**VirtualImages** | Pointer to [**BillingVirtualImages**](billing_virtualImages.md) |  | [optional] 
**Snapshots** | Pointer to [**BillingSnapshots**](billing_snapshots.md) |  | [optional] 
**Price** | Pointer to **float32** |  | [optional] 
**Cost** | Pointer to **float32** |  | [optional] 

## Methods

### NewBillingZones

`func NewBillingZones() *BillingZones`

NewBillingZones instantiates a new BillingZones object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingZonesWithDefaults

`func NewBillingZonesWithDefaults() *BillingZones`

NewBillingZonesWithDefaults instantiates a new BillingZones object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetZoneName

`func (o *BillingZones) GetZoneName() string`

GetZoneName returns the ZoneName field if non-nil, zero value otherwise.

### GetZoneNameOk

`func (o *BillingZones) GetZoneNameOk() (*string, bool)`

GetZoneNameOk returns a tuple with the ZoneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneName

`func (o *BillingZones) SetZoneName(v string)`

SetZoneName sets ZoneName field to given value.

### HasZoneName

`func (o *BillingZones) HasZoneName() bool`

HasZoneName returns a boolean if a field has been set.

### GetZoneId

`func (o *BillingZones) GetZoneId() int64`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *BillingZones) GetZoneIdOk() (*int64, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *BillingZones) SetZoneId(v int64)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *BillingZones) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetZoneUUID

`func (o *BillingZones) GetZoneUUID() string`

GetZoneUUID returns the ZoneUUID field if non-nil, zero value otherwise.

### GetZoneUUIDOk

`func (o *BillingZones) GetZoneUUIDOk() (*string, bool)`

GetZoneUUIDOk returns a tuple with the ZoneUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneUUID

`func (o *BillingZones) SetZoneUUID(v string)`

SetZoneUUID sets ZoneUUID field to given value.

### HasZoneUUID

`func (o *BillingZones) HasZoneUUID() bool`

HasZoneUUID returns a boolean if a field has been set.

### GetZoneCode

`func (o *BillingZones) GetZoneCode() string`

GetZoneCode returns the ZoneCode field if non-nil, zero value otherwise.

### GetZoneCodeOk

`func (o *BillingZones) GetZoneCodeOk() (*string, bool)`

GetZoneCodeOk returns a tuple with the ZoneCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneCode

`func (o *BillingZones) SetZoneCode(v string)`

SetZoneCode sets ZoneCode field to given value.

### HasZoneCode

`func (o *BillingZones) HasZoneCode() bool`

HasZoneCode returns a boolean if a field has been set.

### SetZoneCodeNil

`func (o *BillingZones) SetZoneCodeNil(b bool)`

 SetZoneCodeNil sets the value for ZoneCode to be an explicit nil

### UnsetZoneCode
`func (o *BillingZones) UnsetZoneCode()`

UnsetZoneCode ensures that no value is present for ZoneCode, not even an explicit nil
### GetStartDate

`func (o *BillingZones) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *BillingZones) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *BillingZones) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *BillingZones) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *BillingZones) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *BillingZones) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *BillingZones) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *BillingZones) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetPriceUnit

`func (o *BillingZones) GetPriceUnit() string`

GetPriceUnit returns the PriceUnit field if non-nil, zero value otherwise.

### GetPriceUnitOk

`func (o *BillingZones) GetPriceUnitOk() (*string, bool)`

GetPriceUnitOk returns a tuple with the PriceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUnit

`func (o *BillingZones) SetPriceUnit(v string)`

SetPriceUnit sets PriceUnit field to given value.

### HasPriceUnit

`func (o *BillingZones) HasPriceUnit() bool`

HasPriceUnit returns a boolean if a field has been set.

### GetComputeServers

`func (o *BillingZones) GetComputeServers() BillingComputeServers`

GetComputeServers returns the ComputeServers field if non-nil, zero value otherwise.

### GetComputeServersOk

`func (o *BillingZones) GetComputeServersOk() (*BillingComputeServers, bool)`

GetComputeServersOk returns a tuple with the ComputeServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeServers

`func (o *BillingZones) SetComputeServers(v BillingComputeServers)`

SetComputeServers sets ComputeServers field to given value.

### HasComputeServers

`func (o *BillingZones) HasComputeServers() bool`

HasComputeServers returns a boolean if a field has been set.

### GetInstances

`func (o *BillingZones) GetInstances() BillingInstances`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *BillingZones) GetInstancesOk() (*BillingInstances, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *BillingZones) SetInstances(v BillingInstances)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *BillingZones) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### GetDiscoveredServers

`func (o *BillingZones) GetDiscoveredServers() BillingComputeServers`

GetDiscoveredServers returns the DiscoveredServers field if non-nil, zero value otherwise.

### GetDiscoveredServersOk

`func (o *BillingZones) GetDiscoveredServersOk() (*BillingComputeServers, bool)`

GetDiscoveredServersOk returns a tuple with the DiscoveredServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredServers

`func (o *BillingZones) SetDiscoveredServers(v BillingComputeServers)`

SetDiscoveredServers sets DiscoveredServers field to given value.

### HasDiscoveredServers

`func (o *BillingZones) HasDiscoveredServers() bool`

HasDiscoveredServers returns a boolean if a field has been set.

### GetLoadBalancers

`func (o *BillingZones) GetLoadBalancers() BillingLoadBalancers`

GetLoadBalancers returns the LoadBalancers field if non-nil, zero value otherwise.

### GetLoadBalancersOk

`func (o *BillingZones) GetLoadBalancersOk() (*BillingLoadBalancers, bool)`

GetLoadBalancersOk returns a tuple with the LoadBalancers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancers

`func (o *BillingZones) SetLoadBalancers(v BillingLoadBalancers)`

SetLoadBalancers sets LoadBalancers field to given value.

### HasLoadBalancers

`func (o *BillingZones) HasLoadBalancers() bool`

HasLoadBalancers returns a boolean if a field has been set.

### GetVirtualImages

`func (o *BillingZones) GetVirtualImages() BillingVirtualImages`

GetVirtualImages returns the VirtualImages field if non-nil, zero value otherwise.

### GetVirtualImagesOk

`func (o *BillingZones) GetVirtualImagesOk() (*BillingVirtualImages, bool)`

GetVirtualImagesOk returns a tuple with the VirtualImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualImages

`func (o *BillingZones) SetVirtualImages(v BillingVirtualImages)`

SetVirtualImages sets VirtualImages field to given value.

### HasVirtualImages

`func (o *BillingZones) HasVirtualImages() bool`

HasVirtualImages returns a boolean if a field has been set.

### GetSnapshots

`func (o *BillingZones) GetSnapshots() BillingSnapshots`

GetSnapshots returns the Snapshots field if non-nil, zero value otherwise.

### GetSnapshotsOk

`func (o *BillingZones) GetSnapshotsOk() (*BillingSnapshots, bool)`

GetSnapshotsOk returns a tuple with the Snapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshots

`func (o *BillingZones) SetSnapshots(v BillingSnapshots)`

SetSnapshots sets Snapshots field to given value.

### HasSnapshots

`func (o *BillingZones) HasSnapshots() bool`

HasSnapshots returns a boolean if a field has been set.

### GetPrice

`func (o *BillingZones) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *BillingZones) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *BillingZones) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *BillingZones) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetCost

`func (o *BillingZones) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *BillingZones) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *BillingZones) SetCost(v float32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *BillingZones) HasCost() bool`

HasCost returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


