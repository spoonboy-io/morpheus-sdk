package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class TaskPuppetConfigTaskOptions {
    
    String port
    
    String host
    
    String username
    
    String puppetEnvironment
    
    String puppetNodeName
    
    String sshKey
    
    String localScriptGitId
    
    String localScriptGitRef
    
    String password
    
    String passwordHash
}
