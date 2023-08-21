package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.BlueprintTerraformCreateTerraformGit;

@Canonical
class BlueprintTerraformCreateTerraform {
    /* Configuration Type */
    String configType
    /* Terraform definition in JSON for `configType` `json` */
    String json
    /* Terraform definition for `configType` `tf` */
    String tf
    
    BlueprintTerraformCreateTerraformGit git
    /* tfvar secret from Morpheus Cypher */
    String tfvarSecret
}
