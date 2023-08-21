package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.DashboardActivity;
import org.openapitools.model.DashboardBackups;
import org.openapitools.model.DashboardInstanceStats;
import org.openapitools.model.DashboardLogStats;
import org.openapitools.model.DashboardMonitoring;
import org.openapitools.model.DashboardProvisioning;

@Canonical
class Dashboard {
    
    Boolean success
    
    DashboardInstanceStats instanceStats
    
    DashboardProvisioning provisioning
    
    DashboardMonitoring monitoring
    
    DashboardBackups backups
    
    List<DashboardActivity> activity = new ArrayList<DashboardActivity>()
    
    DashboardLogStats logStats
}
