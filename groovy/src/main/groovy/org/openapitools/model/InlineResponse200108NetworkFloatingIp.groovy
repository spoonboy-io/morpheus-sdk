package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.InlineResponse200108NetworkFloatingIpCloud;
import org.openapitools.model.InlineResponse200108NetworkFloatingIpCreatedBy;
import org.openapitools.model.InlineResponse200108NetworkFloatingIpNetworkDomain;
import org.openapitools.model.InlineResponse200108NetworkFloatingIpServer;

@Canonical
class InlineResponse200108NetworkFloatingIp {
    
    Long id
    
    String externalId
    
    InlineResponse200108NetworkFloatingIpCloud cloud
    
    InlineResponse200108NetworkFloatingIpServer server
    
    String ipStatus
    /* IP Address */
    String ipAddress
    
    String ipRange
    
    String ptrId
    
    InlineResponse200108NetworkFloatingIpNetworkDomain networkDomain
    
    InlineResponse200108NetworkFloatingIpCreatedBy createdBy
    
    Date dateCreated
    
    Date lastUpdated
}
