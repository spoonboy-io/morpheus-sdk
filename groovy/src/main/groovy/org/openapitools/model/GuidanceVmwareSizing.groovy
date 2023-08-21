package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.GuidanceVmwareSizingConfig;
import org.openapitools.model.GuidanceVmwareSizingPlanAfterAction;
import org.openapitools.model.GuidanceVmwareSizingPlanBeforeAction;
import org.openapitools.model.GuidanceVmwareSizingResource;
import org.openapitools.model.GuidanceVmwareSizingSavings;
import org.openapitools.model.GuidanceVmwareSizingType;
import org.openapitools.model.GuidanceVmwareSizingZone;

@Canonical
class GuidanceVmwareSizing {
    
    Long id
    
    Date dateCreated
    
    Date lastUpdated
    
    String actionCategory
    
    String actionMessage
    
    String actionTitle
    
    String actionType
    
    String actionValue
    
    String actionValueType
    
    Long actionPlanId
    
    String statusMessage
    
    Long accountId
    
    String userId
    
    Long siteId
    
    GuidanceVmwareSizingZone zone
    
    String state
    
    String stateMessage
    
    String severity
    
    Boolean resolved
    
    String resolvedMessage
    
    String refType
    
    Long refId
    
    String refName
    
    GuidanceVmwareSizingType type
    
    GuidanceVmwareSizingSavings savings
    
    GuidanceVmwareSizingResource resource
    
    GuidanceVmwareSizingPlanBeforeAction planBeforeAction
    
    GuidanceVmwareSizingPlanAfterAction planAfterAction
    
    GuidanceVmwareSizingConfig config
}
