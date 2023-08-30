# AddSecurityGroupRulesRequestRule


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**protocol** | **str** | Either tcp, udp, icmp. Required if customRule is true, otherwise, ignored. | 
**rule_type** | **str** | Either &#x60;customRule&#x60; or an &#x60;instance type&#x60; code. | defaults to "customRule"
**name** | **str** | A name for the rule | [optional] 
**direction** | **str** | Either &#x60;ingress&#x60; or &#x60;egress&#x60; | [optional]  if omitted the server will use the default value of "ingress"
**source_type** | **str** | Either &#x60;cidr&#x60;, &#x60;group&#x60;, &#x60;tier&#x60;, &#x60;all&#x60; | [optional]  if omitted the server will use the default value of "cidr"
**source** | **str** | CIDR representing the source IP(s) which should receive access. Required for &#x60;sourceType&#x60;&#x3D;cidr | [optional] 
**source_group** | [**AddSecurityGroupRulesRequestRuleSourceGroup**](AddSecurityGroupRulesRequestRuleSourceGroup.md) |  | [optional] 
**source_tier** | [**AddSecurityGroupRulesRequestRuleSourceTier**](AddSecurityGroupRulesRequestRuleSourceTier.md) |  | [optional] 
**port_range** | **str** | Either a single value (i.e. 55) or a port range (i.e. 1-65535) for which to open access to the source. Required if customRule is true, otherwise, ignored.  | [optional] 
**destination_type** | **str** | Either cidr, group, tier, instance. | [optional]  if omitted the server will use the default value of "cidr"
**destination** | **str** | CIDR representing the destination IP(s) which should receive access. Required for &#x60;destinationType&#x60;&#x3D;cidr. | [optional] 
**destination_group** | [**AddSecurityGroupRulesRequestRuleDestinationGroup**](AddSecurityGroupRulesRequestRuleDestinationGroup.md) |  | [optional] 
**destination_tier** | [**AddSecurityGroupRulesRequestRuleDestinationTier**](AddSecurityGroupRulesRequestRuleDestinationTier.md) |  | [optional] 
**policy** | **str** | Either &#x60;accept&#x60; or &#x60;deny&#x60;. | [optional] 
**instance_type_id** | **int** | The id of an Instance Type. If specified, the source CIDR will have access to all ports exposed by the particular instance in the cloud, app, or instance. Required if customRule is false, otherwise ignored.  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


