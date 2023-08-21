package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.InlineResponse20082LoadBalancerInstanceSslCert;

@Canonical
class SecurityGroupLocations {
    
    Long id
    
    String name
    
    String description
    
    String externalId
    
    String iacId
    
    InlineResponse20082LoadBalancerInstanceSslCert zone
    
    InlineResponse20082LoadBalancerInstanceSslCert zonePool
    
    String status
    
    String priority
    
    String groupLayer
}
