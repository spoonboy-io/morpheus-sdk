package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.math.BigDecimal;
import java.util.ArrayList;
import java.util.Arrays;
import org.openapitools.jackson.nullable.JsonNullable;
import org.openapitools.model.BillingInstancesInstancesInnerContainersInnerUsagesInnerApplicablePricesInner;
import org.openapitools.model.BillingInstancesInstancesInnerContainersInnerUsagesInnerPricesUsedInner;
import org.openapitools.model.BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInner;

@Canonical
class BillingInstancesInstancesInnerContainersInnerUsagesInner {
    
    String name
    
    String instanceName
    
    String zoneName
    
    String accountName
    
    List<BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInner> volumes
    
    Long maxMemory
    
    String maxCpu
    
    Long maxCores
    
    String serverExternalId
    
    String serverInternalId
    
    String planName
    
    BigDecimal hourlyPrice
    
    BigDecimal hourlyCost
    
    String currency
    
    List<BillingInstancesInstancesInnerContainersInnerUsagesInnerPricesUsedInner> pricesUsed
    
    BigDecimal cost
    
    BigDecimal price
    
    String createdByUser
    
    Long createdByUserId
    
    Long siteId
    
    String siteName
    
    String siteUUID
    
    String siteCode
    
    Date startDate
    
    Date endDate
    
    String status
    
    List<Object> tags
    
    List<BillingInstancesInstancesInnerContainersInnerUsagesInnerApplicablePricesInner> applicablePrices
    
    Long servicePlanId
    
    String servicePlanName
    
    Long resourcePoolId
    
    String resourcePoolName
}
