# GetNetworkFirewallRules200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rules** | [**GetNetworkFirewallRules200ResponseAllOfRulesInner[]**](GetNetworkFirewallRules200ResponseAllOfRulesInner.md) |  | [optional] 
**Meta** | [**MetaMeta**](MetaMeta.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GetNetworkFirewallRules200Response = Initialize-PSOpenAPIToolsGetNetworkFirewallRules200Response  -Rules null `
 -Meta null
```

- Convert the resource to JSON
```powershell
$GetNetworkFirewallRules200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

