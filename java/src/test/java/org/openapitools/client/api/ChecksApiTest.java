/*
 * Morpheus API
 * Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.
 *
 * The version of the OpenAPI document: 6.1.1
 * Contact: dev@morpheusdata.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package org.openapitools.client.api;

import org.openapitools.client.ApiException;
import org.openapitools.client.model.AddCheckApps200Response;
import org.openapitools.client.model.AddCheckAppsRequest;
import org.openapitools.client.model.DefaultError;
import org.openapitools.client.model.ListCheckApps200Response;
import java.time.OffsetDateTime;
import org.junit.jupiter.api.Disabled;
import org.junit.jupiter.api.Test;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * API tests for ChecksApi
 */
@Disabled
public class ChecksApiTest {

    private final ChecksApi api = new ChecksApi();

    /**
     * Create a New Check App
     *
     * Create a new check app.
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void addCheckAppsTest() throws ApiException {
        AddCheckAppsRequest addCheckAppsRequest = null;
        AddCheckApps200Response response = api.addCheckApps(addCheckAppsRequest);
        // TODO: test validations
    }

    /**
     * List All Check Apps
     *
     * Get a list of check apps.
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void listCheckAppsTest() throws ApiException {
        Long max = null;
        Long offset = null;
        String sort = null;
        String name = null;
        String phrase = null;
        String status = null;
        OffsetDateTime lastUpdated = null;
        Boolean deleted = null;
        ListCheckApps200Response response = api.listCheckApps(max, offset, sort, name, phrase, status, lastUpdated, deleted);
        // TODO: test validations
    }

}