# OptionTypeListUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name | [optional] 
**Description** | Pointer to **NullableString** | Description | [optional] 
**Labels** | Pointer to **[]string** | Array of label strings, can be used for filtering. | [optional] 
**Type** | Pointer to **string** | Option List Type eg. &#x60;rest&#x60;, &#x60;api&#x60;, &#x60;ldap&#x60; or &#x60;manual&#x60;. | [optional] [default to "rest"]
**SourceUrl** | Pointer to **string** | Source URL, the http(s) URL to request data from. Required when type is rest. | [optional] 
**Visibility** | Pointer to **string** | Visibility | [optional] [default to "private"]
**SourceMethod** | Pointer to **string** | Source Method, the HTTP method. | [optional] [default to "GET"]
**ApiType** | Pointer to **NullableString** | Api Type, The code of the api option list to use, eg. clouds, environments, groups, instances, instance-wiki, networks, servicePlans, resourcePools, securityGroups, servers, server-wiki. Required when type is api. | [optional] 
**IgnoreSSLErrors** | Pointer to **bool** | Ignore SSL Errors. | [optional] [default to false]
**RealTime** | Pointer to **bool** | Real Time. | [optional] [default to false]
**Credential** | Pointer to [**OptionTypeListCreateCredential**](optionTypeListCreate_credential.md) |  | [optional] 
**ServiceUsername** | Pointer to **NullableString** | Username for authenticating with Basic Auth when type is rest or ldap username. | [optional] 
**ServicePassword** | Pointer to **NullableString** | Password for authenticating with Basic Auth when type is rest or ldap password. | [optional] 
**InitialDataset** | Pointer to **NullableString** | Initial Dataset. Create an initial JSON or CSV dataset to be used as the collection for this option list. It should be a list containing objects with properties &#39;name&#39;, and &#39;value&#39;. Required when type is manual. | [optional] 
**TranslationScript** | Pointer to **NullableString** | Translation Script. Create a js script to translate the result data object into an Array containing objects with properties &#39;name&#39; and &#39;value&#39;. The input data is provided as data and the result should be put on the global variable results. | [optional] 
**RequestScript** | Pointer to **NullableString** | Request Script. Create a js script to prepare the request. Return a data object as the body for a post, and return an array containing properties &#39;name&#39; and &#39;value&#39; for a get. The input data is provided as data and the result should be put on the global variable results. | [optional] 
**Config** | Pointer to [**OptionTypeListCreateConfig**](optionTypeListCreate_config.md) |  | [optional] 

## Methods

### NewOptionTypeListUpdate

`func NewOptionTypeListUpdate() *OptionTypeListUpdate`

NewOptionTypeListUpdate instantiates a new OptionTypeListUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptionTypeListUpdateWithDefaults

`func NewOptionTypeListUpdateWithDefaults() *OptionTypeListUpdate`

NewOptionTypeListUpdateWithDefaults instantiates a new OptionTypeListUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OptionTypeListUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OptionTypeListUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OptionTypeListUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OptionTypeListUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *OptionTypeListUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OptionTypeListUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OptionTypeListUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OptionTypeListUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *OptionTypeListUpdate) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *OptionTypeListUpdate) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLabels

`func (o *OptionTypeListUpdate) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *OptionTypeListUpdate) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *OptionTypeListUpdate) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *OptionTypeListUpdate) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *OptionTypeListUpdate) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *OptionTypeListUpdate) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetType

`func (o *OptionTypeListUpdate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OptionTypeListUpdate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OptionTypeListUpdate) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OptionTypeListUpdate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSourceUrl

`func (o *OptionTypeListUpdate) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *OptionTypeListUpdate) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *OptionTypeListUpdate) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.

### HasSourceUrl

`func (o *OptionTypeListUpdate) HasSourceUrl() bool`

HasSourceUrl returns a boolean if a field has been set.

### GetVisibility

`func (o *OptionTypeListUpdate) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *OptionTypeListUpdate) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *OptionTypeListUpdate) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *OptionTypeListUpdate) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetSourceMethod

`func (o *OptionTypeListUpdate) GetSourceMethod() string`

GetSourceMethod returns the SourceMethod field if non-nil, zero value otherwise.

### GetSourceMethodOk

`func (o *OptionTypeListUpdate) GetSourceMethodOk() (*string, bool)`

