package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class ReportConfig {
    
    String reportType
    
    String startDate
    
    String endDate
    
    String cloudId
}
