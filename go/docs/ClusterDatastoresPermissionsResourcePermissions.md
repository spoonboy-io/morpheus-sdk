# ClusterDatastoresPermissionsResourcePermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultStore** | Pointer to **bool** |  | [optional] 
**AllPlans** | Pointer to **bool** |  | [optional] 
**DefaultTarget** | Pointer to **bool** |  | [optional] 
**CanManage** | Pointer to **bool** |  | [optional] 
**All** | Pointer to **bool** |  | [optional] 
**Account** | Pointer to [**ApiBlueprintsIdUpdatePermissionsResourcePermissionSites**](_api_blueprints__id__update_permissions_resourcePermission_sites.md) |  | [optional] 
**Sites** | Pointer to [**[]ResourcePermissionsSites**](ResourcePermissionsSites.md) |  | [optional] 
**Plans** | Pointer to [**[]ResourcePermissionsSites**](ResourcePermissionsSites.md) |  | [optional] 

## Methods

### NewClusterDatastoresPermissionsResourcePermissions

`func NewClusterDatastoresPermissionsResourcePermissions() *ClusterDatastoresPermissionsResourcePermissions`

NewClusterDatastoresPermissionsResourcePermissions instantiates a new ClusterDatastoresPermissionsResourcePermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterDatastoresPermissionsResourcePermissionsWithDefaults

`func NewClusterDatastoresPermissionsResourcePermissionsWithDefaults() *ClusterDatastoresPermissionsResourcePermissions`

NewClusterDatastoresPermissionsResourcePermissionsWithDefaults instantiates a new ClusterDatastoresPermissionsResourcePermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultStore

`func (o *ClusterDatastoresPermissionsResourcePermissions) GetDefaultStore() bool`

GetDefaultStore returns the DefaultStore field if non-nil, zero value otherwise.

### GetDefaultStoreOk

`func (o *ClusterDatastoresPermissionsResourcePermissions) GetDefaultStoreOk() (*bool, bool)`

GetDefaultStoreOk returns a tuple with the DefaultStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultStore

`func (o *ClusterDatastoresPermissionsResourcePermissions) SetDefaultStore(v bool)`

SetDefaultStore sets DefaultStore field to given value.

### HasDefaultStore

`func (o *ClusterDatastoresPermissionsResourcePermissions) HasDefaultStore() bool`

HasDefaultStore returns a boolean if a field has been set.

### GetAllPlans

`func (o *ClusterDatastoresPermissionsResourcePermissions) GetAllPlans() bool`

GetAllPlans returns the AllPlans field if non-nil, zero value otherwise.

### GetAllPlansOk

`func (o *ClusterDatastoresPermissionsResourcePermissions) GetAllPlansOk() (*bool, bool)`

GetAllPlansOk returns a tuple with the AllPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllPlans

`func (o *ClusterDatastoresPermissionsResourcePermissions) SetAllPlans(v bool)`

SetAllPlans sets AllPlans field to given value.

### HasAllPlans

`func (o *ClusterDatastoresPermissionsResourcePermissions) HasAllPlans() bool`

HasAllPlans returns a boolean if a field has been set.

### GetDefaultTarget

`func (o *ClusterDatastoresPermissionsResourcePermissions) GetDefaultTarget() bool`

GetDefaultTarget returns the DefaultTarget field if non-nil, zero value otherwise.

### GetDefaultTargetOk

`func (o *ClusterDatastoresPermissionsResourcePermissions) GetDefaultTargetOk() (*bool, bool)`

GetDefaultTargetOk returns a tuple with the DefaultTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTarget

`func (o *ClusterDatastoresPermissionsResourcePermissions) SetDefaultTarget(v bool)`

SetDefaultTarget sets DefaultTarget field to given value.

### HasDefaultTarget

`func (o *ClusterDatastoresPermissionsResourcePermissions) HasDefaultTarget() bool`

HasDefaultTarget returns a boolean if a field has been set.

### GetCanManage

`func (o *ClusterDatastoresPermissionsResourcePermissions) GetCanManage() bool`

GetCanManage returns the CanManage field if non-nil, zero value otherwise.

### GetCanManageOk

`func (o *ClusterDatastoresPermissionsResourcePermissions) GetCanManageOk() (*bool, bool)`

GetCanManageOk returns a tuple with the CanManage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManage

`func (o *ClusterDatastoresPermissionsResourcePermissions) SetCanManage(v bool)`

SetCanManage sets CanManage field to given value.

### HasCanManage

`func (o *ClusterDatastoresPermissionsResourcePermissions) HasCanManage() bool`

HasCanManage returns a boolean if a field has been set.

### GetAll

`func (o *ClusterDatastoresPermissionsResourcePermissions) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *ClusterDatastoresPermissionsResourcePermissions) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *ClusterDatastoresPermissionsResourcePermissions) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *ClusterDatastoresPermissionsResourcePermissions) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetAccount

`func (o *ClusterDatastoresPermissionsResourcePermissions) GetAccount() ApiBlueprintsIdUpdatePermissionsResourcePermissionSites`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ClusterDatastoresPermissionsResourcePermissions) GetAccountOk() (*ApiBlueprintsIdUpdatePermissionsResourcePermissionSites, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ClusterDatastoresPermissionsResourcePermissions) SetAccount(v ApiBlueprintsIdUpdatePermissionsResourcePermissionSites)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ClusterDatastoresPermissionsResourcePermissions) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetSites

`func (o *ClusterDatastoresPermissionsResourcePermissions) GetSites() []ResourcePermissionsSites`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *ClusterDatastoresPermissionsResourcePermissions) GetSitesOk() (*[]ResourcePermissionsSites, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *ClusterDatastoresPermissionsResourcePermissions) SetSites(v []ResourcePermissionsSites)`

SetSites sets Sites field to given value.

### HasSites

`func (o *ClusterDatastoresPermissionsResourcePermissions) HasSites() bool`

HasSites returns a boolean if a field has been set.

### SetSitesNil

`func (o *ClusterDatastoresPermissionsResourcePermissions) SetSitesNil(b bool)`

 SetSitesNil sets the value for Sites to be an explicit nil

### UnsetSites
`func (o *ClusterDatastoresPermissionsResourcePermissions) UnsetSites()`

UnsetSites ensures that no value is present for Sites, not even an explicit nil
### GetPlans

`func (o *ClusterDatastoresPermissionsResourcePermissions) GetPlans() []ResourcePermissionsSites`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *ClusterDatastoresPermissionsResourcePermissions) GetPlansOk() (*[]ResourcePermissionsSites, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *ClusterDatastoresPermissionsResourcePermissions) SetPlans(v []ResourcePermissionsSites)`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *ClusterDatastoresPermissionsResourcePermissions) HasPlans() bool`

HasPlans returns a boolean if a field has been set.

### SetPlansNil

`func (o *ClusterDatastoresPermissionsResourcePermissions) SetPlansNil(b bool)`

 SetPlansNil sets the value for Plans to be an explicit nil

### UnsetPlans
`func (o *ClusterDatastoresPermissionsResourcePermissions) UnsetPlans()`

UnsetPlans ensures that no value is present for Plans, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


