# DashboardLogStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**Data** | Pointer to [**[]DashboardLogStatsData**](DashboardLogStatsData.md) |  | [optional] 
**StartMs** | Pointer to **int64** |  | [optional] 
**Earliest** | Pointer to **int64** |  | [optional] 
**EndMs** | Pointer to **int64** |  | [optional] 
**Interval** | Pointer to **int64** |  | [optional] 

## Methods

### NewDashboardLogStats

`func NewDashboardLogStats() *DashboardLogStats`

NewDashboardLogStats instantiates a new DashboardLogStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardLogStatsWithDefaults

`func NewDashboardLogStatsWithDefaults() *DashboardLogStats`

NewDashboardLogStatsWithDefaults instantiates a new DashboardLogStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *DashboardLogStats) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *DashboardLogStats) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *DashboardLogStats) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *DashboardLogStats) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetData

`func (o *DashboardLogStats) GetData() []DashboardLogStatsData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DashboardLogStats) GetDataOk() (*[]DashboardLogStatsData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DashboardLogStats) SetData(v []DashboardLogStatsData)`

SetData sets Data field to given value.

### HasData

`func (o *DashboardLogStats) HasData() bool`

HasData returns a boolean if a field has been set.

### GetStartMs

`func (o *DashboardLogStats) GetStartMs() int64`

GetStartMs returns the StartMs field if non-nil, zero value otherwise.

### GetStartMsOk

`func (o *DashboardLogStats) GetStartMsOk() (*int64, bool)`

GetStartMsOk returns a tuple with the StartMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartMs

`func (o *DashboardLogStats) SetStartMs(v int64)`

SetStartMs sets StartMs field to given value.

### HasStartMs

`func (o *DashboardLogStats) HasStartMs() bool`

HasStartMs returns a boolean if a field has been set.

### GetEarliest

`func (o *DashboardLogStats) GetEarliest() int64`

GetEarliest returns the Earliest field if non-nil, zero value otherwise.

### GetEarliestOk

`func (o *DashboardLogStats) GetEarliestOk() (*int64, bool)`

GetEarliestOk returns a tuple with the Earliest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarliest

`func (o *DashboardLogStats) SetEarliest(v int64)`

SetEarliest sets Earliest field to given value.

### HasEarliest

`func (o *DashboardLogStats) HasEarliest() bool`

HasEarliest returns a boolean if a field has been set.

### GetEndMs

`func (o *DashboardLogStats) GetEndMs() int64`

GetEndMs returns the EndMs field if non-nil, zero value otherwise.

### GetEndMsOk

`func (o *DashboardLogStats) GetEndMsOk() (*int64, bool)`

GetEndMsOk returns a tuple with the EndMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndMs

`func (o *DashboardLogStats) SetEndMs(v int64)`

SetEndMs sets EndMs field to given value.

### HasEndMs

`func (o *DashboardLogStats) HasEndMs() bool`

HasEndMs returns a boolean if a field has been set.

### GetInterval

`func (o *DashboardLogStats) GetInterval() int64`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *DashboardLogStats) GetIntervalOk() (*int64, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *DashboardLogStats) SetInterval(v int64)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *DashboardLogStats) HasInterval() bool`

HasInterval returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


