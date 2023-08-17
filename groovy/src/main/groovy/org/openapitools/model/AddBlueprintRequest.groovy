package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.Arrays;
import org.openapitools.jackson.nullable.JsonNullable;
import org.openapitools.model.BlueprintARMCreate;
import org.openapitools.model.BlueprintARMCreateArm;
import org.openapitools.model.BlueprintCFTCreate;
import org.openapitools.model.BlueprintCFTCreateCloudFormation;
import org.openapitools.model.BlueprintHelmCreate;
import org.openapitools.model.BlueprintHelmCreateHelm;
import org.openapitools.model.BlueprintKubernetesCreate;
import org.openapitools.model.BlueprintKubernetesCreateKubernetes;
import org.openapitools.model.BlueprintMorpheusCreate;
import org.openapitools.model.BlueprintTerraformCreate;
import org.openapitools.model.BlueprintTerraformCreateConfig;
import org.openapitools.model.BlueprintTerraformCreateTerraform;

@Canonical
class AddBlueprintRequest {
    /* A name for the blueprint */
    String name
    /* Path to display image. Defaults to an internal Morpheus image. */
    String image

    enum TypeEnum {
    
        ARM("arm"),
        
        CLOUDFORMATION("cloudFormation"),
        
        HELM("helm"),
        
        KUBERNETES("kubernetes"),
        
        MORPHEUS("morpheus"),
        
        TERRAFORM("terraform")
    
        private final String value
    
        TypeEnum(String value) {
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

    /* Blueprint Type */
    TypeEnum type
    /* Array of label strings, can be used for filtering. */
    List<String> labels
    
    BlueprintARMCreateArm arm
    
    BlueprintCFTCreateCloudFormation cloudFormation
    
    BlueprintHelmCreateHelm helm
    
    BlueprintKubernetesCreateKubernetes kubernetes
    
    BlueprintTerraformCreateConfig config
    /* Tier definitions - Create in UI to view a baseline for object */
    Object tiers
    
    BlueprintTerraformCreateTerraform terraform
}
