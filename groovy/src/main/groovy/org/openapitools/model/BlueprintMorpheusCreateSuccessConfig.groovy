package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class BlueprintMorpheusCreateSuccessConfig {
    /* Path to display image. Defaults to an internal Morpheus image. */
    String image
    /* A name for the blueprint */
    String name
    /* Blueprint Type */
    String type
    /* Tier definitions - Create in UI to view a baseline for object */
    Object tiers
}
