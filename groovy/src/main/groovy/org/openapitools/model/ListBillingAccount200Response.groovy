package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.Billing;

@Canonical
class ListBillingAccount200Response {
    
    Boolean success
    
    Billing billingInfo
}
