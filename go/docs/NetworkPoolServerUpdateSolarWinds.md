# NetworkPoolServerUpdateSolarWinds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name | [optional] 
**Enabled** | Pointer to **bool** | Can be used to enable / disable the network pool server. | [optional] [default to true]
**ServiceUrl** | Pointer to **NullableString** | URL | [optional] 
**ServiceUsername** | Pointer to **NullableString** | Username | [optional] 
**ServicePassword** | Pointer to **NullableString** | Password | [optional] 
**ServiceThrottleRate** | Pointer to **NullableInt64** | Throttle Rate | [optional] [default to 0]
**IgnoreSsl** | Pointer to **bool** | Disable SSL SNI Verification | [optional] 
**Config** | Pointer to [**NetworkPoolServerCreateBluecatConfig**](networkPoolServerCreateBluecat_config.md) |  | [optional] 
**Credential** | Pointer to [**NetworkPoolServerCreateBluecatCredential**](networkPoolServerCreateBluecat_credential.md) |  | [optional] 

## Methods

### NewNetworkPoolServerUpdateSolarWinds

`func NewNetworkPoolServerUpdateSolarWinds() *NetworkPoolServerUpdateSolarWinds`

NewNetworkPoolServerUpdateSolarWinds instantiates a new NetworkPoolServerUpdateSolarWinds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkPoolServerUpdateSolarWindsWithDefaults

`func NewNetworkPoolServerUpdateSolarWindsWithDefaults() *NetworkPoolServerUpdateSolarWinds`

NewNetworkPoolServerUpdateSolarWindsWithDefaults instantiates a new NetworkPoolServerUpdateSolarWinds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NetworkPoolServerUpdateSolarWinds) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkPoolServerUpdateSolarWinds) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkPoolServerUpdateSolarWinds) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkPoolServerUpdateSolarWinds) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnabled

