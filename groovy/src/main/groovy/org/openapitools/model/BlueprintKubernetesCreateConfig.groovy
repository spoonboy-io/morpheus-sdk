package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.Arrays;
import org.openapitools.model.BlueprintKubernetesCreateConfigSpecsInner;

@Canonical
class BlueprintKubernetesCreateConfig {
    /* Array of Kubernetes specs in Morpheus */
    List<BlueprintKubernetesCreateConfigSpecsInner> specs
}
