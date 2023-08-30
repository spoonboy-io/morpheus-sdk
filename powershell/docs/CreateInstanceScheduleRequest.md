# CreateInstanceScheduleRequest
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceSchedule** | [**InstanceScheduleCreate**](InstanceScheduleCreate.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$CreateInstanceScheduleRequest = Initialize-PSOpenAPIToolsCreateInstanceScheduleRequest  -InstanceSchedule null
```

- Convert the resource to JSON
```powershell
$CreateInstanceScheduleRequest | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

