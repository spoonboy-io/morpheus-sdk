package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.math.BigDecimal;
import java.util.ArrayList;
import java.util.List;

@Canonical
class DashboardMonitoring {
    
    BigDecimal avgHealth
    
    BigDecimal avgResponseTime
    
    BigDecimal warningApps
    
    BigDecimal warningChecks
    
    BigDecimal failApps
    
    BigDecimal totalApps
    
    BigDecimal failChecks
    
    BigDecimal successApps
    
    BigDecimal mutedApps
    
    BigDecimal successChecks
    
    BigDecimal totalChecks
    
    BigDecimal mutedChecks
    
    List<BigDecimal> responseTimes = new ArrayList<BigDecimal>()
    
    Boolean allSuccess
    
    BigDecimal openIncidents
}
