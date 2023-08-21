package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;

@Canonical
class Creds {
    
    Long id
    
    String name
    
    String type
    
    List<String> types = new ArrayList<String>()
}
