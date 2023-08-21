package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.OptionType;

@Canonical
class PolicyType {
    
    Long id
    
    String code
    
    String name
    
    String description
    
    String category
    
    String loadMethod
    
    String enforceMethod
    
    String prepareMethod
    
    String validateMethod
    
    Boolean enforceOnProvision
    
    Boolean enforceOnManaged
    
    List<OptionType> optionTypes = new ArrayList<OptionType>()
}
