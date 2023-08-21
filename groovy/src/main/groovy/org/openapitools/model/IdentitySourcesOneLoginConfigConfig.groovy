package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class IdentitySourcesOneLoginConfigConfig {
    
    String subdomain
    
    String region
    
    String clientSecret
    
    String clientId
    
    String requiredRole
    
    String requiredRoleId
    
    String clientSecretHash
}
