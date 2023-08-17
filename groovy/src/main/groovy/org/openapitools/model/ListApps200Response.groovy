package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.Arrays;
import org.openapitools.model.App;
import org.openapitools.model.AppStats;
import org.openapitools.model.MetaMeta;

@Canonical
class ListApps200Response {
    
    AppStats stats
    
    MetaMeta meta
    
    List<App> apps
}
