package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;

@Canonical
class Client {
    
    Long id
    
    String clientId
    
    Long accessTokenValiditySeconds
    
    Long refreshTokenValiditySeconds
    
    List<String> authorities = new ArrayList<String>()
    
    List<String> authorizedGrantTypes = new ArrayList<String>()
    
    List<String> scopes = new ArrayList<String>()
}
