# AddPowerSchedules200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schedule** | [**AddPowerSchedules200ResponseAllOfSchedule**](AddPowerSchedules200ResponseAllOfSchedule.md) |  | [optional] 
**Success** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AddPowerSchedules200Response = Initialize-PSOpenAPIToolsAddPowerSchedules200Response  -Schedule null `
 -Success null
```

- Convert the resource to JSON
```powershell
$AddPowerSchedules200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

