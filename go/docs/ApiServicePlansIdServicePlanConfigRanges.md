# ApiServicePlansIdServicePlanConfigRanges

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinStorage** | Pointer to **string** | Custom min storage in GB (unless &#x60;storageSizeType&#x60; set to mb) | [optional] 
**MaxStorage** | Pointer to **string** | Custom max storage in GB (unless &#x60;storageSizeType&#x60; set to mb) | [optional] 
**MinMemory** | Pointer to **int64** | Custom min memory in bytes | [optional] 
**MaxMemory** | Pointer to **int64** | Custom max memory in bytes | [optional] 
**MinCores** | Pointer to **string** | Custom min cores | [optional] 
**MaxCores** | Pointer to **string** | Custom max cores | [optional] 
**MinSockets** | Pointer to **string** | Custom min sockets | [optional] 
**MaxSockets** | Pointer to **string** | Custom max sockets | [optional] 
**MinCoresPerSocket** | Pointer to **string** | Custom min cores allowed per socket | [optional] 
**MaxCoresPerSocket** | Pointer to **string** | Custom max cores allowed per socket | [optional] 

## Methods

### NewApiServicePlansIdServicePlanConfigRanges

`func NewApiServicePlansIdServicePlanConfigRanges() *ApiServicePlansIdServicePlanConfigRanges`

NewApiServicePlansIdServicePlanConfigRanges instantiates a new ApiServicePlansIdServicePlanConfigRanges object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiServicePlansIdServicePlanConfigRangesWithDefaults

`func NewApiServicePlansIdServicePlanConfigRangesWithDefaults() *ApiServicePlansIdServicePlanConfigRanges`

NewApiServicePlansIdServicePlanConfigRangesWithDefaults instantiates a new ApiServicePlansIdServicePlanConfigRanges object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinStorage

`func (o *ApiServicePlansIdServicePlanConfigRanges) GetMinStorage() string`

GetMinStorage returns the MinStorage field if non-nil, zero value otherwise.

### GetMinStorageOk

`func (o *ApiServicePlansIdServicePlanConfigRanges) GetMinStorageOk() (*string, bool)`

GetMinStorageOk returns a tuple with the MinStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinStorage

`func (o *ApiServicePlansIdServicePlanConfigRanges) SetMinStorage(v string)`

SetMinStorage sets MinStorage field to given value.

### HasMinStorage

`func (o *ApiServicePlansIdServicePlanConfigRanges) HasMinStorage() bool`

HasMinStorage returns a boolean if a field has been set.

### GetMaxStorage

`func (o *ApiServicePlansIdServicePlanConfigRanges) GetMaxStorage() string`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *ApiServicePlansIdServicePlanConfigRanges) GetMaxStorageOk() (*string, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *ApiServicePlansIdServicePlanConfigRanges) SetMaxStorage(v string)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *ApiServicePlansIdServicePlanConfigRanges) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### GetMinMemory

`func (o *ApiServicePlansIdServicePlanConfigRanges) GetMinMemory() int64`

GetMinMemory returns the MinMemory field if non-nil, zero value otherwise.

### GetMinMemoryOk

`func (o *ApiServicePlansIdServicePlanConfigRanges) GetMinMemoryOk() (*int64, bool)`

GetMinMemoryOk returns a tuple with the MinMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinMemory

`func (o *ApiServicePlansIdServicePlanConfigRanges) SetMinMemory(v int64)`

SetMinMemory sets MinMemory field to given value.

### HasMinMemory

`func (o *ApiServicePlansIdServicePlanConfigRanges) HasMinMemory() bool`

HasMinMemory returns a boolean if a field has been set.

### GetMaxMemory

`func (o *ApiServicePlansIdServicePlanConfigRanges) GetMaxMemory() int64`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *ApiServicePlansIdServicePlanConfigRanges) GetMaxMemoryOk() (*int64, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *ApiServicePlansIdServicePlanConfigRanges) SetMaxMemory(v int64)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *ApiServicePlansIdServicePlanConfigRanges) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### GetMinCores

`func (o *ApiServicePlansIdServicePlanConfigRanges) GetMinCores() string`

GetMinCores returns the MinCores field if non-nil, zero value otherwise.

### GetMinCoresOk

`func (o *ApiServicePlansIdServicePlanConfigRanges) GetMinCoresOk() (*string, bool)`

GetMinCoresOk returns a tuple with the MinCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinCores

