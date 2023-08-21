# ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**All** | Pointer to **bool** | Pass &#x60;true&#x60; to allow access all groups | [optional] [default to true]
**Sites** | Pointer to [**[]ApiZonesZoneIdDataStoresIdDatastoreResourcePermissionsSites**](ApiZonesZoneIdDataStoresIdDatastoreResourcePermissionsSites.md) | Array of groups that are allowed access | [optional] 
**AllPlans** | Pointer to **bool** | Pass true to allow access all plans | [optional] [default to true]
**Plans** | Pointer to [**[]ApiZonesZoneIdDataStoresIdDatastoreResourcePermissionsSites**](ApiZonesZoneIdDataStoresIdDatastoreResourcePermissionsSites.md) | Array of plans that are allowed access | [optional] 

## Methods

### NewApiZonesZoneIdDataStoresIdDatastoreResourcePermissions

`func NewApiZonesZoneIdDataStoresIdDatastoreResourcePermissions() *ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions`

NewApiZonesZoneIdDataStoresIdDatastoreResourcePermissions instantiates a new ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiZonesZoneIdDataStoresIdDatastoreResourcePermissionsWithDefaults

`func NewApiZonesZoneIdDataStoresIdDatastoreResourcePermissionsWithDefaults() *ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions`

NewApiZonesZoneIdDataStoresIdDatastoreResourcePermissionsWithDefaults instantiates a new ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAll

`func (o *ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetSites

`func (o *ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions) GetSites() []ApiZonesZoneIdDataStoresIdDatastoreResourcePermissionsSites`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions) GetSitesOk() (*[]ApiZonesZoneIdDataStoresIdDatastoreResourcePermissionsSites, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions) SetSites(v []ApiZonesZoneIdDataStoresIdDatastoreResourcePermissionsSites)`

SetSites sets Sites field to given value.

### HasSites

`func (o *ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions) HasSites() bool`

HasSites returns a boolean if a field has been set.

### GetAllPlans

`func (o *ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions) GetAllPlans() bool`

GetAllPlans returns the AllPlans field if non-nil, zero value otherwise.

### GetAllPlansOk

`func (o *ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions) GetAllPlansOk() (*bool, bool)`

GetAllPlansOk returns a tuple with the AllPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllPlans

`func (o *ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions) SetAllPlans(v bool)`

SetAllPlans sets AllPlans field to given value.

### HasAllPlans

`func (o *ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions) HasAllPlans() bool`

HasAllPlans returns a boolean if a field has been set.

### GetPlans

`func (o *ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions) GetPlans() []ApiZonesZoneIdDataStoresIdDatastoreResourcePermissionsSites`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions) GetPlansOk() (*[]ApiZonesZoneIdDataStoresIdDatastoreResourcePermissionsSites, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions) SetPlans(v []ApiZonesZoneIdDataStoresIdDatastoreResourcePermissionsSites)`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions) HasPlans() bool`

HasPlans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