`func (o *NetworkPoolServerUpdateSolarWinds) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NetworkPoolServerUpdateSolarWinds) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *NetworkPoolServerUpdateSolarWinds) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *NetworkPoolServerUpdateSolarWinds) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetServiceUrl

`func (o *NetworkPoolServerUpdateSolarWinds) GetServiceUrl() string`

GetServiceUrl returns the ServiceUrl field if non-nil, zero value otherwise.

### GetServiceUrlOk

`func (o *NetworkPoolServerUpdateSolarWinds) GetServiceUrlOk() (*string, bool)`

GetServiceUrlOk returns a tuple with the ServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUrl

`func (o *NetworkPoolServerUpdateSolarWinds) SetServiceUrl(v string)`

SetServiceUrl sets ServiceUrl field to given value.

### HasServiceUrl

`func (o *NetworkPoolServerUpdateSolarWinds) HasServiceUrl() bool`

HasServiceUrl returns a boolean if a field has been set.

### SetServiceUrlNil

`func (o *NetworkPoolServerUpdateSolarWinds) SetServiceUrlNil(b bool)`

 SetServiceUrlNil sets the value for ServiceUrl to be an explicit nil

### UnsetServiceUrl
`func (o *NetworkPoolServerUpdateSolarWinds) UnsetServiceUrl()`

UnsetServiceUrl ensures that no value is present for ServiceUrl, not even an explicit nil
### GetServiceUsername

`func (o *NetworkPoolServerUpdateSolarWinds) GetServiceUsername() string`

GetServiceUsername returns the ServiceUsername field if non-nil, zero value otherwise.

### GetServiceUsernameOk

`func (o *NetworkPoolServerUpdateSolarWinds) GetServiceUsernameOk() (*string, bool)`

GetServiceUsernameOk returns a tuple with the ServiceUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUsername

`func (o *NetworkPoolServerUpdateSolarWinds) SetServiceUsername(v string)`

SetServiceUsername sets ServiceUsername field to given value.

### HasServiceUsername

`func (o *NetworkPoolServerUpdateSolarWinds) HasServiceUsername() bool`

HasServiceUsername returns a boolean if a field has been set.

### SetServiceUsernameNil

`func (o *NetworkPoolServerUpdateSolarWinds) SetServiceUsernameNil(b bool)`

 SetServiceUsernameNil sets the value for ServiceUsername to be an explicit nil

### UnsetServiceUsername
`func (o *NetworkPoolServerUpdateSolarWinds) UnsetServiceUsername()`

UnsetServiceUsername ensures that no value is present for ServiceUsername, not even an explicit nil
### GetServicePassword

`func (o *NetworkPoolServerUpdateSolarWinds) GetServicePassword() string`

GetServicePassword returns the ServicePassword field if non-nil, zero value otherwise.

### GetServicePasswordOk

`func (o *NetworkPoolServerUpdateSolarWinds) GetServicePasswordOk() (*string, bool)`

GetServicePasswordOk returns a tuple with the ServicePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePassword

`func (o *NetworkPoolServerUpdateSolarWinds) SetServicePassword(v string)`

SetServicePassword sets ServicePassword field to given value.

### HasServicePassword

`func (o *NetworkPoolServerUpdateSolarWinds) HasServicePassword() bool`

HasServicePassword returns a boolean if a field has been set.

### SetServicePasswordNil

`func (o *NetworkPoolServerUpdateSolarWinds) SetServicePasswordNil(b bool)`

 SetServicePasswordNil sets the value for ServicePassword to be an explicit nil

### UnsetServicePassword
`func (o *NetworkPoolServerUpdateSolarWinds) UnsetServicePassword()`

UnsetServicePassword ensures that no value is present for ServicePassword, not even an explicit nil
### GetServiceThrottleRate

`func (o *NetworkPoolServerUpdateSolarWinds) GetServiceThrottleRate() int64`

GetServiceThrottleRate returns the ServiceThrottleRate field if non-nil, zero value otherwise.

### GetServiceThrottleRateOk

`func (o *NetworkPoolServerUpdateSolarWinds) GetServiceThrottleRateOk() (*int64, bool)`

GetServiceThrottleRateOk returns a tuple with the ServiceThrottleRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceThrottleRate

`func (o *NetworkPoolServerUpdateSolarWinds) SetServiceThrottleRate(v int64)`

SetServiceThrottleRate sets ServiceThrottleRate field to given value.

### HasServiceThrottleRate

`func (o *NetworkPoolServerUpdateSolarWinds) HasServiceThrottleRate() bool`

HasServiceThrottleRate returns a boolean if a field has been set.

### SetServiceThrottleRateNil

`func (o *NetworkPoolServerUpdateSolarWinds) SetServiceThrottleRateNil(b bool)`

 SetServiceThrottleRateNil sets the value for ServiceThrottleRate to be an explicit nil

### UnsetServiceThrottleRate
`func (o *NetworkPoolServerUpdateSolarWinds) UnsetServiceThrottleRate()`

UnsetServiceThrottleRate ensures that no value is present for ServiceThrottleRate, not even an explicit nil
### GetIgnoreSsl

`func (o *NetworkPoolServerUpdateSolarWinds) GetIgnoreSsl() bool`

GetIgnoreSsl returns the IgnoreSsl field if non-nil, zero value otherwise.

### GetIgnoreSslOk

`func (o *NetworkPoolServerUpdateSolarWinds) GetIgnoreSslOk() (*bool, bool)`

GetIgnoreSslOk returns a tuple with the IgnoreSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreSsl

`func (o *NetworkPoolServerUpdateSolarWinds) SetIgnoreSsl(v bool)`

SetIgnoreSsl sets IgnoreSsl field to given value.

### HasIgnoreSsl

`func (o *NetworkPoolServerUpdateSolarWinds) HasIgnoreSsl() bool`

HasIgnoreSsl returns a boolean if a field has been set.

### GetConfig

`func (o *NetworkPoolServerUpdateSolarWinds) GetConfig() NetworkPoolServerCreateBluecatConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *NetworkPoolServerUpdateSolarWinds) GetConfigOk() (*NetworkPoolServerCreateBluecatConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *NetworkPoolServerUpdateSolarWinds) SetConfig(v NetworkPoolServerCreateBluecatConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *NetworkPoolServerUpdateSolarWinds) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCredential

`func (o *NetworkPoolServerUpdateSolarWinds) GetCredential() NetworkPoolServerCreateBluecatCredential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *NetworkPoolServerUpdateSolarWinds) GetCredentialOk() (*NetworkPoolServerCreateBluecatCredential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *NetworkPoolServerUpdateSolarWinds) SetCredential(v NetworkPoolServerCreateBluecatCredential)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *NetworkPoolServerUpdateSolarWinds) HasCredential() bool`

HasCredential returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


