# GetAllNetworkFloatingIps200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkFloatingIps** | [**GetAllNetworkFloatingIps200ResponseAllOfNetworkFloatingIpsInner[]**](GetAllNetworkFloatingIps200ResponseAllOfNetworkFloatingIpsInner.md) |  | [optional] 
**Meta** | [**MetaMeta**](MetaMeta.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GetAllNetworkFloatingIps200Response = Initialize-PSOpenAPIToolsGetAllNetworkFloatingIps200Response  -NetworkFloatingIps null `
 -Meta null
```

- Convert the resource to JSON
```powershell
$GetAllNetworkFloatingIps200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

