package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class IdentitySourcesCustomSSOConfigConfig {
    
    Boolean noninteractive
    
    String loginUrl
    
    String logoutUrl
    
    String encryptionAlgo
    
    String encryptionKey
    
    String requiredAttributeValue
}
