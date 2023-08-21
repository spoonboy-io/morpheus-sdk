package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class IdentitySourcesADConfigConfig {
    
    String url
    
    String domain
    
    String useSSL
    
    String bindingUsername
    
    String bindingPassword
    
    String requiredGroup
    
    String requiredGroupDN
    
    Boolean searchMemberGroups
    
    String bindingPasswordHash
}
