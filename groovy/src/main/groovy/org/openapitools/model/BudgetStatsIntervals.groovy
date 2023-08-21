package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.math.BigDecimal;

@Canonical
class BudgetStatsIntervals {
    
    Long index
    
    String month
    
    String shortName
    
    String chartName
    
    Long budget
    
    BigDecimal cost
    
    Date startDate
    
    Date endDate
}
