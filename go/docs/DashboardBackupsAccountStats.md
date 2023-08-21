# DashboardBackupsAccountStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalSizeByDay** | Pointer to **[]float32** |  | [optional] 
**TotalSizeByDay7Days** | Pointer to **[]float32** |  | [optional] 
**FormattedTotalSize** | Pointer to [**DashboardBackupsAccountStatsFormattedTotalSize**](dashboard_backups_accountStats_formattedTotalSize.md) |  | [optional] 
**BackupCount** | Pointer to **float32** |  | [optional] 
**TotalSize** | Pointer to **float32** |  | [optional] 
**Success** | Pointer to **float32** |  | [optional] 
**Failed** | Pointer to **float32** |  | [optional] 
**TotalCompleted** | Pointer to **float32** |  | [optional] 
**LastSevenDays** | Pointer to [**DashboardBackupsAccountStatsLastSevenDays**](dashboard_backups_accountStats_lastSevenDays.md) |  | [optional] 
**AvgSize** | Pointer to **float32** |  | [optional] 
**FailedRate** | Pointer to **float32** |  | [optional] 
**SuccessRate** | Pointer to **float32** |  | [optional] 
**NextFireTotal** | Pointer to **float32** |  | [optional] 
**BackupDayCount** | Pointer to **[]float32** |  | [optional] 
**BackupDayCountTotal** | Pointer to **float32** |  | [optional] 

## Methods

### NewDashboardBackupsAccountStats

`func NewDashboardBackupsAccountStats() *DashboardBackupsAccountStats`

NewDashboardBackupsAccountStats instantiates a new DashboardBackupsAccountStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardBackupsAccountStatsWithDefaults

`func NewDashboardBackupsAccountStatsWithDefaults() *DashboardBackupsAccountStats`

NewDashboardBackupsAccountStatsWithDefaults instantiates a new DashboardBackupsAccountStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalSizeByDay

`func (o *DashboardBackupsAccountStats) GetTotalSizeByDay() []float32`

GetTotalSizeByDay returns the TotalSizeByDay field if non-nil, zero value otherwise.

### GetTotalSizeByDayOk

`func (o *DashboardBackupsAccountStats) GetTotalSizeByDayOk() (*[]float32, bool)`

GetTotalSizeByDayOk returns a tuple with the TotalSizeByDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSizeByDay

`func (o *DashboardBackupsAccountStats) SetTotalSizeByDay(v []float32)`

SetTotalSizeByDay sets TotalSizeByDay field to given value.

### HasTotalSizeByDay

`func (o *DashboardBackupsAccountStats) HasTotalSizeByDay() bool`

HasTotalSizeByDay returns a boolean if a field has been set.

### GetTotalSizeByDay7Days

`func (o *DashboardBackupsAccountStats) GetTotalSizeByDay7Days() []float32`

GetTotalSizeByDay7Days returns the TotalSizeByDay7Days field if non-nil, zero value otherwise.

### GetTotalSizeByDay7DaysOk

`func (o *DashboardBackupsAccountStats) GetTotalSizeByDay7DaysOk() (*[]float32, bool)`

GetTotalSizeByDay7DaysOk returns a tuple with the TotalSizeByDay7Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSizeByDay7Days

`func (o *DashboardBackupsAccountStats) SetTotalSizeByDay7Days(v []float32)`

SetTotalSizeByDay7Days sets TotalSizeByDay7Days field to given value.

### HasTotalSizeByDay7Days

`func (o *DashboardBackupsAccountStats) HasTotalSizeByDay7Days() bool`

HasTotalSizeByDay7Days returns a boolean if a field has been set.

### GetFormattedTotalSize

`func (o *DashboardBackupsAccountStats) GetFormattedTotalSize() DashboardBackupsAccountStatsFormattedTotalSize`

GetFormattedTotalSize returns the FormattedTotalSize field if non-nil, zero value otherwise.

### GetFormattedTotalSizeOk

`func (o *DashboardBackupsAccountStats) GetFormattedTotalSizeOk() (*DashboardBackupsAccountStatsFormattedTotalSize, bool)`

