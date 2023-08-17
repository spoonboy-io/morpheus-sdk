package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.math.BigDecimal;
import java.util.ArrayList;
import java.util.Arrays;
import org.openapitools.jackson.nullable.JsonNullable;
import org.openapitools.model.BillingServersServersInnerUsagesInnerApplicablePricesInner;
import org.openapitools.model.BillingServersServersInnerUsagesInnerPricesUsedInner;
import org.openapitools.model.BillingServersServersInnerUsagesInnerVolumesInner;

@Canonical
class BillingServersServersInnerUsagesInner {
    
    String name
    
    String zoneName
    
    String accountName
    
    List<BillingServersServersInnerUsagesInnerVolumesInner> volumes
    
    Long maxMemory
    
    String maxCpu
    
    Long maxCores
    
    String serverExternalId
    
    String serverInternalId
    
    String planName
    
    BigDecimal hourlyPrice
    
    BigDecimal hourlyCost
    
    String currency
    
    List<BillingServersServersInnerUsagesInnerPricesUsedInner> pricesUsed
    
    BigDecimal cost
    
    BigDecimal price
    
    String createdByUser
    
    Long createdByUserId
    
    String siteId
    
    String siteName
    
    String siteUUID
    
    String siteCode
    
    Date startDate
    
    Date endDate
    
    String status
    
    List<Object> tags
    
    List<BillingServersServersInnerUsagesInnerApplicablePricesInner> applicablePrices
    
    Long servicePlanId
    
    String servicePlanName
    
    String resourcePoolId
    
    String resourcePoolName
}
