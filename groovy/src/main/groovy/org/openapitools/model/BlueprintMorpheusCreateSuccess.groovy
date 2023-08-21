package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.BlueprintMorpheusCreateSuccessConfig;

@Canonical
class BlueprintMorpheusCreateSuccess {
    /* A name for the blueprint */
    String name
    /* Blueprint Type */
    String type
    
    BlueprintMorpheusCreateSuccessConfig config
    /* Private or Public Access */
    String visibility = VisibilityEnum.PRIVATE
    /* Resource Permission Block */
    Object resourcePermission
    /* Owner */
    Object owner
    /* Tenant */
    Object tenant
}
