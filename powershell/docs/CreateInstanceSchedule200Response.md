# CreateInstanceSchedule200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceSchedule** | [**InstanceSchedule**](InstanceSchedule.md) |  | [optional] 
**Success** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$CreateInstanceSchedule200Response = Initialize-PSOpenAPIToolsCreateInstanceSchedule200Response  -InstanceSchedule null `
 -Success null
```

- Convert the resource to JSON
```powershell
$CreateInstanceSchedule200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

