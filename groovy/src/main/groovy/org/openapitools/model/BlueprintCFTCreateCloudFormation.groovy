package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.BlueprintCFTCreateCloudFormationGit;

@Canonical
class BlueprintCFTCreateCloudFormation {

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
    /* CloudFormation Template in JSON */
    String json
    /* CloudFormation Template in YAML */
    String yaml
    
    BlueprintCFTCreateCloudFormationGit git
    /* CloudFormation Attribute CAPABILITY_IAM */
    Boolean IAM = false
    /* CloudFormation Attribute CAPABILITY_NAMED_IAM */
    Boolean CAPABILITY_NAMED_IAM = false
    /* CloudFormation Attribute CAPABILITY_AUTO_EXPAND */
    Boolean CAPABILITY_AUTO_EXPAND = false
    /* Install Morpheus Agent */
    Boolean installAgent = false
    /* Cloud Init Enabled */
    Boolean cloudInitEnabled = false
}
