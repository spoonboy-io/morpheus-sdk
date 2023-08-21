# ApiZonesIdZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A unique name scoped to your account for the cloud | 
**Description** | Pointer to **string** | Optional description field if you want to put more info there | [optional] 
**Code** | Pointer to **string** | Optional code for use with policies | [optional] 
**Location** | Pointer to **string** | Optional location for your cloud | [optional] 
**Visibility** | Pointer to **string** | private or public | [optional] [default to "private"]
**ZoneType** | **string** | Map containing code or id of the cloud type | [default to "standard"]
**GroupId** | **int64** | Specifies which Server group this cloud should be assigned to | 
**AccountId** | Pointer to **int64** | Specifies which Tenant this cloud should be assigned to | [optional] 
**Enabled** | Pointer to **bool** | Can be used to disable the cloud | [optional] [default to true]
**AutoRecoverPowerState** | Pointer to **bool** | Automatically Power on VMs | [optional] [default to false]
**ScalePriority** | Pointer to **int64** | Scale Priority | [optional] [default to 1]
**LinkedAccountId** | Pointer to **int64** | Linked Account ID (enter commercial ID to get costing for AWS Govcloud) | [optional] 
**Config** | Pointer to **map[string]interface{}** | Map containing zone configuration settings. See the section on specific zone types for details. | [optional] 
**SecurityMode** | Pointer to **string** | host firewall. &#x60;off&#x60; or &#x60;internal&#x60;. a.k.a. \&quot;local firewall\&quot; | [optional] [default to "off"]
**DefaultCloudLogos** | Pointer to **bool** | Can be used to clear any custom logo and darkLogo, reverting to the defaults for the cloud type | [optional] 
**Credential** | **map[string]interface{}** | Map containing Credential ID. &#x60;local&#x60; means use the values set in the local cloud config instead of associating a credential. | [default to {"type":"local"}]

## Methods

### NewApiZonesIdZone

`func NewApiZonesIdZone(name string, zoneType string, groupId int64, credential map[string]interface{}, ) *ApiZonesIdZone`

NewApiZonesIdZone instantiates a new ApiZonesIdZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiZonesIdZoneWithDefaults

`func NewApiZonesIdZoneWithDefaults() *ApiZonesIdZone`

NewApiZonesIdZoneWithDefaults instantiates a new ApiZonesIdZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiZonesIdZone) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiZonesIdZone) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiZonesIdZone) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ApiZonesIdZone) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiZonesIdZone) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiZonesIdZone) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiZonesIdZone) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCode

`func (o *ApiZonesIdZone) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ApiZonesIdZone) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ApiZonesIdZone) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ApiZonesIdZone) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetLocation

`func (o *ApiZonesIdZone) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ApiZonesIdZone) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ApiZonesIdZone) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ApiZonesIdZone) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetVisibility

`func (o *ApiZonesIdZone) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ApiZonesIdZone) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ApiZonesIdZone) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ApiZonesIdZone) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetZoneType

`func (o *ApiZonesIdZone) GetZoneType() string`

GetZoneType returns the ZoneType field if non-nil, zero value otherwise.

### GetZoneTypeOk

`func (o *ApiZonesIdZone) GetZoneTypeOk() (*string, bool)`

GetZoneTypeOk returns a tuple with the ZoneType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneType

`func (o *ApiZonesIdZone) SetZoneType(v string)`

SetZoneType sets ZoneType field to given value.


### GetGroupId

`func (o *ApiZonesIdZone) GetGroupId() int64`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ApiZonesIdZone) GetGroupIdOk() (*int64, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ApiZonesIdZone) SetGroupId(v int64)`

SetGroupId sets GroupId field to given value.


### GetAccountId

`func (o *ApiZonesIdZone) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ApiZonesIdZone) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ApiZonesIdZone) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ApiZonesIdZone) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetEnabled

`func (o *ApiZonesIdZone) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ApiZonesIdZone) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ApiZonesIdZone) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ApiZonesIdZone) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAutoRecoverPowerState

`func (o *ApiZonesIdZone) GetAutoRecoverPowerState() bool`

