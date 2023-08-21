package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.InlineResponse20040AppDeployInstance;
import org.openapitools.model.InlineResponse20079LoadBalancerMonitorLoadBalancerType;

@Canonical
class ProvisioningLicense {
    
    Long id
    
    String name
    
    String description
    
    InlineResponse20079LoadBalancerMonitorLoadBalancerType licenseType
    
    String licenseKey
    
    String orgName
    
    String fullName
    
    String licenseVersion
    
    Long copies
    
    Long reservationCount
    
    List<Object> tenants = new ArrayList<Object>()
    
    List<InlineResponse20040AppDeployInstance> virtualImages = new ArrayList<InlineResponse20040AppDeployInstance>()
    
    InlineResponse20040AppDeployInstance account
}
