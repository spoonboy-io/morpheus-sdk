# ClusterDatastoresPermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourcePermissions** | Pointer to [**ClusterDatastoresPermissionsResourcePermissions**](clusterDatastores_permissions_resourcePermissions.md) |  | [optional] 
**Tenants** | Pointer to [**[]InlineResponse20040AppDeployInstance**](InlineResponse20040AppDeployInstance.md) |  | [optional] 

## Methods

### NewClusterDatastoresPermissions

`func NewClusterDatastoresPermissions() *ClusterDatastoresPermissions`

NewClusterDatastoresPermissions instantiates a new ClusterDatastoresPermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterDatastoresPermissionsWithDefaults

`func NewClusterDatastoresPermissionsWithDefaults() *ClusterDatastoresPermissions`

NewClusterDatastoresPermissionsWithDefaults instantiates a new ClusterDatastoresPermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourcePermissions

`func (o *ClusterDatastoresPermissions) GetResourcePermissions() ClusterDatastoresPermissionsResourcePermissions`

GetResourcePermissions returns the ResourcePermissions field if non-nil, zero value otherwise.

### GetResourcePermissionsOk

`func (o *ClusterDatastoresPermissions) GetResourcePermissionsOk() (*ClusterDatastoresPermissionsResourcePermissions, bool)`

GetResourcePermissionsOk returns a tuple with the ResourcePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermissions

`func (o *ClusterDatastoresPermissions) SetResourcePermissions(v ClusterDatastoresPermissionsResourcePermissions)`

SetResourcePermissions sets ResourcePermissions field to given value.

### HasResourcePermissions

`func (o *ClusterDatastoresPermissions) HasResourcePermissions() bool`

HasResourcePermissions returns a boolean if a field has been set.

### GetTenants

`func (o *ClusterDatastoresPermissions) GetTenants() []InlineResponse20040AppDeployInstance`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *ClusterDatastoresPermissions) GetTenantsOk() (*[]InlineResponse20040AppDeployInstance, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *ClusterDatastoresPermissions) SetTenants(v []InlineResponse20040AppDeployInstance)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *ClusterDatastoresPermissions) HasTenants() bool`

HasTenants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


