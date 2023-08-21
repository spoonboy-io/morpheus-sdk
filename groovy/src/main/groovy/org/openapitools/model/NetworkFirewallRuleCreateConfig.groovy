package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;

@Canonical
class NetworkFirewallRuleCreateConfig {
    
    List<String> application = new ArrayList<String>()
    
    List<String> profile = new ArrayList<String>()
}
