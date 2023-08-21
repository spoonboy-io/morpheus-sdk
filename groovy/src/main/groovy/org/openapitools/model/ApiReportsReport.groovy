package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.math.BigDecimal;

@Canonical
class ApiReportsReport {
    /* Code value for the report type you want to run */
    String type
    
    String startDate
    
    String endDate
    
    String startMonth
    
    String endMonth
    /* The Group ID filter for the report */
    BigDecimal groupId
    /* The Cloud ID filter for the report */
    BigDecimal cloudId
}
