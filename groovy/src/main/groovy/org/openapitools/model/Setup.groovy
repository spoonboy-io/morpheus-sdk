package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class Setup {
    
    Boolean success
    /* Morpheus build version that the server is running. */
    String buildVersion
    /* The Appliance Server URL as defined under Appliance Settings. */
    String applianceUrl
    /* Flag to determine if the appliance has been setup, only true when appliance is a fresh install and has not been initialized. */
    Boolean setupNeeded
}
