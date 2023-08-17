package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.Arrays;
import org.openapitools.jackson.nullable.JsonNullable;
import org.openapitools.model.ActivityActivityInnerUser;
import org.openapitools.model.AppBlueprint;
import org.openapitools.model.AppStats;
import org.openapitools.model.ApplianceSettingsEnabledZoneTypesInner;

@Canonical
class App {
    
    Long id
    
    String name
    
    String description
    
    List<String> labels
    
    String environment
    
    Long accountId
    
    ApplianceSettingsEnabledZoneTypesInner account
    
    ActivityActivityInnerUser owner
    
    Long siteId
    
    ApplianceSettingsEnabledZoneTypesInner group
    
    AppBlueprint blueprint
    
    String type
    
    Date dateCreated
    
    Date lastUpdated
    
    Date removalDate
    
    String appContext
    
    String status
    
    String appStatus
    
    Long instanceCount
    
    Long containerCount
    
    List<Object> appTiers
    
    List<ApplianceSettingsEnabledZoneTypesInner> instances
    
    AppStats stats
}
