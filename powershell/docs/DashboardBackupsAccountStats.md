# DashboardBackupsAccountStats
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalSizeByDay** | **Decimal[]** |  | [optional] 
**TotalSizeByDay7Days** | **Decimal[]** |  | [optional] 
**FormattedTotalSize** | [**DashboardBackupsAccountStatsFormattedTotalSize**](DashboardBackupsAccountStatsFormattedTotalSize.md) |  | [optional] 
**BackupCount** | **Decimal** |  | [optional] 
**TotalSize** | **Decimal** |  | [optional] 
**Success** | **Decimal** |  | [optional] 
**Failed** | **Decimal** |  | [optional] 
**TotalCompleted** | **Decimal** |  | [optional] 
**LastSevenDays** | [**DashboardBackupsAccountStatsLastSevenDays**](DashboardBackupsAccountStatsLastSevenDays.md) |  | [optional] 
**AvgSize** | **Decimal** |  | [optional] 
**FailedRate** | **Decimal** |  | [optional] 
**SuccessRate** | **Decimal** |  | [optional] 
**NextFireTotal** | **Decimal** |  | [optional] 
**BackupDayCount** | **Decimal[]** |  | [optional] 
**BackupDayCountTotal** | **Decimal** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$DashboardBackupsAccountStats = Initialize-PSOpenAPIToolsDashboardBackupsAccountStats  -TotalSizeByDay null `
 -TotalSizeByDay7Days null `
 -FormattedTotalSize null `
 -BackupCount null `
 -TotalSize null `
 -Success null `
 -Failed null `
 -TotalCompleted null `
 -LastSevenDays null `
 -AvgSize null `
 -FailedRate null `
 -SuccessRate null `
 -NextFireTotal null `
 -BackupDayCount null `
 -BackupDayCountTotal null
```

- Convert the resource to JSON
```powershell
$DashboardBackupsAccountStats | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

