package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.BlueprintARMCreateArm;

@Canonical
class BlueprintARMCreate {
    /* A name for the blueprint */
    String name
    /* Path to display image. Defaults to an internal Morpheus image. */
    String image
    /* Blueprint Type */
    String type
    /* Array of label strings, can be used for filtering. */
    List<String> labels = new ArrayList<String>()
    
    BlueprintARMCreateArm arm
}
