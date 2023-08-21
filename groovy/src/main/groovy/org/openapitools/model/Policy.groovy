package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.InlineResponse20040AppDeployInstance;
import org.openapitools.model.InlineResponse20079LoadBalancerMonitorLoadBalancerType;
import org.openapitools.model.InlineResponse20082LoadBalancerInstanceSslCert;
import org.openapitools.model.InlineResponse20083LoadBalancerNodeCreatedBy;
import org.openapitools.model.OneOfobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobject;
import org.openapitools.model.PolicyRole;

@Canonical
class Policy {
    
    Long id
    
    String name
    
    String description
    
    InlineResponse20079LoadBalancerMonitorLoadBalancerType policyType
    
    InlineResponse20082LoadBalancerInstanceSslCert zone
    
    InlineResponse20082LoadBalancerInstanceSslCert site
    
    InlineResponse20083LoadBalancerNodeCreatedBy user
    
    PolicyRole role
    
    String refType
    
    String refId
    
    Boolean eachUser
    
    OneOfobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobject config = null
    
    Boolean enabled
    
    InlineResponse20040AppDeployInstance owner
    
    List<InlineResponse20040AppDeployInstance> accounts = new ArrayList<InlineResponse20040AppDeployInstance>()
}
