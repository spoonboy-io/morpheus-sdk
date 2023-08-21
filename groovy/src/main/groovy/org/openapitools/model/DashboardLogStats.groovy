package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.DashboardLogStatsData;

@Canonical
class DashboardLogStats {
    
    Boolean success
    
    List<DashboardLogStatsData> data = new ArrayList<DashboardLogStatsData>()
    
    Long startMs
    
    Long earliest
    
    Long endMs
    
    Long interval
}
