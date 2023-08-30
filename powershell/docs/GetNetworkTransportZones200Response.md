# GetNetworkTransportZones200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkScopes** | [**GetNetworkTransportZones200ResponseAllOfNetworkScopesInner[]**](GetNetworkTransportZones200ResponseAllOfNetworkScopesInner.md) |  | [optional] 
**Meta** | [**MetaMeta**](MetaMeta.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GetNetworkTransportZones200Response = Initialize-PSOpenAPIToolsGetNetworkTransportZones200Response  -NetworkScopes null `
 -Meta null
```

- Convert the resource to JSON
```powershell
$GetNetworkTransportZones200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

