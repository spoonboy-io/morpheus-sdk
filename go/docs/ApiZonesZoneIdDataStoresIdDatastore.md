# ApiZonesZoneIdDataStoresIdDatastore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Activate &#x60;true&#x60; or disable &#x60;false&#x60; the datastore | [optional] 
**Visibility** | Pointer to **string** | Setting &#x60;private&#x60; or &#x60;public&#x60; | [optional] [default to "private"]
**TenantPermissions** | Pointer to [**[]ApiZonesZoneIdDataStoresIdDatastoreTenantPermissions**](ApiZonesZoneIdDataStoresIdDatastoreTenantPermissions.md) |  | [optional] 
**ResourcePermissions** | Pointer to [**ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions**](_api_zones__zoneId__data_stores__id__datastore_resourcePermissions.md) |  | [optional] 

## Methods

### NewApiZonesZoneIdDataStoresIdDatastore

`func NewApiZonesZoneIdDataStoresIdDatastore() *ApiZonesZoneIdDataStoresIdDatastore`

NewApiZonesZoneIdDataStoresIdDatastore instantiates a new ApiZonesZoneIdDataStoresIdDatastore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiZonesZoneIdDataStoresIdDatastoreWithDefaults

`func NewApiZonesZoneIdDataStoresIdDatastoreWithDefaults() *ApiZonesZoneIdDataStoresIdDatastore`

NewApiZonesZoneIdDataStoresIdDatastoreWithDefaults instantiates a new ApiZonesZoneIdDataStoresIdDatastore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *ApiZonesZoneIdDataStoresIdDatastore) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ApiZonesZoneIdDataStoresIdDatastore) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ApiZonesZoneIdDataStoresIdDatastore) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ApiZonesZoneIdDataStoresIdDatastore) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetVisibility

`func (o *ApiZonesZoneIdDataStoresIdDatastore) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ApiZonesZoneIdDataStoresIdDatastore) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ApiZonesZoneIdDataStoresIdDatastore) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ApiZonesZoneIdDataStoresIdDatastore) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetTenantPermissions

`func (o *ApiZonesZoneIdDataStoresIdDatastore) GetTenantPermissions() []ApiZonesZoneIdDataStoresIdDatastoreTenantPermissions`

GetTenantPermissions returns the TenantPermissions field if non-nil, zero value otherwise.

### GetTenantPermissionsOk

`func (o *ApiZonesZoneIdDataStoresIdDatastore) GetTenantPermissionsOk() (*[]ApiZonesZoneIdDataStoresIdDatastoreTenantPermissions, bool)`

GetTenantPermissionsOk returns a tuple with the TenantPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantPermissions

`func (o *ApiZonesZoneIdDataStoresIdDatastore) SetTenantPermissions(v []ApiZonesZoneIdDataStoresIdDatastoreTenantPermissions)`

SetTenantPermissions sets TenantPermissions field to given value.

### HasTenantPermissions

`func (o *ApiZonesZoneIdDataStoresIdDatastore) HasTenantPermissions() bool`

HasTenantPermissions returns a boolean if a field has been set.

### GetResourcePermissions

`func (o *ApiZonesZoneIdDataStoresIdDatastore) GetResourcePermissions() ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions`

GetResourcePermissions returns the ResourcePermissions field if non-nil, zero value otherwise.

### GetResourcePermissionsOk

`func (o *ApiZonesZoneIdDataStoresIdDatastore) GetResourcePermissionsOk() (*ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions, bool)`

GetResourcePermissionsOk returns a tuple with the ResourcePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermissions

`func (o *ApiZonesZoneIdDataStoresIdDatastore) SetResourcePermissions(v ApiZonesZoneIdDataStoresIdDatastoreResourcePermissions)`

SetResourcePermissions sets ResourcePermissions field to given value.

### HasResourcePermissions

`func (o *ApiZonesZoneIdDataStoresIdDatastore) HasResourcePermissions() bool`

HasResourcePermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


