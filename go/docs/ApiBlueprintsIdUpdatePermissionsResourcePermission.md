# ApiBlueprintsIdUpdatePermissionsResourcePermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**All** | Pointer to **bool** | Set to true to grant access to all groups | [optional] 
**Sites** | Pointer to [**[]ApiBlueprintsIdUpdatePermissionsResourcePermissionSites**](ApiBlueprintsIdUpdatePermissionsResourcePermissionSites.md) | Array of objects identifying groups with access | [optional] 
**OwnerId** | Pointer to **int64** | User ID, can be used to change blueprint owner. | [optional] 

## Methods

### NewApiBlueprintsIdUpdatePermissionsResourcePermission

`func NewApiBlueprintsIdUpdatePermissionsResourcePermission() *ApiBlueprintsIdUpdatePermissionsResourcePermission`

NewApiBlueprintsIdUpdatePermissionsResourcePermission instantiates a new ApiBlueprintsIdUpdatePermissionsResourcePermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiBlueprintsIdUpdatePermissionsResourcePermissionWithDefaults

`func NewApiBlueprintsIdUpdatePermissionsResourcePermissionWithDefaults() *ApiBlueprintsIdUpdatePermissionsResourcePermission`

NewApiBlueprintsIdUpdatePermissionsResourcePermissionWithDefaults instantiates a new ApiBlueprintsIdUpdatePermissionsResourcePermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAll

`func (o *ApiBlueprintsIdUpdatePermissionsResourcePermission) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *ApiBlueprintsIdUpdatePermissionsResourcePermission) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *ApiBlueprintsIdUpdatePermissionsResourcePermission) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *ApiBlueprintsIdUpdatePermissionsResourcePermission) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetSites

`func (o *ApiBlueprintsIdUpdatePermissionsResourcePermission) GetSites() []ApiBlueprintsIdUpdatePermissionsResourcePermissionSites`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *ApiBlueprintsIdUpdatePermissionsResourcePermission) GetSitesOk() (*[]ApiBlueprintsIdUpdatePermissionsResourcePermissionSites, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *ApiBlueprintsIdUpdatePermissionsResourcePermission) SetSites(v []ApiBlueprintsIdUpdatePermissionsResourcePermissionSites)`

SetSites sets Sites field to given value.

### HasSites

`func (o *ApiBlueprintsIdUpdatePermissionsResourcePermission) HasSites() bool`

HasSites returns a boolean if a field has been set.

### GetOwnerId

`func (o *ApiBlueprintsIdUpdatePermissionsResourcePermission) GetOwnerId() int64`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *ApiBlueprintsIdUpdatePermissionsResourcePermission) GetOwnerIdOk() (*int64, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *ApiBlueprintsIdUpdatePermissionsResourcePermission) SetOwnerId(v int64)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *ApiBlueprintsIdUpdatePermissionsResourcePermission) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


