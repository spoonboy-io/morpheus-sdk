# GetInstanceThreshold200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instance** | [**ListDeploys200ResponseAllOfAppDeploysInnerInstance**](ListDeploys200ResponseAllOfAppDeploysInnerInstance.md) |  | [optional] 
**InstanceThreshold** | [**InstanceThreshold**](InstanceThreshold.md) |  | [optional] 
**InstanceSchedules** | [**InstanceSchedule[]**](InstanceSchedule.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GetInstanceThreshold200Response = Initialize-PSOpenAPIToolsGetInstanceThreshold200Response  -Instance null `
 -InstanceThreshold null `
 -InstanceSchedules null
```

- Convert the resource to JSON
```powershell
$GetInstanceThreshold200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

