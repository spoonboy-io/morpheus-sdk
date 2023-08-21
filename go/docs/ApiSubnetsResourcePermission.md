# ApiSubnetsResourcePermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**All** | Pointer to **bool** | Pass true to allow access all groups | [optional] 
**Sites** | Pointer to [**[]ApiBlueprintsIdUpdatePermissionsResourcePermissionSites**](ApiBlueprintsIdUpdatePermissionsResourcePermissionSites.md) | Array of groups ID objects that are allowed access | [optional] 
**AllPlans** | Pointer to **bool** |  | [optional] 
**Plans** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewApiSubnetsResourcePermission

`func NewApiSubnetsResourcePermission() *ApiSubnetsResourcePermission`

NewApiSubnetsResourcePermission instantiates a new ApiSubnetsResourcePermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiSubnetsResourcePermissionWithDefaults

`func NewApiSubnetsResourcePermissionWithDefaults() *ApiSubnetsResourcePermission`

NewApiSubnetsResourcePermissionWithDefaults instantiates a new ApiSubnetsResourcePermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAll

`func (o *ApiSubnetsResourcePermission) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *ApiSubnetsResourcePermission) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *ApiSubnetsResourcePermission) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *ApiSubnetsResourcePermission) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetSites

`func (o *ApiSubnetsResourcePermission) GetSites() []ApiBlueprintsIdUpdatePermissionsResourcePermissionSites`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *ApiSubnetsResourcePermission) GetSitesOk() (*[]ApiBlueprintsIdUpdatePermissionsResourcePermissionSites, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *ApiSubnetsResourcePermission) SetSites(v []ApiBlueprintsIdUpdatePermissionsResourcePermissionSites)`

SetSites sets Sites field to given value.

### HasSites

`func (o *ApiSubnetsResourcePermission) HasSites() bool`

HasSites returns a boolean if a field has been set.

### GetAllPlans

`func (o *ApiSubnetsResourcePermission) GetAllPlans() bool`

GetAllPlans returns the AllPlans field if non-nil, zero value otherwise.

### GetAllPlansOk

`func (o *ApiSubnetsResourcePermission) GetAllPlansOk() (*bool, bool)`

GetAllPlansOk returns a tuple with the AllPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllPlans

`func (o *ApiSubnetsResourcePermission) SetAllPlans(v bool)`

SetAllPlans sets AllPlans field to given value.

### HasAllPlans

`func (o *ApiSubnetsResourcePermission) HasAllPlans() bool`

HasAllPlans returns a boolean if a field has been set.

### GetPlans

`func (o *ApiSubnetsResourcePermission) GetPlans() []map[string]interface{}`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *ApiSubnetsResourcePermission) GetPlansOk() (*[]map[string]interface{}, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *ApiSubnetsResourcePermission) SetPlans(v []map[string]interface{})`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *ApiSubnetsResourcePermission) HasPlans() bool`

HasPlans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


