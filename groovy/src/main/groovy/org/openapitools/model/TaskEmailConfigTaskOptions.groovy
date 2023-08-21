package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class TaskEmailConfigTaskOptions {
    
    String localScriptGitId
    
    String emailSkipTemplate
    
    String username
    
    String emailSubject
    
    String host
    
    String password
    
    String passwordHash
    
    String emailAddress
    
    String port
    
    String sshKey
    
    String localScriptGitRef
}
