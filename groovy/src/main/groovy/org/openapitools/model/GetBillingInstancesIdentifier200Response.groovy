package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.BillingInstance;

@Canonical
class GetBillingInstancesIdentifier200Response {
    
    Boolean success
    
    BillingInstance billingInfo
}
