package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.StorageServerTypeOptionTypes;

@Canonical
class StorageVolumeType {
    
    Long id
    
    String code
    
    String name
    
    String description
    
    Long displayOrder
    
    Boolean defaultType
    
    Boolean customLabel
    
    Boolean customSize
    
    String customSizeOptions
    
    Boolean configurableIOPS
    
    Boolean hasDatastore
    
    String category
    
    Boolean enabled
    
    List<StorageServerTypeOptionTypes> optionTypes = new ArrayList<StorageServerTypeOptionTypes>()
}
