package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.BlueprintHelmCreateHelmGit;

@Canonical
class BlueprintHelmCreateHelm {
    /* Configuration Type */
    String configType
    
    BlueprintHelmCreateHelmGit git
}
