package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class TaskPowershellConfigTaskOptions {
    
    String host
    
    String localScriptGitId
    
    String username
    
    String port
    
    String localScriptGitRef
    
    String password
    
    String passwordHash
    
    String sshKey
}
