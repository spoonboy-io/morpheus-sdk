package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.InlineResponse20040AppDeployInstance;

@Canonical
class UsersAvailableRolesRoles {
    
    Long id
    
    String authority
    
    String description
    
    String roleType
    
    InlineResponse20040AppDeployInstance owner
}