`func (o *ApiServicePlansIdServicePlanConfigRanges) SetMinCores(v string)`

SetMinCores sets MinCores field to given value.

### HasMinCores

`func (o *ApiServicePlansIdServicePlanConfigRanges) HasMinCores() bool`

HasMinCores returns a boolean if a field has been set.

### GetMaxCores

`func (o *ApiServicePlansIdServicePlanConfigRanges) GetMaxCores() string`

GetMaxCores returns the MaxCores field if non-nil, zero value otherwise.

### GetMaxCoresOk

`func (o *ApiServicePlansIdServicePlanConfigRanges) GetMaxCoresOk() (*string, bool)`

GetMaxCoresOk returns a tuple with the MaxCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCores

`func (o *ApiServicePlansIdServicePlanConfigRanges) SetMaxCores(v string)`

SetMaxCores sets MaxCores field to given value.

### HasMaxCores

`func (o *ApiServicePlansIdServicePlanConfigRanges) HasMaxCores() bool`

HasMaxCores returns a boolean if a field has been set.

### GetMinSockets

`func (o *ApiServicePlansIdServicePlanConfigRanges) GetMinSockets() string`

GetMinSockets returns the MinSockets field if non-nil, zero value otherwise.

### GetMinSocketsOk

`func (o *ApiServicePlansIdServicePlanConfigRanges) GetMinSocketsOk() (*string, bool)`

GetMinSocketsOk returns a tuple with the MinSockets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSockets

`func (o *ApiServicePlansIdServicePlanConfigRanges) SetMinSockets(v string)`

SetMinSockets sets MinSockets field to given value.

### HasMinSockets

`func (o *ApiServicePlansIdServicePlanConfigRanges) HasMinSockets() bool`

HasMinSockets returns a boolean if a field has been set.

### GetMaxSockets

`func (o *ApiServicePlansIdServicePlanConfigRanges) GetMaxSockets() string`

GetMaxSockets returns the MaxSockets field if non-nil, zero value otherwise.

### GetMaxSocketsOk

`func (o *ApiServicePlansIdServicePlanConfigRanges) GetMaxSocketsOk() (*string, bool)`

GetMaxSocketsOk returns a tuple with the MaxSockets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSockets

`func (o *ApiServicePlansIdServicePlanConfigRanges) SetMaxSockets(v string)`

SetMaxSockets sets MaxSockets field to given value.

### HasMaxSockets

`func (o *ApiServicePlansIdServicePlanConfigRanges) HasMaxSockets() bool`

HasMaxSockets returns a boolean if a field has been set.

### GetMinCoresPerSocket

`func (o *ApiServicePlansIdServicePlanConfigRanges) GetMinCoresPerSocket() string`

GetMinCoresPerSocket returns the MinCoresPerSocket field if non-nil, zero value otherwise.

### GetMinCoresPerSocketOk

`func (o *ApiServicePlansIdServicePlanConfigRanges) GetMinCoresPerSocketOk() (*string, bool)`

GetMinCoresPerSocketOk returns a tuple with the MinCoresPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinCoresPerSocket

`func (o *ApiServicePlansIdServicePlanConfigRanges) SetMinCoresPerSocket(v string)`

SetMinCoresPerSocket sets MinCoresPerSocket field to given value.

### HasMinCoresPerSocket

`func (o *ApiServicePlansIdServicePlanConfigRanges) HasMinCoresPerSocket() bool`

HasMinCoresPerSocket returns a boolean if a field has been set.

### GetMaxCoresPerSocket

`func (o *ApiServicePlansIdServicePlanConfigRanges) GetMaxCoresPerSocket() string`

GetMaxCoresPerSocket returns the MaxCoresPerSocket field if non-nil, zero value otherwise.

### GetMaxCoresPerSocketOk

`func (o *ApiServicePlansIdServicePlanConfigRanges) GetMaxCoresPerSocketOk() (*string, bool)`

GetMaxCoresPerSocketOk returns a tuple with the MaxCoresPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCoresPerSocket

`func (o *ApiServicePlansIdServicePlanConfigRanges) SetMaxCoresPerSocket(v string)`

SetMaxCoresPerSocket sets MaxCoresPerSocket field to given value.

### HasMaxCoresPerSocket

`func (o *ApiServicePlansIdServicePlanConfigRanges) HasMaxCoresPerSocket() bool`

HasMaxCoresPerSocket returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


