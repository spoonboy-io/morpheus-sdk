package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.SpecTemplateCreateConfig;
import org.openapitools.model.SpecTemplateCreateFile;
import org.openapitools.model.SpecTemplateCreateType;

@Canonical
class SpecTemplateCreate {
    /* Spec template name */
    String name
    /* Array of label strings, can be used for filtering. */
    List<String> labels = new ArrayList<String>()
    
    SpecTemplateCreateType type
    
    SpecTemplateCreateFile file
    
    SpecTemplateCreateConfig config
}
