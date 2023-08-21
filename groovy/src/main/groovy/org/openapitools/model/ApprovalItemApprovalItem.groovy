package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.ApiBlueprintsIdUpdatePermissionsResourcePermissionSites;
import org.openapitools.model.ApprovalItemApprovalItemReference;

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
    
    ApiBlueprintsIdUpdatePermissionsResourcePermissionSites approval
    
    ApprovalItemApprovalItemReference reference
}
