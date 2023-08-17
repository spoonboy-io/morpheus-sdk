package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.Arrays;
import org.openapitools.model.Alert;
import org.openapitools.model.Check;
import org.openapitools.model.CheckApp;
import org.openapitools.model.CheckGroup;

@Canonical
class GetAlerts200Response {
    
    Alert alert
    
    List<Check> checks
    
    List<CheckGroup> checkGroups
    
    List<CheckApp> apps
}
