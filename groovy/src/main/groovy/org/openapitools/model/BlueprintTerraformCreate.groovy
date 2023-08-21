package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.BlueprintTerraformCreateConfig;
import org.openapitools.model.BlueprintTerraformCreateTerraform;

@Canonical
class BlueprintTerraformCreate {
    /* A name for the blueprint */
    String name
    /* Path to display image. Defaults to an internal Morpheus image. */
    String image
    /* Blueprint Type */
    String type
    /* Array of label strings, can be used for filtering. */
    List<String> labels = new ArrayList<String>()
    
    BlueprintTerraformCreateTerraform terraform
    
    BlueprintTerraformCreateConfig config
}
