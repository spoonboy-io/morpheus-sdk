# DashboardMonitoring

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvgHealth** | Pointer to **float32** |  | [optional] 
**AvgResponseTime** | Pointer to **float32** |  | [optional] 
**WarningApps** | Pointer to **float32** |  | [optional] 
**WarningChecks** | Pointer to **float32** |  | [optional] 
**FailApps** | Pointer to **float32** |  | [optional] 
**TotalApps** | Pointer to **float32** |  | [optional] 
**FailChecks** | Pointer to **float32** |  | [optional] 
**SuccessApps** | Pointer to **float32** |  | [optional] 
**MutedApps** | Pointer to **float32** |  | [optional] 
**SuccessChecks** | Pointer to **float32** |  | [optional] 
**TotalChecks** | Pointer to **float32** |  | [optional] 
**MutedChecks** | Pointer to **float32** |  | [optional] 
**ResponseTimes** | Pointer to **[]float32** |  | [optional] 
**AllSuccess** | Pointer to **bool** |  | [optional] 
**OpenIncidents** | Pointer to **float32** |  | [optional] 

## Methods

### NewDashboardMonitoring

`func NewDashboardMonitoring() *DashboardMonitoring`

NewDashboardMonitoring instantiates a new DashboardMonitoring object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardMonitoringWithDefaults

`func NewDashboardMonitoringWithDefaults() *DashboardMonitoring`

NewDashboardMonitoringWithDefaults instantiates a new DashboardMonitoring object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvgHealth

`func (o *DashboardMonitoring) GetAvgHealth() float32`

GetAvgHealth returns the AvgHealth field if non-nil, zero value otherwise.

### GetAvgHealthOk

`func (o *DashboardMonitoring) GetAvgHealthOk() (*float32, bool)`

GetAvgHealthOk returns a tuple with the AvgHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgHealth

`func (o *DashboardMonitoring) SetAvgHealth(v float32)`

SetAvgHealth sets AvgHealth field to given value.

### HasAvgHealth

`func (o *DashboardMonitoring) HasAvgHealth() bool`

HasAvgHealth returns a boolean if a field has been set.

### GetAvgResponseTime

`func (o *DashboardMonitoring) GetAvgResponseTime() float32`

GetAvgResponseTime returns the AvgResponseTime field if non-nil, zero value otherwise.

### GetAvgResponseTimeOk

`func (o *DashboardMonitoring) GetAvgResponseTimeOk() (*float32, bool)`

GetAvgResponseTimeOk returns a tuple with the AvgResponseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgResponseTime

`func (o *DashboardMonitoring) SetAvgResponseTime(v float32)`

SetAvgResponseTime sets AvgResponseTime field to given value.

### HasAvgResponseTime

`func (o *DashboardMonitoring) HasAvgResponseTime() bool`

HasAvgResponseTime returns a boolean if a field has been set.

### GetWarningApps

`func (o *DashboardMonitoring) GetWarningApps() float32`

GetWarningApps returns the WarningApps field if non-nil, zero value otherwise.

### GetWarningAppsOk

`func (o *DashboardMonitoring) GetWarningAppsOk() (*float32, bool)`

GetWarningAppsOk returns a tuple with the WarningApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningApps

`func (o *DashboardMonitoring) SetWarningApps(v float32)`

SetWarningApps sets WarningApps field to given value.

### HasWarningApps

`func (o *DashboardMonitoring) HasWarningApps() bool`

HasWarningApps returns a boolean if a field has been set.

### GetWarningChecks

`func (o *DashboardMonitoring) GetWarningChecks() float32`

GetWarningChecks returns the WarningChecks field if non-nil, zero value otherwise.

### GetWarningChecksOk

`func (o *DashboardMonitoring) GetWarningChecksOk() (*float32, bool)`

GetWarningChecksOk returns a tuple with the WarningChecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningChecks

`func (o *DashboardMonitoring) SetWarningChecks(v float32)`

SetWarningChecks sets WarningChecks field to given value.

### HasWarningChecks

`func (o *DashboardMonitoring) HasWarningChecks() bool`

HasWarningChecks returns a boolean if a field has been set.

### GetFailApps

`func (o *DashboardMonitoring) GetFailApps() float32`

GetFailApps returns the FailApps field if non-nil, zero value otherwise.

### GetFailAppsOk

`func (o *DashboardMonitoring) GetFailAppsOk() (*float32, bool)`

GetFailAppsOk returns a tuple with the FailApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailApps

`func (o *DashboardMonitoring) SetFailApps(v float32)`

SetFailApps sets FailApps field to given value.

### HasFailApps

`func (o *DashboardMonitoring) HasFailApps() bool`

HasFailApps returns a boolean if a field has been set.

### GetTotalApps

`func (o *DashboardMonitoring) GetTotalApps() float32`

GetTotalApps returns the TotalApps field if non-nil, zero value otherwise.

### GetTotalAppsOk

`func (o *DashboardMonitoring) GetTotalAppsOk() (*float32, bool)`

GetTotalAppsOk returns a tuple with the TotalApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalApps

`func (o *DashboardMonitoring) SetTotalApps(v float32)`

SetTotalApps sets TotalApps field to given value.

### HasTotalApps

`func (o *DashboardMonitoring) HasTotalApps() bool`

HasTotalApps returns a boolean if a field has been set.

### GetFailChecks

`func (o *DashboardMonitoring) GetFailChecks() float32`

GetFailChecks returns the FailChecks field if non-nil, zero value otherwise.

### GetFailChecksOk

`func (o *DashboardMonitoring) GetFailChecksOk() (*float32, bool)`

