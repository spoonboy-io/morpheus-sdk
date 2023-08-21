package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;

@Canonical
class FileTemplateUpdate {
    /* File template name */
    String name
    /* Array of label strings, can be used for filtering. */
    List<String> labels = new ArrayList<String>()
    /* Filename for the file template */
    String fileName
    /* Path for the file template */
    String filePath
    /* Category */
    String category
    /* Template Phase, provision, start, etc. */
    String templatePhase
    /* Template content, that is, the file template content itself. */
    String template
    /* File Owner */
    Long fileOwner
    /* Setting Name */
    String settingName
    /* Setting Category */
    String settingCategory
}
