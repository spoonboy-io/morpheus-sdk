package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class TaskJavaConfigTaskOptions {
    
    String username
    
    String port
    
    String jsScript
    
    String host
    
    String localScriptGitRef
    
    String password
    
    String passwordHash
    
    String sshKey
    
    String localScriptGitId
}
