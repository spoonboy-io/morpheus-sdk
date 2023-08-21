package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.math.BigDecimal;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.BillingInstancesPrices;

@Canonical
class BillingInstancesApplicablePrices {
    
    Date startDate
    
    Date endDate
    
    BigDecimal numUnits
    
    BigDecimal cost
    
    BigDecimal price
    
    String currency
    
    List<BillingInstancesPrices> prices = new ArrayList<BillingInstancesPrices>()
}
