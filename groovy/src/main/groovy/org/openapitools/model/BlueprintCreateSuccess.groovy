package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.Arrays;
import org.openapitools.jackson.nullable.JsonNullable;
import org.openapitools.model.BlueprintCreateSuccessConfig;

@Canonical
class BlueprintCreateSuccess {
    /* Blueprint ID */
    Long id
    /* A name for the blueprint */
    String name
    /* A description for the blueprint */
    String description
    
    List<String> labels
    /* Category */
    String category
    
    BlueprintCreateSuccessConfig config
}
