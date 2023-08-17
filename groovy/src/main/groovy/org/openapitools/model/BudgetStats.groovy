package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.Arrays;
import org.openapitools.model.BudgetStatsCurrent;
import org.openapitools.model.BudgetStatsIntervalsInner;

@Canonical
class BudgetStats {
    
    String currency
    
    Long conversionRate
    
    List<BudgetStatsIntervalsInner> intervals
    
    BudgetStatsCurrent current
}
