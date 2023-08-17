package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.math.BigDecimal;
import java.util.ArrayList;
import java.util.Arrays;
import org.openapitools.model.BillingInstancesInstancesInnerContainersInnerUsagesInnerApplicablePricesInnerPricesInner;

@Canonical
class BillingInstancesInstancesInnerContainersInnerUsagesInnerApplicablePricesInner {
    
    Date startDate
    
    Date endDate
    
    BigDecimal numUnits
    
    BigDecimal cost
    
    BigDecimal price
    
    String currency
    
    List<BillingInstancesInstancesInnerContainersInnerUsagesInnerApplicablePricesInnerPricesInner> prices
}