GetAutoRecoverPowerState returns the AutoRecoverPowerState field if non-nil, zero value otherwise.

### GetAutoRecoverPowerStateOk

`func (o *ApiZonesIdZone) GetAutoRecoverPowerStateOk() (*bool, bool)`

GetAutoRecoverPowerStateOk returns a tuple with the AutoRecoverPowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRecoverPowerState

`func (o *ApiZonesIdZone) SetAutoRecoverPowerState(v bool)`

SetAutoRecoverPowerState sets AutoRecoverPowerState field to given value.

### HasAutoRecoverPowerState

`func (o *ApiZonesIdZone) HasAutoRecoverPowerState() bool`

HasAutoRecoverPowerState returns a boolean if a field has been set.

### GetScalePriority

`func (o *ApiZonesIdZone) GetScalePriority() int64`

GetScalePriority returns the ScalePriority field if non-nil, zero value otherwise.

### GetScalePriorityOk

`func (o *ApiZonesIdZone) GetScalePriorityOk() (*int64, bool)`

GetScalePriorityOk returns a tuple with the ScalePriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScalePriority

`func (o *ApiZonesIdZone) SetScalePriority(v int64)`

SetScalePriority sets ScalePriority field to given value.

### HasScalePriority

`func (o *ApiZonesIdZone) HasScalePriority() bool`

HasScalePriority returns a boolean if a field has been set.

### GetLinkedAccountId

`func (o *ApiZonesIdZone) GetLinkedAccountId() int64`

GetLinkedAccountId returns the LinkedAccountId field if non-nil, zero value otherwise.

### GetLinkedAccountIdOk

`func (o *ApiZonesIdZone) GetLinkedAccountIdOk() (*int64, bool)`

GetLinkedAccountIdOk returns a tuple with the LinkedAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedAccountId

`func (o *ApiZonesIdZone) SetLinkedAccountId(v int64)`

SetLinkedAccountId sets LinkedAccountId field to given value.

### HasLinkedAccountId

`func (o *ApiZonesIdZone) HasLinkedAccountId() bool`

HasLinkedAccountId returns a boolean if a field has been set.

### GetConfig

`func (o *ApiZonesIdZone) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ApiZonesIdZone) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ApiZonesIdZone) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ApiZonesIdZone) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetSecurityMode

`func (o *ApiZonesIdZone) GetSecurityMode() string`

GetSecurityMode returns the SecurityMode field if non-nil, zero value otherwise.

### GetSecurityModeOk

`func (o *ApiZonesIdZone) GetSecurityModeOk() (*string, bool)`

GetSecurityModeOk returns a tuple with the SecurityMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityMode

`func (o *ApiZonesIdZone) SetSecurityMode(v string)`

SetSecurityMode sets SecurityMode field to given value.

### HasSecurityMode

`func (o *ApiZonesIdZone) HasSecurityMode() bool`

HasSecurityMode returns a boolean if a field has been set.

### GetDefaultCloudLogos

`func (o *ApiZonesIdZone) GetDefaultCloudLogos() bool`

GetDefaultCloudLogos returns the DefaultCloudLogos field if non-nil, zero value otherwise.

### GetDefaultCloudLogosOk

`func (o *ApiZonesIdZone) GetDefaultCloudLogosOk() (*bool, bool)`

GetDefaultCloudLogosOk returns a tuple with the DefaultCloudLogos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCloudLogos

`func (o *ApiZonesIdZone) SetDefaultCloudLogos(v bool)`

SetDefaultCloudLogos sets DefaultCloudLogos field to given value.

### HasDefaultCloudLogos

`func (o *ApiZonesIdZone) HasDefaultCloudLogos() bool`

HasDefaultCloudLogos returns a boolean if a field has been set.

### GetCredential

`func (o *ApiZonesIdZone) GetCredential() map[string]interface{}`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *ApiZonesIdZone) GetCredentialOk() (*map[string]interface{}, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *ApiZonesIdZone) SetCredential(v map[string]interface{})`

SetCredential sets Credential field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


