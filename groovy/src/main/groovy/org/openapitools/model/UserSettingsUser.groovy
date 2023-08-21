package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.InlineResponse20040AppDeployInstance;
import org.openapitools.model.InlineResponse20079LoadBalancerMonitorLoadBalancerType;

@Canonical
class UserSettingsUser {
    
    Long id
    
    String username
    
    String firstName
    
    String lastName
    
    String email
    
    String linuxUsername
    
    String linuxPassword
    
    Long linuxKeyPairId
    
    String windowsUsername
    
    String windowsPassword
    
    String avatar
    
    String desktopBackground
    
    Boolean receiveNotifications
    
    InlineResponse20040AppDeployInstance defaultGroup
    
    InlineResponse20040AppDeployInstance defaultCloud
    
    InlineResponse20079LoadBalancerMonitorLoadBalancerType defaultPersona
    
    Boolean isUsing2FA
    
    InlineResponse20040AppDeployInstance tenant
}
