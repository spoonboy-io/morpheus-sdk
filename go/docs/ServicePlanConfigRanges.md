# ServicePlanConfigRanges

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinStorage** | Pointer to **NullableString** |  | [optional] 
**MaxStorage** | Pointer to **NullableString** |  | [optional] 
**MinPerDiskSize** | Pointer to **NullableString** |  | [optional] 
**MaxPerDiskSize** | Pointer to **NullableString** |  | [optional] 
**MinMemory** | Pointer to **NullableFloat32** |  | [optional] 
**MaxMemory** | Pointer to **NullableFloat32** |  | [optional] 
**MinCores** | Pointer to **NullableString** |  | [optional] 
**MaxCores** | Pointer to **NullableString** |  | [optional] 
**MinSockets** | Pointer to **NullableString** |  | [optional] 
**MaxSockets** | Pointer to **NullableString** |  | [optional] 
**MinCoresPerSocket** | Pointer to **NullableString** |  | [optional] 
**MaxCoresPerSocket** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewServicePlanConfigRanges

`func NewServicePlanConfigRanges() *ServicePlanConfigRanges`

NewServicePlanConfigRanges instantiates a new ServicePlanConfigRanges object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicePlanConfigRangesWithDefaults

`func NewServicePlanConfigRangesWithDefaults() *ServicePlanConfigRanges`

NewServicePlanConfigRangesWithDefaults instantiates a new ServicePlanConfigRanges object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinStorage

`func (o *ServicePlanConfigRanges) GetMinStorage() string`

GetMinStorage returns the MinStorage field if non-nil, zero value otherwise.

### GetMinStorageOk

`func (o *ServicePlanConfigRanges) GetMinStorageOk() (*string, bool)`

GetMinStorageOk returns a tuple with the MinStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinStorage

`func (o *ServicePlanConfigRanges) SetMinStorage(v string)`

SetMinStorage sets MinStorage field to given value.

### HasMinStorage

`func (o *ServicePlanConfigRanges) HasMinStorage() bool`

HasMinStorage returns a boolean if a field has been set.

### SetMinStorageNil

`func (o *ServicePlanConfigRanges) SetMinStorageNil(b bool)`

 SetMinStorageNil sets the value for MinStorage to be an explicit nil

### UnsetMinStorage
`func (o *ServicePlanConfigRanges) UnsetMinStorage()`

UnsetMinStorage ensures that no value is present for MinStorage, not even an explicit nil
### GetMaxStorage

`func (o *ServicePlanConfigRanges) GetMaxStorage() string`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *ServicePlanConfigRanges) GetMaxStorageOk() (*string, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *ServicePlanConfigRanges) SetMaxStorage(v string)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *ServicePlanConfigRanges) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### SetMaxStorageNil

`func (o *ServicePlanConfigRanges) SetMaxStorageNil(b bool)`

 SetMaxStorageNil sets the value for MaxStorage to be an explicit nil

### UnsetMaxStorage
`func (o *ServicePlanConfigRanges) UnsetMaxStorage()`

UnsetMaxStorage ensures that no value is present for MaxStorage, not even an explicit nil
### GetMinPerDiskSize

`func (o *ServicePlanConfigRanges) GetMinPerDiskSize() string`

GetMinPerDiskSize returns the MinPerDiskSize field if non-nil, zero value otherwise.

### GetMinPerDiskSizeOk

`func (o *ServicePlanConfigRanges) GetMinPerDiskSizeOk() (*string, bool)`

GetMinPerDiskSizeOk returns a tuple with the MinPerDiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPerDiskSize

`func (o *ServicePlanConfigRanges) SetMinPerDiskSize(v string)`

SetMinPerDiskSize sets MinPerDiskSize field to given value.

### HasMinPerDiskSize

`func (o *ServicePlanConfigRanges) HasMinPerDiskSize() bool`

HasMinPerDiskSize returns a boolean if a field has been set.

### SetMinPerDiskSizeNil

`func (o *ServicePlanConfigRanges) SetMinPerDiskSizeNil(b bool)`

 SetMinPerDiskSizeNil sets the value for MinPerDiskSize to be an explicit nil

### UnsetMinPerDiskSize
`func (o *ServicePlanConfigRanges) UnsetMinPerDiskSize()`

UnsetMinPerDiskSize ensures that no value is present for MinPerDiskSize, not even an explicit nil
### GetMaxPerDiskSize

`func (o *ServicePlanConfigRanges) GetMaxPerDiskSize() string`

GetMaxPerDiskSize returns the MaxPerDiskSize field if non-nil, zero value otherwise.

### GetMaxPerDiskSizeOk

`func (o *ServicePlanConfigRanges) GetMaxPerDiskSizeOk() (*string, bool)`

