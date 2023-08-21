package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class ClusterServerCreatePlan {
    /* The id for the memory and storage option pre-configured within Morpheus. */
    Long id
    /* The code for the memory and storage option pre-configured within Morpheus. */
    String code
    /* Map of custom options depending on selected service plan . An example would be `maxMemory`, or `maxCores`. */
    Object options
}
