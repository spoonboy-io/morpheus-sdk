# CreateNetworkServerGroupRequest
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | [**NetworkServerGroupCreate**](NetworkServerGroupCreate.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$CreateNetworkServerGroupRequest = Initialize-PSOpenAPIToolsCreateNetworkServerGroupRequest  -Group null
```

- Convert the resource to JSON
```powershell
$CreateNetworkServerGroupRequest | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

