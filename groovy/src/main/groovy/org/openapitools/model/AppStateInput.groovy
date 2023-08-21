package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.AppStateInputData;
import org.openapitools.model.AppStateInputProviders;
import org.openapitools.model.AppStateInputVariables;

@Canonical
class AppStateInput {
    
    List<AppStateInputVariables> variables = new ArrayList<AppStateInputVariables>()
    
    List<AppStateInputProviders> providers = new ArrayList<AppStateInputProviders>()
    
    List<AppStateInputData> data = new ArrayList<AppStateInputData>()
}
