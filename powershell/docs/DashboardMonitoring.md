# DashboardMonitoring
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvgHealth** | **Decimal** |  | [optional] 
**AvgResponseTime** | **Decimal** |  | [optional] 
**WarningApps** | **Decimal** |  | [optional] 
**WarningChecks** | **Decimal** |  | [optional] 
**FailApps** | **Decimal** |  | [optional] 
**TotalApps** | **Decimal** |  | [optional] 
**FailChecks** | **Decimal** |  | [optional] 
**SuccessApps** | **Decimal** |  | [optional] 
**MutedApps** | **Decimal** |  | [optional] 
**SuccessChecks** | **Decimal** |  | [optional] 
**TotalChecks** | **Decimal** |  | [optional] 
**MutedChecks** | **Decimal** |  | [optional] 
**ResponseTimes** | **Decimal[]** |  | [optional] 
**AllSuccess** | **Boolean** |  | [optional] 
**OpenIncidents** | **Decimal** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$DashboardMonitoring = Initialize-PSOpenAPIToolsDashboardMonitoring  -AvgHealth null `
 -AvgResponseTime null `
 -WarningApps null `
 -WarningChecks null `
 -FailApps null `
 -TotalApps null `
 -FailChecks null `
 -SuccessApps null `
 -MutedApps null `
 -SuccessChecks null `
 -TotalChecks null `
 -MutedChecks null `
 -ResponseTimes null `
 -AllSuccess null `
 -OpenIncidents null
```

- Convert the resource to JSON
```powershell
$DashboardMonitoring | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

