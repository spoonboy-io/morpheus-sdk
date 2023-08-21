package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class TaskRestartConfigTaskOptions {
    
    String password
    
    String passwordHash
    
    String localScriptGitRef
    
    String host
    
    String username
    
    String localScriptGitId
    
    String sshKey
    
    String port
}
