package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.AppBlueprint;
import org.openapitools.model.InlineResponse200107NetworkPoolCreatedBy;
import org.openapitools.model.InlineResponse20040AppDeployInstance;

@Canonical
class AppCreateResponse {
    
    Long id
    
    String name
    
    String description
    
    List<String> labels = new ArrayList<String>()
    
    String environment
    
    Long accountId
    
    InlineResponse20040AppDeployInstance account
    
    InlineResponse200107NetworkPoolCreatedBy owner
    
    Long siteId
    
    InlineResponse20040AppDeployInstance group
    
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
    
    List<Object> appTiers = new ArrayList<Object>()
    
    List<InlineResponse20040AppDeployInstance> instances = new ArrayList<InlineResponse20040AppDeployInstance>()
}
