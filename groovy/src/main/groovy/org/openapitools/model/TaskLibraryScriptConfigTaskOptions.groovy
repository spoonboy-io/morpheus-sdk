package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class TaskLibraryScriptConfigTaskOptions {
    
    String host
    
    String username
    
    String password
    
    String passwordHash
    
    String localScriptGitRef
    
    String localScriptGitId
    
    String containerScriptId
    
    String sshKey
    
    String containerScript
    
    String port
}
