package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.ApiJobsIdJobTargets;
import org.openapitools.model.ApiJobsIdJobTask;
import org.openapitools.model.OneOfstringlong;

@Canonical
class ApiJobsIdJob {
    /* A name for the Job */
    String name
    /* Array of label strings, can be used for filtering. */
    List<String> labels = new ArrayList<String>()
    /* Use this to set enabled state */
    Boolean enabled = true
    
    ApiJobsIdJobTask task
    
    ApiJobsIdJobTask workflow
    /* Scan Checklist. Only applies to type scap-package. */
    String scanPath
    /* Security Profile. Only applies to type scap-package. */
    String securityProfile
    /* Target type where job will execute */
    String targetType
    
    List<ApiJobsIdJobTargets> targets = new ArrayList<ApiJobsIdJobTargets>()
    /* Instance Label. Only applicable if `targetType` is `instance-label`. */
    String instanceLabel
    /* Server Label. Only applicable if `targetType` is `server-label`. */
    String serverLabel
    
    OneOfstringlong scheduleMode = null
    /* Map of options to be used as values in the workflow tasks. These correspond to option types. */
    Object customOptions
    /* Job custom configuration (String in JSON format) */
    String customConfig
    /* Date and Time to execute the job. Use UTC time in the format 2020-02-15T05:00:00Z. Required when scheduleMode is 'dateTime'. */
    Date dateTime
    /* If true, executes job */
    Boolean run
}
