package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;

@Canonical
class NetworkRouterPermissionsUpdateTenantPermissions {
    /* Array of tenant account IDs */
    List<Long> accounts = new ArrayList<Long>()
}
