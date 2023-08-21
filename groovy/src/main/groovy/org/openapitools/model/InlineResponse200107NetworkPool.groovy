package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.InlineResponse200107NetworkPoolCreatedBy;

@Canonical
class InlineResponse200107NetworkPool {
    
    Long id
    
    Long networkPoolId
    
    String ipType
    
    String ipAddress
    
    String gatewayAddress
    
    String subnetMask
    
    String dnsServer
    
    String interfaceName
    
    String description
    
    Boolean active
    
    Boolean staticIp
    
    String fqdn
    
    String domainName
    
    String hostname
    
    String internalId
    
    String externalId
    
    String ptrId
    
    Date dateCreated
    
    Date lastUpdated
    
    Date startDate
    
    Date endDate
    
    String refType
    
    String refId
    
    String subRefId
    
    String networkDomain
    
    InlineResponse200107NetworkPoolCreatedBy createdBy
}
