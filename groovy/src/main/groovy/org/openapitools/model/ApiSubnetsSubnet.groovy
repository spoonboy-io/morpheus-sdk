package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.ApiBlueprintsIdUpdatePermissionsResourcePermissionSites;
import org.openapitools.model.ApiSubnetsSubnetType;

@Canonical
class ApiSubnetsSubnet {
    
    ApiSubnetsSubnetType type
    /* Configuration object. Settings vary by type. */
    Object config
    /* Array of tenant account ID objects that are allowed access */
    List<ApiBlueprintsIdUpdatePermissionsResourcePermissionSites> tenants = new ArrayList<ApiBlueprintsIdUpdatePermissionsResourcePermissionSites>()
    /* private or public */
    String visibility = "private"
    /* Array of label strings, can be used for filtering. */
    List<String> labels = new ArrayList<String>()
}
