package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.math.BigDecimal;
import java.util.ArrayList;
import java.util.Arrays;
import org.openapitools.jackson.nullable.JsonNullable;
import org.openapitools.model.BillingServerUsagesInnerApplicablePricesInner;

@Canonical
class BillingServerUsagesInner {
    
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
    
    List<Object> tags
    
    List<BillingServerUsagesInnerApplicablePricesInner> applicablePrices
    
    Long servicePlanId
    
    String servicePlanName
    
    Long resourcePoolId
    
    String resourcePoolName
}
