package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;

@Canonical
class SubnetResourcePermission {
    
    Boolean all
    
    List<Object> sites = new ArrayList<Object>()
    
    Boolean allPlans
    
    List<Object> plans = new ArrayList<Object>()
}
