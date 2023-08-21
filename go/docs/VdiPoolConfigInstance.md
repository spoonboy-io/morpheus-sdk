# VdiPoolConfigInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserGroup** | Pointer to [**ZoneVcenterConfigNetworkServer**](zoneVcenterConfig_networkServer.md) |  | [optional] 
**NetworkDomain** | Pointer to [**ApiBlueprintsIdUpdatePermissionsResourcePermissionSites**](_api_blueprints__id__update_permissions_resourcePermission_sites.md) |  | [optional] 

## Methods

### NewVdiPoolConfigInstance

`func NewVdiPoolConfigInstance() *VdiPoolConfigInstance`

NewVdiPoolConfigInstance instantiates a new VdiPoolConfigInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVdiPoolConfigInstanceWithDefaults

`func NewVdiPoolConfigInstanceWithDefaults() *VdiPoolConfigInstance`

NewVdiPoolConfigInstanceWithDefaults instantiates a new VdiPoolConfigInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserGroup

`func (o *VdiPoolConfigInstance) GetUserGroup() ZoneVcenterConfigNetworkServer`

GetUserGroup returns the UserGroup field if non-nil, zero value otherwise.

### GetUserGroupOk

`func (o *VdiPoolConfigInstance) GetUserGroupOk() (*ZoneVcenterConfigNetworkServer, bool)`

GetUserGroupOk returns a tuple with the UserGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroup

`func (o *VdiPoolConfigInstance) SetUserGroup(v ZoneVcenterConfigNetworkServer)`

SetUserGroup sets UserGroup field to given value.

### HasUserGroup

`func (o *VdiPoolConfigInstance) HasUserGroup() bool`

HasUserGroup returns a boolean if a field has been set.

### GetNetworkDomain

`func (o *VdiPoolConfigInstance) GetNetworkDomain() ApiBlueprintsIdUpdatePermissionsResourcePermissionSites`

GetNetworkDomain returns the NetworkDomain field if non-nil, zero value otherwise.

### GetNetworkDomainOk

`func (o *VdiPoolConfigInstance) GetNetworkDomainOk() (*ApiBlueprintsIdUpdatePermissionsResourcePermissionSites, bool)`

GetNetworkDomainOk returns a tuple with the NetworkDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDomain

`func (o *VdiPoolConfigInstance) SetNetworkDomain(v ApiBlueprintsIdUpdatePermissionsResourcePermissionSites)`

SetNetworkDomain sets NetworkDomain field to given value.

### HasNetworkDomain

`func (o *VdiPoolConfigInstance) HasNetworkDomain() bool`

HasNetworkDomain returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


