package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.BlueprintTerraformCreateTerraformGit;

@Canonical
class BlueprintTerraformCreateTerraform {

    enum ConfigTypeEnum {
    
        TF("tf"),
        
        SPEC("spec"),
        
        GIT("git"),
        
        JSON("json")
    
        private final String value
    
        ConfigTypeEnum(String value) {
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

    /* Configuration Type */
    ConfigTypeEnum configType
    /* Terraform definition in JSON for `configType` `json` */
    String json
    /* Terraform definition for `configType` `tf` */
    String tf
    
    BlueprintTerraformCreateTerraformGit git
    /* tfvar secret from Morpheus Cypher */
    String tfvarSecret
}
