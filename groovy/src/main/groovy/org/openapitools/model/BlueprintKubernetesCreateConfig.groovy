package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.BlueprintKubernetesCreateConfigSpecs;

@Canonical
class BlueprintKubernetesCreateConfig {
    /* Array of Kubernetes specs in Morpheus */
    List<BlueprintKubernetesCreateConfigSpecs> specs = new ArrayList<BlueprintKubernetesCreateConfigSpecs>()
}
