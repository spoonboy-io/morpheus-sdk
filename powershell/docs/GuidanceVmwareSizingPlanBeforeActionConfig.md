# GuidanceVmwareSizingPlanBeforeActionConfig
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StorageSizeType** | **String** |  | [optional] 
**MemorySizeType** | **String** |  | [optional] 
**Ranges** | [**GuidanceVmwareSizingPlanBeforeActionConfigRanges**](GuidanceVmwareSizingPlanBeforeActionConfigRanges.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GuidanceVmwareSizingPlanBeforeActionConfig = Initialize-PSOpenAPIToolsGuidanceVmwareSizingPlanBeforeActionConfig  -StorageSizeType null `
 -MemorySizeType null `
 -Ranges null
```

- Convert the resource to JSON
```powershell
$GuidanceVmwareSizingPlanBeforeActionConfig | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

