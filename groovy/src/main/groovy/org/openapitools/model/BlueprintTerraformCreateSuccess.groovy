package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.BlueprintTerraformCreateConfig;
import org.openapitools.model.BlueprintTerraformCreateTerraform;

@Canonical
class BlueprintTerraformCreateSuccess {
    /* A name for the blueprint */
    String name
    /* Path to display image. Defaults to an internal Morpheus image. */
    String image
    /* Blueprint Type */
    String type
    
    BlueprintTerraformCreateTerraform terraform
    
    BlueprintTerraformCreateConfig config
    /* Private or Public Access */
    String visibility = VisibilityEnum.PRIVATE
    /* Resource Permission Block */
    Object resourcePermission
    /* Owner */
    Object owner
    /* Tenant */
    Object tenant
}
