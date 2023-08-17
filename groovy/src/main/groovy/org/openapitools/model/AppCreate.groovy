package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.Arrays;
import org.openapitools.jackson.nullable.JsonNullable;
import org.openapitools.model.AppCreateBlueprintId;
import org.openapitools.model.AppCreateDefaultCloud;
import org.openapitools.model.AppCreateGroup;

@Canonical
class AppCreate {
    
    Long templateId
    
    AppCreateBlueprintId blueprintId
    /* A unique name for the app */
    String name
    /* Description */
    String description
    /* Array of label strings, can be used for filtering. */
    List<String> labels
    
    AppCreateGroup group
    
    AppCreateDefaultCloud defaultCloud
    /* Environment code (appContext) */
    String environment
    /* Configuration of app elements */
    Object tiers
}