GetSourceMethodOk returns a tuple with the SourceMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceMethod

`func (o *OptionTypeListUpdate) SetSourceMethod(v string)`

SetSourceMethod sets SourceMethod field to given value.

### HasSourceMethod

`func (o *OptionTypeListUpdate) HasSourceMethod() bool`

HasSourceMethod returns a boolean if a field has been set.

### GetApiType

`func (o *OptionTypeListUpdate) GetApiType() string`

GetApiType returns the ApiType field if non-nil, zero value otherwise.

### GetApiTypeOk

`func (o *OptionTypeListUpdate) GetApiTypeOk() (*string, bool)`

GetApiTypeOk returns a tuple with the ApiType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiType

`func (o *OptionTypeListUpdate) SetApiType(v string)`

SetApiType sets ApiType field to given value.

### HasApiType

`func (o *OptionTypeListUpdate) HasApiType() bool`

HasApiType returns a boolean if a field has been set.

### SetApiTypeNil

`func (o *OptionTypeListUpdate) SetApiTypeNil(b bool)`

 SetApiTypeNil sets the value for ApiType to be an explicit nil

### UnsetApiType
`func (o *OptionTypeListUpdate) UnsetApiType()`

UnsetApiType ensures that no value is present for ApiType, not even an explicit nil
### GetIgnoreSSLErrors

`func (o *OptionTypeListUpdate) GetIgnoreSSLErrors() bool`

GetIgnoreSSLErrors returns the IgnoreSSLErrors field if non-nil, zero value otherwise.

### GetIgnoreSSLErrorsOk

`func (o *OptionTypeListUpdate) GetIgnoreSSLErrorsOk() (*bool, bool)`

GetIgnoreSSLErrorsOk returns a tuple with the IgnoreSSLErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreSSLErrors

`func (o *OptionTypeListUpdate) SetIgnoreSSLErrors(v bool)`

SetIgnoreSSLErrors sets IgnoreSSLErrors field to given value.

### HasIgnoreSSLErrors

`func (o *OptionTypeListUpdate) HasIgnoreSSLErrors() bool`

HasIgnoreSSLErrors returns a boolean if a field has been set.

### GetRealTime

`func (o *OptionTypeListUpdate) GetRealTime() bool`

GetRealTime returns the RealTime field if non-nil, zero value otherwise.

### GetRealTimeOk

`func (o *OptionTypeListUpdate) GetRealTimeOk() (*bool, bool)`

GetRealTimeOk returns a tuple with the RealTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealTime

`func (o *OptionTypeListUpdate) SetRealTime(v bool)`

SetRealTime sets RealTime field to given value.

### HasRealTime

`func (o *OptionTypeListUpdate) HasRealTime() bool`

HasRealTime returns a boolean if a field has been set.

### GetCredential

`func (o *OptionTypeListUpdate) GetCredential() OptionTypeListCreateCredential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *OptionTypeListUpdate) GetCredentialOk() (*OptionTypeListCreateCredential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *OptionTypeListUpdate) SetCredential(v OptionTypeListCreateCredential)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *OptionTypeListUpdate) HasCredential() bool`

HasCredential returns a boolean if a field has been set.

### GetServiceUsername

`func (o *OptionTypeListUpdate) GetServiceUsername() string`

GetServiceUsername returns the ServiceUsername field if non-nil, zero value otherwise.

### GetServiceUsernameOk

`func (o *OptionTypeListUpdate) GetServiceUsernameOk() (*string, bool)`

GetServiceUsernameOk returns a tuple with the ServiceUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUsername

`func (o *OptionTypeListUpdate) SetServiceUsername(v string)`

SetServiceUsername sets ServiceUsername field to given value.

### HasServiceUsername

`func (o *OptionTypeListUpdate) HasServiceUsername() bool`

HasServiceUsername returns a boolean if a field has been set.

### SetServiceUsernameNil

`func (o *OptionTypeListUpdate) SetServiceUsernameNil(b bool)`

 SetServiceUsernameNil sets the value for ServiceUsername to be an explicit nil

### UnsetServiceUsername
`func (o *OptionTypeListUpdate) UnsetServiceUsername()`

UnsetServiceUsername ensures that no value is present for ServiceUsername, not even an explicit nil
### GetServicePassword

`func (o *OptionTypeListUpdate) GetServicePassword() string`

GetServicePassword returns the ServicePassword field if non-nil, zero value otherwise.

