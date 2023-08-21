package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.InlineResponse20040AppDeployInstance;

@Canonical
class SecurityGroupLocation {
    
    Long id
    
    String name
    
    String description
    
    String externalId
    
    String iacId
    
    InlineResponse20040AppDeployInstance zone
    
    InlineResponse20040AppDeployInstance zonePool
    
    String status
    
    String priority
    
    String groupLayer
}
