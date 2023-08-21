package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.ApiServicePlansServicePlanConfigRanges;

@Canonical
class ApiServicePlansServicePlanConfig {
    /* Specifies range min / max storage multiplier */
    String storageSizeType = StorageSizeTypeEnum.GB
    /* Specifies range min / max memory multiplier */
    String memorySizeType = MemorySizeTypeEnum.MB
    
    ApiServicePlansServicePlanConfigRanges ranges
}
