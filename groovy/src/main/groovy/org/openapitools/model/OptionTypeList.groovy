package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.InlineResponse20040AppDeployInstance;
import org.openapitools.model.OptionTypeListConfig;
import org.openapitools.model.OptionTypeListCredential;

@Canonical
class OptionTypeList {
    
    Long id
    
    String name
    
    String description
    
    List<String> labels = new ArrayList<String>()
    
    String type
    
    String sourceUrl
    
    String sourceMethod
    
    String apiType
    
    Boolean ignoreSSLErrors
    
    Boolean realTime
    
    String visibility
    
    OptionTypeListConfig config
    
    OptionTypeListCredential credential
    
    String serviceUsername
    
    String servicePassword
    
    String initialDataset
    
    String translationScript
    
    String requestScript
    
    InlineResponse20040AppDeployInstance account
}
