package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.InlineResponse20040AppDeployInstance;

@Canonical
class InlineResponse20095NetworkRouterFirewallRules {
    
    Long id
    
    String name
    
    String code
    
    Boolean enabled
    
    String groupName
    
    String direction
    
    String ruleType
    
    String policy
    
    List<String> source = new ArrayList<String>()
    
    String sourceType
    
    List<String> destination = new ArrayList<String>()
    
    String destinationType
    
    List<String> profiles = new ArrayList<String>()
    
    String protocol
    
    String application
    
    String applicationType
    
    String portRange
    
    String sourcePortRange
    
    String sourceGroup
    
    String sourceTier
    
    List<InlineResponse20040AppDeployInstance> applications = new ArrayList<InlineResponse20040AppDeployInstance>()
}
