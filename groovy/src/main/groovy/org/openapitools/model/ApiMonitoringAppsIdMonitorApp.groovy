package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;

@Canonical
class ApiMonitoringAppsIdMonitorApp {
    /* Unique name scoped to your account for the check app */
    String name
    /* Optional description field */
    String description
    /* Used to determine if check should affect account wide availability calculations */
    Boolean inUptime = true
    /* Severity level of incidents that are created when this check fails */
    String severity = SeverityEnum.CRITICAL
    /* Used to determine if check app is active */
    Boolean active = true
    
    List<Integer> checks = new ArrayList<Integer>()
    
    List<Integer> checkGroups = new ArrayList<Integer>()
}
