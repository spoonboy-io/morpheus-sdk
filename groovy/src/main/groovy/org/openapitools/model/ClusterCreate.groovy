package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.ClusterCreateCloud;
import org.openapitools.model.ClusterCreateGroup;
import org.openapitools.model.ClusterCreateLayout;
import org.openapitools.model.ClusterServerCreate;
import org.openapitools.model.OneOfstringobject;

@Canonical
class ClusterCreate {
    
    OneOfstringobject type = null
    /* Name of the cluster to be created */
    String name
    /* Description of the cluster to be created */
    String description
    /* Array of strings (keywords). This will override labels passed under the `server` object. */
    List<String> labels = new ArrayList<String>()
    
    ClusterCreateGroup group
    
    ClusterCreateCloud cloud
    
    ClusterCreateLayout layout
    
    ClusterServerCreate server
}
