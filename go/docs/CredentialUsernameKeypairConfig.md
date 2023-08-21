# CredentialUsernameKeypairConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Credential Type Code | 
**Name** | **string** | A unique name scoped to your account for the credential | 
**Description** | Pointer to **string** | Optional Description | [optional] 
**Enabled** | Pointer to **bool** | Credential enabled | [optional] [default to true]
**Integration** | Pointer to [**NullableCredentialAccessSecretKeyConfigIntegration**](credentialAccessSecretKeyConfig_integration.md) |  | [optional] 
**Username** | **string** | Username | 
**AuthKey** | [**CredentialEmailPrivateKeyConfigAuthKey**](credentialEmailPrivateKeyConfig_authKey.md) |  | 

## Methods

### NewCredentialUsernameKeypairConfig

`func NewCredentialUsernameKeypairConfig(type_ string, name string, username string, authKey CredentialEmailPrivateKeyConfigAuthKey, ) *CredentialUsernameKeypairConfig`

NewCredentialUsernameKeypairConfig instantiates a new CredentialUsernameKeypairConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialUsernameKeypairConfigWithDefaults

`func NewCredentialUsernameKeypairConfigWithDefaults() *CredentialUsernameKeypairConfig`

NewCredentialUsernameKeypairConfigWithDefaults instantiates a new CredentialUsernameKeypairConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CredentialUsernameKeypairConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CredentialUsernameKeypairConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CredentialUsernameKeypairConfig) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *CredentialUsernameKeypairConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CredentialUsernameKeypairConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CredentialUsernameKeypairConfig) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CredentialUsernameKeypairConfig) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CredentialUsernameKeypairConfig) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CredentialUsernameKeypairConfig) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CredentialUsernameKeypairConfig) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *CredentialUsernameKeypairConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CredentialUsernameKeypairConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CredentialUsernameKeypairConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CredentialUsernameKeypairConfig) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetIntegration

`func (o *CredentialUsernameKeypairConfig) GetIntegration() CredentialAccessSecretKeyConfigIntegration`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *CredentialUsernameKeypairConfig) GetIntegrationOk() (*CredentialAccessSecretKeyConfigIntegration, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *CredentialUsernameKeypairConfig) SetIntegration(v CredentialAccessSecretKeyConfigIntegration)`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *CredentialUsernameKeypairConfig) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### SetIntegrationNil

`func (o *CredentialUsernameKeypairConfig) SetIntegrationNil(b bool)`

 SetIntegrationNil sets the value for Integration to be an explicit nil

### UnsetIntegration
`func (o *CredentialUsernameKeypairConfig) UnsetIntegration()`

UnsetIntegration ensures that no value is present for Integration, not even an explicit nil
### GetUsername

`func (o *CredentialUsernameKeypairConfig) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CredentialUsernameKeypairConfig) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CredentialUsernameKeypairConfig) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetAuthKey

`func (o *CredentialUsernameKeypairConfig) GetAuthKey() CredentialEmailPrivateKeyConfigAuthKey`

GetAuthKey returns the AuthKey field if non-nil, zero value otherwise.

### GetAuthKeyOk

`func (o *CredentialUsernameKeypairConfig) GetAuthKeyOk() (*CredentialEmailPrivateKeyConfigAuthKey, bool)`

GetAuthKeyOk returns a tuple with the AuthKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthKey

`func (o *CredentialUsernameKeypairConfig) SetAuthKey(v CredentialEmailPrivateKeyConfigAuthKey)`

SetAuthKey sets AuthKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


