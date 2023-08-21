package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.math.BigDecimal;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.BillingServerApplicablePrices;

@Canonical
class BillingServerUsages {
    
    BigDecimal cost
    
    BigDecimal price
    
    String createdByUser
    
    Long createdByUserId
    
    String siteId
    
    String siteName
    
    String siteUUID
    
    String siteCode
    
    String currency
    
    Date startDate
    
    Date endDate
    
    String status
    
    List<Object> tags = new ArrayList<Object>()
    
    List<BillingServerApplicablePrices> applicablePrices = new ArrayList<BillingServerApplicablePrices>()
    
    Long servicePlanId
    
    String servicePlanName
    
    Long resourcePoolId
    
    String resourcePoolName
}
