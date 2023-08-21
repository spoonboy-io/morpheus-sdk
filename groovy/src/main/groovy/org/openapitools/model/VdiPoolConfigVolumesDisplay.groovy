package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class VdiPoolConfigVolumesDisplay {
    
    String storage
    
    String name
    
    String controller
    
    String datastore
    
    String displayOrder
    
    Long size
    
    String mountPoint
}
