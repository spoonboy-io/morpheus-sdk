package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.ApiMonitoringAlertsAlertContacts;

@Canonical
class Alert {
    
    Long id
    
    String name
    
    Boolean allApps
    
    Boolean allChecks
    
    Boolean allGroups
    
    Boolean active
    
    String minSeverity
    
    Long minDuration
    
    Date dateCreated
    
    Date lastUpdated
    
    List<Integer> checks = new ArrayList<Integer>()
    
    List<Integer> checkGroups = new ArrayList<Integer>()
    
    List<Integer> apps = new ArrayList<Integer>()
    
    List<ApiMonitoringAlertsAlertContacts> contacts = new ArrayList<ApiMonitoringAlertsAlertContacts>()
}
