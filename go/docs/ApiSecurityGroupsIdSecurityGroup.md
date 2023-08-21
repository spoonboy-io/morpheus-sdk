# ApiSecurityGroupsIdSecurityGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name for your security group | [optional] 
**Description** | Pointer to **string** | Optional description field | [optional] 
**Active** | Pointer to **bool** | Set to &#x60;false&#x60; to disable a security group. | [optional] 
**TenantPermissions** | Pointer to [**[]ApiSecurityGroupsSecurityGroupTenantPermissions**](ApiSecurityGroupsSecurityGroupTenantPermissions.md) |  | [optional] 
**ResourcePermissions** | Pointer to [**ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions**](_api_zones__zoneId__data_stores__id__datastore_resourcePermissions.md) |  | [optional] 

## Methods

### NewApiSecurityGroupsIdSecurityGroup

`func NewApiSecurityGroupsIdSecurityGroup() *ApiSecurityGroupsIdSecurityGroup`

NewApiSecurityGroupsIdSecurityGroup instantiates a new ApiSecurityGroupsIdSecurityGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiSecurityGroupsIdSecurityGroupWithDefaults

`func NewApiSecurityGroupsIdSecurityGroupWithDefaults() *ApiSecurityGroupsIdSecurityGroup`

NewApiSecurityGroupsIdSecurityGroupWithDefaults instantiates a new ApiSecurityGroupsIdSecurityGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiSecurityGroupsIdSecurityGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiSecurityGroupsIdSecurityGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiSecurityGroupsIdSecurityGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiSecurityGroupsIdSecurityGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ApiSecurityGroupsIdSecurityGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiSecurityGroupsIdSecurityGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiSecurityGroupsIdSecurityGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiSecurityGroupsIdSecurityGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetActive

`func (o *ApiSecurityGroupsIdSecurityGroup) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ApiSecurityGroupsIdSecurityGroup) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ApiSecurityGroupsIdSecurityGroup) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ApiSecurityGroupsIdSecurityGroup) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetTenantPermissions

`func (o *ApiSecurityGroupsIdSecurityGroup) GetTenantPermissions() []ApiSecurityGroupsSecurityGroupTenantPermissions`

GetTenantPermissions returns the TenantPermissions field if non-nil, zero value otherwise.

### GetTenantPermissionsOk

`func (o *ApiSecurityGroupsIdSecurityGroup) GetTenantPermissionsOk() (*[]ApiSecurityGroupsSecurityGroupTenantPermissions, bool)`

GetTenantPermissionsOk returns a tuple with the TenantPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantPermissions

`func (o *ApiSecurityGroupsIdSecurityGroup) SetTenantPermissions(v []ApiSecurityGroupsSecurityGroupTenantPermissions)`

SetTenantPermissions sets TenantPermissions field to given value.

### HasTenantPermissions

`func (o *ApiSecurityGroupsIdSecurityGroup) HasTenantPermissions() bool`

HasTenantPermissions returns a boolean if a field has been set.

### GetResourcePermissions

`func (o *ApiSecurityGroupsIdSecurityGroup) GetResourcePermissions() ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions`

GetResourcePermissions returns the ResourcePermissions field if non-nil, zero value otherwise.

### GetResourcePermissionsOk

`func (o *ApiSecurityGroupsIdSecurityGroup) GetResourcePermissionsOk() (*ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions, bool)`

GetResourcePermissionsOk returns a tuple with the ResourcePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermissions

`func (o *ApiSecurityGroupsIdSecurityGroup) SetResourcePermissions(v ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions)`

SetResourcePermissions sets ResourcePermissions field to given value.

### HasResourcePermissions

`func (o *ApiSecurityGroupsIdSecurityGroup) HasResourcePermissions() bool`

HasResourcePermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


