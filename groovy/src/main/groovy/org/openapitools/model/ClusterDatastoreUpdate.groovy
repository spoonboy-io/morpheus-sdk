package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.ClusterUpdatePermissions;

@Canonical
class ClusterDatastoreUpdate {
    /* Datastore active */
    Boolean active = true
    
    ClusterUpdatePermissions permissions
    /* Visibility for datastore */
    String visibility = "private"
}
