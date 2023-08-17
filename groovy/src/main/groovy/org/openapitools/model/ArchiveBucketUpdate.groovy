package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.UpdateBlueprintPermissionsRequestResourcePermissionSitesInner;

@Canonical
class ArchiveBucketUpdate {
    /* A name for the archive bucket. Must be globally unique. */
    String name
    /* A description for the archive bucket */
    String description

    enum VisibilityEnum {
    
        PUBLIC("public"),
        
        PRIVATE("private")
    
        private final String value
    
        VisibilityEnum(String value) {
            this.value = value
        }
    
        String getValue() {
            value
        }
    
        @Override
        String toString() {
            String.valueOf(value)
        }
    }

    /* Visibility - Set to public to allow all tenants */
    VisibilityEnum visibility = VisibilityEnum.PRIVATE
    /* Public URL - Set to true to allow anonymous access */
    Boolean isPublic = false
    
    UpdateBlueprintPermissionsRequestResourcePermissionSitesInner accounts
}
