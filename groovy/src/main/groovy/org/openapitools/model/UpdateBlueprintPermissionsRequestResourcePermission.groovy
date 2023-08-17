package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.Arrays;
import org.openapitools.model.UpdateBlueprintPermissionsRequestResourcePermissionSitesInner;

@Canonical
class UpdateBlueprintPermissionsRequestResourcePermission {
    /* Set to true to grant access to all groups */
    Boolean all
    /* Array of objects identifying groups with access */
    List<UpdateBlueprintPermissionsRequestResourcePermissionSitesInner> sites
    /* User ID, can be used to change blueprint owner. */
    Long ownerId
}
