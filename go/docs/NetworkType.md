# NetworkType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Network Type ID | [optional] 
**Name** | Pointer to **string** | Network Name | [optional] 
**Code** | Pointer to **string** | Network Code | [optional] 

## Methods

### NewNetworkType

`func NewNetworkType() *NetworkType`

NewNetworkType instantiates a new NetworkType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkTypeWithDefaults

`func NewNetworkTypeWithDefaults() *NetworkType`

NewNetworkTypeWithDefaults instantiates a new NetworkType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NetworkType) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkType) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkType) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *NetworkType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *NetworkType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *NetworkType) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *NetworkType) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *NetworkType) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *NetworkType) HasCode() bool`

HasCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


