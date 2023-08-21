package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;

@Canonical
class ApiMonitoringGroupsIdCheckGroup {
    /* Unique name scoped to your account for the check group */
    String name
    /* Optional description field */
    String description
    /* This specifies the minimum number of checks within the group that must be happy to keep the group from becoming unhealthy. */
    Integer minHappy = 1
    /* Used to determine if check should affect account wide availability calculations */
    Boolean inUptime = true
    /* Determines the maximum severity level this group can incur on an incident when failing */
    String severity = SeverityEnum.CRITICAL
    /* Used to determine if check group is active */
    Boolean active = true
    
    List<Integer> checks = new ArrayList<Integer>()
}
