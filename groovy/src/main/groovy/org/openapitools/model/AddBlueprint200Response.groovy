package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.jackson.nullable.JsonNullable;
import org.openapitools.model.BlueprintCreateSuccess;

@Canonical
class AddBlueprint200Response {
    
    Boolean success
    
    BlueprintCreateSuccess blueprint
    
    String msg
    
    String errors
    
    String errorCode
    
    Boolean inProgress
}
