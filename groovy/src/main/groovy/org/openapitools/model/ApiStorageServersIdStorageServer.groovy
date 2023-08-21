package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.ApiBlueprintsIdUpdatePermissionsResourcePermissionSites;

@Canonical
class ApiStorageServersIdStorageServer {
    /* Name */
    String name
    /* The `Storage Type` Code or ID */
    String type
    /* description */
    String description
    /* The enabled flag */
    Boolean enabled = true
    /* Configuration object with parameters that vary by `type` */
    Object config
    /* private or public */
    String visibility = VisibilityEnum.PRIVATE
    /* Array of tenant account ids that are allowed access */
    List<ApiBlueprintsIdUpdatePermissionsResourcePermissionSites> tenants = new ArrayList<ApiBlueprintsIdUpdatePermissionsResourcePermissionSites>()
}
