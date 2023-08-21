# GroupStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceCounts** | Pointer to [**GroupStatsInstanceCounts**](group_stats_instanceCounts.md) |  | [optional] 
**ServerCounts** | Pointer to [**ZoneStatsServerCounts**](zone_stats_serverCounts.md) |  | [optional] 

## Methods

### NewGroupStats

`func NewGroupStats() *GroupStats`

NewGroupStats instantiates a new GroupStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupStatsWithDefaults

`func NewGroupStatsWithDefaults() *GroupStats`

NewGroupStatsWithDefaults instantiates a new GroupStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceCounts

`func (o *GroupStats) GetInstanceCounts() GroupStatsInstanceCounts`

GetInstanceCounts returns the InstanceCounts field if non-nil, zero value otherwise.

### GetInstanceCountsOk

`func (o *GroupStats) GetInstanceCountsOk() (*GroupStatsInstanceCounts, bool)`

GetInstanceCountsOk returns a tuple with the InstanceCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceCounts

`func (o *GroupStats) SetInstanceCounts(v GroupStatsInstanceCounts)`

SetInstanceCounts sets InstanceCounts field to given value.

### HasInstanceCounts

`func (o *GroupStats) HasInstanceCounts() bool`

HasInstanceCounts returns a boolean if a field has been set.

### GetServerCounts

`func (o *GroupStats) GetServerCounts() ZoneStatsServerCounts`

GetServerCounts returns the ServerCounts field if non-nil, zero value otherwise.

### GetServerCountsOk

`func (o *GroupStats) GetServerCountsOk() (*ZoneStatsServerCounts, bool)`

GetServerCountsOk returns a tuple with the ServerCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCounts

`func (o *GroupStats) SetServerCounts(v ZoneStatsServerCounts)`

SetServerCounts sets ServerCounts field to given value.

### HasServerCounts

`func (o *GroupStats) HasServerCounts() bool`

HasServerCounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


