# AddAppInstanceRequest
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceId** | **Int64** | The ID of the instance being added | 
**TierName** | **String** | The Name of the Tier | 

## Examples

- Prepare the resource
```powershell
$AddAppInstanceRequest = Initialize-PSOpenAPIToolsAddAppInstanceRequest  -InstanceId null `
 -TierName null
```

- Convert the resource to JSON
```powershell
$AddAppInstanceRequest | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

