package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class TaskJrubyConfigTaskOptions {
    
    String host
    
    String username
    
    String localScriptGitId
    
    String password
    
    String passwordHash
    
    String sshKey
    
    String port
    
    String localScriptGitRef
}
