package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class VdiApp {
    
    Long id
    
    String name
    
    String description
    
    String launchPrefix
    
    String iconPath
    
    String logo
    
    Date dateCreated
    
    Date lastUpdated
}
