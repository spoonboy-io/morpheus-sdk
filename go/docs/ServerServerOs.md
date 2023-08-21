# ServerServerOs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Vendor** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**OsFamily** | Pointer to **NullableString** |  | [optional] 
**OsVersion** | Pointer to **string** |  | [optional] 
**BitCount** | Pointer to **int64** |  | [optional] 
**Platform** | Pointer to **string** |  | [optional] 

## Methods

### NewServerServerOs

`func NewServerServerOs() *ServerServerOs`

NewServerServerOs instantiates a new ServerServerOs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerServerOsWithDefaults

`func NewServerServerOsWithDefaults() *ServerServerOs`

NewServerServerOsWithDefaults instantiates a new ServerServerOs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServerServerOs) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServerServerOs) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServerServerOs) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ServerServerOs) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCode

`func (o *ServerServerOs) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ServerServerOs) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ServerServerOs) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ServerServerOs) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *ServerServerOs) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServerServerOs) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServerServerOs) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServerServerOs) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ServerServerOs) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServerServerOs) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServerServerOs) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServerServerOs) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ServerServerOs) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ServerServerOs) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetVendor

`func (o *ServerServerOs) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *ServerServerOs) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *ServerServerOs) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *ServerServerOs) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetCategory

`func (o *ServerServerOs) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ServerServerOs) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ServerServerOs) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ServerServerOs) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetOsFamily

`func (o *ServerServerOs) GetOsFamily() string`

GetOsFamily returns the OsFamily field if non-nil, zero value otherwise.

### GetOsFamilyOk

`func (o *ServerServerOs) GetOsFamilyOk() (*string, bool)`

GetOsFamilyOk returns a tuple with the OsFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsFamily

`func (o *ServerServerOs) SetOsFamily(v string)`

SetOsFamily sets OsFamily field to given value.

### HasOsFamily

`func (o *ServerServerOs) HasOsFamily() bool`

HasOsFamily returns a boolean if a field has been set.

### SetOsFamilyNil

`func (o *ServerServerOs) SetOsFamilyNil(b bool)`

 SetOsFamilyNil sets the value for OsFamily to be an explicit nil

### UnsetOsFamily
`func (o *ServerServerOs) UnsetOsFamily()`

UnsetOsFamily ensures that no value is present for OsFamily, not even an explicit nil
### GetOsVersion

`func (o *ServerServerOs) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *ServerServerOs) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *ServerServerOs) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *ServerServerOs) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetBitCount

`func (o *ServerServerOs) GetBitCount() int64`

GetBitCount returns the BitCount field if non-nil, zero value otherwise.

### GetBitCountOk

`func (o *ServerServerOs) GetBitCountOk() (*int64, bool)`

GetBitCountOk returns a tuple with the BitCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitCount

`func (o *ServerServerOs) SetBitCount(v int64)`

SetBitCount sets BitCount field to given value.

### HasBitCount

`func (o *ServerServerOs) HasBitCount() bool`

HasBitCount returns a boolean if a field has been set.

### GetPlatform

`func (o *ServerServerOs) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ServerServerOs) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ServerServerOs) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *ServerServerOs) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