GetFormattedTotalSizeOk returns a tuple with the FormattedTotalSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedTotalSize

`func (o *DashboardBackupsAccountStats) SetFormattedTotalSize(v DashboardBackupsAccountStatsFormattedTotalSize)`

SetFormattedTotalSize sets FormattedTotalSize field to given value.

### HasFormattedTotalSize

`func (o *DashboardBackupsAccountStats) HasFormattedTotalSize() bool`

HasFormattedTotalSize returns a boolean if a field has been set.

### GetBackupCount

`func (o *DashboardBackupsAccountStats) GetBackupCount() float32`

GetBackupCount returns the BackupCount field if non-nil, zero value otherwise.

### GetBackupCountOk

`func (o *DashboardBackupsAccountStats) GetBackupCountOk() (*float32, bool)`

GetBackupCountOk returns a tuple with the BackupCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupCount

`func (o *DashboardBackupsAccountStats) SetBackupCount(v float32)`

SetBackupCount sets BackupCount field to given value.

### HasBackupCount

`func (o *DashboardBackupsAccountStats) HasBackupCount() bool`

HasBackupCount returns a boolean if a field has been set.

### GetTotalSize

`func (o *DashboardBackupsAccountStats) GetTotalSize() float32`

GetTotalSize returns the TotalSize field if non-nil, zero value otherwise.

### GetTotalSizeOk

`func (o *DashboardBackupsAccountStats) GetTotalSizeOk() (*float32, bool)`

GetTotalSizeOk returns a tuple with the TotalSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSize

`func (o *DashboardBackupsAccountStats) SetTotalSize(v float32)`

SetTotalSize sets TotalSize field to given value.

### HasTotalSize

`func (o *DashboardBackupsAccountStats) HasTotalSize() bool`

HasTotalSize returns a boolean if a field has been set.

### GetSuccess

`func (o *DashboardBackupsAccountStats) GetSuccess() float32`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *DashboardBackupsAccountStats) GetSuccessOk() (*float32, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *DashboardBackupsAccountStats) SetSuccess(v float32)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *DashboardBackupsAccountStats) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetFailed

`func (o *DashboardBackupsAccountStats) GetFailed() float32`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *DashboardBackupsAccountStats) GetFailedOk() (*float32, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *DashboardBackupsAccountStats) SetFailed(v float32)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *DashboardBackupsAccountStats) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### GetTotalCompleted

`func (o *DashboardBackupsAccountStats) GetTotalCompleted() float32`

GetTotalCompleted returns the TotalCompleted field if non-nil, zero value otherwise.

### GetTotalCompletedOk

`func (o *DashboardBackupsAccountStats) GetTotalCompletedOk() (*float32, bool)`

GetTotalCompletedOk returns a tuple with the TotalCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCompleted

`func (o *DashboardBackupsAccountStats) SetTotalCompleted(v float32)`

SetTotalCompleted sets TotalCompleted field to given value.

### HasTotalCompleted

`func (o *DashboardBackupsAccountStats) HasTotalCompleted() bool`

HasTotalCompleted returns a boolean if a field has been set.

### GetLastSevenDays

`func (o *DashboardBackupsAccountStats) GetLastSevenDays() DashboardBackupsAccountStatsLastSevenDays`

GetLastSevenDays returns the LastSevenDays field if non-nil, zero value otherwise.

### GetLastSevenDaysOk

`func (o *DashboardBackupsAccountStats) GetLastSevenDaysOk() (*DashboardBackupsAccountStatsLastSevenDays, bool)`

GetLastSevenDaysOk returns a tuple with the LastSevenDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSevenDays

`func (o *DashboardBackupsAccountStats) SetLastSevenDays(v DashboardBackupsAccountStatsLastSevenDays)`

SetLastSevenDays sets LastSevenDays field to given value.

### HasLastSevenDays

`func (o *DashboardBackupsAccountStats) HasLastSevenDays() bool`

HasLastSevenDays returns a boolean if a field has been set.

### GetAvgSize

`func (o *DashboardBackupsAccountStats) GetAvgSize() float32`

GetAvgSize returns the AvgSize field if non-nil, zero value otherwise.

