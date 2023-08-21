package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class TaskVroConfigTaskOptions {
    
    String sshKey
    
    String localScriptGitRef
    
    String port
    
    String vroBody
    
    String password
    
    String passwordHash
    
    String host
    
    String localScriptGitId
    
    String username
}
