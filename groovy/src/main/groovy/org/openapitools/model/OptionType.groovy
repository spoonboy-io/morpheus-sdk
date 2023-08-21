package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.InlineResponse20082LoadBalancerInstanceSslCert;

@Canonical
class OptionType {
    
    Long id
    
    String name
    
    List<String> labels = new ArrayList<String>()
    
    String description
    
    String code
    
    String fieldName
    
    String fieldLabel
    
    String fieldCode
    
    String fieldContext
    
    String fieldGroup
    
    String fieldClass
    
    String fieldAddOn
    
    String fieldComponent
    
    String fieldInput
    
    String placeHolder
    
    String verifyPattern
    
    String helpBlock
    
    String helpBlockFieldCode
    
    String defaultValue
    
    String optionSource
    
    String optionSourceType
    
    InlineResponse20082LoadBalancerInstanceSslCert optionList
    
    String type
    
    Boolean advanced
    
    Boolean required
    
    Boolean exportMeta
    
    Boolean editable
    
    Boolean creatable
    
    Object config
    
    Long displayOrder
    
    String wrapperClass
    
    Boolean enabled
    
    Boolean noBlank
    
    String dependsOnCode
    
    String visibleOnCode
    
    String requireOnCode
    
    Boolean contextualDefault
    
    Boolean displayValueOnDetails
    
    Boolean showOnCreate
    
    Boolean showOnEdit
    
    Boolean localCredential
}
