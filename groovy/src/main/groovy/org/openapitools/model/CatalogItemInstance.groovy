package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;

@Canonical
class CatalogItemInstance {
    
    Long id
    
    String name
    
    String status
    
    List<String> locations = new ArrayList<String>()
    
    Long virtualMachines
    
    String version
}
