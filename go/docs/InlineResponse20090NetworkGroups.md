# InlineResponse20090NetworkGroups

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Networks** | Pointer to **[]int64** |  | [optional] 
**Subnets** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Tenants** | Pointer to [**[]InlineResponse20040AppDeployInstance**](InlineResponse20040AppDeployInstance.md) |  | [optional] 

## Methods

### NewInlineResponse20090NetworkGroups

`func NewInlineResponse20090NetworkGroups() *InlineResponse20090NetworkGroups`

NewInlineResponse20090NetworkGroups instantiates a new InlineResponse20090NetworkGroups object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20090NetworkGroupsWithDefaults

`func NewInlineResponse20090NetworkGroupsWithDefaults() *InlineResponse20090NetworkGroups`

NewInlineResponse20090NetworkGroupsWithDefaults instantiates a new InlineResponse20090NetworkGroups object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse20090NetworkGroups) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse20090NetworkGroups) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse20090NetworkGroups) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse20090NetworkGroups) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse20090NetworkGroups) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse20090NetworkGroups) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse20090NetworkGroups) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse20090NetworkGroups) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *InlineResponse20090NetworkGroups) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineResponse20090NetworkGroups) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineResponse20090NetworkGroups) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineResponse20090NetworkGroups) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVisibility

`func (o *InlineResponse20090NetworkGroups) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *InlineResponse20090NetworkGroups) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *InlineResponse20090NetworkGroups) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *InlineResponse20090NetworkGroups) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetActive

`func (o *InlineResponse20090NetworkGroups) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *InlineResponse20090NetworkGroups) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *InlineResponse20090NetworkGroups) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *InlineResponse20090NetworkGroups) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetNetworks

`func (o *InlineResponse20090NetworkGroups) GetNetworks() []int64`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *InlineResponse20090NetworkGroups) GetNetworksOk() (*[]int64, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *InlineResponse20090NetworkGroups) SetNetworks(v []int64)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *InlineResponse20090NetworkGroups) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetSubnets

`func (o *InlineResponse20090NetworkGroups) GetSubnets() []map[string]interface{}`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *InlineResponse20090NetworkGroups) GetSubnetsOk() (*[]map[string]interface{}, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *InlineResponse20090NetworkGroups) SetSubnets(v []map[string]interface{})`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *InlineResponse20090NetworkGroups) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.

### GetTenants

`func (o *InlineResponse20090NetworkGroups) GetTenants() []InlineResponse20040AppDeployInstance`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *InlineResponse20090NetworkGroups) GetTenantsOk() (*[]InlineResponse20040AppDeployInstance, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *InlineResponse20090NetworkGroups) SetTenants(v []InlineResponse20040AppDeployInstance)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *InlineResponse20090NetworkGroups) HasTenants() bool`

HasTenants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


