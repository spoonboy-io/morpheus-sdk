# InlineResponse200131ResourcePoolGroups

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Mode** | Pointer to **string** | Pool selection mode. Valid values are &#x60;roundrobin&#x60; or &#x60;availablecapacity&#x60;. | [optional] 
**Pools** | Pointer to **[]int64** | Array of Resource Pool IDs | [optional] 
**Tenants** | Pointer to [**[]InlineResponse20040AppDeployInstance**](InlineResponse20040AppDeployInstance.md) |  | [optional] 
**ResourcePermission** | Pointer to [**InlineResponse200131ResourcePermission**](inline_response_200_131_resourcePermission.md) |  | [optional] 

## Methods

### NewInlineResponse200131ResourcePoolGroups

`func NewInlineResponse200131ResourcePoolGroups() *InlineResponse200131ResourcePoolGroups`

NewInlineResponse200131ResourcePoolGroups instantiates a new InlineResponse200131ResourcePoolGroups object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200131ResourcePoolGroupsWithDefaults

`func NewInlineResponse200131ResourcePoolGroupsWithDefaults() *InlineResponse200131ResourcePoolGroups`

NewInlineResponse200131ResourcePoolGroupsWithDefaults instantiates a new InlineResponse200131ResourcePoolGroups object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse200131ResourcePoolGroups) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse200131ResourcePoolGroups) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse200131ResourcePoolGroups) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse200131ResourcePoolGroups) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200131ResourcePoolGroups) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200131ResourcePoolGroups) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200131ResourcePoolGroups) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200131ResourcePoolGroups) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *InlineResponse200131ResourcePoolGroups) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineResponse200131ResourcePoolGroups) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineResponse200131ResourcePoolGroups) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineResponse200131ResourcePoolGroups) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVisibility

`func (o *InlineResponse200131ResourcePoolGroups) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *InlineResponse200131ResourcePoolGroups) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *InlineResponse200131ResourcePoolGroups) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *InlineResponse200131ResourcePoolGroups) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetMode

`func (o *InlineResponse200131ResourcePoolGroups) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *InlineResponse200131ResourcePoolGroups) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *InlineResponse200131ResourcePoolGroups) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *InlineResponse200131ResourcePoolGroups) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetPools

`func (o *InlineResponse200131ResourcePoolGroups) GetPools() []int64`

GetPools returns the Pools field if non-nil, zero value otherwise.

### GetPoolsOk

`func (o *InlineResponse200131ResourcePoolGroups) GetPoolsOk() (*[]int64, bool)`

GetPoolsOk returns a tuple with the Pools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPools

`func (o *InlineResponse200131ResourcePoolGroups) SetPools(v []int64)`

SetPools sets Pools field to given value.

### HasPools

`func (o *InlineResponse200131ResourcePoolGroups) HasPools() bool`

HasPools returns a boolean if a field has been set.

### GetTenants

`func (o *InlineResponse200131ResourcePoolGroups) GetTenants() []InlineResponse20040AppDeployInstance`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *InlineResponse200131ResourcePoolGroups) GetTenantsOk() (*[]InlineResponse20040AppDeployInstance, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *InlineResponse200131ResourcePoolGroups) SetTenants(v []InlineResponse20040AppDeployInstance)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *InlineResponse200131ResourcePoolGroups) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetResourcePermission

`func (o *InlineResponse200131ResourcePoolGroups) GetResourcePermission() InlineResponse200131ResourcePermission`

GetResourcePermission returns the ResourcePermission field if non-nil, zero value otherwise.

### GetResourcePermissionOk

`func (o *InlineResponse200131ResourcePoolGroups) GetResourcePermissionOk() (*InlineResponse200131ResourcePermission, bool)`

GetResourcePermissionOk returns a tuple with the ResourcePermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermission

`func (o *InlineResponse200131ResourcePoolGroups) SetResourcePermission(v InlineResponse200131ResourcePermission)`

SetResourcePermission sets ResourcePermission field to given value.

### HasResourcePermission

`func (o *InlineResponse200131ResourcePoolGroups) HasResourcePermission() bool`

HasResourcePermission returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


