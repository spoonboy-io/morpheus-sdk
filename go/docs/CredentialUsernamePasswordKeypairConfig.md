# CredentialUsernamePasswordKeypairConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Credential Type Code | 
**Name** | **string** | A unique name scoped to your account for the credential | 
**Description** | Pointer to **string** | Optional Description | [optional] 
**Enabled** | Pointer to **bool** | Credential enabled | [optional] [default to true]
**Integration** | Pointer to [**NullableCredentialAccessSecretKeyConfigIntegration**](credentialAccessSecretKeyConfig_integration.md) |  | [optional] 
**Username** | **string** | Username | 
**Password** | **string** | User password, API Key, or applicable secret | 
**AuthKey** | [**CredentialEmailPrivateKeyConfigAuthKey**](credentialEmailPrivateKeyConfig_authKey.md) |  | 

## Methods

### NewCredentialUsernamePasswordKeypairConfig

`func NewCredentialUsernamePasswordKeypairConfig(type_ string, name string, username string, password string, authKey CredentialEmailPrivateKeyConfigAuthKey, ) *CredentialUsernamePasswordKeypairConfig`

NewCredentialUsernamePasswordKeypairConfig instantiates a new CredentialUsernamePasswordKeypairConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialUsernamePasswordKeypairConfigWithDefaults

`func NewCredentialUsernamePasswordKeypairConfigWithDefaults() *CredentialUsernamePasswordKeypairConfig`

NewCredentialUsernamePasswordKeypairConfigWithDefaults instantiates a new CredentialUsernamePasswordKeypairConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CredentialUsernamePasswordKeypairConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CredentialUsernamePasswordKeypairConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CredentialUsernamePasswordKeypairConfig) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *CredentialUsernamePasswordKeypairConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CredentialUsernamePasswordKeypairConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CredentialUsernamePasswordKeypairConfig) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CredentialUsernamePasswordKeypairConfig) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CredentialUsernamePasswordKeypairConfig) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CredentialUsernamePasswordKeypairConfig) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CredentialUsernamePasswordKeypairConfig) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *CredentialUsernamePasswordKeypairConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CredentialUsernamePasswordKeypairConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CredentialUsernamePasswordKeypairConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CredentialUsernamePasswordKeypairConfig) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetIntegration

`func (o *CredentialUsernamePasswordKeypairConfig) GetIntegration() CredentialAccessSecretKeyConfigIntegration`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *CredentialUsernamePasswordKeypairConfig) GetIntegrationOk() (*CredentialAccessSecretKeyConfigIntegration, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *CredentialUsernamePasswordKeypairConfig) SetIntegration(v CredentialAccessSecretKeyConfigIntegration)`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *CredentialUsernamePasswordKeypairConfig) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### SetIntegrationNil

`func (o *CredentialUsernamePasswordKeypairConfig) SetIntegrationNil(b bool)`

 SetIntegrationNil sets the value for Integration to be an explicit nil

### UnsetIntegration
`func (o *CredentialUsernamePasswordKeypairConfig) UnsetIntegration()`

UnsetIntegration ensures that no value is present for Integration, not even an explicit nil
### GetUsername

`func (o *CredentialUsernamePasswordKeypairConfig) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CredentialUsernamePasswordKeypairConfig) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CredentialUsernamePasswordKeypairConfig) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *CredentialUsernamePasswordKeypairConfig) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CredentialUsernamePasswordKeypairConfig) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CredentialUsernamePasswordKeypairConfig) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetAuthKey

`func (o *CredentialUsernamePasswordKeypairConfig) GetAuthKey() CredentialEmailPrivateKeyConfigAuthKey`

GetAuthKey returns the AuthKey field if non-nil, zero value otherwise.

### GetAuthKeyOk

`func (o *CredentialUsernamePasswordKeypairConfig) GetAuthKeyOk() (*CredentialEmailPrivateKeyConfigAuthKey, bool)`

GetAuthKeyOk returns a tuple with the AuthKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthKey

`func (o *CredentialUsernamePasswordKeypairConfig) SetAuthKey(v CredentialEmailPrivateKeyConfigAuthKey)`

SetAuthKey sets AuthKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


