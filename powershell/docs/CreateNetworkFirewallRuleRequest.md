# CreateNetworkFirewallRuleRequest
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rule** | [**NetworkFirewallRuleCreate**](NetworkFirewallRuleCreate.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$CreateNetworkFirewallRuleRequest = Initialize-PSOpenAPIToolsCreateNetworkFirewallRuleRequest  -Rule null
```

- Convert the resource to JSON
```powershell
$CreateNetworkFirewallRuleRequest | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

