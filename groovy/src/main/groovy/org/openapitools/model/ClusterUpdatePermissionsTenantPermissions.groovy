package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;

@Canonical
class ClusterUpdatePermissionsTenantPermissions {
    /* Array of tenant account ids that are allowed access */
    List<Long> accounts = new ArrayList<Long>()
}
