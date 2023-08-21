package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.math.BigDecimal;
import org.openapitools.model.BillingComputeServers;
import org.openapitools.model.BillingInstances;
import org.openapitools.model.BillingLoadBalancers;
import org.openapitools.model.BillingSnapshots;
import org.openapitools.model.BillingVirtualImages;

@Canonical
class BillingZone {
    
    String zoneName
    
    Long zoneId
    
    String zoneUUID
    
    String zoneCode
    
    Date startDate
    
    Date endDate
    
    String priceUnit
    
    BillingComputeServers computeServers
    
    BillingInstances instances
    
    BillingComputeServers discoveredServers
    
    BillingLoadBalancers loadBalancers
    
    BillingVirtualImages virtualImages
    
    BillingSnapshots snapshots
    
    BigDecimal price
    
    BigDecimal cost
}
