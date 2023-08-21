package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class CheckSshConfig {
    
    Long sshPort
    
    String checkUser
    
    String tunnelOn
    
    String textCheckOn
    
    String checkPassword
    
    String sshHost
    
    String sshUser
    
    String webTextMatch
    
    String checkPasswordHash
}
