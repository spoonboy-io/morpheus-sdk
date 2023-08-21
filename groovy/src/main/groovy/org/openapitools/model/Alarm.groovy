package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class Alarm {
    
    Long id
    
    String name
    
    Date dateCreated
    
    Date lastUpdated
}
