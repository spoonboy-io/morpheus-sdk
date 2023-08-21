package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.BlueprintKubernetesCreateConfig;
import org.openapitools.model.BlueprintKubernetesCreateKubernetes;

@Canonical
class BlueprintKubernetesCreateSuccess {
    /* A name for the blueprint */
    String name
    /* Path to display image. Defaults to an internal Morpheus image. */
    String image
    /* Blueprint Type */
    String type
    
    BlueprintKubernetesCreateKubernetes kubernetes
    
    BlueprintKubernetesCreateConfig config
    /* Private or Public Access */
    String visibility = VisibilityEnum.PRIVATE
    /* Resource Permission Block */
    Object resourcePermission
    /* Owner */
    Object owner
    /* Tenant */
    Object tenant
}
