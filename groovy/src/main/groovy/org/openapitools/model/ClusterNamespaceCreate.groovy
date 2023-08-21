package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.ClusterNamespaceCreateResourcePermissions;

@Canonical
class ClusterNamespaceCreate {
    /* Namespace name */
    String name
    /* Namespace description */
    String description
    /* Namespace active */
    Boolean active = false
    
    ClusterNamespaceCreateResourcePermissions resourcePermissions
}
