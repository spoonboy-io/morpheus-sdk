# ApiStorageServersStorageServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name | 
**Type** | **string** | The &#x60;Storage Type&#x60; Code or ID | 
**Description** | Pointer to **string** | description | [optional] 
**Enabled** | Pointer to **bool** | The enabled flag | [optional] [default to true]
**Config** | **map[string]interface{}** | Configuration object with parameters that vary by &#x60;type&#x60; | 
**Visibility** | Pointer to **string** | private or public | [optional] [default to "private"]
**Tenants** | Pointer to [**[]ApiBlueprintsIdUpdatePermissionsResourcePermissionSites**](ApiBlueprintsIdUpdatePermissionsResourcePermissionSites.md) | Array of tenant account ids that are allowed access | [optional] 

## Methods

### NewApiStorageServersStorageServer

`func NewApiStorageServersStorageServer(name string, type_ string, config map[string]interface{}, ) *ApiStorageServersStorageServer`

NewApiStorageServersStorageServer instantiates a new ApiStorageServersStorageServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiStorageServersStorageServerWithDefaults

`func NewApiStorageServersStorageServerWithDefaults() *ApiStorageServersStorageServer`

NewApiStorageServersStorageServerWithDefaults instantiates a new ApiStorageServersStorageServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiStorageServersStorageServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiStorageServersStorageServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiStorageServersStorageServer) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *ApiStorageServersStorageServer) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiStorageServersStorageServer) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiStorageServersStorageServer) SetType(v string)`

SetType sets Type field to given value.


### GetDescription

`func (o *ApiStorageServersStorageServer) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiStorageServersStorageServer) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiStorageServersStorageServer) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiStorageServersStorageServer) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *ApiStorageServersStorageServer) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ApiStorageServersStorageServer) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ApiStorageServersStorageServer) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ApiStorageServersStorageServer) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetConfig

`func (o *ApiStorageServersStorageServer) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ApiStorageServersStorageServer) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ApiStorageServersStorageServer) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.


### GetVisibility

`func (o *ApiStorageServersStorageServer) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ApiStorageServersStorageServer) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ApiStorageServersStorageServer) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ApiStorageServersStorageServer) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetTenants

`func (o *ApiStorageServersStorageServer) GetTenants() []ApiBlueprintsIdUpdatePermissionsResourcePermissionSites`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *ApiStorageServersStorageServer) GetTenantsOk() (*[]ApiBlueprintsIdUpdatePermissionsResourcePermissionSites, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *ApiStorageServersStorageServer) SetTenants(v []ApiBlueprintsIdUpdatePermissionsResourcePermissionSites)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *ApiStorageServersStorageServer) HasTenants() bool`

HasTenants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


