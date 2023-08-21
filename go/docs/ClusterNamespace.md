# ClusterNamespace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**Permissions** | Pointer to [**ClusterNamespacePermissions**](clusterNamespace_permissions.md) |  | [optional] 

## Methods

### NewClusterNamespace

`func NewClusterNamespace() *ClusterNamespace`

NewClusterNamespace instantiates a new ClusterNamespace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterNamespaceWithDefaults

`func NewClusterNamespaceWithDefaults() *ClusterNamespace`

NewClusterNamespaceWithDefaults instantiates a new ClusterNamespace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClusterNamespace) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClusterNamespace) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClusterNamespace) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ClusterNamespace) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVisibility

`func (o *ClusterNamespace) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ClusterNamespace) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ClusterNamespace) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ClusterNamespace) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetName

`func (o *ClusterNamespace) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterNamespace) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterNamespace) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClusterNamespace) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ClusterNamespace) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ClusterNamespace) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ClusterNamespace) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ClusterNamespace) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *ClusterNamespace) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClusterNamespace) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClusterNamespace) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ClusterNamespace) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetExternalId

`func (o *ClusterNamespace) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ClusterNamespace) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ClusterNamespace) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ClusterNamespace) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetPermissions

`func (o *ClusterNamespace) GetPermissions() ClusterNamespacePermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *ClusterNamespace) GetPermissionsOk() (*ClusterNamespacePermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *ClusterNamespace) SetPermissions(v ClusterNamespacePermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *ClusterNamespace) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


