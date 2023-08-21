package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class TaskShellConfigTaskOptions {
    
    String localScriptGitRef
    
    String port
    
    String password
    
    String passwordHash
    
    String username
    
    String host
    
    String sshKey
    
    String localScriptGitId
}
