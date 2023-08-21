# InlineResponse200131ResourcePermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**All** | Pointer to **bool** | Pass &#x60;true&#x60; to allow access all groups | [optional] [default to true]
**Sites** | Pointer to [**[]InlineResponse200131ResourcePermissionSites**](InlineResponse200131ResourcePermissionSites.md) | Array of groups that are allowed access | [optional] 

## Methods

### NewInlineResponse200131ResourcePermission

`func NewInlineResponse200131ResourcePermission() *InlineResponse200131ResourcePermission`

NewInlineResponse200131ResourcePermission instantiates a new InlineResponse200131ResourcePermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200131ResourcePermissionWithDefaults

`func NewInlineResponse200131ResourcePermissionWithDefaults() *InlineResponse200131ResourcePermission`

NewInlineResponse200131ResourcePermissionWithDefaults instantiates a new InlineResponse200131ResourcePermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAll

`func (o *InlineResponse200131ResourcePermission) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *InlineResponse200131ResourcePermission) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *InlineResponse200131ResourcePermission) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *InlineResponse200131ResourcePermission) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetSites

`func (o *InlineResponse200131ResourcePermission) GetSites() []InlineResponse200131ResourcePermissionSites`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *InlineResponse200131ResourcePermission) GetSitesOk() (*[]InlineResponse200131ResourcePermissionSites, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *InlineResponse200131ResourcePermission) SetSites(v []InlineResponse200131ResourcePermissionSites)`

SetSites sets Sites field to given value.

### HasSites

`func (o *InlineResponse200131ResourcePermission) HasSites() bool`

HasSites returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


