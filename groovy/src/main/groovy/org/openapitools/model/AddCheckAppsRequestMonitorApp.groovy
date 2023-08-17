package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.Arrays;

@Canonical
class AddCheckAppsRequestMonitorApp {
    /* Unique name scoped to your account for the check app */
    String name
    /* Optional description field */
    String description
    /* Used to determine if check should affect account wide availability calculations */
    Boolean inUptime = true

    enum SeverityEnum {
    
        INFO("info"),
        
        WARNING("warning"),
        
        CRITICAL("critical")
    
        private final String value
    
        SeverityEnum(String value) {
            this.value = value
        }
    
        String getValue() {
            value
        }
    
        @Override
        String toString() {
            String.valueOf(value)
        }
    }

    /* Severity level of incidents that are created when this check fails */
    SeverityEnum severity = SeverityEnum.CRITICAL
    /* Used to determine if check app is active */
    Boolean active = true
    
    List<Integer> checks
    
    List<Integer> checkGroups
}
