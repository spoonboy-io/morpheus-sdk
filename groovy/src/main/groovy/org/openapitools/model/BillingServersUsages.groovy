package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.math.BigDecimal;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.BillingServersApplicablePrices;
import org.openapitools.model.BillingServersPricesUsed;
import org.openapitools.model.BillingServersVolumes;

@Canonical
class BillingServersUsages {
    
    String name
    
    String zoneName
    
    String accountName
    
    List<BillingServersVolumes> volumes = new ArrayList<BillingServersVolumes>()
    
    Long maxMemory
    
    String maxCpu
    
    Long maxCores
    
    String serverExternalId
    
    String serverInternalId
    
    String planName
    
    BigDecimal hourlyPrice
    
    BigDecimal hourlyCost
    
    String currency
    
    List<BillingServersPricesUsed> pricesUsed = new ArrayList<BillingServersPricesUsed>()
    
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
    
    List<Object> tags = new ArrayList<Object>()
    
    List<BillingServersApplicablePrices> applicablePrices = new ArrayList<BillingServersApplicablePrices>()
    
    Long servicePlanId
    
    String servicePlanName
    
    String resourcePoolId
    
    String resourcePoolName
}