### GetAvgSizeOk

`func (o *DashboardBackupsAccountStats) GetAvgSizeOk() (*float32, bool)`

GetAvgSizeOk returns a tuple with the AvgSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgSize

`func (o *DashboardBackupsAccountStats) SetAvgSize(v float32)`

SetAvgSize sets AvgSize field to given value.

### HasAvgSize

`func (o *DashboardBackupsAccountStats) HasAvgSize() bool`

HasAvgSize returns a boolean if a field has been set.

### GetFailedRate

`func (o *DashboardBackupsAccountStats) GetFailedRate() float32`

GetFailedRate returns the FailedRate field if non-nil, zero value otherwise.

### GetFailedRateOk

`func (o *DashboardBackupsAccountStats) GetFailedRateOk() (*float32, bool)`

GetFailedRateOk returns a tuple with the FailedRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedRate

`func (o *DashboardBackupsAccountStats) SetFailedRate(v float32)`

SetFailedRate sets FailedRate field to given value.

### HasFailedRate

`func (o *DashboardBackupsAccountStats) HasFailedRate() bool`

HasFailedRate returns a boolean if a field has been set.

### GetSuccessRate

`func (o *DashboardBackupsAccountStats) GetSuccessRate() float32`

GetSuccessRate returns the SuccessRate field if non-nil, zero value otherwise.

### GetSuccessRateOk

`func (o *DashboardBackupsAccountStats) GetSuccessRateOk() (*float32, bool)`

GetSuccessRateOk returns a tuple with the SuccessRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessRate

`func (o *DashboardBackupsAccountStats) SetSuccessRate(v float32)`

SetSuccessRate sets SuccessRate field to given value.

### HasSuccessRate

`func (o *DashboardBackupsAccountStats) HasSuccessRate() bool`

HasSuccessRate returns a boolean if a field has been set.

### GetNextFireTotal

`func (o *DashboardBackupsAccountStats) GetNextFireTotal() float32`

GetNextFireTotal returns the NextFireTotal field if non-nil, zero value otherwise.

### GetNextFireTotalOk

`func (o *DashboardBackupsAccountStats) GetNextFireTotalOk() (*float32, bool)`

GetNextFireTotalOk returns a tuple with the NextFireTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextFireTotal

`func (o *DashboardBackupsAccountStats) SetNextFireTotal(v float32)`

SetNextFireTotal sets NextFireTotal field to given value.

### HasNextFireTotal

`func (o *DashboardBackupsAccountStats) HasNextFireTotal() bool`

HasNextFireTotal returns a boolean if a field has been set.

### GetBackupDayCount

`func (o *DashboardBackupsAccountStats) GetBackupDayCount() []float32`

GetBackupDayCount returns the BackupDayCount field if non-nil, zero value otherwise.

### GetBackupDayCountOk

`func (o *DashboardBackupsAccountStats) GetBackupDayCountOk() (*[]float32, bool)`

GetBackupDayCountOk returns a tuple with the BackupDayCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupDayCount

`func (o *DashboardBackupsAccountStats) SetBackupDayCount(v []float32)`

SetBackupDayCount sets BackupDayCount field to given value.

### HasBackupDayCount

`func (o *DashboardBackupsAccountStats) HasBackupDayCount() bool`

HasBackupDayCount returns a boolean if a field has been set.

### GetBackupDayCountTotal

`func (o *DashboardBackupsAccountStats) GetBackupDayCountTotal() float32`

GetBackupDayCountTotal returns the BackupDayCountTotal field if non-nil, zero value otherwise.

### GetBackupDayCountTotalOk

`func (o *DashboardBackupsAccountStats) GetBackupDayCountTotalOk() (*float32, bool)`

GetBackupDayCountTotalOk returns a tuple with the BackupDayCountTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupDayCountTotal

`func (o *DashboardBackupsAccountStats) SetBackupDayCountTotal(v float32)`

SetBackupDayCountTotal sets BackupDayCountTotal field to given value.

### HasBackupDayCountTotal

`func (o *DashboardBackupsAccountStats) HasBackupDayCountTotal() bool`

HasBackupDayCountTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


