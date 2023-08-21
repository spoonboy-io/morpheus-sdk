# NetworkPoolServerCreateSolarWinds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type Code (SolarWinds) | 
**Name** | **string** | Name | 
**Enabled** | Pointer to **bool** | Can be used to enable / disable the network pool server. | [optional] [default to true]
**ServiceUrl** | **NullableString** | URL | 
**ServiceUsername** | Pointer to **NullableString** | Username | [optional] 
**ServicePassword** | Pointer to **NullableString** | Password | [optional] 
**ServiceThrottleRate** | Pointer to **NullableInt64** | Throttle Rate | [optional] [default to 0]
**IgnoreSsl** | Pointer to **bool** | Disable SSL SNI Verification | [optional] [default to true]
**Config** | Pointer to [**NetworkPoolServerCreateBluecatConfig**](networkPoolServerCreateBluecat_config.md) |  | [optional] 
**Credential** | Pointer to [**NetworkPoolServerCreateBluecatCredential**](networkPoolServerCreateBluecat_credential.md) |  | [optional] 

## Methods

### NewNetworkPoolServerCreateSolarWinds

`func NewNetworkPoolServerCreateSolarWinds(type_ string, name string, serviceUrl NullableString, ) *NetworkPoolServerCreateSolarWinds`

NewNetworkPoolServerCreateSolarWinds instantiates a new NetworkPoolServerCreateSolarWinds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkPoolServerCreateSolarWindsWithDefaults

`func NewNetworkPoolServerCreateSolarWindsWithDefaults() *NetworkPoolServerCreateSolarWinds`

NewNetworkPoolServerCreateSolarWindsWithDefaults instantiates a new NetworkPoolServerCreateSolarWinds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *NetworkPoolServerCreateSolarWinds) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NetworkPoolServerCreateSolarWinds) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NetworkPoolServerCreateSolarWinds) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *NetworkPoolServerCreateSolarWinds) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkPoolServerCreateSolarWinds) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkPoolServerCreateSolarWinds) SetName(v string)`

SetName sets Name field to given value.


### GetEnabled

`func (o *NetworkPoolServerCreateSolarWinds) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NetworkPoolServerCreateSolarWinds) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *NetworkPoolServerCreateSolarWinds) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *NetworkPoolServerCreateSolarWinds) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetServiceUrl

`func (o *NetworkPoolServerCreateSolarWinds) GetServiceUrl() string`

GetServiceUrl returns the ServiceUrl field if non-nil, zero value otherwise.

### GetServiceUrlOk

`func (o *NetworkPoolServerCreateSolarWinds) GetServiceUrlOk() (*string, bool)`

GetServiceUrlOk returns a tuple with the ServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUrl

`func (o *NetworkPoolServerCreateSolarWinds) SetServiceUrl(v string)`

SetServiceUrl sets ServiceUrl field to given value.


### SetServiceUrlNil

`func (o *NetworkPoolServerCreateSolarWinds) SetServiceUrlNil(b bool)`

 SetServiceUrlNil sets the value for ServiceUrl to be an explicit nil

### UnsetServiceUrl
`func (o *NetworkPoolServerCreateSolarWinds) UnsetServiceUrl()`

UnsetServiceUrl ensures that no value is present for ServiceUrl, not even an explicit nil
### GetServiceUsername

`func (o *NetworkPoolServerCreateSolarWinds) GetServiceUsername() string`

GetServiceUsername returns the ServiceUsername field if non-nil, zero value otherwise.

### GetServiceUsernameOk

`func (o *NetworkPoolServerCreateSolarWinds) GetServiceUsernameOk() (*string, bool)`

GetServiceUsernameOk returns a tuple with the ServiceUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUsername

`func (o *NetworkPoolServerCreateSolarWinds) SetServiceUsername(v string)`

SetServiceUsername sets ServiceUsername field to given value.

### HasServiceUsername

`func (o *NetworkPoolServerCreateSolarWinds) HasServiceUsername() bool`

HasServiceUsername returns a boolean if a field has been set.

### SetServiceUsernameNil

`func (o *NetworkPoolServerCreateSolarWinds) SetServiceUsernameNil(b bool)`

 SetServiceUsernameNil sets the value for ServiceUsername to be an explicit nil

### UnsetServiceUsername
`func (o *NetworkPoolServerCreateSolarWinds) UnsetServiceUsername()`

