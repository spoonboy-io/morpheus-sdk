package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.Budget;

@Canonical
class GetBudgets200Response {
    
    Boolean success
    
    Budget budget
}
