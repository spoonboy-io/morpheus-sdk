package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.BlueprintARMCreateArmCloudInitEnabled;
import org.openapitools.model.BlueprintARMCreateArmGit;
import org.openapitools.model.BlueprintARMCreateArmInstallAgent;

@Canonical
class BlueprintARMCreateArm {

    enum ConfigTypeEnum {
    
        JSON("json"),
        
        YAML("yaml"),
        
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
    /* ARM Template in JSON */
    String json
    /* ARM Template in YAML */
    String yaml
    
    BlueprintARMCreateArmGit git

    enum OsTypeEnum {
    
        LINUX("linux"),
        
        WINDOWS("windows")
    
        private final String value
    
        OsTypeEnum(String value) {
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

    /* OS Type */
    OsTypeEnum osType
    
    BlueprintARMCreateArmInstallAgent installAgent = false
    
    BlueprintARMCreateArmCloudInitEnabled cloudInitEnabled = false
}
