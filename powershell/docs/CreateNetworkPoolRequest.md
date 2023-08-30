# CreateNetworkPoolRequest
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkPool** | [**NetworkPoolCreate**](NetworkPoolCreate.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$CreateNetworkPoolRequest = Initialize-PSOpenAPIToolsCreateNetworkPoolRequest  -NetworkPool null
```

- Convert the resource to JSON
```powershell
$CreateNetworkPoolRequest | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

