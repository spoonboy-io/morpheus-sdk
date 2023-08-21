# NetworkRoutersCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name | 
**Type** | [**NetworkRoutersCreateType**](networkRoutersCreate_type.md) |  | 
**Site** | [**NetworkRoutersCreateSite**](networkRoutersCreate_site.md) |  | 
**Enabled** | Pointer to **bool** | Can be used to enable / disable the network router (true, false). Default is on | [optional] 
**Zone** | Pointer to [**NetworkRoutersCreateZone**](networkRoutersCreate_zone.md) |  | [optional] 
**NetworkServer** | Pointer to [**NetworkRoutersCreateNetworkServer**](networkRoutersCreate_networkServer.md) |  | [optional] 

## Methods

### NewNetworkRoutersCreate

`func NewNetworkRoutersCreate(name string, type_ NetworkRoutersCreateType, site NetworkRoutersCreateSite, ) *NetworkRoutersCreate`

NewNetworkRoutersCreate instantiates a new NetworkRoutersCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkRoutersCreateWithDefaults

`func NewNetworkRoutersCreateWithDefaults() *NetworkRoutersCreate`

NewNetworkRoutersCreateWithDefaults instantiates a new NetworkRoutersCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NetworkRoutersCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkRoutersCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkRoutersCreate) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *NetworkRoutersCreate) GetType() NetworkRoutersCreateType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NetworkRoutersCreate) GetTypeOk() (*NetworkRoutersCreateType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NetworkRoutersCreate) SetType(v NetworkRoutersCreateType)`

SetType sets Type field to given value.


### GetSite

`func (o *NetworkRoutersCreate) GetSite() NetworkRoutersCreateSite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *NetworkRoutersCreate) GetSiteOk() (*NetworkRoutersCreateSite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *NetworkRoutersCreate) SetSite(v NetworkRoutersCreateSite)`

SetSite sets Site field to given value.


### GetEnabled

`func (o *NetworkRoutersCreate) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NetworkRoutersCreate) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *NetworkRoutersCreate) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *NetworkRoutersCreate) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetZone

`func (o *NetworkRoutersCreate) GetZone() NetworkRoutersCreateZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *NetworkRoutersCreate) GetZoneOk() (*NetworkRoutersCreateZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *NetworkRoutersCreate) SetZone(v NetworkRoutersCreateZone)`

SetZone sets Zone field to given value.

### HasZone

`func (o *NetworkRoutersCreate) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetNetworkServer

`func (o *NetworkRoutersCreate) GetNetworkServer() NetworkRoutersCreateNetworkServer`

GetNetworkServer returns the NetworkServer field if non-nil, zero value otherwise.

### GetNetworkServerOk

`func (o *NetworkRoutersCreate) GetNetworkServerOk() (*NetworkRoutersCreateNetworkServer, bool)`

GetNetworkServerOk returns a tuple with the NetworkServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkServer

`func (o *NetworkRoutersCreate) SetNetworkServer(v NetworkRoutersCreateNetworkServer)`

SetNetworkServer sets NetworkServer field to given value.

### HasNetworkServer

`func (o *NetworkRoutersCreate) HasNetworkServer() bool`

HasNetworkServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


