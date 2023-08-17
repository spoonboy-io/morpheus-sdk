package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.math.BigDecimal;
import java.util.ArrayList;
import java.util.Arrays;
import org.openapitools.model.BillingInstancesInstancesInnerContainersInnerUsagesInner;

@Canonical
class BillingInstancesInstancesInnerContainersInner {
    
    String refType
    
    String refUUID
    
    Long refId
    
    Date startDate
    
    Date endDate
    
    BigDecimal cost
    
    BigDecimal price
    
    BigDecimal numUnits
    
    String unit
    
    String currency
    
    List<BillingInstancesInstancesInnerContainersInnerUsagesInner> usages
    
    Long numUsages
    
    Long totalUsages
    
    Boolean hasMoreUsages
    
    Boolean foundPricing
    
    String name
    
    Long serverId
    
    String serverUUID
    
    String serverUniqueId
    
    String serverExternalId
    
    String serverInternalId
    
    Long resourcePoolId
    
    String resourcePoolName
}
