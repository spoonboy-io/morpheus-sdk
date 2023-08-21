package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.Check;
import org.openapitools.model.CheckApp;
import org.openapitools.model.CheckGroup;
import org.openapitools.model.Incident;

@Canonical
class InlineResponse20016 {
    
    CheckApp monitorApp
    
    List<CheckGroup> checkGroups = new ArrayList<CheckGroup>()
    
    List<Check> checks = new ArrayList<Check>()
    
    List<Incident> openIncidents = new ArrayList<Incident>()
}
