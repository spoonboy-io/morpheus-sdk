# DashboardLogStats
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **Boolean** |  | [optional] 
**VarData** | [**DashboardLogStatsDataInner[]**](DashboardLogStatsDataInner.md) |  | [optional] 
**StartMs** | **Int64** |  | [optional] 
**Earliest** | **Int64** |  | [optional] 
**EndMs** | **Int64** |  | [optional] 
**Interval** | **Int64** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$DashboardLogStats = Initialize-PSOpenAPIToolsDashboardLogStats  -Success null `
 -VarData null `
 -StartMs null `
 -Earliest null `
 -EndMs null `
 -Interval null
```

- Convert the resource to JSON
```powershell
$DashboardLogStats | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

