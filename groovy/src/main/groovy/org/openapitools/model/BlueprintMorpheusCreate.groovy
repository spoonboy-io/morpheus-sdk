package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.Arrays;
import org.openapitools.jackson.nullable.JsonNullable;

@Canonical
class BlueprintMorpheusCreate {
    /* A name for the blueprint */
    String name

    enum TypeEnum {
    
        MORPHEUS("morpheus")
    
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
    /* Tier definitions - Create in UI to view a baseline for object */
    Object tiers
}
