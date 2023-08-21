package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;

@Canonical
class IntegrationSNOWConfigIntegrationConfig {
    
    List<Object> serviceNowCustomCmdbMapping = new ArrayList<Object>()
    
    List<Object> serviceNowCmdbClassMapping = new ArrayList<Object>()
    
    List<Object> serviceNowCMDBBusinessObject = new ArrayList<Object>()
    /* Ignore SSL Errors. */
    Boolean ignoreCertErrors = false
}
