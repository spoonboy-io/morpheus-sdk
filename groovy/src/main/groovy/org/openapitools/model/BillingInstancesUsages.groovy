package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.math.BigDecimal;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.BillingInstancesApplicablePrices;
import org.openapitools.model.BillingInstancesPricesUsed;
import org.openapitools.model.BillingInstancesVolumes;

@Canonical
class BillingInstancesUsages {
    
    String name
    
    String instanceName
    
    String zoneName
    
    String accountName
    
    List<BillingInstancesVolumes> volumes = new ArrayList<BillingInstancesVolumes>()
    
    Long maxMemory
    
    String maxCpu
    
    Long maxCores
    
    String serverExternalId
    
    String serverInternalId
    
    String planName
    
    BigDecimal hourlyPrice
    
    BigDecimal hourlyCost
    
    String currency
    
    List<BillingInstancesPricesUsed> pricesUsed = new ArrayList<BillingInstancesPricesUsed>()
    
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
    
    List<Object> tags = new ArrayList<Object>()
    
    List<BillingInstancesApplicablePrices> applicablePrices = new ArrayList<BillingInstancesApplicablePrices>()
    
    Long servicePlanId
    
    String servicePlanName
    
    Long resourcePoolId
    
    String resourcePoolName
}
