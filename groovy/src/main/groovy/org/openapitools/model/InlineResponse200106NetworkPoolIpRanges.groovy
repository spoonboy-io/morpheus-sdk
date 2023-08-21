package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class InlineResponse200106NetworkPoolIpRanges {
    
    Long id
    
    String startAddress
    
    String endAddress
    
    String internalId
    
    String externalId
    
    String description
    
    Long addressCount
    
    Boolean active
    
    Date dateCreated
    
    Date lastUpdated
    
    String cidr
    
    String cidrIPv6
}