GetMaxPerDiskSizeOk returns a tuple with the MaxPerDiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPerDiskSize

`func (o *ServicePlanConfigRanges) SetMaxPerDiskSize(v string)`

SetMaxPerDiskSize sets MaxPerDiskSize field to given value.

### HasMaxPerDiskSize

`func (o *ServicePlanConfigRanges) HasMaxPerDiskSize() bool`

HasMaxPerDiskSize returns a boolean if a field has been set.

### SetMaxPerDiskSizeNil

`func (o *ServicePlanConfigRanges) SetMaxPerDiskSizeNil(b bool)`

 SetMaxPerDiskSizeNil sets the value for MaxPerDiskSize to be an explicit nil

### UnsetMaxPerDiskSize
`func (o *ServicePlanConfigRanges) UnsetMaxPerDiskSize()`

UnsetMaxPerDiskSize ensures that no value is present for MaxPerDiskSize, not even an explicit nil
### GetMinMemory

`func (o *ServicePlanConfigRanges) GetMinMemory() float32`

GetMinMemory returns the MinMemory field if non-nil, zero value otherwise.

### GetMinMemoryOk

`func (o *ServicePlanConfigRanges) GetMinMemoryOk() (*float32, bool)`

GetMinMemoryOk returns a tuple with the MinMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinMemory

`func (o *ServicePlanConfigRanges) SetMinMemory(v float32)`

SetMinMemory sets MinMemory field to given value.

### HasMinMemory

`func (o *ServicePlanConfigRanges) HasMinMemory() bool`

HasMinMemory returns a boolean if a field has been set.

### SetMinMemoryNil

`func (o *ServicePlanConfigRanges) SetMinMemoryNil(b bool)`

 SetMinMemoryNil sets the value for MinMemory to be an explicit nil

### UnsetMinMemory
`func (o *ServicePlanConfigRanges) UnsetMinMemory()`

UnsetMinMemory ensures that no value is present for MinMemory, not even an explicit nil
### GetMaxMemory

`func (o *ServicePlanConfigRanges) GetMaxMemory() float32`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *ServicePlanConfigRanges) GetMaxMemoryOk() (*float32, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *ServicePlanConfigRanges) SetMaxMemory(v float32)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *ServicePlanConfigRanges) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### SetMaxMemoryNil

`func (o *ServicePlanConfigRanges) SetMaxMemoryNil(b bool)`

 SetMaxMemoryNil sets the value for MaxMemory to be an explicit nil

### UnsetMaxMemory
`func (o *ServicePlanConfigRanges) UnsetMaxMemory()`

UnsetMaxMemory ensures that no value is present for MaxMemory, not even an explicit nil
### GetMinCores

`func (o *ServicePlanConfigRanges) GetMinCores() string`

GetMinCores returns the MinCores field if non-nil, zero value otherwise.

### GetMinCoresOk

`func (o *ServicePlanConfigRanges) GetMinCoresOk() (*string, bool)`

GetMinCoresOk returns a tuple with the MinCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinCores

`func (o *ServicePlanConfigRanges) SetMinCores(v string)`

SetMinCores sets MinCores field to given value.

### HasMinCores

`func (o *ServicePlanConfigRanges) HasMinCores() bool`

HasMinCores returns a boolean if a field has been set.

### SetMinCoresNil

`func (o *ServicePlanConfigRanges) SetMinCoresNil(b bool)`

 SetMinCoresNil sets the value for MinCores to be an explicit nil

### UnsetMinCores
`func (o *ServicePlanConfigRanges) UnsetMinCores()`

UnsetMinCores ensures that no value is present for MinCores, not even an explicit nil
### GetMaxCores

`func (o *ServicePlanConfigRanges) GetMaxCores() string`

GetMaxCores returns the MaxCores field if non-nil, zero value otherwise.

### GetMaxCoresOk

`func (o *ServicePlanConfigRanges) GetMaxCoresOk() (*string, bool)`

GetMaxCoresOk returns a tuple with the MaxCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCores

`func (o *ServicePlanConfigRanges) SetMaxCores(v string)`

SetMaxCores sets MaxCores field to given value.

### HasMaxCores

`func (o *ServicePlanConfigRanges) HasMaxCores() bool`

HasMaxCores returns a boolean if a field has been set.

### SetMaxCoresNil

`func (o *ServicePlanConfigRanges) SetMaxCoresNil(b bool)`

 SetMaxCoresNil sets the value for MaxCores to be an explicit nil

### UnsetMaxCores
`func (o *ServicePlanConfigRanges) UnsetMaxCores()`

UnsetMaxCores ensures that no value is present for MaxCores, not even an explicit nil
### GetMinSockets

