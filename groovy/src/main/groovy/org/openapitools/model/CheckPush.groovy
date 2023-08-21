package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.CheckPushCheckType;

@Canonical
class CheckPush {
    /* Unique name scoped to your account for the check */
    String name
    /* Optional description field */
    String description
    
    CheckPushCheckType checkType
    /* Number of seconds you want between check executions (minimum value is 60, depending on your subscription plan) */
    Integer checkInterval = 300
    /* Used to determine if check should affect account wide availability calculations */
    Boolean inUptime = true
    /* Used to determine if check should be scheduled to execute */
    Boolean active = true
    /* Severity level threshold for sending notifications. */
    String severity = SeverityEnum.CRITICAL
}
