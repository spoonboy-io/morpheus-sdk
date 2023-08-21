package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.HashMap;
import java.util.List;

@Canonical
class DashboardLogStatsData {
    
    String key
    
    Map<String, Integer> values = new HashMap<String, Integer>()
    
    Long count
}
