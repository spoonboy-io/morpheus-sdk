package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class IdentitySourcesJumpCloudConfigConfig {
    
    String organizationId
    
    String bindingUsername
    
    String bindingPassword
    
    String requiredRole
    
    String bindingPasswordHash
}
