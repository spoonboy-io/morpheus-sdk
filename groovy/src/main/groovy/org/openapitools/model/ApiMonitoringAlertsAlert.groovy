package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.ApiMonitoringAlertsAlertContacts;

@Canonical
class ApiMonitoringAlertsAlert {
    /* Unique name scoped to your account for the alert */
    String name
    /* Duration in minutes of the delay before sending notification(s) */
    Integer minDuration = 0
    /* Severity level threshold for sending notifications. */
    String minSeverity = MinSeverityEnum.CRITICAL
    /* Set to false to disable notifications */
    Boolean active = true
    /* Trigger for all checks */
    Boolean allChecks = false
    /* Trigger for all check groups */
    Boolean allGroups = false
    /* Trigger for all monitor apps */
    Boolean allApps = false
    
    List<Integer> checks = new ArrayList<Integer>()
    
    List<Integer> groups = new ArrayList<Integer>()
    
    List<Integer> apps = new ArrayList<Integer>()
    
    List<ApiMonitoringAlertsAlertContacts> contacts = new ArrayList<ApiMonitoringAlertsAlertContacts>()
}
