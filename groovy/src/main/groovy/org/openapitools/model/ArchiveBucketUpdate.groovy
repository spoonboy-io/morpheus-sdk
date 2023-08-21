package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.ApiBlueprintsIdUpdatePermissionsResourcePermissionSites;

@Canonical
class ArchiveBucketUpdate {
    /* A name for the archive bucket. Must be globally unique. */
    String name
    /* A description for the archive bucket */
    String description
    /* Visibility - Set to public to allow all tenants */
    String visibility = VisibilityEnum.PRIVATE
    /* Public URL - Set to true to allow anonymous access */
    Boolean isPublic = false
    
    ApiBlueprintsIdUpdatePermissionsResourcePermissionSites accounts
}