`func (o *ServicePlanConfigRanges) GetMinSockets() string`

GetMinSockets returns the MinSockets field if non-nil, zero value otherwise.

### GetMinSocketsOk

`func (o *ServicePlanConfigRanges) GetMinSocketsOk() (*string, bool)`

GetMinSocketsOk returns a tuple with the MinSockets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSockets

`func (o *ServicePlanConfigRanges) SetMinSockets(v string)`

SetMinSockets sets MinSockets field to given value.

### HasMinSockets

`func (o *ServicePlanConfigRanges) HasMinSockets() bool`

HasMinSockets returns a boolean if a field has been set.

### SetMinSocketsNil

`func (o *ServicePlanConfigRanges) SetMinSocketsNil(b bool)`

 SetMinSocketsNil sets the value for MinSockets to be an explicit nil

### UnsetMinSockets
`func (o *ServicePlanConfigRanges) UnsetMinSockets()`

UnsetMinSockets ensures that no value is present for MinSockets, not even an explicit nil
### GetMaxSockets

`func (o *ServicePlanConfigRanges) GetMaxSockets() string`

GetMaxSockets returns the MaxSockets field if non-nil, zero value otherwise.

### GetMaxSocketsOk

`func (o *ServicePlanConfigRanges) GetMaxSocketsOk() (*string, bool)`

GetMaxSocketsOk returns a tuple with the MaxSockets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSockets

`func (o *ServicePlanConfigRanges) SetMaxSockets(v string)`

SetMaxSockets sets MaxSockets field to given value.

### HasMaxSockets

`func (o *ServicePlanConfigRanges) HasMaxSockets() bool`

HasMaxSockets returns a boolean if a field has been set.

### SetMaxSocketsNil

`func (o *ServicePlanConfigRanges) SetMaxSocketsNil(b bool)`

 SetMaxSocketsNil sets the value for MaxSockets to be an explicit nil

### UnsetMaxSockets
`func (o *ServicePlanConfigRanges) UnsetMaxSockets()`

UnsetMaxSockets ensures that no value is present for MaxSockets, not even an explicit nil
### GetMinCoresPerSocket

`func (o *ServicePlanConfigRanges) GetMinCoresPerSocket() string`

GetMinCoresPerSocket returns the MinCoresPerSocket field if non-nil, zero value otherwise.

### GetMinCoresPerSocketOk

`func (o *ServicePlanConfigRanges) GetMinCoresPerSocketOk() (*string, bool)`

GetMinCoresPerSocketOk returns a tuple with the MinCoresPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinCoresPerSocket

`func (o *ServicePlanConfigRanges) SetMinCoresPerSocket(v string)`

SetMinCoresPerSocket sets MinCoresPerSocket field to given value.

### HasMinCoresPerSocket

`func (o *ServicePlanConfigRanges) HasMinCoresPerSocket() bool`

HasMinCoresPerSocket returns a boolean if a field has been set.

### SetMinCoresPerSocketNil

`func (o *ServicePlanConfigRanges) SetMinCoresPerSocketNil(b bool)`

 SetMinCoresPerSocketNil sets the value for MinCoresPerSocket to be an explicit nil

### UnsetMinCoresPerSocket
`func (o *ServicePlanConfigRanges) UnsetMinCoresPerSocket()`

UnsetMinCoresPerSocket ensures that no value is present for MinCoresPerSocket, not even an explicit nil
### GetMaxCoresPerSocket

`func (o *ServicePlanConfigRanges) GetMaxCoresPerSocket() string`

GetMaxCoresPerSocket returns the MaxCoresPerSocket field if non-nil, zero value otherwise.

### GetMaxCoresPerSocketOk

`func (o *ServicePlanConfigRanges) GetMaxCoresPerSocketOk() (*string, bool)`

GetMaxCoresPerSocketOk returns a tuple with the MaxCoresPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCoresPerSocket

`func (o *ServicePlanConfigRanges) SetMaxCoresPerSocket(v string)`

SetMaxCoresPerSocket sets MaxCoresPerSocket field to given value.

### HasMaxCoresPerSocket

`func (o *ServicePlanConfigRanges) HasMaxCoresPerSocket() bool`

HasMaxCoresPerSocket returns a boolean if a field has been set.

### SetMaxCoresPerSocketNil

`func (o *ServicePlanConfigRanges) SetMaxCoresPerSocketNil(b bool)`

 SetMaxCoresPerSocketNil sets the value for MaxCoresPerSocket to be an explicit nil

### UnsetMaxCoresPerSocket
`func (o *ServicePlanConfigRanges) UnsetMaxCoresPerSocket()`

UnsetMaxCoresPerSocket ensures that no value is present for MaxCoresPerSocket, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


