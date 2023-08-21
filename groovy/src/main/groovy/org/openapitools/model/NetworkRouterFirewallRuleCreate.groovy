package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class NetworkRouterFirewallRuleCreate {
    /* Firewall rule name */
    String name
    /* Can be used to enable / disable the rule (true, false). Default is on */
    Boolean enabled = true
    /* Priority for rule */
    Long priority
}
