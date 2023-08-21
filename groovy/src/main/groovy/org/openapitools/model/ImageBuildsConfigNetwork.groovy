package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.AppPrepareApplyDataGroup;

@Canonical
class ImageBuildsConfigNetwork {
    
    String idName
    
    AppPrepareApplyDataGroup pool
    
    String id
    
    Boolean hasPool
}
