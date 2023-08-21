package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;

@Canonical
class BlueprintMorpheusCreate {
    /* A name for the blueprint */
    String name
    /* Blueprint Type */
    String type
    /* Array of label strings, can be used for filtering. */
    List<String> labels = new ArrayList<String>()
    /* Tier definitions - Create in UI to view a baseline for object */
    Object tiers
}
