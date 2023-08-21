# SearchHits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID | [optional] 
**Uuid** | Pointer to **string** | UUID | [optional] 
**Name** | Pointer to **string** | Name | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **NullableTime** |  | [optional] 
**Score** | Pointer to **NullableFloat32** |  | [optional] 

## Methods

### NewSearchHits

`func NewSearchHits() *SearchHits`

NewSearchHits instantiates a new SearchHits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchHitsWithDefaults

`func NewSearchHitsWithDefaults() *SearchHits`

NewSearchHitsWithDefaults instantiates a new SearchHits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SearchHits) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SearchHits) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SearchHits) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SearchHits) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *SearchHits) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *SearchHits) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *SearchHits) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *SearchHits) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *SearchHits) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SearchHits) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SearchHits) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SearchHits) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *SearchHits) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SearchHits) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SearchHits) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SearchHits) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SearchHits) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SearchHits) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetType

`func (o *SearchHits) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SearchHits) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SearchHits) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SearchHits) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDateCreated

`func (o *SearchHits) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *SearchHits) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *SearchHits) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *SearchHits) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *SearchHits) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *SearchHits) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetScore

`func (o *SearchHits) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *SearchHits) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *SearchHits) SetScore(v float32)`

SetScore sets Score field to given value.

### HasScore

`func (o *SearchHits) HasScore() bool`

HasScore returns a boolean if a field has been set.

### SetScoreNil

`func (o *SearchHits) SetScoreNil(b bool)`

 SetScoreNil sets the value for Score to be an explicit nil

### UnsetScore
`func (o *SearchHits) UnsetScore()`

UnsetScore ensures that no value is present for Score, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


