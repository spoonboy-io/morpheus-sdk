package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.CheckSocketCheckType;
import org.openapitools.model.CheckSocketConfig;

@Canonical
class CheckSocket {
    /* Unique name scoped to your account for the check */
    String name
    /* Optional description field */
    String description
    
    CheckSocketCheckType checkType
    /* Number of seconds you want between check executions (minimum value is 60, depending on your subscription plan) */
    Integer checkInterval = 300
    /* Used to determine if check should affect account wide availability calculations */
    Boolean inUptime = true
    /* Used to determine if check should be scheduled to execute */
    Boolean active = true
    /* Severity level threshold for sending notifications. */
    String severity = SeverityEnum.CRITICAL
    
    CheckSocketConfig config = null
}
