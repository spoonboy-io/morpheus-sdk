package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.Arrays;
import org.openapitools.model.Activity;
import org.openapitools.model.MetaMeta;

@Canonical
class ListActivity200Response {
    
    MetaMeta meta
    
    List<Activity> activity
}
