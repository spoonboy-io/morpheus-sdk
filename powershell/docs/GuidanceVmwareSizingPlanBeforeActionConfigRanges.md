# GuidanceVmwareSizingPlanBeforeActionConfigRanges
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinStorage** | **String** |  | [optional] 
**MaxStorage** | **String** |  | [optional] 
**MinMemory** | **String** |  | [optional] 
**MaxMemory** | **String** |  | [optional] 
**MinCores** | **String** |  | [optional] 
**MaxCores** | **String** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GuidanceVmwareSizingPlanBeforeActionConfigRanges = Initialize-PSOpenAPIToolsGuidanceVmwareSizingPlanBeforeActionConfigRanges  -MinStorage null `
 -MaxStorage null `
 -MinMemory null `
 -MaxMemory null `
 -MinCores null `
 -MaxCores null
```

- Convert the resource to JSON
```powershell
$GuidanceVmwareSizingPlanBeforeActionConfigRanges | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

