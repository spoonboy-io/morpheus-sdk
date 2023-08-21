package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class TaskChefBootstrapConfigTaskOptions {
    
    String chefDataKey
    
    String chefDataKeyHash
    
    String chefRunList
    
    String localScriptGitRef
    
    String chefDataKeyPath
    
    String localScriptGitId
    
    String port
    
    String chefEnv
    
    String chefNodeName
    
    String host
    
    String sshKey
    
    String username
    
    String password
    
    String passwordHash
    
    String chefAttributes
}
