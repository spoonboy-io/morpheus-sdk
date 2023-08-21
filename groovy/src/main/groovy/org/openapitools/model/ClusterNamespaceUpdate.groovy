package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.ClusterNamespaceUpdatePermissions;

@Canonical
class ClusterNamespaceUpdate {
    /* Namespace name */
    String name
    /* Namespace description */
    String description
    /* Namespace active */
    Boolean active = false
    
    ClusterNamespaceUpdatePermissions permissions
}
