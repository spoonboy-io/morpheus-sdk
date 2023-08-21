# Dashboard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**InstanceStats** | Pointer to [**DashboardInstanceStats**](dashboard_instanceStats.md) |  | [optional] 
**Provisioning** | Pointer to [**DashboardProvisioning**](dashboard_provisioning.md) |  | [optional] 
**Monitoring** | Pointer to [**DashboardMonitoring**](dashboard_monitoring.md) |  | [optional] 
**Backups** | Pointer to [**DashboardBackups**](dashboard_backups.md) |  | [optional] 
**Activity** | Pointer to [**[]DashboardActivity**](DashboardActivity.md) |  | [optional] 
**LogStats** | Pointer to [**DashboardLogStats**](dashboard_logStats.md) |  | [optional] 

## Methods

### NewDashboard

`func NewDashboard() *Dashboard`

NewDashboard instantiates a new Dashboard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardWithDefaults

`func NewDashboardWithDefaults() *Dashboard`

NewDashboardWithDefaults instantiates a new Dashboard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *Dashboard) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *Dashboard) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *Dashboard) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *Dashboard) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetInstanceStats

`func (o *Dashboard) GetInstanceStats() DashboardInstanceStats`

GetInstanceStats returns the InstanceStats field if non-nil, zero value otherwise.

### GetInstanceStatsOk

`func (o *Dashboard) GetInstanceStatsOk() (*DashboardInstanceStats, bool)`

GetInstanceStatsOk returns a tuple with the InstanceStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceStats

`func (o *Dashboard) SetInstanceStats(v DashboardInstanceStats)`

SetInstanceStats sets InstanceStats field to given value.

### HasInstanceStats

`func (o *Dashboard) HasInstanceStats() bool`

HasInstanceStats returns a boolean if a field has been set.

### GetProvisioning

`func (o *Dashboard) GetProvisioning() DashboardProvisioning`

GetProvisioning returns the Provisioning field if non-nil, zero value otherwise.

### GetProvisioningOk

`func (o *Dashboard) GetProvisioningOk() (*DashboardProvisioning, bool)`

GetProvisioningOk returns a tuple with the Provisioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioning

`func (o *Dashboard) SetProvisioning(v DashboardProvisioning)`

SetProvisioning sets Provisioning field to given value.

### HasProvisioning

`func (o *Dashboard) HasProvisioning() bool`

HasProvisioning returns a boolean if a field has been set.

### GetMonitoring

`func (o *Dashboard) GetMonitoring() DashboardMonitoring`

GetMonitoring returns the Monitoring field if non-nil, zero value otherwise.

### GetMonitoringOk

`func (o *Dashboard) GetMonitoringOk() (*DashboardMonitoring, bool)`

GetMonitoringOk returns a tuple with the Monitoring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoring

`func (o *Dashboard) SetMonitoring(v DashboardMonitoring)`

SetMonitoring sets Monitoring field to given value.

### HasMonitoring

`func (o *Dashboard) HasMonitoring() bool`

HasMonitoring returns a boolean if a field has been set.

### GetBackups

`func (o *Dashboard) GetBackups() DashboardBackups`

GetBackups returns the Backups field if non-nil, zero value otherwise.

### GetBackupsOk

`func (o *Dashboard) GetBackupsOk() (*DashboardBackups, bool)`

GetBackupsOk returns a tuple with the Backups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackups

`func (o *Dashboard) SetBackups(v DashboardBackups)`

SetBackups sets Backups field to given value.

### HasBackups

`func (o *Dashboard) HasBackups() bool`

HasBackups returns a boolean if a field has been set.

### GetActivity

`func (o *Dashboard) GetActivity() []DashboardActivity`

GetActivity returns the Activity field if non-nil, zero value otherwise.

### GetActivityOk

`func (o *Dashboard) GetActivityOk() (*[]DashboardActivity, bool)`

GetActivityOk returns a tuple with the Activity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivity

`func (o *Dashboard) SetActivity(v []DashboardActivity)`

SetActivity sets Activity field to given value.

### HasActivity

`func (o *Dashboard) HasActivity() bool`

HasActivity returns a boolean if a field has been set.

### GetLogStats

`func (o *Dashboard) GetLogStats() DashboardLogStats`

GetLogStats returns the LogStats field if non-nil, zero value otherwise.

### GetLogStatsOk

`func (o *Dashboard) GetLogStatsOk() (*DashboardLogStats, bool)`

GetLogStatsOk returns a tuple with the LogStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogStats

`func (o *Dashboard) SetLogStats(v DashboardLogStats)`

SetLogStats sets LogStats field to given value.

### HasLogStats

`func (o *Dashboard) HasLogStats() bool`

HasLogStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


