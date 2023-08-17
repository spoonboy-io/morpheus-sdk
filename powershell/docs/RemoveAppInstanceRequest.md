# RemoveAppInstanceRequest
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceId** | **Int64** | The ID of the instance being removed | 

## Examples

- Prepare the resource
```powershell
$RemoveAppInstanceRequest = Initialize-PSOpenAPIToolsRemoveAppInstanceRequest  -InstanceId null
```

- Convert the resource to JSON
```powershell
$RemoveAppInstanceRequest | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

