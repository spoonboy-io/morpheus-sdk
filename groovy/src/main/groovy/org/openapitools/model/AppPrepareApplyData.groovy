package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.AppPrepareApplyDataGroup;
import org.openapitools.model.AppPrepareApplyDataTerraform;

@Canonical
class AppPrepareApplyData {
    
    String image
    
    String name
    
    Boolean autoValidate
    
    AppPrepareApplyDataTerraform terraform
    
    String type
    
    Object config
    
    String blueprintName
    
    String description
    
    Long templateId
    
    Long blueprintId
    
    AppPrepareApplyDataGroup group
}
