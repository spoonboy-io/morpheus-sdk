package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.InlineResponse20040AppDeployInstance;

@Canonical
class Approvals {
    
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
}
