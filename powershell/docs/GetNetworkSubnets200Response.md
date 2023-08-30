# GetNetworkSubnets200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subnets** | [**Subnet[]**](Subnet.md) |  | [optional] 
**Meta** | [**MetaMeta**](MetaMeta.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GetNetworkSubnets200Response = Initialize-PSOpenAPIToolsGetNetworkSubnets200Response  -Subnets null `
 -Meta null
```

- Convert the resource to JSON
```powershell
$GetNetworkSubnets200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

