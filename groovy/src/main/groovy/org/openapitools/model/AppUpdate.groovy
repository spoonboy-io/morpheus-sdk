package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;

@Canonical
class AppUpdate {
    /* A unique name for the app */
    String name
    /* Description */
    String description
    /* Array of label strings, can be used for filtering. */
    List<String> labels = new ArrayList<String>()
    /* Environment code (appContext) */
    String environment
    /* User ID, can be used to change app owner. This also changes the owner for each instance in the app. */
    Long ownerId
}
