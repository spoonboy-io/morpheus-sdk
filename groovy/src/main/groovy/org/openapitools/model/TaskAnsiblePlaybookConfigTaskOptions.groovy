package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class TaskAnsiblePlaybookConfigTaskOptions {
    
    String ansibleOptions
    
    String ansiblePlaybook
    
    String sshKey
    
    String port
    
    String localScriptGitRef
    
    String password
    
    String passwordHash
    
    String localScriptGitId
    
    String ansibleGitId
    
    String host
    
    String ansibleSkipTags
    
    String ansibleTags
    
    String username
    
    String ansibleGitRef
}
