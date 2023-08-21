package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class SecurityGroupRules {
    
    Long id
    
    String name
    
    String ruleType
    
    Boolean customRule
    
    String instanceTypeId
    
    String direction
    
    String policy
    
    String sourceType
    
    String source
    
    String sourceGroup
    
    String sourceTier
    
    String portRange
    
    String protocol
    
    String destinationType
    
    String destination
    
    String destinationGroup
    
    String destinationTier
    
    String externalId
    
    String enabled
}
