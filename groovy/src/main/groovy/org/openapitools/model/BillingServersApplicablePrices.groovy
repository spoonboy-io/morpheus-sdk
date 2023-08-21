package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.math.BigDecimal;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.BillingServersPrices;

@Canonical
class BillingServersApplicablePrices {
    
    Date startDate
    
    Date endDate
    
    BigDecimal numUnits
    
    BigDecimal cost
    
    BigDecimal price
    
    String currency
    
    List<BillingServersPrices> prices = new ArrayList<BillingServersPrices>()
}
