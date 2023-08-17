package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.Arrays;
import org.openapitools.model.BlueprintKubernetesCreateConfigSpecsInner;

@Canonical
class BlueprintTerraformCreateConfig {
    /* Array of Terraform specs in Morpheus */
    List<BlueprintKubernetesCreateConfigSpecsInner> specs
}