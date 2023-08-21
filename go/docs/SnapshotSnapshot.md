# SnapshotSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**State** | Pointer to **NullableString** |  | [optional] 
**SnapshotType** | Pointer to **string** |  | [optional] 
**SnapshotCreated** | Pointer to **time.Time** |  | [optional] 
**Zone** | Pointer to [**SnapshotSnapshotZone**](snapshot_snapshot_zone.md) |  | [optional] 
**Datastore** | Pointer to **NullableString** |  | [optional] 
**ParentSnapshot** | Pointer to **NullableString** |  | [optional] 
**CurrentlyActive** | Pointer to **bool** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewSnapshotSnapshot

`func NewSnapshotSnapshot() *SnapshotSnapshot`

NewSnapshotSnapshot instantiates a new SnapshotSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotSnapshotWithDefaults

`func NewSnapshotSnapshotWithDefaults() *SnapshotSnapshot`

NewSnapshotSnapshotWithDefaults instantiates a new SnapshotSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SnapshotSnapshot) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SnapshotSnapshot) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SnapshotSnapshot) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SnapshotSnapshot) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SnapshotSnapshot) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SnapshotSnapshot) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SnapshotSnapshot) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SnapshotSnapshot) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *SnapshotSnapshot) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SnapshotSnapshot) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SnapshotSnapshot) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SnapshotSnapshot) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SnapshotSnapshot) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SnapshotSnapshot) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetExternalId

`func (o *SnapshotSnapshot) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *SnapshotSnapshot) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *SnapshotSnapshot) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *SnapshotSnapshot) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetStatus

`func (o *SnapshotSnapshot) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SnapshotSnapshot) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SnapshotSnapshot) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SnapshotSnapshot) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetState

`func (o *SnapshotSnapshot) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SnapshotSnapshot) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SnapshotSnapshot) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *SnapshotSnapshot) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *SnapshotSnapshot) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *SnapshotSnapshot) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetSnapshotType

`func (o *SnapshotSnapshot) GetSnapshotType() string`

GetSnapshotType returns the SnapshotType field if non-nil, zero value otherwise.

### GetSnapshotTypeOk

`func (o *SnapshotSnapshot) GetSnapshotTypeOk() (*string, bool)`

GetSnapshotTypeOk returns a tuple with the SnapshotType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotType

`func (o *SnapshotSnapshot) SetSnapshotType(v string)`

SetSnapshotType sets SnapshotType field to given value.

### HasSnapshotType

`func (o *SnapshotSnapshot) HasSnapshotType() bool`

HasSnapshotType returns a boolean if a field has been set.

### GetSnapshotCreated

`func (o *SnapshotSnapshot) GetSnapshotCreated() time.Time`

GetSnapshotCreated returns the SnapshotCreated field if non-nil, zero value otherwise.

### GetSnapshotCreatedOk

`func (o *SnapshotSnapshot) GetSnapshotCreatedOk() (*time.Time, bool)`

GetSnapshotCreatedOk returns a tuple with the SnapshotCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotCreated

`func (o *SnapshotSnapshot) SetSnapshotCreated(v time.Time)`

SetSnapshotCreated sets SnapshotCreated field to given value.

### HasSnapshotCreated

`func (o *SnapshotSnapshot) HasSnapshotCreated() bool`

HasSnapshotCreated returns a boolean if a field has been set.

### GetZone

`func (o *SnapshotSnapshot) GetZone() SnapshotSnapshotZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *SnapshotSnapshot) GetZoneOk() (*SnapshotSnapshotZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *SnapshotSnapshot) SetZone(v SnapshotSnapshotZone)`

SetZone sets Zone field to given value.

### HasZone

`func (o *SnapshotSnapshot) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetDatastore

`func (o *SnapshotSnapshot) GetDatastore() string`

GetDatastore returns the Datastore field if non-nil, zero value otherwise.

### GetDatastoreOk

`func (o *SnapshotSnapshot) GetDatastoreOk() (*string, bool)`

GetDatastoreOk returns a tuple with the Datastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastore

`func (o *SnapshotSnapshot) SetDatastore(v string)`

SetDatastore sets Datastore field to given value.

### HasDatastore

`func (o *SnapshotSnapshot) HasDatastore() bool`

HasDatastore returns a boolean if a field has been set.

### SetDatastoreNil

`func (o *SnapshotSnapshot) SetDatastoreNil(b bool)`

 SetDatastoreNil sets the value for Datastore to be an explicit nil

### UnsetDatastore
`func (o *SnapshotSnapshot) UnsetDatastore()`

UnsetDatastore ensures that no value is present for Datastore, not even an explicit nil
### GetParentSnapshot

`func (o *SnapshotSnapshot) GetParentSnapshot() string`

GetParentSnapshot returns the ParentSnapshot field if non-nil, zero value otherwise.

### GetParentSnapshotOk

`func (o *SnapshotSnapshot) GetParentSnapshotOk() (*string, bool)`

GetParentSnapshotOk returns a tuple with the ParentSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSnapshot

`func (o *SnapshotSnapshot) SetParentSnapshot(v string)`

SetParentSnapshot sets ParentSnapshot field to given value.

### HasParentSnapshot

`func (o *SnapshotSnapshot) HasParentSnapshot() bool`

HasParentSnapshot returns a boolean if a field has been set.

### SetParentSnapshotNil

`func (o *SnapshotSnapshot) SetParentSnapshotNil(b bool)`

 SetParentSnapshotNil sets the value for ParentSnapshot to be an explicit nil

### UnsetParentSnapshot
`func (o *SnapshotSnapshot) UnsetParentSnapshot()`

UnsetParentSnapshot ensures that no value is present for ParentSnapshot, not even an explicit nil
### GetCurrentlyActive

`func (o *SnapshotSnapshot) GetCurrentlyActive() bool`

GetCurrentlyActive returns the CurrentlyActive field if non-nil, zero value otherwise.

### GetCurrentlyActiveOk

`func (o *SnapshotSnapshot) GetCurrentlyActiveOk() (*bool, bool)`

GetCurrentlyActiveOk returns a tuple with the CurrentlyActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentlyActive

`func (o *SnapshotSnapshot) SetCurrentlyActive(v bool)`

SetCurrentlyActive sets CurrentlyActive field to given value.

### HasCurrentlyActive

`func (o *SnapshotSnapshot) HasCurrentlyActive() bool`

HasCurrentlyActive returns a boolean if a field has been set.

### GetDateCreated

`func (o *SnapshotSnapshot) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *SnapshotSnapshot) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *SnapshotSnapshot) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *SnapshotSnapshot) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


