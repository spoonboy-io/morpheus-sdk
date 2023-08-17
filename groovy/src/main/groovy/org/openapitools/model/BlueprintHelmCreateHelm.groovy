package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.BlueprintHelmCreateHelmGit;

@Canonical
class BlueprintHelmCreateHelm {

    enum ConfigTypeEnum {
    
        GIT("git")
    
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
    
    BlueprintHelmCreateHelmGit git
}
