package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class TaskGroovyConfigTaskOptions {
    
    String localScriptGitRef
    
    String password
    
    String passwordHash
    
    String host
    
    String localScriptGitId
    
    String sshKey
    
    String port
    
    String username
}
