package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.AppCreateDefaultCloud;
import org.openapitools.model.AppCreateGroup;
import org.openapitools.model.OneOflongstring;

@Canonical
class AppCreate {
    
    Long templateId
    /* The ID of the Blueprint. Use \"existing\" to create a blank app. */
    OneOflongstring blueprintId = null
    /* A unique name for the app */
    String name
    /* Description */
    String description
    /* Array of label strings, can be used for filtering. */
    List<String> labels = new ArrayList<String>()
    
    AppCreateGroup group
    
    AppCreateDefaultCloud defaultCloud
    /* Environment code (appContext) */
    String environment
    /* Configuration of app elements */
    Object tiers
}
