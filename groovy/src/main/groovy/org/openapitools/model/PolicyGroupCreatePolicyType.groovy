package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.OneOfobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobject;

@Canonical
class PolicyGroupCreatePolicyType {
    /* The policy type */
    String code
    /* A map of config values. The expected values vary by policyType. */
    OneOfobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobject config = null
    /* Set to false to disable */
    Boolean enabled = true
    /* Scope object type */
    String refType
    /* Scope object ID (`group`) */
    Long refId
    /* Array of tenants to scope the policy to */
    List<Long> accounts = new ArrayList<Long>()
    /* Apply individually to each user in role.  Only when `refType` equals `Role` */
    Boolean eachUser
}
