# GuidanceStats
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **Int64** |  | [optional] 
**Savings** | [**GuidanceVmwareSizingSavings**](GuidanceVmwareSizingSavings.md) |  | [optional] 
**Severity** | [**GuidanceStatsSeverity**](GuidanceStatsSeverity.md) |  | [optional] 
**Type** | [**GuidanceStatsType**](GuidanceStatsType.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GuidanceStats = Initialize-PSOpenAPIToolsGuidanceStats  -Total null `
 -Savings null `
 -Severity null `
 -Type null
```

- Convert the resource to JSON
```powershell
$GuidanceStats | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

