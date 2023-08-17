package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.math.BigDecimal;
import org.openapitools.jackson.nullable.JsonNullable;

@Canonical
class BillingServerUsagesInnerApplicablePricesInnerPricesInner {
    
    String type
    
    BigDecimal pricePerUnit
    
    BigDecimal costPerUnit
    
    BigDecimal cost
    
    BigDecimal price
    
    Long quantity
    
    String datastoreId
    
    String volumeType
    
    String datastore
}
