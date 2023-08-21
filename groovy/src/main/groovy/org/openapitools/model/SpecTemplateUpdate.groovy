package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.SpecTemplateUpdateConfig;
import org.openapitools.model.SpecTemplateUpdateFile;
import org.openapitools.model.SpecTemplateUpdateType;

@Canonical
class SpecTemplateUpdate {
    /* Spec template name */
    String name
    /* Array of label strings, can be used for filtering. */
    List<String> labels = new ArrayList<String>()
    
    SpecTemplateUpdateType type
    
    SpecTemplateUpdateFile file
    
    SpecTemplateUpdateConfig config
}
