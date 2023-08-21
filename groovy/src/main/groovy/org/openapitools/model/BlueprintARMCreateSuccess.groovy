package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.BlueprintARMCreateArm;

@Canonical
class BlueprintARMCreateSuccess {
    /* A name for the blueprint */
    String name
    /* Path to display image. Defaults to an internal Morpheus image. */
    String image
    /* Blueprint Type */
    String type
    
    BlueprintARMCreateArm arm
    /* Private or Public Access */
    String visibility = VisibilityEnum.PRIVATE
    /* Resource Permission Block */
    Object resourcePermission
    /* Owner */
    Object owner
    /* Tenant */
    Object tenant
}
