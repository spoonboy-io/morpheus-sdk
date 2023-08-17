package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.Arrays;

@Canonical
class Client {
    
    Long id
    
    String clientId
    
    Long accessTokenValiditySeconds
    
    Long refreshTokenValiditySeconds
    
    List<String> authorities
    
    List<String> authorizedGrantTypes
    
    List<String> scopes
}
