package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.Arrays;
import org.openapitools.model.AddAlertsRequestAlertContactsInner;

@Canonical
class UpdateAlertsRequestAlert {
    /* Unique name scoped to your account for the alert */
    String name
    /* Duration in minutes of the delay before sending notification(s) */
    Integer minDuration = 0

    enum MinSeverityEnum {
    
        INFO("info"),
        
        WARNING("warning"),
        
        CRITICAL("critical")
    
        private final String value
    
        MinSeverityEnum(String value) {
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

    /* Severity level threshold for sending notifications. */
    MinSeverityEnum minSeverity = MinSeverityEnum.CRITICAL
    /* Set to false to disable notifications */
    Boolean active = true
    /* Trigger for all checks */
    Boolean allChecks = false
    /* Trigger for all check groups */
    Boolean allGroups = false
    /* Trigger for all monitor apps */
    Boolean allApps = false
    
    List<Integer> checks
    
    List<Integer> groups
    
    List<Integer> apps
    
    List<AddAlertsRequestAlertContactsInner> contacts
}
