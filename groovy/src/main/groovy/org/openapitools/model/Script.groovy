package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;

@Canonical
class Script {
    
    Long id
    
    String code
    
    String account
    
    String name
    
    List<String> labels = new ArrayList<String>()
    
    String category
    
    Long sortOrder
    
    String scriptVersion
    
    String scriptPhase
    
    String scriptType
    
    String script
    
    String scriptService
    
    String scriptMethod
    
    String runAsUser
    
    String runAsPassword
    
    Boolean sudoUser
    
    Boolean failOnError
}
