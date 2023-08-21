package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.InlineResponse20040AppDeployInstance;

@Canonical
class Roles {
    
    Long id
    /* a unique name of the role */
    String name
    /* Alias for name */
    String authority
    
    String description
    
    String scope
    
    String roleType
    
    Boolean multitenant
    
    Boolean multitenantLocked
    
    String parentRoleId
    
    Boolean diverged
    
    Long ownerId
    
    InlineResponse20040AppDeployInstance owner
    
    String defaultPersona
    
    Date dateCreated
    
    Date lastUpdated
}
