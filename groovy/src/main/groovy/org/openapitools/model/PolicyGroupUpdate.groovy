package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.PolicyGroupUpdatePolicyType;

@Canonical
class PolicyGroupUpdate {
    /* A name for the policy */
    String name
    /* A description for the policy */
    String description
    
    PolicyGroupUpdatePolicyType policyType
}
