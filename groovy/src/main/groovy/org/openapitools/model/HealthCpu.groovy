package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.math.BigDecimal;

@Canonical
class HealthCpu {
    
    Boolean success
    
    Long cpuLoad
    
    Long cpuTotalLoad
    
    Long processorCount
    
    Long processTime
    
    BigDecimal systemLoad
    
    String status
}
