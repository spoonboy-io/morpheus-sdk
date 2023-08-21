package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.BillingInstancesDatastore;

@Canonical
class BillingInstancesVolumes {
    
    Long size
    
    String typeCode
    
    BillingInstancesDatastore datastore
}
