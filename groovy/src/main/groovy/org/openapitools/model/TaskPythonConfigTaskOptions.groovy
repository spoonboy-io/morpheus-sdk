package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class TaskPythonConfigTaskOptions {
    
    String pythonArgs
    
    String pythonBinary
    
    String pythonAdditionalPackages
    
    String port
    
    String host
    
    String username
    
    String sshKey
    
    String password
    
    String passwordHash
    
    String localScriptGitId
    
    String localScriptGitRef
}
