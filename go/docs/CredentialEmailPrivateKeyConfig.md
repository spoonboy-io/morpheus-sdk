# CredentialEmailPrivateKeyConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Credential Type Code | 
**Name** | **string** | A unique name scoped to your account for the credential | 
**Description** | Pointer to **string** | Optional Description | [optional] 
**Enabled** | Pointer to **bool** | Credential enabled | [optional] [default to true]
**Integration** | Pointer to [**NullableCredentialAccessSecretKeyConfigIntegration**](credentialAccessSecretKeyConfig_integration.md) |  | [optional] 
**Username** | **string** | Email | 
**AuthKey** | [**CredentialEmailPrivateKeyConfigAuthKey**](credentialEmailPrivateKeyConfig_authKey.md) |  | 

## Methods

### NewCredentialEmailPrivateKeyConfig

`func NewCredentialEmailPrivateKeyConfig(type_ string, name string, username string, authKey CredentialEmailPrivateKeyConfigAuthKey, ) *CredentialEmailPrivateKeyConfig`

NewCredentialEmailPrivateKeyConfig instantiates a new CredentialEmailPrivateKeyConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialEmailPrivateKeyConfigWithDefaults

`func NewCredentialEmailPrivateKeyConfigWithDefaults() *CredentialEmailPrivateKeyConfig`

NewCredentialEmailPrivateKeyConfigWithDefaults instantiates a new CredentialEmailPrivateKeyConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CredentialEmailPrivateKeyConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CredentialEmailPrivateKeyConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CredentialEmailPrivateKeyConfig) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *CredentialEmailPrivateKeyConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CredentialEmailPrivateKeyConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CredentialEmailPrivateKeyConfig) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CredentialEmailPrivateKeyConfig) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CredentialEmailPrivateKeyConfig) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CredentialEmailPrivateKeyConfig) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CredentialEmailPrivateKeyConfig) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *CredentialEmailPrivateKeyConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CredentialEmailPrivateKeyConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CredentialEmailPrivateKeyConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CredentialEmailPrivateKeyConfig) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetIntegration

`func (o *CredentialEmailPrivateKeyConfig) GetIntegration() CredentialAccessSecretKeyConfigIntegration`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *CredentialEmailPrivateKeyConfig) GetIntegrationOk() (*CredentialAccessSecretKeyConfigIntegration, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *CredentialEmailPrivateKeyConfig) SetIntegration(v CredentialAccessSecretKeyConfigIntegration)`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *CredentialEmailPrivateKeyConfig) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### SetIntegrationNil

`func (o *CredentialEmailPrivateKeyConfig) SetIntegrationNil(b bool)`

 SetIntegrationNil sets the value for Integration to be an explicit nil

### UnsetIntegration
`func (o *CredentialEmailPrivateKeyConfig) UnsetIntegration()`

UnsetIntegration ensures that no value is present for Integration, not even an explicit nil
### GetUsername

`func (o *CredentialEmailPrivateKeyConfig) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CredentialEmailPrivateKeyConfig) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CredentialEmailPrivateKeyConfig) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetAuthKey

`func (o *CredentialEmailPrivateKeyConfig) GetAuthKey() CredentialEmailPrivateKeyConfigAuthKey`

GetAuthKey returns the AuthKey field if non-nil, zero value otherwise.

### GetAuthKeyOk

`func (o *CredentialEmailPrivateKeyConfig) GetAuthKeyOk() (*CredentialEmailPrivateKeyConfigAuthKey, bool)`

GetAuthKeyOk returns a tuple with the AuthKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthKey

`func (o *CredentialEmailPrivateKeyConfig) SetAuthKey(v CredentialEmailPrivateKeyConfigAuthKey)`

SetAuthKey sets AuthKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


