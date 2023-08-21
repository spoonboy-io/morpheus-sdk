package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.math.BigDecimal;
import java.util.ArrayList;
import java.util.List;

@Canonical
class BillingComputeServers {
    
    BigDecimal price
    
    BigDecimal cost
    
    List<Object> servers = new ArrayList<Object>()
    
    Long count
}
