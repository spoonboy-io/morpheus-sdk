package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class AddAlertsRequestAlertContactsInner {
    
    Long id
    
    String name
    
    String method
    
    Boolean notify
    
    Boolean close
}
