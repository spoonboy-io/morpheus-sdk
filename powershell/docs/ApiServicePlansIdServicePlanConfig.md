# ApiServicePlansIdServicePlanConfig
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StorageSizeType** | **String** | Specifies range min / max storage multiplier | [optional] [default to "gb"]
**MemorySizeType** | **String** | Specifies range min / max memory multiplier | [optional] [default to "mb"]
**Ranges** | [**ApiServicePlansIdServicePlanConfigRanges**](ApiServicePlansIdServicePlanConfigRanges.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ApiServicePlansIdServicePlanConfig = Initialize-PSOpenAPIToolsApiServicePlansIdServicePlanConfig  -StorageSizeType null `
 -MemorySizeType null `
 -Ranges null
```

- Convert the resource to JSON
```powershell
$ApiServicePlansIdServicePlanConfig | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