### GetServicePasswordOk

`func (o *OptionTypeListUpdate) GetServicePasswordOk() (*string, bool)`

GetServicePasswordOk returns a tuple with the ServicePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePassword

`func (o *OptionTypeListUpdate) SetServicePassword(v string)`

SetServicePassword sets ServicePassword field to given value.

### HasServicePassword

`func (o *OptionTypeListUpdate) HasServicePassword() bool`

HasServicePassword returns a boolean if a field has been set.

### SetServicePasswordNil

`func (o *OptionTypeListUpdate) SetServicePasswordNil(b bool)`

 SetServicePasswordNil sets the value for ServicePassword to be an explicit nil

### UnsetServicePassword
`func (o *OptionTypeListUpdate) UnsetServicePassword()`

UnsetServicePassword ensures that no value is present for ServicePassword, not even an explicit nil
### GetInitialDataset

`func (o *OptionTypeListUpdate) GetInitialDataset() string`

GetInitialDataset returns the InitialDataset field if non-nil, zero value otherwise.

### GetInitialDatasetOk

`func (o *OptionTypeListUpdate) GetInitialDatasetOk() (*string, bool)`

GetInitialDatasetOk returns a tuple with the InitialDataset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialDataset

`func (o *OptionTypeListUpdate) SetInitialDataset(v string)`

SetInitialDataset sets InitialDataset field to given value.

### HasInitialDataset

`func (o *OptionTypeListUpdate) HasInitialDataset() bool`

HasInitialDataset returns a boolean if a field has been set.

### SetInitialDatasetNil

`func (o *OptionTypeListUpdate) SetInitialDatasetNil(b bool)`

 SetInitialDatasetNil sets the value for InitialDataset to be an explicit nil

### UnsetInitialDataset
`func (o *OptionTypeListUpdate) UnsetInitialDataset()`

UnsetInitialDataset ensures that no value is present for InitialDataset, not even an explicit nil
### GetTranslationScript

`func (o *OptionTypeListUpdate) GetTranslationScript() string`

GetTranslationScript returns the TranslationScript field if non-nil, zero value otherwise.

### GetTranslationScriptOk

`func (o *OptionTypeListUpdate) GetTranslationScriptOk() (*string, bool)`

GetTranslationScriptOk returns a tuple with the TranslationScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslationScript

`func (o *OptionTypeListUpdate) SetTranslationScript(v string)`

SetTranslationScript sets TranslationScript field to given value.

### HasTranslationScript

`func (o *OptionTypeListUpdate) HasTranslationScript() bool`

HasTranslationScript returns a boolean if a field has been set.

### SetTranslationScriptNil

`func (o *OptionTypeListUpdate) SetTranslationScriptNil(b bool)`

 SetTranslationScriptNil sets the value for TranslationScript to be an explicit nil

### UnsetTranslationScript
`func (o *OptionTypeListUpdate) UnsetTranslationScript()`

UnsetTranslationScript ensures that no value is present for TranslationScript, not even an explicit nil
### GetRequestScript

`func (o *OptionTypeListUpdate) GetRequestScript() string`

GetRequestScript returns the RequestScript field if non-nil, zero value otherwise.

### GetRequestScriptOk

`func (o *OptionTypeListUpdate) GetRequestScriptOk() (*string, bool)`

GetRequestScriptOk returns a tuple with the RequestScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestScript

`func (o *OptionTypeListUpdate) SetRequestScript(v string)`

SetRequestScript sets RequestScript field to given value.

### HasRequestScript

`func (o *OptionTypeListUpdate) HasRequestScript() bool`

HasRequestScript returns a boolean if a field has been set.

### SetRequestScriptNil

`func (o *OptionTypeListUpdate) SetRequestScriptNil(b bool)`

 SetRequestScriptNil sets the value for RequestScript to be an explicit nil

### UnsetRequestScript
`func (o *OptionTypeListUpdate) UnsetRequestScript()`

UnsetRequestScript ensures that no value is present for RequestScript, not even an explicit nil
### GetConfig

`func (o *OptionTypeListUpdate) GetConfig() OptionTypeListCreateConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *OptionTypeListUpdate) GetConfigOk() (*OptionTypeListCreateConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *OptionTypeListUpdate) SetConfig(v OptionTypeListCreateConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *OptionTypeListUpdate) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


