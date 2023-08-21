# ZoneCreateCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "local"]
**Id** | Pointer to **int32** |  | [optional] 

## Methods

### NewZoneCreateCredential

`func NewZoneCreateCredential() *ZoneCreateCredential`

NewZoneCreateCredential instantiates a new ZoneCreateCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneCreateCredentialWithDefaults

`func NewZoneCreateCredentialWithDefaults() *ZoneCreateCredential`

NewZoneCreateCredentialWithDefaults instantiates a new ZoneCreateCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ZoneCreateCredential) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ZoneCreateCredential) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ZoneCreateCredential) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ZoneCreateCredential) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *ZoneCreateCredential) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ZoneCreateCredential) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ZoneCreateCredential) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ZoneCreateCredential) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

