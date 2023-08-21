# ClusterDatastorePermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourcePermissions** | Pointer to [**ClusterDatastorePermissionsResourcePermissions**](clusterDatastore_permissions_resourcePermissions.md) |  | [optional] 
**TenantPermissions** | Pointer to [**InlineResponse20095NetworkRouterPermissionsTenantPermissions**](inline_response_200_95_networkRouter_permissions_tenantPermissions.md) |  | [optional] 

## Methods

### NewClusterDatastorePermissions

`func NewClusterDatastorePermissions() *ClusterDatastorePermissions`

NewClusterDatastorePermissions instantiates a new ClusterDatastorePermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterDatastorePermissionsWithDefaults

`func NewClusterDatastorePermissionsWithDefaults() *ClusterDatastorePermissions`

NewClusterDatastorePermissionsWithDefaults instantiates a new ClusterDatastorePermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourcePermissions

`func (o *ClusterDatastorePermissions) GetResourcePermissions() ClusterDatastorePermissionsResourcePermissions`

GetResourcePermissions returns the ResourcePermissions field if non-nil, zero value otherwise.

### GetResourcePermissionsOk

`func (o *ClusterDatastorePermissions) GetResourcePermissionsOk() (*ClusterDatastorePermissionsResourcePermissions, bool)`

GetResourcePermissionsOk returns a tuple with the ResourcePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermissions

`func (o *ClusterDatastorePermissions) SetResourcePermissions(v ClusterDatastorePermissionsResourcePermissions)`

SetResourcePermissions sets ResourcePermissions field to given value.

### HasResourcePermissions

`func (o *ClusterDatastorePermissions) HasResourcePermissions() bool`

HasResourcePermissions returns a boolean if a field has been set.

### GetTenantPermissions

`func (o *ClusterDatastorePermissions) GetTenantPermissions() InlineResponse20095NetworkRouterPermissionsTenantPermissions`

GetTenantPermissions returns the TenantPermissions field if non-nil, zero value otherwise.

### GetTenantPermissionsOk

`func (o *ClusterDatastorePermissions) GetTenantPermissionsOk() (*InlineResponse20095NetworkRouterPermissionsTenantPermissions, bool)`

GetTenantPermissionsOk returns a tuple with the TenantPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantPermissions

`func (o *ClusterDatastorePermissions) SetTenantPermissions(v InlineResponse20095NetworkRouterPermissionsTenantPermissions)`

SetTenantPermissions sets TenantPermissions field to given value.

### HasTenantPermissions

`func (o *ClusterDatastorePermissions) HasTenantPermissions() bool`

HasTenantPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


