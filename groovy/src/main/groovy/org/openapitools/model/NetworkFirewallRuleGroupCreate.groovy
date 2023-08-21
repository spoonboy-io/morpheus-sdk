package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class NetworkFirewallRuleGroupCreate {
    /* Network firewall rule group name */
    String name
    /* Network firewall rule group description */
    String description
    /* Network firewall rule group priority */
    Long priority
    /* Use SecurityPolicy */
    String externalType
    
    String groupLayer
}
