package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.AppStateOutput;
import org.openapitools.model.InstanceStateInput;

@Canonical
class InstanceState {
    
    List<Object> workloads = new ArrayList<Object>()
    
    Boolean iacDrift
    
    List<Object> planResources = new ArrayList<Object>()
    
    List<Object> specs = new ArrayList<Object>()
    
    String stateData
    
    String planData
    
    InstanceStateInput input
    
    AppStateOutput output
}
