package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.math.BigDecimal;

@Canonical
class BillingServersPricesUsed {
    
    String type
    
    BigDecimal pricePerUnit
    
    BigDecimal costPerUnit
    
    BigDecimal quantity
}
