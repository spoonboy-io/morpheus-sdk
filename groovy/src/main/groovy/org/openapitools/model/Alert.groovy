package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.Arrays;
import org.openapitools.jackson.nullable.JsonNullable;
import org.openapitools.model.AddAlertsRequestAlertContactsInner;

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
    
    List<Integer> checks
    
    List<Integer> checkGroups
    
    List<Integer> apps
    
    List<AddAlertsRequestAlertContactsInner> contacts
}
