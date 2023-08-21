package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.SpecTemplateCreateFileRepository;

@Canonical
class SpecTemplateCreateFile {
    /* File Source i.e. local, repository, url. */
    String sourceType = "local"
    /* File content, the template text. Only required when sourceType is `local`. */
    String content
    /* Content Path, the repo file location or url. Required when sourceType is repository or url. */
    String contentPath
    /* Content Ref, the branch/tag. Only used when sourceType is repo. */
    String contentRef
    
    SpecTemplateCreateFileRepository repository
}
