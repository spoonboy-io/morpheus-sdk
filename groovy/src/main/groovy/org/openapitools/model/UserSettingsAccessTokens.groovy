package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class UserSettingsAccessTokens {
    
    String clientId
    
    String username
    
    Date expiration
    
    String tokenType
    
    String maskedAccessToken
    
    String maskedRefreshToken
}
