package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;

@Canonical
class NetworkGroupsCreate {
    
    String name
    
    String description
    
    List<Long> networks = new ArrayList<Long>()
    
    List<Object> subnets = new ArrayList<Object>()
}
