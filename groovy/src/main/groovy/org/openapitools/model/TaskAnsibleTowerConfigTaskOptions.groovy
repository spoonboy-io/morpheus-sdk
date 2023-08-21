package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class TaskAnsibleTowerConfigTaskOptions {
    
    String password
    
    String passwordHash
    
    String ansibleTowerGitRef
    
    String localScriptGitId
    
    String host
    
    String username
    
    String sshKey
    
    String ansibleGroup
    
    String ansibleTowerExecuteMode
    
    String localScriptGitRef
    
    String port
}
