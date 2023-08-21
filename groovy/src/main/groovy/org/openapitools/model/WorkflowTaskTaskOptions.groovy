package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class WorkflowTaskTaskOptions {
    
    String localScriptGitRef
    
    String username
    
    String localScriptGitId
    
    String host
    
    String sshKey
    
    String port
    
    String password
    
    String passwordHash
    
    String shellSudo
}
