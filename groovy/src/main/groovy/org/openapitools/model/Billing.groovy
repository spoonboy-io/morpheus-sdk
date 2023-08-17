package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.math.BigDecimal;
import java.util.ArrayList;
import java.util.Arrays;
import org.openapitools.model.BillingZonesInner;

@Canonical
class Billing {
    
    Long accountId
    
    String accountUUID
    
    String name
    
    Date startDate
    
    Date endDate
    
    String priceUnit
    
    BigDecimal price
    
    BigDecimal cost
    
    List<BillingZonesInner> zones
}
