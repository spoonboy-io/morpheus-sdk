package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.OptionTypeListCreateConfig;
import org.openapitools.model.OptionTypeListCreateCredential;

@Canonical
class OptionTypeListCreate {
    /* Name */
    String name
    /* Description */
    String description
    /* Array of label strings, can be used for filtering. */
    List<String> labels = new ArrayList<String>()
    /* Option List Type eg. `rest`, `api`, `ldap` or `manual`. */
    String type = TypeEnum.REST
    /* Source URL, the http(s) URL to request data from. Required when type is rest. */
    String sourceUrl
    /* Visibility */
    String visibility = VisibilityEnum.PRIVATE
    /* Source Method, the HTTP method. */
    String sourceMethod = SourceMethodEnum.GET
    /* Api Type, The code of the api option list to use, eg. clouds, environments, groups, instances, instance-wiki, networks, servicePlans, resourcePools, securityGroups, servers, server-wiki. Required when type is api. */
    String apiType
    /* Ignore SSL Errors. */
    Boolean ignoreSSLErrors = false
    /* Real Time. */
    Boolean realTime = false
    
    OptionTypeListCreateCredential credential
    /* Username for authenticating with Basic Auth when type is rest or ldap username. */
    String serviceUsername
    /* Password for authenticating with Basic Auth when type is rest or ldap password. */
    String servicePassword
    /* Initial Dataset. Create an initial JSON or CSV dataset to be used as the collection for this option list. It should be a list containing objects with properties 'name', and 'value'. Required when type is manual. */
    String initialDataset
    /* Translation Script. Create a js script to translate the result data object into an Array containing objects with properties 'name' and 'value'. The input data is provided as data and the result should be put on the global variable results. */
    String translationScript
    /* Request Script. Create a js script to prepare the request. Return a data object as the body for a post, and return an array containing properties 'name' and 'value' for a get. The input data is provided as data and the result should be put on the global variable results. */
    String requestScript
    
    OptionTypeListCreateConfig config
}
