package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class CredentialConfig {
    
    String clientSecret
    
    String clientId
    
    String clientAuth
    
    String scope
    
    String grantType
    
    String accessTokenUrl
    
    String clientSecretHash
}
