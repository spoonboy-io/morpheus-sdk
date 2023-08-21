package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.InlineResponse20082LoadBalancerInstanceSslCert;
import org.openapitools.model.InlineResponse20083LoadBalancerNodeCreatedBy;

@Canonical
class InlineResponse20083LoadBalancerNode {
    
    Long id
    
    String name
    
    String visibility
    
    String description
    
    String ipAddress
    
    Long port
    
    String portType
    
    String monitorPort
    
    Long weight
    
    String nodeState
    
    String internalId
    
    String externalId
    
    Boolean enabled
    
    String status
    
    String statusMessage
    
    Date statusDate
    
    InlineResponse20082LoadBalancerInstanceSslCert server
    
    Long instanceId
    
    Long containerId
    
    String nodeSource
    
    InlineResponse20082LoadBalancerInstanceSslCert monitor
    
    Long maxConnections
    
    String externalRefType
    
    String externalRefId
    
    String externalRefName
    
    InlineResponse20083LoadBalancerNodeCreatedBy createdBy
    
    Date dateCreated
    
    Date lastUpdated
}
