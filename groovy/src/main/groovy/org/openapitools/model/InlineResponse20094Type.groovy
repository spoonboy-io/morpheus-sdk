package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;

@Canonical
class InlineResponse20094Type {
    
    Long id
    
    String code
    
    String name
    
    String description
    
    Boolean enabled
    
    Boolean creatable
    
    Boolean selectable
    
    Boolean hasFirewall
    
    Boolean hasDhcp
    
    Boolean hasRouting
    
    Boolean hasNetworkServer
    
    List<Object> optionTypes = new ArrayList<Object>()
    
    List<Object> ruleOptionTypes = new ArrayList<Object>()
}
