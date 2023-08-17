package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.Arrays;
import org.openapitools.model.AppStateInputDataInner;
import org.openapitools.model.AppStateInputProvidersInner;
import org.openapitools.model.AppStateInputVariablesInner;

@Canonical
class AppStateInput {
    
    List<AppStateInputVariablesInner> variables
    
    List<AppStateInputProvidersInner> providers
    
    List<AppStateInputDataInner> data
}
