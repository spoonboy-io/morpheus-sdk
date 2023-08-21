package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.ClusterUpdatePermissionsResourcePermissionsSites;

@Canonical
class ClusterNamespaceCreateResourcePermissions {
    /* Pass true to allow access to all groups */
    Boolean all
    /* Array of groups that are allowed access */
    List<ClusterUpdatePermissionsResourcePermissionsSites> sites = new ArrayList<ClusterUpdatePermissionsResourcePermissionsSites>()
    /* Pass true to allow access to all plans */
    Boolean allPlans
    /* Array of plans that are allowed access */
    List<ClusterUpdatePermissionsResourcePermissionsSites> plans = new ArrayList<ClusterUpdatePermissionsResourcePermissionsSites>()
}
