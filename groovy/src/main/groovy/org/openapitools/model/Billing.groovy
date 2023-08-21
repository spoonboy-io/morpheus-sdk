package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.math.BigDecimal;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.BillingZones;

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
    
    List<BillingZones> zones = new ArrayList<BillingZones>()
}
