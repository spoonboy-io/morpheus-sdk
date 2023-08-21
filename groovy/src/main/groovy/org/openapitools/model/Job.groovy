package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.InlineResponse20094Network;
import org.openapitools.model.JobCreatedBy;
import org.openapitools.model.JobCustomOptions;
import org.openapitools.model.JobSecurityPackage;
import org.openapitools.model.JobTargets;
import org.openapitools.model.JobTask;
import org.openapitools.model.JobWorkflow;
import org.openapitools.model.OneOfstringlong;

@Canonical
class Job {
    
    Long id
    
    String name
    
    List<String> labels = new ArrayList<String>()
    
    InlineResponse20094Network type
    
    JobWorkflow workflow
    
    JobTask task
    
    JobSecurityPackage securityPackage
    
    String jobSummary
    
    OneOfstringlong scheduleMode = null
    
    String dateTime
    
    String status
    
    String namespace
    
    String category
    
    String description
    
    Boolean enabled
    
    Date dateCreated
    
    Date lastUpdated
    
    Date lastRun
    
    String lastResult
    
    JobCreatedBy createdBy
    
    String targetType
    
    List<JobTargets> targets = new ArrayList<JobTargets>()
    /* Scan Checklist. Only applies to type scap-package. */
    String scanPath
    /* Security Profile. Only applies to type scap-package. */
    String securityProfile
    
    String customConfig
    
    JobCustomOptions customOptions
}
