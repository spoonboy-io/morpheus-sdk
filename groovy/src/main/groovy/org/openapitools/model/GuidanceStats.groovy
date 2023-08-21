package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.GuidanceStatsSeverity;
import org.openapitools.model.GuidanceStatsType;
import org.openapitools.model.GuidanceVmwareSizingSavings;

@Canonical
class GuidanceStats {
    
    Long total
    
    GuidanceVmwareSizingSavings savings
    
    GuidanceStatsSeverity severity
    
    GuidanceStatsType type
}
