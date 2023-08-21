package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.InlineResponse20040AppDeployInstance;
import org.openapitools.model.InlineResponse20079LoadBalancerMonitorLoadBalancerType;
import org.openapitools.model.UserAccess;
import org.openapitools.model.UserRoles;

@Canonical
class User {
    
    Long id
    
    Long accountId
    
    String username
    
    String displayName
    
    String email
    
    String firstName
    
    String lastName
    
    Boolean enabled
    
    Boolean receiveNotifications
    
    Boolean isUsing2FA
    
    Boolean accountExpired
    
    Boolean accountLocked
    
    Boolean passwordExpired
    
    Long loginCount
    
    Long loginAttempts
    
    Date lastLoginDate
    
    List<UserRoles> roles = new ArrayList<UserRoles>()
    
    InlineResponse20040AppDeployInstance account
    
    String linuxUsername
    
    String linuxPassword
    
    Long linuxKeyPairId
    
    String windowsUsername
    
    String windowsPassword
    
    InlineResponse20079LoadBalancerMonitorLoadBalancerType defaultPersona
    
    Date dateCreated
    
    Date lastUpdated
    
    UserAccess access
}
