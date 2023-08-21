package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.InlineResponse20040AppDeployInstance;

@Canonical
class Environment {
    
    Long id
    
    InlineResponse20040AppDeployInstance account
    
    String code
    
    String name
    
    String description
    
    String visibility
    
    Boolean active
    
    Long sortOrder
    
    Date dateCreated
    
    Date lastUpdated
}
