# CredentialOauth2ConfigConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GrantType** | **string** | OAuth 2.0 grant type | 
**AccessTokenUrl** | **string** | Token endpoint | 
**ClientId** | Pointer to **string** | Client ID | [optional] 
**ClientSecret** | Pointer to **string** | Client Secret | [optional] 
**Scope** | Pointer to **string** | Scope | [optional] 
**ClientAuth** | **string** | Client Authentication Method | 

## Methods

### NewCredentialOauth2ConfigConfig

`func NewCredentialOauth2ConfigConfig(grantType string, accessTokenUrl string, clientAuth string, ) *CredentialOauth2ConfigConfig`

NewCredentialOauth2ConfigConfig instantiates a new CredentialOauth2ConfigConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialOauth2ConfigConfigWithDefaults

`func NewCredentialOauth2ConfigConfigWithDefaults() *CredentialOauth2ConfigConfig`

NewCredentialOauth2ConfigConfigWithDefaults instantiates a new CredentialOauth2ConfigConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrantType

`func (o *CredentialOauth2ConfigConfig) GetGrantType() string`

GetGrantType returns the GrantType field if non-nil, zero value otherwise.

### GetGrantTypeOk

`func (o *CredentialOauth2ConfigConfig) GetGrantTypeOk() (*string, bool)`

GetGrantTypeOk returns a tuple with the GrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantType

`func (o *CredentialOauth2ConfigConfig) SetGrantType(v string)`

SetGrantType sets GrantType field to given value.


### GetAccessTokenUrl

`func (o *CredentialOauth2ConfigConfig) GetAccessTokenUrl() string`

GetAccessTokenUrl returns the AccessTokenUrl field if non-nil, zero value otherwise.

### GetAccessTokenUrlOk

`func (o *CredentialOauth2ConfigConfig) GetAccessTokenUrlOk() (*string, bool)`

GetAccessTokenUrlOk returns a tuple with the AccessTokenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenUrl

`func (o *CredentialOauth2ConfigConfig) SetAccessTokenUrl(v string)`

SetAccessTokenUrl sets AccessTokenUrl field to given value.


### GetClientId

`func (o *CredentialOauth2ConfigConfig) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *CredentialOauth2ConfigConfig) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *CredentialOauth2ConfigConfig) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *CredentialOauth2ConfigConfig) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *CredentialOauth2ConfigConfig) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *CredentialOauth2ConfigConfig) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *CredentialOauth2ConfigConfig) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *CredentialOauth2ConfigConfig) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetScope

`func (o *CredentialOauth2ConfigConfig) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *CredentialOauth2ConfigConfig) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *CredentialOauth2ConfigConfig) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *CredentialOauth2ConfigConfig) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetClientAuth

`func (o *CredentialOauth2ConfigConfig) GetClientAuth() string`

GetClientAuth returns the ClientAuth field if non-nil, zero value otherwise.

### GetClientAuthOk

`func (o *CredentialOauth2ConfigConfig) GetClientAuthOk() (*string, bool)`

GetClientAuthOk returns a tuple with the ClientAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientAuth

`func (o *CredentialOauth2ConfigConfig) SetClientAuth(v string)`

SetClientAuth sets ClientAuth field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


