package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.jackson.nullable.JsonNullable;
import org.openapitools.model.ApplianceSettingsEnabledZoneTypesInner;

@Canonical
class Approvals {
    
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
}
