package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.BillingInstances;

@Canonical
class ListBillingInstances200Response {
    
    Boolean success
    
    BillingInstances billingInfo
}
