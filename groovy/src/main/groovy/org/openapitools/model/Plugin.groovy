package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.AppStateTemplate;
import org.openapitools.model.OptionType;

@Canonical
class Plugin {
    
    Long id
    
    String name
    
    String code
    
    String description
    
    String version
    
    Boolean enabled
    
    String author
    
    String websiteUrl
    
    String sourceCodeLocationUrl
    
    String issueTrackerUrl
    
    Boolean valid
    
    Boolean hasValidUpdate
    
    String status
    
    String statusMessage
    
    List<AppStateTemplate> providers = new ArrayList<AppStateTemplate>()
    
    Object config
    
    List<OptionType> optionTypes = new ArrayList<OptionType>()
    
    Date dateCreated
    
    Date lastUpdated
}
