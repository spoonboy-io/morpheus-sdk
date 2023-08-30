# CreateNetworkDomainRequest
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkDomain** | [**NetworkDomainCreate**](NetworkDomainCreate.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$CreateNetworkDomainRequest = Initialize-PSOpenAPIToolsCreateNetworkDomainRequest  -NetworkDomain null
```

- Convert the resource to JSON
```powershell
$CreateNetworkDomainRequest | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

