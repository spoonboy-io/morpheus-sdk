package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.Arrays;
import org.openapitools.model.Blueprint;
import org.openapitools.model.MetaMeta;

@Canonical
class ListBlueprints200Response {
    
    MetaMeta meta
    
    List<Blueprint> blueprints
}
