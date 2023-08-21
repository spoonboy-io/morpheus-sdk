package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.math.BigDecimal;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.BillingInstanceContainers;

@Canonical
class BillingInstance {
    
    Long instanceId
    
    String instanceUUID
    
    Date startDate
    
    Date endDate
    
    String name
    
    BigDecimal price
    
    BigDecimal cost
    
    String currency
    
    List<BillingInstanceContainers> containers = new ArrayList<BillingInstanceContainers>()
}
