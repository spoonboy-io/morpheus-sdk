# GetNetworkRouter200ResponseNetworkRouterType
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Code** | **String** |  | [optional] 
**Name** | **String** |  | [optional] 
**Description** | **String** |  | [optional] 
**Enabled** | **Boolean** |  | [optional] 
**Creatable** | **Boolean** |  | [optional] 
**Selectable** | **Boolean** |  | [optional] 
**HasFirewall** | **Boolean** |  | [optional] 
**HasDhcp** | **Boolean** |  | [optional] 
**HasRouting** | **Boolean** |  | [optional] 
**HasNat** | **Boolean** |  | [optional] 
**HasNetworkServer** | **Boolean** |  | [optional] 
**HasFirewallGroups** | **Boolean** |  | [optional] 
**HasSecurityGroupPriority** | **Boolean** |  | [optional] 
**OptionTypes** | [**OptionType[]**](OptionType.md) |  | [optional] 
**RuleOptionTypes** | [**OptionType[]**](OptionType.md) |  | [optional] 
**FirewallGroupOptionTypes** | [**OptionType[]**](OptionType.md) |  | [optional] 
**NatOptionTypes** | [**OptionType[]**](OptionType.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GetNetworkRouter200ResponseNetworkRouterType = Initialize-PSOpenAPIToolsGetNetworkRouter200ResponseNetworkRouterType  -Id null `
 -Code null `
 -Name null `
 -Description null `
 -Enabled null `
 -Creatable null `
 -Selectable null `
 -HasFirewall null `
 -HasDhcp null `
 -HasRouting null `
 -HasNat null `
 -HasNetworkServer null `
 -HasFirewallGroups null `
 -HasSecurityGroupPriority null `
 -OptionTypes null `
 -RuleOptionTypes null `
 -FirewallGroupOptionTypes null `
 -NatOptionTypes null
```

- Convert the resource to JSON
```powershell
$GetNetworkRouter200ResponseNetworkRouterType | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

