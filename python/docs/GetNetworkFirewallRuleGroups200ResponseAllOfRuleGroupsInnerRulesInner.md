# GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInnerRulesInner


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **int** |  | [optional] 
**direction** | **str** |  | [optional] 
**source_type** | **str** |  | [optional] 
**destination_type** | **str** |  | [optional] 
**name** | **str** |  | [optional] 
**policy** | **str** |  | [optional] 
**priority** | **int** |  | [optional] 
**enabled** | **bool** |  | [optional] 
**rule_group** | [**ListDeploys200ResponseAllOfAppDeploysInnerInstance**](ListDeploys200ResponseAllOfAppDeploysInnerInstance.md) |  | [optional] 
**group_name** | **str** |  | [optional] 
**config** | **{str: (bool, date, datetime, dict, float, int, list, str, none_type)}** |  | [optional] 
**sources** | [**[GetNetworkFirewallRules200ResponseAllOfRulesInnerSourcesInner]**](GetNetworkFirewallRules200ResponseAllOfRulesInnerSourcesInner.md) |  | [optional] 
**destinations** | [**[GetNetworkFirewallRules200ResponseAllOfRulesInnerSourcesInner]**](GetNetworkFirewallRules200ResponseAllOfRulesInnerSourcesInner.md) |  | [optional] 
**applications** | [**[GetNetworkFirewallRules200ResponseAllOfRulesInnerSourcesInner]**](GetNetworkFirewallRules200ResponseAllOfRulesInnerSourcesInner.md) |  | [optional] 
**scopes** | [**[GetNetworkFirewallRules200ResponseAllOfRulesInnerSourcesInner]**](GetNetworkFirewallRules200ResponseAllOfRulesInnerSourcesInner.md) |  | [optional] 
**profiles** | [**[GetNetworkFirewallRules200ResponseAllOfRulesInnerSourcesInner]**](GetNetworkFirewallRules200ResponseAllOfRulesInnerSourcesInner.md) |  | [optional] 
**applied_targets** | **[{str: (bool, date, datetime, dict, float, int, list, str, none_type)}]** |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


