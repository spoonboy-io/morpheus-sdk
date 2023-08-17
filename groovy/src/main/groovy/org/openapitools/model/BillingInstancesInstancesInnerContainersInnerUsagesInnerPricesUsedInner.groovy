package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.math.BigDecimal;

@Canonical
class BillingInstancesInstancesInnerContainersInnerUsagesInnerPricesUsedInner {
    
    String type
    
    BigDecimal pricePerUnit
    
    BigDecimal costPerUnit
    
    Long quantity
}
