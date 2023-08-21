package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.GuidanceVmwareSizingPlanBeforeActionConfigRanges;

@Canonical
class GuidanceVmwareSizingPlanBeforeActionConfig {
    
    String storageSizeType
    
    String memorySizeType
    
    GuidanceVmwareSizingPlanBeforeActionConfigRanges ranges
}
