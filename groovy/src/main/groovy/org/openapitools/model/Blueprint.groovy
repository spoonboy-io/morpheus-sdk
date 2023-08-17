package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.Arrays;
import org.openapitools.jackson.nullable.JsonNullable;
import org.openapitools.model.ActivityActivityInnerUser;
import org.openapitools.model.ApplianceSettingsEnabledZoneTypesInner;

@Canonical
class Blueprint {
    
    Long id
    
    String name
    
    List<String> labels
    
    String type
    
    String description
    
    String category
    
    Object config
    
    String visibility
    
    Object resourcePermission
    
    ActivityActivityInnerUser owner
    
    ApplianceSettingsEnabledZoneTypesInner tenant
}
