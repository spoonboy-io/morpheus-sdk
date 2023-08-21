# GuidanceStatsType
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Size** | **Int64** |  | [optional] 
**Shutdown** | **Int64** |  | [optional] 
**Move** | **Int64** |  | [optional] 
**Schedule** | **Int64** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GuidanceStatsType = Initialize-PSOpenAPIToolsGuidanceStatsType  -Size null `
 -Shutdown null `
 -Move null `
 -Schedule null
```

- Convert the resource to JSON
```powershell
$GuidanceStatsType | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

