package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.math.BigDecimal;
import java.util.ArrayList;
import java.util.Arrays;
import org.openapitools.model.BillingInstancesInstancesInner;

@Canonical
class BillingInstances {
    
    BigDecimal price
    
    BigDecimal cost
    
    Date startDate
    
    Date endDate
    
    List<BillingInstancesInstancesInner> instances
}
