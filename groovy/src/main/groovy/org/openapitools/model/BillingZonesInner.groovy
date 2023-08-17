package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.math.BigDecimal;
import org.openapitools.jackson.nullable.JsonNullable;
import org.openapitools.model.BillingZonesInnerComputeServers;
import org.openapitools.model.BillingZonesInnerInstances;
import org.openapitools.model.BillingZonesInnerLoadBalancers;
import org.openapitools.model.BillingZonesInnerSnapshots;
import org.openapitools.model.BillingZonesInnerVirtualImages;

@Canonical
class BillingZonesInner {
    
    String zoneName
    
    Long zoneId
    
    String zoneUUID
    
    String zoneCode
    
    Date startDate
    
    Date endDate
    
    String priceUnit
    
    BillingZonesInnerComputeServers computeServers
    
    BillingZonesInnerInstances instances
    
    BillingZonesInnerComputeServers discoveredServers
    
    BillingZonesInnerLoadBalancers loadBalancers
    
    BillingZonesInnerVirtualImages virtualImages
    
    BillingZonesInnerSnapshots snapshots
    
    BigDecimal price
    
    BigDecimal cost
}
