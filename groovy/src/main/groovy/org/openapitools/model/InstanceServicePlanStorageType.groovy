package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;

@Canonical
class InstanceServicePlanStorageType {
    
    Integer id
    
    Boolean editable
    
    List<Object> optionTypes = new ArrayList<Object>()
    
    Integer displayOrder
    
    String code
    
    String volumeType
    
    String minStorage
    
    Boolean deletable
    
    Boolean defaultType
    
    String createDatastore
    
    Boolean resizable
    
    String storageType
    
    Boolean allowSearch
    
    String volumeOptionSource
    
    String displayName
    
    String minIOPS
    
    String maxIOPS
    
    Boolean hasDatastore
    
    Boolean customSize
    
    Boolean autoDelete
    
    String name
    
    Boolean configurableIOPS
    
    Boolean customLabel
    
    Boolean enabled
    
    String description
    
    String volumeCategory
    
    String externalId
    
    String maxStorage
}
