package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class StorageType {
    
    Long id
    
    String code
    
    String name
    
    Long displayOrder
    
    Boolean defaultType
    
    Boolean customLabel
    
    Boolean customSize
    
    String customSizeOptions
}