UnsetServiceUsername ensures that no value is present for ServiceUsername, not even an explicit nil
### GetServicePassword

`func (o *NetworkPoolServerCreateSolarWinds) GetServicePassword() string`

GetServicePassword returns the ServicePassword field if non-nil, zero value otherwise.

### GetServicePasswordOk

`func (o *NetworkPoolServerCreateSolarWinds) GetServicePasswordOk() (*string, bool)`

GetServicePasswordOk returns a tuple with the ServicePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePassword

`func (o *NetworkPoolServerCreateSolarWinds) SetServicePassword(v string)`

SetServicePassword sets ServicePassword field to given value.

### HasServicePassword

`func (o *NetworkPoolServerCreateSolarWinds) HasServicePassword() bool`

HasServicePassword returns a boolean if a field has been set.

### SetServicePasswordNil

`func (o *NetworkPoolServerCreateSolarWinds) SetServicePasswordNil(b bool)`

 SetServicePasswordNil sets the value for ServicePassword to be an explicit nil

### UnsetServicePassword
`func (o *NetworkPoolServerCreateSolarWinds) UnsetServicePassword()`

UnsetServicePassword ensures that no value is present for ServicePassword, not even an explicit nil
### GetServiceThrottleRate

`func (o *NetworkPoolServerCreateSolarWinds) GetServiceThrottleRate() int64`

GetServiceThrottleRate returns the ServiceThrottleRate field if non-nil, zero value otherwise.

### GetServiceThrottleRateOk

`func (o *NetworkPoolServerCreateSolarWinds) GetServiceThrottleRateOk() (*int64, bool)`

GetServiceThrottleRateOk returns a tuple with the ServiceThrottleRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceThrottleRate

`func (o *NetworkPoolServerCreateSolarWinds) SetServiceThrottleRate(v int64)`

SetServiceThrottleRate sets ServiceThrottleRate field to given value.

### HasServiceThrottleRate

`func (o *NetworkPoolServerCreateSolarWinds) HasServiceThrottleRate() bool`

HasServiceThrottleRate returns a boolean if a field has been set.

### SetServiceThrottleRateNil

`func (o *NetworkPoolServerCreateSolarWinds) SetServiceThrottleRateNil(b bool)`

 SetServiceThrottleRateNil sets the value for ServiceThrottleRate to be an explicit nil

### UnsetServiceThrottleRate
`func (o *NetworkPoolServerCreateSolarWinds) UnsetServiceThrottleRate()`

UnsetServiceThrottleRate ensures that no value is present for ServiceThrottleRate, not even an explicit nil
### GetIgnoreSsl

`func (o *NetworkPoolServerCreateSolarWinds) GetIgnoreSsl() bool`

GetIgnoreSsl returns the IgnoreSsl field if non-nil, zero value otherwise.

### GetIgnoreSslOk

`func (o *NetworkPoolServerCreateSolarWinds) GetIgnoreSslOk() (*bool, bool)`

GetIgnoreSslOk returns a tuple with the IgnoreSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreSsl

`func (o *NetworkPoolServerCreateSolarWinds) SetIgnoreSsl(v bool)`

SetIgnoreSsl sets IgnoreSsl field to given value.

### HasIgnoreSsl

`func (o *NetworkPoolServerCreateSolarWinds) HasIgnoreSsl() bool`

HasIgnoreSsl returns a boolean if a field has been set.

### GetConfig

`func (o *NetworkPoolServerCreateSolarWinds) GetConfig() NetworkPoolServerCreateBluecatConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *NetworkPoolServerCreateSolarWinds) GetConfigOk() (*NetworkPoolServerCreateBluecatConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *NetworkPoolServerCreateSolarWinds) SetConfig(v NetworkPoolServerCreateBluecatConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *NetworkPoolServerCreateSolarWinds) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCredential

`func (o *NetworkPoolServerCreateSolarWinds) GetCredential() NetworkPoolServerCreateBluecatCredential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *NetworkPoolServerCreateSolarWinds) GetCredentialOk() (*NetworkPoolServerCreateBluecatCredential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *NetworkPoolServerCreateSolarWinds) SetCredential(v NetworkPoolServerCreateBluecatCredential)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *NetworkPoolServerCreateSolarWinds) HasCredential() bool`

HasCredential returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


