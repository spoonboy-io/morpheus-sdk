package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class TaskHttpConfigTaskOptions {
    
    String webPassword
    
    String webPasswordHash
    
    String localScriptGitId
    
    String localScriptGitRef
    
    String webUser
    
    String webBody
    
    String webHeaders
    
    String password
    
    String passwordHash
    
    String username
    
    String ignoreSSL
    
    String webMethod
    
    String webUrl
    
    String host
    
    String port
    
    String sshKey
}
