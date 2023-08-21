# NetworkPoolServerCreatePhpIpam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type Code (phpIPAM) | 
**Name** | **string** | Name | 
**Enabled** | Pointer to **bool** | Can be used to enable / disable the network pool server. | [optional] [default to true]
**ServiceUrl** | **NullableString** | URL | 
**ServiceUsername** | Pointer to **NullableString** | Username | [optional] 
**ServicePassword** | Pointer to **NullableString** | Password | [optional] 
**ServiceThrottleRate** | Pointer to **NullableInt64** | Throttle Rate | [optional] [default to 0]
**IgnoreSsl** | Pointer to **bool** | Disable SSL SNI Verification | [optional] [default to true]
**NetworkFilter** | Pointer to **NullableString** | Network Filter | [optional] 
**Config** | [**NetworkPoolServerCreatePhpIpamConfig**](networkPoolServerCreatePhpIpam_config.md) |  | 
**Credential** | Pointer to [**NetworkPoolServerCreateBluecatCredential**](networkPoolServerCreateBluecat_credential.md) |  | [optional] 

## Methods

### NewNetworkPoolServerCreatePhpIpam

`func NewNetworkPoolServerCreatePhpIpam(type_ string, name string, serviceUrl NullableString, config NetworkPoolServerCreatePhpIpamConfig, ) *NetworkPoolServerCreatePhpIpam`

NewNetworkPoolServerCreatePhpIpam instantiates a new NetworkPoolServerCreatePhpIpam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkPoolServerCreatePhpIpamWithDefaults

`func NewNetworkPoolServerCreatePhpIpamWithDefaults() *NetworkPoolServerCreatePhpIpam`

NewNetworkPoolServerCreatePhpIpamWithDefaults instantiates a new NetworkPoolServerCreatePhpIpam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *NetworkPoolServerCreatePhpIpam) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NetworkPoolServerCreatePhpIpam) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NetworkPoolServerCreatePhpIpam) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *NetworkPoolServerCreatePhpIpam) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkPoolServerCreatePhpIpam) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkPoolServerCreatePhpIpam) SetName(v string)`

SetName sets Name field to given value.


### GetEnabled

`func (o *NetworkPoolServerCreatePhpIpam) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NetworkPoolServerCreatePhpIpam) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *NetworkPoolServerCreatePhpIpam) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *NetworkPoolServerCreatePhpIpam) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetServiceUrl

`func (o *NetworkPoolServerCreatePhpIpam) GetServiceUrl() string`

GetServiceUrl returns the ServiceUrl field if non-nil, zero value otherwise.

### GetServiceUrlOk

`func (o *NetworkPoolServerCreatePhpIpam) GetServiceUrlOk() (*string, bool)`

GetServiceUrlOk returns a tuple with the ServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUrl

`func (o *NetworkPoolServerCreatePhpIpam) SetServiceUrl(v string)`

SetServiceUrl sets ServiceUrl field to given value.


### SetServiceUrlNil

`func (o *NetworkPoolServerCreatePhpIpam) SetServiceUrlNil(b bool)`

 SetServiceUrlNil sets the value for ServiceUrl to be an explicit nil

### UnsetServiceUrl
`func (o *NetworkPoolServerCreatePhpIpam) UnsetServiceUrl()`

UnsetServiceUrl ensures that no value is present for ServiceUrl, not even an explicit nil
### GetServiceUsername

`func (o *NetworkPoolServerCreatePhpIpam) GetServiceUsername() string`

GetServiceUsername returns the ServiceUsername field if non-nil, zero value otherwise.

### GetServiceUsernameOk

`func (o *NetworkPoolServerCreatePhpIpam) GetServiceUsernameOk() (*string, bool)`

GetServiceUsernameOk returns a tuple with the ServiceUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUsername

`func (o *NetworkPoolServerCreatePhpIpam) SetServiceUsername(v string)`

SetServiceUsername sets ServiceUsername field to given value.

### HasServiceUsername

`func (o *NetworkPoolServerCreatePhpIpam) HasServiceUsername() bool`

HasServiceUsername returns a boolean if a field has been set.

### SetServiceUsernameNil

`func (o *NetworkPoolServerCreatePhpIpam) SetServiceUsernameNil(b bool)`

 SetServiceUsernameNil sets the value for ServiceUsername to be an explicit nil

### UnsetServiceUsername
`func (o *NetworkPoolServerCreatePhpIpam) UnsetServiceUsername()`

UnsetServiceUsername ensures that no value is present for ServiceUsername, not even an explicit nil
### GetServicePassword

`func (o *NetworkPoolServerCreatePhpIpam) GetServicePassword() string`

GetServicePassword returns the ServicePassword field if non-nil, zero value otherwise.

### GetServicePasswordOk

`func (o *NetworkPoolServerCreatePhpIpam) GetServicePasswordOk() (*string, bool)`

GetServicePasswordOk returns a tuple with the ServicePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePassword

