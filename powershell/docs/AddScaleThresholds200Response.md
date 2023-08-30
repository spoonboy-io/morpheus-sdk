# AddScaleThresholds200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScaleThreshold** | [**ScaleThreshold**](ScaleThreshold.md) |  | [optional] 
**Success** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AddScaleThresholds200Response = Initialize-PSOpenAPIToolsAddScaleThresholds200Response  -ScaleThreshold null `
 -Success null
```

- Convert the resource to JSON
```powershell
$AddScaleThresholds200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

