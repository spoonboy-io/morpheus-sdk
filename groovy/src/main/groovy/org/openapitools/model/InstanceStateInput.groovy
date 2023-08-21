package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;

@Canonical
class InstanceStateInput {
    
    List<Object> variables = new ArrayList<Object>()
    
    List<Object> providers = new ArrayList<Object>()
    
    List<Object> data = new ArrayList<Object>()
}
