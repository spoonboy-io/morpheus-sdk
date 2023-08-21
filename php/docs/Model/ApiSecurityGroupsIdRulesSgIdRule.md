# # ApiSecurityGroupsIdRulesSgIdRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **string** | A name for the rule | [optional]
**direction** | **string** | Either &#x60;ingress&#x60; or &#x60;egress&#x60; | [optional] [default to 'ingress']
**source_type** | **string** | Either &#x60;cidr&#x60;, &#x60;group&#x60;, &#x60;tier&#x60;, &#x60;all&#x60;. | [optional] [default to 'cidr']
**source** | **string** | CIDR representing the source IP(s) which should receive access. Required for &#x60;sourceType&#x60;&#x3D;cidr | [optional]
**source_group** | [**\OpenAPI\Client\Model\ApiSecurityGroupsIdRulesRuleSourceGroup**](ApiSecurityGroupsIdRulesRuleSourceGroup.md) |  | [optional]
**source_tier** | [**\OpenAPI\Client\Model\ApiSecurityGroupsIdRulesRuleSourceTier**](ApiSecurityGroupsIdRulesRuleSourceTier.md) |  | [optional]
**port_range** | **string** | Either a single value (i.e. 55) or a port range (i.e. 1-65535) for which to open access to the source. Required if customRule is true, otherwise, ignored. | [optional]
**protocol** | **string** | Either tcp, udp, icmp. Required if customRule is true, otherwise, ignored. |
**destination_type** | **string** | Either cidr, group, tier, instance. | [optional] [default to 'cidr']
**destination** | **string** | CIDR representing the destination IP(s) which should receive access. Required for &#x60;destinationType&#x60;&#x3D;cidr. | [optional]
**destination_group** | [**\OpenAPI\Client\Model\ApiSecurityGroupsIdRulesRuleDestinationGroup**](ApiSecurityGroupsIdRulesRuleDestinationGroup.md) |  | [optional]
**destination_tier** | [**\OpenAPI\Client\Model\ApiSecurityGroupsIdRulesRuleDestinationTier**](ApiSecurityGroupsIdRulesRuleDestinationTier.md) |  | [optional]
**rule_type** | **string** | Either &#x60;customRule&#x60; or an &#x60;instance type&#x60; code. | [default to 'customRule']
**policy** | **string** | Either &#x60;accept&#x60; or &#x60;deny&#x60;. | [optional]
**instance_type_id** | **int** | The id of an Instance Type. If specified, the source CIDR will have access to all ports exposed by the particular instance in the cloud, app, or instance. Required if customRule is false, otherwise ignored. | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
