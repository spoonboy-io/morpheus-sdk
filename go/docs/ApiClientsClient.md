# ApiClientsClient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** | ClientId | 
**ClientSecret** | Pointer to **string** | ClientSecret | [optional] 
**AccessTokenValiditySeconds** | **NullableInt32** | Length of time accessToken is valid in seconds. | 
**RefreshTokenValiditySeconds** | **NullableInt32** | Length of time refreshToken is valid in seconds. | 

## Methods

### NewApiClientsClient

`func NewApiClientsClient(clientId string, accessTokenValiditySeconds NullableInt32, refreshTokenValiditySeconds NullableInt32, ) *ApiClientsClient`

NewApiClientsClient instantiates a new ApiClientsClient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiClientsClientWithDefaults

`func NewApiClientsClientWithDefaults() *ApiClientsClient`

NewApiClientsClientWithDefaults instantiates a new ApiClientsClient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *ApiClientsClient) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ApiClientsClient) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ApiClientsClient) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *ApiClientsClient) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *ApiClientsClient) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *ApiClientsClient) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *ApiClientsClient) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetAccessTokenValiditySeconds

`func (o *ApiClientsClient) GetAccessTokenValiditySeconds() int32`

GetAccessTokenValiditySeconds returns the AccessTokenValiditySeconds field if non-nil, zero value otherwise.

### GetAccessTokenValiditySecondsOk

`func (o *ApiClientsClient) GetAccessTokenValiditySecondsOk() (*int32, bool)`

GetAccessTokenValiditySecondsOk returns a tuple with the AccessTokenValiditySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenValiditySeconds

`func (o *ApiClientsClient) SetAccessTokenValiditySeconds(v int32)`

SetAccessTokenValiditySeconds sets AccessTokenValiditySeconds field to given value.


### SetAccessTokenValiditySecondsNil

`func (o *ApiClientsClient) SetAccessTokenValiditySecondsNil(b bool)`

 SetAccessTokenValiditySecondsNil sets the value for AccessTokenValiditySeconds to be an explicit nil

### UnsetAccessTokenValiditySeconds
`func (o *ApiClientsClient) UnsetAccessTokenValiditySeconds()`

UnsetAccessTokenValiditySeconds ensures that no value is present for AccessTokenValiditySeconds, not even an explicit nil
### GetRefreshTokenValiditySeconds

`func (o *ApiClientsClient) GetRefreshTokenValiditySeconds() int32`

GetRefreshTokenValiditySeconds returns the RefreshTokenValiditySeconds field if non-nil, zero value otherwise.

### GetRefreshTokenValiditySecondsOk

`func (o *ApiClientsClient) GetRefreshTokenValiditySecondsOk() (*int32, bool)`

GetRefreshTokenValiditySecondsOk returns a tuple with the RefreshTokenValiditySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenValiditySeconds

`func (o *ApiClientsClient) SetRefreshTokenValiditySeconds(v int32)`

SetRefreshTokenValiditySeconds sets RefreshTokenValiditySeconds field to given value.


### SetRefreshTokenValiditySecondsNil

`func (o *ApiClientsClient) SetRefreshTokenValiditySecondsNil(b bool)`

 SetRefreshTokenValiditySecondsNil sets the value for RefreshTokenValiditySeconds to be an explicit nil

### UnsetRefreshTokenValiditySeconds
`func (o *ApiClientsClient) UnsetRefreshTokenValiditySeconds()`

UnsetRefreshTokenValiditySeconds ensures that no value is present for RefreshTokenValiditySeconds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

