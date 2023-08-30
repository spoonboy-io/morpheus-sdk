# CreateResourcePoolGroupRequest
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourcePoolGroup** | [**ResourcePoolGroupsCreateInput**](ResourcePoolGroupsCreateInput.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$CreateResourcePoolGroupRequest = Initialize-PSOpenAPIToolsCreateResourcePoolGroupRequest  -ResourcePoolGroup null
```

- Convert the resource to JSON
```powershell
$CreateResourcePoolGroupRequest | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

