package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.AppStateInput;
import org.openapitools.model.AppStateOutput;
import org.openapitools.model.AppStateSpecs;
import org.openapitools.model.AppStateWorkloads;

@Canonical
class AppState {
    
    List<AppStateWorkloads> workloads = new ArrayList<AppStateWorkloads>()
    
    Boolean iacDrift
    
    List<Object> planResources = new ArrayList<Object>()
    
    List<AppStateSpecs> specs = new ArrayList<AppStateSpecs>()
    
    String stateData
    
    String planData
    
    AppStateInput input
    
    AppStateOutput output
}
