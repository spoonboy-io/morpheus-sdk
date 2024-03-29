package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.math.BigDecimal;

@Canonical
class BillingServerPrices {
    
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
