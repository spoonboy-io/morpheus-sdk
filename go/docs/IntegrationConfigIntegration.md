# IntegrationConfigIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name, a unique identifier for the integration | 
**Type** | **string** | Integration Type Code | 
**Enabled** | Pointer to **bool** | Set &#x60;true&#x60; to enable integration | [optional] 
**Refresh** | Pointer to **bool** | Pass &#x60;false&#x60; to skip refresh.  By default, refresh is done on update, when it is supported by the integration type.  | [optional] [default to true]
**Credential** | Pointer to [**OneOfobjectobject**](oneOf&lt;object,object&gt;.md) | Map containing Credential ID or the default {\&quot;type\&quot;: \&quot;local\&quot;}  which means use the values set in the local task options username and password instead of associating a credential.  | [optional] 

## Methods

### NewIntegrationConfigIntegration

`func NewIntegrationConfigIntegration(name string, type_ string, ) *IntegrationConfigIntegration`

NewIntegrationConfigIntegration instantiates a new IntegrationConfigIntegration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationConfigIntegrationWithDefaults

`func NewIntegrationConfigIntegrationWithDefaults() *IntegrationConfigIntegration`

NewIntegrationConfigIntegrationWithDefaults instantiates a new IntegrationConfigIntegration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IntegrationConfigIntegration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IntegrationConfigIntegration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IntegrationConfigIntegration) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *IntegrationConfigIntegration) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IntegrationConfigIntegration) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IntegrationConfigIntegration) SetType(v string)`

SetType sets Type field to given value.


### GetEnabled

`func (o *IntegrationConfigIntegration) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *IntegrationConfigIntegration) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *IntegrationConfigIntegration) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *IntegrationConfigIntegration) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRefresh

`func (o *IntegrationConfigIntegration) GetRefresh() bool`

GetRefresh returns the Refresh field if non-nil, zero value otherwise.

### GetRefreshOk

`func (o *IntegrationConfigIntegration) GetRefreshOk() (*bool, bool)`

GetRefreshOk returns a tuple with the Refresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefresh

`func (o *IntegrationConfigIntegration) SetRefresh(v bool)`

SetRefresh sets Refresh field to given value.

### HasRefresh

`func (o *IntegrationConfigIntegration) HasRefresh() bool`

HasRefresh returns a boolean if a field has been set.

### GetCredential

`func (o *IntegrationConfigIntegration) GetCredential() OneOfobjectobject`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *IntegrationConfigIntegration) GetCredentialOk() (*OneOfobjectobject, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *IntegrationConfigIntegration) SetCredential(v OneOfobjectobject)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *IntegrationConfigIntegration) HasCredential() bool`

HasCredential returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


