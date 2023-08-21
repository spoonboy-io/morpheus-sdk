package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class IdentitySourcesOktaConfigConfig {
    
    String url
    
    String administratorAPIToken
    
    String requiredGroup
    
    String requiredGroupId
    
    String administratorAPITokenHash
}
