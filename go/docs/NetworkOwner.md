# NetworkOwner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Owner Tenant ID | [optional] 
**Name** | Pointer to **string** | Owner Tenant Name | [optional] 

## Methods

### NewNetworkOwner

`func NewNetworkOwner() *NetworkOwner`

NewNetworkOwner instantiates a new NetworkOwner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkOwnerWithDefaults

`func NewNetworkOwnerWithDefaults() *NetworkOwner`

NewNetworkOwnerWithDefaults instantiates a new NetworkOwner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NetworkOwner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkOwner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkOwner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *NetworkOwner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *NetworkOwner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkOwner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkOwner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkOwner) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


