# ExecuteTasks200ResponseAllOf
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Job** | [**UpdateBlueprintPermissionsRequestResourcePermissionSitesInner**](UpdateBlueprintPermissionsRequestResourcePermissionSitesInner.md) |  | [optional] 
**JobExecution** | [**ExecuteTasks200ResponseAllOfJobExecution**](ExecuteTasks200ResponseAllOfJobExecution.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ExecuteTasks200ResponseAllOf = Initialize-PSOpenAPIToolsExecuteTasks200ResponseAllOf  -Job null `
 -JobExecution null
```

- Convert the resource to JSON
```powershell
$ExecuteTasks200ResponseAllOf | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

