# Health

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**StatusMessage** | Pointer to **string** |  | [optional] 
**ApplianceUrl** | Pointer to **string** |  | [optional] 
**BuildVersion** | Pointer to **string** |  | [optional] 
**SetupNeeded** | Pointer to **bool** |  | [optional] 
**Date** | Pointer to **time.Time** |  | [optional] 
**Cpu** | Pointer to [**HealthCpu**](health_cpu.md) |  | [optional] 
**Memory** | Pointer to [**HealthMemory**](health_memory.md) |  | [optional] 
**Threads** | Pointer to [**HealthThreads**](health_threads.md) |  | [optional] 
**Database** | Pointer to [**HealthDatabase**](health_database.md) |  | [optional] 
**Elastic** | Pointer to [**HealthElastic**](health_elastic.md) |  | [optional] 
**Rabbit** | Pointer to [**HealthRabbit**](health_rabbit.md) |  | [optional] 

## Methods

### NewHealth

`func NewHealth() *Health`

NewHealth instantiates a new Health object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthWithDefaults

`func NewHealthWithDefaults() *Health`

NewHealthWithDefaults instantiates a new Health object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *Health) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *Health) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *Health) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *Health) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetStatusMessage

`func (o *Health) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *Health) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *Health) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *Health) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetApplianceUrl

`func (o *Health) GetApplianceUrl() string`

GetApplianceUrl returns the ApplianceUrl field if non-nil, zero value otherwise.

### GetApplianceUrlOk

`func (o *Health) GetApplianceUrlOk() (*string, bool)`

GetApplianceUrlOk returns a tuple with the ApplianceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceUrl

`func (o *Health) SetApplianceUrl(v string)`

SetApplianceUrl sets ApplianceUrl field to given value.

### HasApplianceUrl

`func (o *Health) HasApplianceUrl() bool`

HasApplianceUrl returns a boolean if a field has been set.

### GetBuildVersion

`func (o *Health) GetBuildVersion() string`

GetBuildVersion returns the BuildVersion field if non-nil, zero value otherwise.

### GetBuildVersionOk

`func (o *Health) GetBuildVersionOk() (*string, bool)`

GetBuildVersionOk returns a tuple with the BuildVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildVersion

`func (o *Health) SetBuildVersion(v string)`

SetBuildVersion sets BuildVersion field to given value.

### HasBuildVersion

`func (o *Health) HasBuildVersion() bool`

HasBuildVersion returns a boolean if a field has been set.

### GetSetupNeeded

`func (o *Health) GetSetupNeeded() bool`

GetSetupNeeded returns the SetupNeeded field if non-nil, zero value otherwise.

### GetSetupNeededOk

`func (o *Health) GetSetupNeededOk() (*bool, bool)`

GetSetupNeededOk returns a tuple with the SetupNeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetupNeeded

`func (o *Health) SetSetupNeeded(v bool)`

SetSetupNeeded sets SetupNeeded field to given value.

### HasSetupNeeded

`func (o *Health) HasSetupNeeded() bool`

HasSetupNeeded returns a boolean if a field has been set.

### GetDate

`func (o *Health) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *Health) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *Health) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *Health) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetCpu

`func (o *Health) GetCpu() HealthCpu`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *Health) GetCpuOk() (*HealthCpu, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *Health) SetCpu(v HealthCpu)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *Health) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetMemory

`func (o *Health) GetMemory() HealthMemory`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *Health) GetMemoryOk() (*HealthMemory, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *Health) SetMemory(v HealthMemory)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *Health) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetThreads

`func (o *Health) GetThreads() HealthThreads`

GetThreads returns the Threads field if non-nil, zero value otherwise.

### GetThreadsOk

`func (o *Health) GetThreadsOk() (*HealthThreads, bool)`

GetThreadsOk returns a tuple with the Threads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreads

`func (o *Health) SetThreads(v HealthThreads)`

SetThreads sets Threads field to given value.

### HasThreads

`func (o *Health) HasThreads() bool`

HasThreads returns a boolean if a field has been set.

### GetDatabase

`func (o *Health) GetDatabase() HealthDatabase`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *Health) GetDatabaseOk() (*HealthDatabase, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *Health) SetDatabase(v HealthDatabase)`

SetDatabase sets Database field to given value.

### HasDatabase

`func (o *Health) HasDatabase() bool`

HasDatabase returns a boolean if a field has been set.

### GetElastic

`func (o *Health) GetElastic() HealthElastic`

GetElastic returns the Elastic field if non-nil, zero value otherwise.

### GetElasticOk

`func (o *Health) GetElasticOk() (*HealthElastic, bool)`

GetElasticOk returns a tuple with the Elastic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElastic

`func (o *Health) SetElastic(v HealthElastic)`

SetElastic sets Elastic field to given value.

### HasElastic

`func (o *Health) HasElastic() bool`

HasElastic returns a boolean if a field has been set.

### GetRabbit

`func (o *Health) GetRabbit() HealthRabbit`

GetRabbit returns the Rabbit field if non-nil, zero value otherwise.

### GetRabbitOk

`func (o *Health) GetRabbitOk() (*HealthRabbit, bool)`

GetRabbitOk returns a tuple with the Rabbit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRabbit

`func (o *Health) SetRabbit(v HealthRabbit)`

SetRabbit sets Rabbit field to given value.

### HasRabbit

`func (o *Health) HasRabbit() bool`

HasRabbit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


