package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;

@Canonical
class ClusterUpdate {
    /* Cluster name */
    String name
    /* Cluster description */
    String description
    /* Array of label strings, can be used for filtering. */
    List<String> labels = new ArrayList<String>()
    /* Cluster enabled */
    Boolean enabled
    /* Cluster API Url */
    String serviceUrl
    /* Cluster API token */
    String serviceToken
    /* Queue cluster refresh */
    Boolean refresh
    /* Cluster managed */
    Boolean managed
}
