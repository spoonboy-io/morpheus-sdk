package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.InlineResponse20095NetworkRouterFirewallRules;

@Canonical
class InlineResponse20095NetworkRouterFirewallRuleGroups {
    
    Long id
    
    String name
    
    String description
    
    String externalId
    
    String iacId
    
    String zone
    
    String zonePool
    
    String status
    
    Long priority
    
    String groupLayer
    
    List<InlineResponse20095NetworkRouterFirewallRules> rules = new ArrayList<InlineResponse20095NetworkRouterFirewallRules>()
}
