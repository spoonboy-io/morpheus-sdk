# InlineObject66

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ttl** | Pointer to [**OneOfstringlong**](oneOf&lt;string,long&gt;.md) | Time to Live. The lease duration in seconds, or a human readable format eg. 15m, 8h, 7d. The default is 0 meaning Never expires. This only is applied if the cypher does not yet exist and is created.  | [optional] 
**Value** | Pointer to **string** | The secret value to be stored. This is only parsed if type is passed as &#x60;string&#x60;. | [optional] 

## Methods

### NewInlineObject66

`func NewInlineObject66() *InlineObject66`

NewInlineObject66 instantiates a new InlineObject66 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject66WithDefaults

`func NewInlineObject66WithDefaults() *InlineObject66`

NewInlineObject66WithDefaults instantiates a new InlineObject66 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTtl

`func (o *InlineObject66) GetTtl() OneOfstringlong`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *InlineObject66) GetTtlOk() (*OneOfstringlong, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *InlineObject66) SetTtl(v OneOfstringlong)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *InlineObject66) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetValue

`func (o *InlineObject66) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *InlineObject66) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *InlineObject66) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *InlineObject66) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


