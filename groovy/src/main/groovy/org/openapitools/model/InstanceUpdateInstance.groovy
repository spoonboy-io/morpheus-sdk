package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.ApiServersIdMakeManagedServerTags;
import org.openapitools.model.InstanceUpdateInstanceRemoveTags;
import org.openapitools.model.InstanceUpdateInstanceSite;

@Canonical
class InstanceUpdateInstance {
    /* Unique name scoped to your account for the instance. */
    String name
    /* Optional description field. */
    String description
    /* Environment */
    String instanceContext
    /* Array of strings (keywords). */
    List<String> labels = new ArrayList<String>()
    /* Metadata tags, Array of objects having a name and value. */
    List<ApiServersIdMakeManagedServerTags> tags = new ArrayList<ApiServersIdMakeManagedServerTags>()
    /* Add or update value of Metadata tags, Array of objects having a name and value. */
    List<ApiServersIdMakeManagedServerTags> addTags = new ArrayList<ApiServersIdMakeManagedServerTags>()
    /* Remove Metadata tags, Array of objects having a name and an optional value. If value is passed, it must match to be removed. */
    List<InstanceUpdateInstanceRemoveTags> removeTags = new ArrayList<InstanceUpdateInstanceRemoveTags>()
    /* Power schedule ID. */
    Long powerScheduleType
    
    InstanceUpdateInstanceSite site
    /* User ID, can be used to change instance owner. */
    Long ownerId
    /* Name used in the UI for display */
    String displayName
}
