package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.Arrays;
import org.openapitools.jackson.nullable.JsonNullable;
import org.openapitools.model.ApplianceSettingsEnabledZoneTypesInner;
import org.openapitools.model.ApprovalItemApprovalItem;

@Canonical
class Approval {
    
    Long id
    
    String name
    
    String internalId
    
    String externalId
    
    String externalName
    
    String requestType
    
    ApplianceSettingsEnabledZoneTypesInner account
    
    ApplianceSettingsEnabledZoneTypesInner approver
    
    String accountIntegration
    
    String status
    
    String errorMessage
    
    Date dateCreated
    
    Date lastUpdated
    
    String requestBy
    
    List<ApprovalItemApprovalItem> approvalItems
}
