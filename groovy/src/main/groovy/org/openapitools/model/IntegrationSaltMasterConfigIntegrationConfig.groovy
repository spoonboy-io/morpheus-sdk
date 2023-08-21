package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class IntegrationSaltMasterConfigIntegrationConfig {
    /* Apply state via Minion instead of Master (salt-call) */
    Boolean saltApplyOnMinion
}
