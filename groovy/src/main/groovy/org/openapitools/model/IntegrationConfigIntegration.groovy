package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.OneOfobjectobject;

@Canonical
class IntegrationConfigIntegration {
    /* Name, a unique identifier for the integration */
    String name
    /* Integration Type Code */
    String type
    /* Set `true` to enable integration */
    Boolean enabled
    /* Pass `false` to skip refresh.  By default, refresh is done on update, when it is supported by the integration type.  */
    Boolean refresh = true
    /* Map containing Credential ID or the default {\"type\": \"local\"}  which means use the values set in the local task options username and password instead of associating a credential.  */
    OneOfobjectobject credential = null
}
