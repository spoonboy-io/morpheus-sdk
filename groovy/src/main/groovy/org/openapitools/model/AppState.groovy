package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.Arrays;
import org.openapitools.model.AppStateInput;
import org.openapitools.model.AppStateOutput;
import org.openapitools.model.AppStateSpecsInner;
import org.openapitools.model.AppStateWorkloadsInner;

@Canonical
class AppState {
    
    List<AppStateWorkloadsInner> workloads
    
    Boolean iacDrift
    
    List<Object> planResources
    
    List<AppStateSpecsInner> specs
    
    String stateData
    
    String planData
    
    AppStateInput input
    
    AppStateOutput output
}
