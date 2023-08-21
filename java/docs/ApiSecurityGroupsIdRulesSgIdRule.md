

# ApiSecurityGroupsIdRulesSgIdRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | A name for the rule |  [optional]
**direction** | [**DirectionEnum**](#DirectionEnum) | Either &#x60;ingress&#x60; or &#x60;egress&#x60; |  [optional]
**sourceType** | [**SourceTypeEnum**](#SourceTypeEnum) | Either &#x60;cidr&#x60;, &#x60;group&#x60;, &#x60;tier&#x60;, &#x60;all&#x60;. |  [optional]
**source** | **String** | CIDR representing the source IP(s) which should receive access. Required for &#x60;sourceType&#x60;&#x3D;cidr |  [optional]
**sourceGroup** | [**ApiSecurityGroupsIdRulesRuleSourceGroup**](ApiSecurityGroupsIdRulesRuleSourceGroup.md) |  |  [optional]
**sourceTier** | [**ApiSecurityGroupsIdRulesRuleSourceTier**](ApiSecurityGroupsIdRulesRuleSourceTier.md) |  |  [optional]
**portRange** | **String** | Either a single value (i.e. 55) or a port range (i.e. 1-65535) for which to open access to the source. Required if customRule is true, otherwise, ignored.  |  [optional]
**protocol** | [**ProtocolEnum**](#ProtocolEnum) | Either tcp, udp, icmp. Required if customRule is true, otherwise, ignored. | 
**destinationType** | [**DestinationTypeEnum**](#DestinationTypeEnum) | Either cidr, group, tier, instance. |  [optional]
**destination** | **String** | CIDR representing the destination IP(s) which should receive access. Required for &#x60;destinationType&#x60;&#x3D;cidr. |  [optional]
**destinationGroup** | [**ApiSecurityGroupsIdRulesRuleDestinationGroup**](ApiSecurityGroupsIdRulesRuleDestinationGroup.md) |  |  [optional]
**destinationTier** | [**ApiSecurityGroupsIdRulesRuleDestinationTier**](ApiSecurityGroupsIdRulesRuleDestinationTier.md) |  |  [optional]
**ruleType** | **String** | Either &#x60;customRule&#x60; or an &#x60;instance type&#x60; code. | 
**policy** | [**PolicyEnum**](#PolicyEnum) | Either &#x60;accept&#x60; or &#x60;deny&#x60;. |  [optional]
**instanceTypeId** | **Long** | The id of an Instance Type. If specified, the source CIDR will have access to all ports exposed by the particular instance in the cloud, app, or instance. Required if customRule is false, otherwise ignored.  |  [optional]



## Enum: DirectionEnum

Name | Value
---- | -----
INGRESS | &quot;ingress&quot;
EGRESS | &quot;egress&quot;



## Enum: SourceTypeEnum

Name | Value
---- | -----
CIDR | &quot;cidr&quot;
GROUP | &quot;group&quot;
TIER | &quot;tier&quot;
ALL | &quot;all&quot;



## Enum: ProtocolEnum

Name | Value
---- | -----
TCP | &quot;tcp&quot;
UDP | &quot;udp&quot;
ICMP | &quot;icmp&quot;



## Enum: DestinationTypeEnum

Name | Value
---- | -----
CIDR | &quot;cidr&quot;
GROUP | &quot;group&quot;
TIER | &quot;tier&quot;
INSTANCE | &quot;instance&quot;



## Enum: PolicyEnum

Name | Value
---- | -----
ACCEPT | &quot;accept&quot;
DENY | &quot;deny&quot;



