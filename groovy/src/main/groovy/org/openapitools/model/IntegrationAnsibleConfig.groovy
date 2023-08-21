package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class IntegrationAnsibleConfig {
    
    String inventory
    
    String defaultBranch
    
    String cacheEnabled
    
    String ansiblePlaybooks
    
    String ansibleRoles
    
    String ansibleGroupVars
    
    String ansibleHostVars
    
    Boolean ansibleCommandBus
    
    Boolean ansibleVerbose
    
    Boolean ansibleGalaxyEnabled
    
    String ansibleDefaultBranch
}
