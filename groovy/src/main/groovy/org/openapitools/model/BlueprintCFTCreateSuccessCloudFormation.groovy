package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.BlueprintCFTCreateCloudFormationGit;
import org.openapitools.model.OneOfbooleanstring;

@Canonical
class BlueprintCFTCreateSuccessCloudFormation {
    /* Configuration Type */
    String configType
    /* CloudFormation Template in JSON */
    String json
    /* CloudFormation Template in YAML */
    String yaml
    
    BlueprintCFTCreateCloudFormationGit git
    /* CloudFormation Attribute CAPABILITY_IAM */
    OneOfbooleanstring IAM = null
    /* CloudFormation Attribute CAPABILITY_NAMED_IAM */
    OneOfbooleanstring CAPABILITY_NAMED_IAM = null
    /* CloudFormation Attribute CAPABILITY_AUTO_EXPAND */
    OneOfbooleanstring CAPABILITY_AUTO_EXPAND = null
    /* Install Morpheus Agent */
    OneOfbooleanstring installAgent = null
    /* Cloud Init Enabled */
    OneOfbooleanstring cloudInitEnabled = null
}
