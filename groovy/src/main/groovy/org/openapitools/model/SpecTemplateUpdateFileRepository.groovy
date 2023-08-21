package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class SpecTemplateUpdateFileRepository {
    /* Code Repository ID, required for type repository. Use `/api/options/codeRepositories` to see available repositories. */
    Long id
}
