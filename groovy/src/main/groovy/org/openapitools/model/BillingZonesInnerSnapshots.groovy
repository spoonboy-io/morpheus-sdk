package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.math.BigDecimal;
import java.util.ArrayList;
import java.util.Arrays;

@Canonical
class BillingZonesInnerSnapshots {
    
    BigDecimal price
    
    BigDecimal cost
    
    List<Object> snapshots
    
    Long count
}
