package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.ApprovalItemApprovalItem;
import org.openapitools.model.InlineResponse20040AppDeployInstance;

@Canonical
class Approval {
    
    Long id
    
    String name
    
    String internalId
    
    String externalId
    
    String externalName
    
    String requestType
    
    InlineResponse20040AppDeployInstance account
    
    InlineResponse20040AppDeployInstance approver
    
    String accountIntegration
    
    String status
    
    String errorMessage
    
    Date dateCreated
    
    Date lastUpdated
    
    String requestBy
    
    List<ApprovalItemApprovalItem> approvalItems = new ArrayList<ApprovalItemApprovalItem>()
}
