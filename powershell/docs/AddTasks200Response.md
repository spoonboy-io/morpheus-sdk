# AddTasks200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Task** | [**AddTasks200ResponseAllOfTask**](AddTasks200ResponseAllOfTask.md) |  | [optional] 
**Success** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AddTasks200Response = Initialize-PSOpenAPIToolsAddTasks200Response  -Task null `
 -Success null
```

- Convert the resource to JSON
```powershell
$AddTasks200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

