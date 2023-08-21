package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.BlueprintARMCreateArmGit;
import org.openapitools.model.OneOfbooleanstring;

@Canonical
class BlueprintARMCreateArm {
    /* Configuration Type */
    String configType
    /* ARM Template in JSON */
    String json
    /* ARM Template in YAML */
    String yaml
    
    BlueprintARMCreateArmGit git
    /* OS Type */
    String osType
    /* Install Morpheus Agent */
    OneOfbooleanstring installAgent = null
    /* Cloud Init Enabled */
    OneOfbooleanstring cloudInitEnabled = null
}
