# HealthRabbitQueues
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** |  | [optional] 
**Count** | **Int64** |  | [optional] 
**Status** | **String** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$HealthRabbitQueues = Initialize-PSOpenAPIToolsHealthRabbitQueues  -Name null `
 -Count null `
 -Status null
```

- Convert the resource to JSON
```powershell
$HealthRabbitQueues | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

