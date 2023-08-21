package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.BlueprintCFTCreateCloudFormation;

@Canonical
class BlueprintCFTCreate {
    /* A name for the blueprint */
    String name
    /* Blueprint Type */
    String type
    /* Array of label strings, can be used for filtering. */
    List<String> labels = new ArrayList<String>()
    
    BlueprintCFTCreateCloudFormation cloudFormation
}
