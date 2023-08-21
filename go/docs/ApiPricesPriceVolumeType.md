# ApiPricesPriceVolumeType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Volume type ID, required for &#x60;storage&#x60; price type. The endpoint /api/prices/volume-types provides a list of available volume type options.  | [optional] 

## Methods

### NewApiPricesPriceVolumeType

`func NewApiPricesPriceVolumeType() *ApiPricesPriceVolumeType`

NewApiPricesPriceVolumeType instantiates a new ApiPricesPriceVolumeType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiPricesPriceVolumeTypeWithDefaults

`func NewApiPricesPriceVolumeTypeWithDefaults() *ApiPricesPriceVolumeType`

NewApiPricesPriceVolumeTypeWithDefaults instantiates a new ApiPricesPriceVolumeType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApiPricesPriceVolumeType) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiPricesPriceVolumeType) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiPricesPriceVolumeType) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ApiPricesPriceVolumeType) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


