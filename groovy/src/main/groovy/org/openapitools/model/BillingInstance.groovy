package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.math.BigDecimal;
import java.util.ArrayList;
import java.util.Arrays;
import org.openapitools.jackson.nullable.JsonNullable;
import org.openapitools.model.BillingInstanceContainersInner;

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
    
    List<BillingInstanceContainersInner> containers
}
