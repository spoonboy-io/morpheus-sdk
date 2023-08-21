package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.BudgetStatsCurrent;
import org.openapitools.model.BudgetStatsIntervals;

@Canonical
class BudgetStats {
    
    String currency
    
    Long conversionRate
    
    List<BudgetStatsIntervals> intervals = new ArrayList<BudgetStatsIntervals>()
    
    BudgetStatsCurrent current
}
