# ApiSecurityGroupsSecurityGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name for your security group | 
**Description** | Pointer to **string** | Optional description field | [optional] 
**ZoneId** | **int64** | Scoped Cloud ID | 
**Active** | Pointer to **bool** | Set to &#x60;false&#x60; to disable a security group. | [optional] 
**CustomOptions** | Pointer to [**AnyOfsecurityGroupLocationAzureCustomOptionssecurityGroupLocationAwsCustomOptionssecurityGroupLocationOpenstackCustomOptions**](anyOf&lt;securityGroupLocationAzureCustomOptions,securityGroupLocationAwsCustomOptions,securityGroupLocationOpenstackCustomOptions&gt;.md) |  | [optional] 
**TenantPermissions** | Pointer to [**[]ApiSecurityGroupsSecurityGroupTenantPermissions**](ApiSecurityGroupsSecurityGroupTenantPermissions.md) |  | [optional] 
**ResourcePermissions** | Pointer to [**ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions**](_api_zones__zoneId__data_stores__id__datastore_resourcePermissions.md) |  | [optional] 

## Methods

### NewApiSecurityGroupsSecurityGroup

`func NewApiSecurityGroupsSecurityGroup(name string, zoneId int64, ) *ApiSecurityGroupsSecurityGroup`

NewApiSecurityGroupsSecurityGroup instantiates a new ApiSecurityGroupsSecurityGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiSecurityGroupsSecurityGroupWithDefaults

`func NewApiSecurityGroupsSecurityGroupWithDefaults() *ApiSecurityGroupsSecurityGroup`

NewApiSecurityGroupsSecurityGroupWithDefaults instantiates a new ApiSecurityGroupsSecurityGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiSecurityGroupsSecurityGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiSecurityGroupsSecurityGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiSecurityGroupsSecurityGroup) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ApiSecurityGroupsSecurityGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiSecurityGroupsSecurityGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiSecurityGroupsSecurityGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiSecurityGroupsSecurityGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetZoneId

`func (o *ApiSecurityGroupsSecurityGroup) GetZoneId() int64`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *ApiSecurityGroupsSecurityGroup) GetZoneIdOk() (*int64, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *ApiSecurityGroupsSecurityGroup) SetZoneId(v int64)`

SetZoneId sets ZoneId field to given value.


### GetActive

`func (o *ApiSecurityGroupsSecurityGroup) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ApiSecurityGroupsSecurityGroup) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ApiSecurityGroupsSecurityGroup) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ApiSecurityGroupsSecurityGroup) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCustomOptions

`func (o *ApiSecurityGroupsSecurityGroup) GetCustomOptions() AnyOfsecurityGroupLocationAzureCustomOptionssecurityGroupLocationAwsCustomOptionssecurityGroupLocationOpenstackCustomOptions`

GetCustomOptions returns the CustomOptions field if non-nil, zero value otherwise.

### GetCustomOptionsOk

`func (o *ApiSecurityGroupsSecurityGroup) GetCustomOptionsOk() (*AnyOfsecurityGroupLocationAzureCustomOptionssecurityGroupLocationAwsCustomOptionssecurityGroupLocationOpenstackCustomOptions, bool)`

GetCustomOptionsOk returns a tuple with the CustomOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomOptions

`func (o *ApiSecurityGroupsSecurityGroup) SetCustomOptions(v AnyOfsecurityGroupLocationAzureCustomOptionssecurityGroupLocationAwsCustomOptionssecurityGroupLocationOpenstackCustomOptions)`

SetCustomOptions sets CustomOptions field to given value.

### HasCustomOptions

`func (o *ApiSecurityGroupsSecurityGroup) HasCustomOptions() bool`

HasCustomOptions returns a boolean if a field has been set.

### GetTenantPermissions

`func (o *ApiSecurityGroupsSecurityGroup) GetTenantPermissions() []ApiSecurityGroupsSecurityGroupTenantPermissions`

GetTenantPermissions returns the TenantPermissions field if non-nil, zero value otherwise.

### GetTenantPermissionsOk

`func (o *ApiSecurityGroupsSecurityGroup) GetTenantPermissionsOk() (*[]ApiSecurityGroupsSecurityGroupTenantPermissions, bool)`

GetTenantPermissionsOk returns a tuple with the TenantPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantPermissions

`func (o *ApiSecurityGroupsSecurityGroup) SetTenantPermissions(v []ApiSecurityGroupsSecurityGroupTenantPermissions)`

SetTenantPermissions sets TenantPermissions field to given value.

### HasTenantPermissions

`func (o *ApiSecurityGroupsSecurityGroup) HasTenantPermissions() bool`

HasTenantPermissions returns a boolean if a field has been set.

### GetResourcePermissions

`func (o *ApiSecurityGroupsSecurityGroup) GetResourcePermissions() ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions`

GetResourcePermissions returns the ResourcePermissions field if non-nil, zero value otherwise.

### GetResourcePermissionsOk

`func (o *ApiSecurityGroupsSecurityGroup) GetResourcePermissionsOk() (*ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions, bool)`

GetResourcePermissionsOk returns a tuple with the ResourcePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermissions

`func (o *ApiSecurityGroupsSecurityGroup) SetResourcePermissions(v ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions)`

SetResourcePermissions sets ResourcePermissions field to given value.

### HasResourcePermissions

`func (o *ApiSecurityGroupsSecurityGroup) HasResourcePermissions() bool`

HasResourcePermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