`func (o *NetworkPoolServerCreatePhpIpam) SetServicePassword(v string)`

SetServicePassword sets ServicePassword field to given value.

### HasServicePassword

`func (o *NetworkPoolServerCreatePhpIpam) HasServicePassword() bool`

HasServicePassword returns a boolean if a field has been set.

### SetServicePasswordNil

`func (o *NetworkPoolServerCreatePhpIpam) SetServicePasswordNil(b bool)`

 SetServicePasswordNil sets the value for ServicePassword to be an explicit nil

### UnsetServicePassword
`func (o *NetworkPoolServerCreatePhpIpam) UnsetServicePassword()`

UnsetServicePassword ensures that no value is present for ServicePassword, not even an explicit nil
### GetServiceThrottleRate

`func (o *NetworkPoolServerCreatePhpIpam) GetServiceThrottleRate() int64`

GetServiceThrottleRate returns the ServiceThrottleRate field if non-nil, zero value otherwise.

### GetServiceThrottleRateOk

`func (o *NetworkPoolServerCreatePhpIpam) GetServiceThrottleRateOk() (*int64, bool)`

GetServiceThrottleRateOk returns a tuple with the ServiceThrottleRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceThrottleRate

`func (o *NetworkPoolServerCreatePhpIpam) SetServiceThrottleRate(v int64)`

SetServiceThrottleRate sets ServiceThrottleRate field to given value.

### HasServiceThrottleRate

`func (o *NetworkPoolServerCreatePhpIpam) HasServiceThrottleRate() bool`

HasServiceThrottleRate returns a boolean if a field has been set.

### SetServiceThrottleRateNil

`func (o *NetworkPoolServerCreatePhpIpam) SetServiceThrottleRateNil(b bool)`

 SetServiceThrottleRateNil sets the value for ServiceThrottleRate to be an explicit nil

### UnsetServiceThrottleRate
`func (o *NetworkPoolServerCreatePhpIpam) UnsetServiceThrottleRate()`

UnsetServiceThrottleRate ensures that no value is present for ServiceThrottleRate, not even an explicit nil
### GetIgnoreSsl

`func (o *NetworkPoolServerCreatePhpIpam) GetIgnoreSsl() bool`

GetIgnoreSsl returns the IgnoreSsl field if non-nil, zero value otherwise.

### GetIgnoreSslOk

`func (o *NetworkPoolServerCreatePhpIpam) GetIgnoreSslOk() (*bool, bool)`

GetIgnoreSslOk returns a tuple with the IgnoreSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreSsl

`func (o *NetworkPoolServerCreatePhpIpam) SetIgnoreSsl(v bool)`

SetIgnoreSsl sets IgnoreSsl field to given value.

### HasIgnoreSsl

`func (o *NetworkPoolServerCreatePhpIpam) HasIgnoreSsl() bool`

HasIgnoreSsl returns a boolean if a field has been set.

### GetNetworkFilter

`func (o *NetworkPoolServerCreatePhpIpam) GetNetworkFilter() string`

GetNetworkFilter returns the NetworkFilter field if non-nil, zero value otherwise.

### GetNetworkFilterOk

`func (o *NetworkPoolServerCreatePhpIpam) GetNetworkFilterOk() (*string, bool)`

GetNetworkFilterOk returns a tuple with the NetworkFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkFilter

`func (o *NetworkPoolServerCreatePhpIpam) SetNetworkFilter(v string)`

SetNetworkFilter sets NetworkFilter field to given value.

### HasNetworkFilter

`func (o *NetworkPoolServerCreatePhpIpam) HasNetworkFilter() bool`

HasNetworkFilter returns a boolean if a field has been set.

### SetNetworkFilterNil

`func (o *NetworkPoolServerCreatePhpIpam) SetNetworkFilterNil(b bool)`

 SetNetworkFilterNil sets the value for NetworkFilter to be an explicit nil

### UnsetNetworkFilter
`func (o *NetworkPoolServerCreatePhpIpam) UnsetNetworkFilter()`

UnsetNetworkFilter ensures that no value is present for NetworkFilter, not even an explicit nil
### GetConfig

`func (o *NetworkPoolServerCreatePhpIpam) GetConfig() NetworkPoolServerCreatePhpIpamConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *NetworkPoolServerCreatePhpIpam) GetConfigOk() (*NetworkPoolServerCreatePhpIpamConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *NetworkPoolServerCreatePhpIpam) SetConfig(v NetworkPoolServerCreatePhpIpamConfig)`

SetConfig sets Config field to given value.


### GetCredential

`func (o *NetworkPoolServerCreatePhpIpam) GetCredential() NetworkPoolServerCreateBluecatCredential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *NetworkPoolServerCreatePhpIpam) GetCredentialOk() (*NetworkPoolServerCreateBluecatCredential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *NetworkPoolServerCreatePhpIpam) SetCredential(v NetworkPoolServerCreateBluecatCredential)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *NetworkPoolServerCreatePhpIpam) HasCredential() bool`

HasCredential returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


