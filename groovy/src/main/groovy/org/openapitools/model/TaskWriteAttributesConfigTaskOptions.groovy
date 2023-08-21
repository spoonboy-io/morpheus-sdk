package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class TaskWriteAttributesConfigTaskOptions {
    
    String host
    
    String sshKey
    
    String localScriptGitRef
    
    String port
    
    String localScriptGitId
    
    String password
    
    String passwordHash
    
    String writeAttributesAttributes
    
    String username
}
