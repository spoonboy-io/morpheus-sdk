package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.jackson.nullable.JsonNullable;
import org.openapitools.model.ApprovalItemApprovalItemReference;
import org.openapitools.model.UpdateBlueprintPermissionsRequestResourcePermissionSitesInner;

@Canonical
class ApprovalItemApprovalItem {
    
    Long id
    
    String name
    
    String externalId
    
    String externalName
    
    String internalId
    
    String approvedBy
    
    String deniedBy
    
    String status
    
    String errorMessage
    
    Date dateCreated
    
    Date lastUpdated
    
    Date dateApproved
    
    Date dateDenied
    
    UpdateBlueprintPermissionsRequestResourcePermissionSitesInner approval
    
    ApprovalItemApprovalItemReference reference
}