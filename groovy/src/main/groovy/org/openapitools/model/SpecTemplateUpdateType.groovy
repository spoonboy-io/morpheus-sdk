package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class SpecTemplateUpdateType {
    /* Spec Template Type. i.e. arm, cloudFormation, helm, kubernetes, oneview, terraform, ucs. */
    String code
}
