package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.CatalogOrderCreateType;

@Canonical
class CatalogOrderCreateItems {
    
    CatalogOrderCreateType type
    /* Config Object, required options depend on the catalog item type's associated option types. The values passed in here are injected into the instance config or app spec or workflow script(s) defined by the type.  */
    Object config
    /* Context Type for running the workflow, determines if a target resource must be selected. `instance`, `server`, or `appliance`. This may only be passed if the type allows it, usually the type determines the context for the user. Only applies to type `workflow`.  */
    String context
    /* Resource (Instance or Server) ID for context when running the `workflow`. Only applies to type `workflow` and only required when context is `instance` or `server`.  */
    Long target
}
