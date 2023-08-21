package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.NetworkFirewallRuleCreateConfig;
import org.openapitools.model.NetworkFirewallRuleCreateRuleGroup;
import org.openapitools.model.NetworkFirewallRuleCreateSources;

@Canonical
class NetworkFirewallRuleCreate {
    
    NetworkFirewallRuleCreateRuleGroup ruleGroup
    /* Network firewall rule name */
    String name
    /* Network firewall rule description */
    String description
    /* Use this to set enabled state */
    Boolean enabled
    /* Network firewall rule priority */
    String priority
    
    String direction
    
    NetworkFirewallRuleCreateSources sources
    
    NetworkFirewallRuleCreateSources destinations
    
    NetworkFirewallRuleCreateConfig config
    
    NetworkFirewallRuleCreateSources scopes
    
    String policy
}
