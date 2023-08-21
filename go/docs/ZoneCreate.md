# ZoneCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A unique name scoped to your account for the cloud | 
**Description** | Pointer to **string** | Optional description field if you want to put more info there | [optional] 
**Code** | Pointer to **string** | Optional code for use with policies | [optional] 
**Location** | Pointer to **NullableString** | Optional location for your cloud | [optional] 
**Visibility** | Pointer to **string** | private or public | [optional] [default to "private"]
**ZoneType** | [**AnyOfobjectobject**](anyOf&lt;object,object&gt;.md) |  | 
**GroupId** | **int64** | Specifies which Server group this cloud should be assigned to | 
**AccountId** | Pointer to **int64** | Specifies which Tenant this cloud should be assigned to | [optional] 
**Enabled** | Pointer to **bool** | Can be used to disable the cloud | [optional] [default to true]
**AutoRecoverPowerState** | Pointer to **bool** | Automatically Power on VMs | [optional] [default to false]
**ScalePriority** | Pointer to **int64** | Scale Priority | [optional] [default to 1]
**LinkedAccountId** | Pointer to **int64** | Linked Account ID (enter commercial ID to get costing for AWS Govcloud) | [optional] 
**Config** | Pointer to **map[string]interface{}** | Map containing zone configuration settings. See the section on specific zone types for details. | [optional] 
**SecurityMode** | Pointer to **string** | host firewall. &#x60;off&#x60; or &#x60;internal&#x60;. a.k.a. \&quot;local firewall\&quot; | [optional] [default to "off"]
**Credential** | Pointer to [**ZoneCreateCredential**](zoneCreate_credential.md) |  | [optional] 

## Methods

### NewZoneCreate

`func NewZoneCreate(name string, zoneType AnyOfobjectobject, groupId int64, ) *ZoneCreate`

NewZoneCreate instantiates a new ZoneCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneCreateWithDefaults

`func NewZoneCreateWithDefaults() *ZoneCreate`

NewZoneCreateWithDefaults instantiates a new ZoneCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ZoneCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ZoneCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ZoneCreate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ZoneCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ZoneCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ZoneCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ZoneCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCode

`func (o *ZoneCreate) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ZoneCreate) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ZoneCreate) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ZoneCreate) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetLocation

`func (o *ZoneCreate) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ZoneCreate) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ZoneCreate) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ZoneCreate) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *ZoneCreate) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *ZoneCreate) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetVisibility

`func (o *ZoneCreate) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ZoneCreate) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ZoneCreate) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ZoneCreate) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetZoneType

`func (o *ZoneCreate) GetZoneType() AnyOfobjectobject`

GetZoneType returns the ZoneType field if non-nil, zero value otherwise.

### GetZoneTypeOk

`func (o *ZoneCreate) GetZoneTypeOk() (*AnyOfobjectobject, bool)`

GetZoneTypeOk returns a tuple with the ZoneType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneType

`func (o *ZoneCreate) SetZoneType(v AnyOfobjectobject)`

SetZoneType sets ZoneType field to given value.


### GetGroupId

`func (o *ZoneCreate) GetGroupId() int64`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ZoneCreate) GetGroupIdOk() (*int64, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ZoneCreate) SetGroupId(v int64)`

SetGroupId sets GroupId field to given value.


### GetAccountId

`func (o *ZoneCreate) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ZoneCreate) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ZoneCreate) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ZoneCreate) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetEnabled

`func (o *ZoneCreate) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ZoneCreate) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ZoneCreate) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ZoneCreate) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAutoRecoverPowerState

`func (o *ZoneCreate) GetAutoRecoverPowerState() bool`

GetAutoRecoverPowerState returns the AutoRecoverPowerState field if non-nil, zero value otherwise.

### GetAutoRecoverPowerStateOk

`func (o *ZoneCreate) GetAutoRecoverPowerStateOk() (*bool, bool)`

GetAutoRecoverPowerStateOk returns a tuple with the AutoRecoverPowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRecoverPowerState

`func (o *ZoneCreate) SetAutoRecoverPowerState(v bool)`

SetAutoRecoverPowerState sets AutoRecoverPowerState field to given value.

### HasAutoRecoverPowerState

`func (o *ZoneCreate) HasAutoRecoverPowerState() bool`

HasAutoRecoverPowerState returns a boolean if a field has been set.

### GetScalePriority

`func (o *ZoneCreate) GetScalePriority() int64`

GetScalePriority returns the ScalePriority field if non-nil, zero value otherwise.

### GetScalePriorityOk

`func (o *ZoneCreate) GetScalePriorityOk() (*int64, bool)`

GetScalePriorityOk returns a tuple with the ScalePriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScalePriority

`func (o *ZoneCreate) SetScalePriority(v int64)`

SetScalePriority sets ScalePriority field to given value.

### HasScalePriority

`func (o *ZoneCreate) HasScalePriority() bool`

HasScalePriority returns a boolean if a field has been set.

### GetLinkedAccountId

`func (o *ZoneCreate) GetLinkedAccountId() int64`

GetLinkedAccountId returns the LinkedAccountId field if non-nil, zero value otherwise.

### GetLinkedAccountIdOk

`func (o *ZoneCreate) GetLinkedAccountIdOk() (*int64, bool)`

GetLinkedAccountIdOk returns a tuple with the LinkedAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedAccountId

`func (o *ZoneCreate) SetLinkedAccountId(v int64)`

SetLinkedAccountId sets LinkedAccountId field to given value.

### HasLinkedAccountId

`func (o *ZoneCreate) HasLinkedAccountId() bool`

HasLinkedAccountId returns a boolean if a field has been set.

### GetConfig

`func (o *ZoneCreate) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ZoneCreate) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ZoneCreate) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ZoneCreate) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetSecurityMode

`func (o *ZoneCreate) GetSecurityMode() string`

GetSecurityMode returns the SecurityMode field if non-nil, zero value otherwise.

### GetSecurityModeOk

`func (o *ZoneCreate) GetSecurityModeOk() (*string, bool)`

GetSecurityModeOk returns a tuple with the SecurityMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityMode

`func (o *ZoneCreate) SetSecurityMode(v string)`

SetSecurityMode sets SecurityMode field to given value.

### HasSecurityMode

`func (o *ZoneCreate) HasSecurityMode() bool`

HasSecurityMode returns a boolean if a field has been set.

### GetCredential

`func (o *ZoneCreate) GetCredential() ZoneCreateCredential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *ZoneCreate) GetCredentialOk() (*ZoneCreateCredential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *ZoneCreate) SetCredential(v ZoneCreateCredential)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *ZoneCreate) HasCredential() bool`

HasCredential returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


