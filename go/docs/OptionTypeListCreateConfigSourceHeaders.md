# OptionTypeListCreateConfigSourceHeaders

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Header name | 
**Value** | Pointer to **string** | Header value | [optional] 
**Masked** | Pointer to **bool** | Can be used to enable / disable masking of value | [optional] [default to false]

## Methods

### NewOptionTypeListCreateConfigSourceHeaders

`func NewOptionTypeListCreateConfigSourceHeaders(name string, ) *OptionTypeListCreateConfigSourceHeaders`

NewOptionTypeListCreateConfigSourceHeaders instantiates a new OptionTypeListCreateConfigSourceHeaders object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptionTypeListCreateConfigSourceHeadersWithDefaults

`func NewOptionTypeListCreateConfigSourceHeadersWithDefaults() *OptionTypeListCreateConfigSourceHeaders`

NewOptionTypeListCreateConfigSourceHeadersWithDefaults instantiates a new OptionTypeListCreateConfigSourceHeaders object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OptionTypeListCreateConfigSourceHeaders) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OptionTypeListCreateConfigSourceHeaders) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OptionTypeListCreateConfigSourceHeaders) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *OptionTypeListCreateConfigSourceHeaders) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *OptionTypeListCreateConfigSourceHeaders) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *OptionTypeListCreateConfigSourceHeaders) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *OptionTypeListCreateConfigSourceHeaders) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetMasked

`func (o *OptionTypeListCreateConfigSourceHeaders) GetMasked() bool`

GetMasked returns the Masked field if non-nil, zero value otherwise.

### GetMaskedOk

`func (o *OptionTypeListCreateConfigSourceHeaders) GetMaskedOk() (*bool, bool)`

GetMaskedOk returns a tuple with the Masked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasked

`func (o *OptionTypeListCreateConfigSourceHeaders) SetMasked(v bool)`

SetMasked sets Masked field to given value.

### HasMasked

`func (o *OptionTypeListCreateConfigSourceHeaders) HasMasked() bool`

HasMasked returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


