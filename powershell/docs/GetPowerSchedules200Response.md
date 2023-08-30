# GetPowerSchedules200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instances** | [**GetPowerSchedules200ResponseAllOfInstancesInner[]**](GetPowerSchedules200ResponseAllOfInstancesInner.md) |  | [optional] 
**Servers** | [**GetPowerSchedules200ResponseAllOfInstancesInner[]**](GetPowerSchedules200ResponseAllOfInstancesInner.md) |  | [optional] 
**Schedule** | [**PowerSchedule**](PowerSchedule.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GetPowerSchedules200Response = Initialize-PSOpenAPIToolsGetPowerSchedules200Response  -Instances null `
 -Servers null `
 -Schedule null
```

- Convert the resource to JSON
```powershell
$GetPowerSchedules200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

