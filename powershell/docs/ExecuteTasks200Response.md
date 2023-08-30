# ExecuteTasks200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Job** | [**UpdateBlueprintPermissionsRequestResourcePermissionSitesInner**](UpdateBlueprintPermissionsRequestResourcePermissionSitesInner.md) |  | [optional] 
**JobExecution** | [**ExecuteTasks200ResponseAllOfJobExecution**](ExecuteTasks200ResponseAllOfJobExecution.md) |  | [optional] 
**Success** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ExecuteTasks200Response = Initialize-PSOpenAPIToolsExecuteTasks200Response  -Job null `
 -JobExecution null `
 -Success null
```

- Convert the resource to JSON
```powershell
$ExecuteTasks200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

