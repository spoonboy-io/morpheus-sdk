package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.Arrays;
import org.openapitools.jackson.nullable.JsonNullable;
import org.openapitools.model.BlueprintKubernetesCreateConfig;
import org.openapitools.model.BlueprintKubernetesCreateKubernetes;

@Canonical
class BlueprintKubernetesCreate {
    /* A name for the blueprint */
    String name
    /* Path to display image. Defaults to an internal Morpheus image. */
    String image

    enum TypeEnum {
    
        KUBERNETES("kubernetes")
    
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
    
    BlueprintKubernetesCreateKubernetes kubernetes
    
    BlueprintKubernetesCreateConfig config
}
