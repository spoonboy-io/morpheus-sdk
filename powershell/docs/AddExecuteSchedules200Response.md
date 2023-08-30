# AddExecuteSchedules200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schedule** | [**ExecuteSchedule**](ExecuteSchedule.md) |  | [optional] 
**Success** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AddExecuteSchedules200Response = Initialize-PSOpenAPIToolsAddExecuteSchedules200Response  -Schedule null `
 -Success null
```

- Convert the resource to JSON
```powershell
$AddExecuteSchedules200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

