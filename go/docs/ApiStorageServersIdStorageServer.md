# ApiStorageServersIdStorageServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name | [optional] 
**Type** | Pointer to **string** | The &#x60;Storage Type&#x60; Code or ID | [optional] 
**Description** | Pointer to **string** | description | [optional] 
**Enabled** | Pointer to **bool** | The enabled flag | [optional] [default to true]
**Config** | Pointer to **map[string]interface{}** | Configuration object with parameters that vary by &#x60;type&#x60; | [optional] 
**Visibility** | Pointer to **string** | private or public | [optional] [default to "private"]
**Tenants** | Pointer to [**[]ApiBlueprintsIdUpdatePermissionsResourcePermissionSites**](ApiBlueprintsIdUpdatePermissionsResourcePermissionSites.md) | Array of tenant account ids that are allowed access | [optional] 

## Methods

### NewApiStorageServersIdStorageServer

`func NewApiStorageServersIdStorageServer() *ApiStorageServersIdStorageServer`

NewApiStorageServersIdStorageServer instantiates a new ApiStorageServersIdStorageServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiStorageServersIdStorageServerWithDefaults

`func NewApiStorageServersIdStorageServerWithDefaults() *ApiStorageServersIdStorageServer`

NewApiStorageServersIdStorageServerWithDefaults instantiates a new ApiStorageServersIdStorageServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiStorageServersIdStorageServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiStorageServersIdStorageServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiStorageServersIdStorageServer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiStorageServersIdStorageServer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *ApiStorageServersIdStorageServer) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiStorageServersIdStorageServer) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiStorageServersIdStorageServer) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApiStorageServersIdStorageServer) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDescription

`func (o *ApiStorageServersIdStorageServer) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiStorageServersIdStorageServer) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiStorageServersIdStorageServer) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiStorageServersIdStorageServer) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *ApiStorageServersIdStorageServer) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ApiStorageServersIdStorageServer) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ApiStorageServersIdStorageServer) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ApiStorageServersIdStorageServer) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetConfig

`func (o *ApiStorageServersIdStorageServer) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ApiStorageServersIdStorageServer) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ApiStorageServersIdStorageServer) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ApiStorageServersIdStorageServer) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetVisibility

`func (o *ApiStorageServersIdStorageServer) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ApiStorageServersIdStorageServer) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ApiStorageServersIdStorageServer) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ApiStorageServersIdStorageServer) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetTenants

`func (o *ApiStorageServersIdStorageServer) GetTenants() []ApiBlueprintsIdUpdatePermissionsResourcePermissionSites`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *ApiStorageServersIdStorageServer) GetTenantsOk() (*[]ApiBlueprintsIdUpdatePermissionsResourcePermissionSites, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *ApiStorageServersIdStorageServer) SetTenants(v []ApiBlueprintsIdUpdatePermissionsResourcePermissionSites)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *ApiStorageServersIdStorageServer) HasTenants() bool`

HasTenants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


