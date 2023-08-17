# AddExecuteSchedules200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **Boolean** |  | [optional] 
**Schedule** | [**ExecuteSchedule**](ExecuteSchedule.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AddExecuteSchedules200Response = Initialize-PSOpenAPIToolsAddExecuteSchedules200Response  -Success null `
 -Schedule null
```

- Convert the resource to JSON
```powershell
$AddExecuteSchedules200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

