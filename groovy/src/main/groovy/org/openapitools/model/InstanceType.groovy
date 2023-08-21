package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.InlineResponse20040AppDeployInstance;
import org.openapitools.model.InstanceTypeLayout;
import org.openapitools.model.OptionType;

@Canonical
class InstanceType {
    
    Long id
    
    InlineResponse20040AppDeployInstance account
    
    String name
    
    List<String> labels = new ArrayList<String>()
    
    String code
    
    String description
    
    String provisionTypeCode
    
    String category
    
    Boolean active
    
    Boolean hasProvisioningStep
    
    Boolean hasDeployment
    
    Boolean hasConfig
    
    Boolean hasSettings
    
    Boolean hasAutoScale
    
    String proxyType
    
    String proxyPort
    
    String proxyProtocol
    
    String environmentPrefix
    
    String backupType
    
    Object config
    
    String visibility
    
    Boolean featured
    
    List<String> versions = new ArrayList<String>()
    
    List<InstanceTypeLayout> instanceTypeLayouts = new ArrayList<InstanceTypeLayout>()
    
    List<OptionType> optionTypes = new ArrayList<OptionType>()
    
    List<Object> environmentVariables = new ArrayList<Object>()
    
    List<Object> priceSets = new ArrayList<Object>()
    /* Logo image URL */
    String imagePath
    /* Dark logo image URL */
    String darkImagePath
}
