package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class Contact {
    
    Long id
    
    String emailAddress
    
    String name
    
    String smsAddress
    
    String slackHook
}
