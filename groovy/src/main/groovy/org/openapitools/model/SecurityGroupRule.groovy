package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.InlineResponse20082LoadBalancerInstanceSslCert;

@Canonical
class SecurityGroupRule {
    
    Long id
    
    String name
    
    String ruleType
    
    Boolean customRule
    
    String instanceTypeId
    
    String direction
    
    String policy
    
    String sourceType
    
    String source
    
    InlineResponse20082LoadBalancerInstanceSslCert sourceGroup
    
    InlineResponse20082LoadBalancerInstanceSslCert sourceTier
    
    String portRange
    
    String protocol
    
    String destinationType
    
    String destination
    
    InlineResponse20082LoadBalancerInstanceSslCert destinationGroup
    
    InlineResponse20082LoadBalancerInstanceSslCert destinationTier
    
    String externalId
    
    String enabled
}
