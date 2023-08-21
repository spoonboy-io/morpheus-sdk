# AppStateWorkloads

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RefType** | Pointer to **string** |  | [optional] 
**RefId** | Pointer to **int64** |  | [optional] 
**RefName** | Pointer to **string** |  | [optional] 
**SubRefName** | Pointer to **NullableString** |  | [optional] 
**StateDate** | Pointer to **time.Time** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**IacDrift** | Pointer to **bool** |  | [optional] 

## Methods

### NewAppStateWorkloads

`func NewAppStateWorkloads() *AppStateWorkloads`

NewAppStateWorkloads instantiates a new AppStateWorkloads object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppStateWorkloadsWithDefaults

`func NewAppStateWorkloadsWithDefaults() *AppStateWorkloads`

NewAppStateWorkloadsWithDefaults instantiates a new AppStateWorkloads object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRefType

`func (o *AppStateWorkloads) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *AppStateWorkloads) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *AppStateWorkloads) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *AppStateWorkloads) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *AppStateWorkloads) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *AppStateWorkloads) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *AppStateWorkloads) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *AppStateWorkloads) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetRefName

`func (o *AppStateWorkloads) GetRefName() string`

GetRefName returns the RefName field if non-nil, zero value otherwise.

### GetRefNameOk

`func (o *AppStateWorkloads) GetRefNameOk() (*string, bool)`

GetRefNameOk returns a tuple with the RefName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefName

`func (o *AppStateWorkloads) SetRefName(v string)`

SetRefName sets RefName field to given value.

### HasRefName

`func (o *AppStateWorkloads) HasRefName() bool`

HasRefName returns a boolean if a field has been set.

### GetSubRefName

`func (o *AppStateWorkloads) GetSubRefName() string`

GetSubRefName returns the SubRefName field if non-nil, zero value otherwise.

### GetSubRefNameOk

`func (o *AppStateWorkloads) GetSubRefNameOk() (*string, bool)`

GetSubRefNameOk returns a tuple with the SubRefName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubRefName

`func (o *AppStateWorkloads) SetSubRefName(v string)`

SetSubRefName sets SubRefName field to given value.

### HasSubRefName

`func (o *AppStateWorkloads) HasSubRefName() bool`

HasSubRefName returns a boolean if a field has been set.

### SetSubRefNameNil

`func (o *AppStateWorkloads) SetSubRefNameNil(b bool)`

 SetSubRefNameNil sets the value for SubRefName to be an explicit nil

### UnsetSubRefName
`func (o *AppStateWorkloads) UnsetSubRefName()`

UnsetSubRefName ensures that no value is present for SubRefName, not even an explicit nil
### GetStateDate

`func (o *AppStateWorkloads) GetStateDate() time.Time`

GetStateDate returns the StateDate field if non-nil, zero value otherwise.

### GetStateDateOk

`func (o *AppStateWorkloads) GetStateDateOk() (*time.Time, bool)`

GetStateDateOk returns a tuple with the StateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateDate

`func (o *AppStateWorkloads) SetStateDate(v time.Time)`

SetStateDate sets StateDate field to given value.

### HasStateDate

`func (o *AppStateWorkloads) HasStateDate() bool`

HasStateDate returns a boolean if a field has been set.

### GetStatus

`func (o *AppStateWorkloads) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AppStateWorkloads) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AppStateWorkloads) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AppStateWorkloads) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetIacDrift

`func (o *AppStateWorkloads) GetIacDrift() bool`

GetIacDrift returns the IacDrift field if non-nil, zero value otherwise.

### GetIacDriftOk

`func (o *AppStateWorkloads) GetIacDriftOk() (*bool, bool)`

GetIacDriftOk returns a tuple with the IacDrift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIacDrift

`func (o *AppStateWorkloads) SetIacDrift(v bool)`

SetIacDrift sets IacDrift field to given value.

### HasIacDrift

`func (o *AppStateWorkloads) HasIacDrift() bool`

HasIacDrift returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


