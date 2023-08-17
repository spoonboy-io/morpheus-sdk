package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.math.BigDecimal;
import java.util.ArrayList;
import java.util.Arrays;
import org.openapitools.model.BillingServerUsagesInnerApplicablePricesInnerPricesInner;

@Canonical
class BillingServerUsagesInnerApplicablePricesInner {
    
    Date startDate
    
    Date endDate
    
    BigDecimal numUnits
    
    BigDecimal cost
    
    BigDecimal price
    
    String currency
    
    List<BillingServerUsagesInnerApplicablePricesInnerPricesInner> prices
}
