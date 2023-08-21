package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.OptionType;

@Canonical
class SecurityPackageType {
    
    Long id
    
    String code
    
    String name
    
    String description
    
    Boolean enabled
    
    List<OptionType> optionTypes = new ArrayList<OptionType>()
}