GetFailChecksOk returns a tuple with the FailChecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailChecks

`func (o *DashboardMonitoring) SetFailChecks(v float32)`

SetFailChecks sets FailChecks field to given value.

### HasFailChecks

`func (o *DashboardMonitoring) HasFailChecks() bool`

HasFailChecks returns a boolean if a field has been set.

### GetSuccessApps

`func (o *DashboardMonitoring) GetSuccessApps() float32`

GetSuccessApps returns the SuccessApps field if non-nil, zero value otherwise.

### GetSuccessAppsOk

`func (o *DashboardMonitoring) GetSuccessAppsOk() (*float32, bool)`

GetSuccessAppsOk returns a tuple with the SuccessApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessApps

`func (o *DashboardMonitoring) SetSuccessApps(v float32)`

SetSuccessApps sets SuccessApps field to given value.

### HasSuccessApps

`func (o *DashboardMonitoring) HasSuccessApps() bool`

HasSuccessApps returns a boolean if a field has been set.

### GetMutedApps

`func (o *DashboardMonitoring) GetMutedApps() float32`

GetMutedApps returns the MutedApps field if non-nil, zero value otherwise.

### GetMutedAppsOk

`func (o *DashboardMonitoring) GetMutedAppsOk() (*float32, bool)`

GetMutedAppsOk returns a tuple with the MutedApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutedApps

`func (o *DashboardMonitoring) SetMutedApps(v float32)`

SetMutedApps sets MutedApps field to given value.

### HasMutedApps

`func (o *DashboardMonitoring) HasMutedApps() bool`

HasMutedApps returns a boolean if a field has been set.

### GetSuccessChecks

`func (o *DashboardMonitoring) GetSuccessChecks() float32`

GetSuccessChecks returns the SuccessChecks field if non-nil, zero value otherwise.

### GetSuccessChecksOk

`func (o *DashboardMonitoring) GetSuccessChecksOk() (*float32, bool)`

GetSuccessChecksOk returns a tuple with the SuccessChecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessChecks

`func (o *DashboardMonitoring) SetSuccessChecks(v float32)`

SetSuccessChecks sets SuccessChecks field to given value.

### HasSuccessChecks

`func (o *DashboardMonitoring) HasSuccessChecks() bool`

HasSuccessChecks returns a boolean if a field has been set.

### GetTotalChecks

`func (o *DashboardMonitoring) GetTotalChecks() float32`

GetTotalChecks returns the TotalChecks field if non-nil, zero value otherwise.

### GetTotalChecksOk

`func (o *DashboardMonitoring) GetTotalChecksOk() (*float32, bool)`

GetTotalChecksOk returns a tuple with the TotalChecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalChecks

`func (o *DashboardMonitoring) SetTotalChecks(v float32)`

SetTotalChecks sets TotalChecks field to given value.

### HasTotalChecks

`func (o *DashboardMonitoring) HasTotalChecks() bool`

HasTotalChecks returns a boolean if a field has been set.

### GetMutedChecks

`func (o *DashboardMonitoring) GetMutedChecks() float32`

GetMutedChecks returns the MutedChecks field if non-nil, zero value otherwise.

### GetMutedChecksOk

`func (o *DashboardMonitoring) GetMutedChecksOk() (*float32, bool)`

GetMutedChecksOk returns a tuple with the MutedChecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutedChecks

`func (o *DashboardMonitoring) SetMutedChecks(v float32)`

SetMutedChecks sets MutedChecks field to given value.

### HasMutedChecks

`func (o *DashboardMonitoring) HasMutedChecks() bool`

HasMutedChecks returns a boolean if a field has been set.

### GetResponseTimes

`func (o *DashboardMonitoring) GetResponseTimes() []float32`

GetResponseTimes returns the ResponseTimes field if non-nil, zero value otherwise.

### GetResponseTimesOk

`func (o *DashboardMonitoring) GetResponseTimesOk() (*[]float32, bool)`

GetResponseTimesOk returns a tuple with the ResponseTimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTimes

`func (o *DashboardMonitoring) SetResponseTimes(v []float32)`

SetResponseTimes sets ResponseTimes field to given value.

### HasResponseTimes

`func (o *DashboardMonitoring) HasResponseTimes() bool`

HasResponseTimes returns a boolean if a field has been set.

### GetAllSuccess

`func (o *DashboardMonitoring) GetAllSuccess() bool`

GetAllSuccess returns the AllSuccess field if non-nil, zero value otherwise.

### GetAllSuccessOk

`func (o *DashboardMonitoring) GetAllSuccessOk() (*bool, bool)`

GetAllSuccessOk returns a tuple with the AllSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllSuccess

`func (o *DashboardMonitoring) SetAllSuccess(v bool)`

SetAllSuccess sets AllSuccess field to given value.

### HasAllSuccess

`func (o *DashboardMonitoring) HasAllSuccess() bool`

HasAllSuccess returns a boolean if a field has been set.

### GetOpenIncidents

`func (o *DashboardMonitoring) GetOpenIncidents() float32`

GetOpenIncidents returns the OpenIncidents field if non-nil, zero value otherwise.

### GetOpenIncidentsOk

`func (o *DashboardMonitoring) GetOpenIncidentsOk() (*float32, bool)`

GetOpenIncidentsOk returns a tuple with the OpenIncidents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenIncidents

`func (o *DashboardMonitoring) SetOpenIncidents(v float32)`

SetOpenIncidents sets OpenIncidents field to given value.

### HasOpenIncidents

`func (o *DashboardMonitoring) HasOpenIncidents() bool`

HasOpenIncidents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


