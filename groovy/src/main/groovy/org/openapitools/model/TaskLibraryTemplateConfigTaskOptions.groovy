package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class TaskLibraryTemplateConfigTaskOptions {
    
    String username
    
    String localScriptGitId
    
    String sshKey
    
    String port
    
    String password
    
    String passwordHash
    
    String containerTemplateId
    
    String containerTemplate
    
    String localScriptGitRef
    
    String host
}
