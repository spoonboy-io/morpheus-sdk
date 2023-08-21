package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.ApiSecurityGroupsIdRulesRuleDestinationGroup;
import org.openapitools.model.ApiSecurityGroupsIdRulesRuleDestinationTier;
import org.openapitools.model.ApiSecurityGroupsIdRulesRuleSourceGroup;
import org.openapitools.model.ApiSecurityGroupsIdRulesRuleSourceTier;

@Canonical
class ApiSecurityGroupsIdRulesSgIdRule {
    /* A name for the rule */
    String name
    /* Either `ingress` or `egress` */
    String direction = DirectionEnum.INGRESS
    /* Either `cidr`, `group`, `tier`, `all`. */
    String sourceType = SourceTypeEnum.CIDR
    /* CIDR representing the source IP(s) which should receive access. Required for `sourceType`=cidr */
    String source
    
    ApiSecurityGroupsIdRulesRuleSourceGroup sourceGroup
    
    ApiSecurityGroupsIdRulesRuleSourceTier sourceTier
    /* Either a single value (i.e. 55) or a port range (i.e. 1-65535) for which to open access to the source. Required if customRule is true, otherwise, ignored.  */
    String portRange
    /* Either tcp, udp, icmp. Required if customRule is true, otherwise, ignored. */
    String protocol
    /* Either cidr, group, tier, instance. */
    String destinationType = DestinationTypeEnum.CIDR
    /* CIDR representing the destination IP(s) which should receive access. Required for `destinationType`=cidr. */
    String destination
    
    ApiSecurityGroupsIdRulesRuleDestinationGroup destinationGroup
    
    ApiSecurityGroupsIdRulesRuleDestinationTier destinationTier
    /* Either `customRule` or an `instance type` code. */
    String ruleType = "customRule"
    /* Either `accept` or `deny`. */
    String policy
    /* The id of an Instance Type. If specified, the source CIDR will have access to all ports exposed by the particular instance in the cloud, app, or instance. Required if customRule is false, otherwise ignored.  */
    Long instanceTypeId
}
