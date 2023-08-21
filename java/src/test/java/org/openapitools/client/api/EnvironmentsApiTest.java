/*
 * Morpheus API
 * Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.
 *
 * The version of the OpenAPI document: 6.2.1
 * Contact: dev@morpheusdata.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package org.openapitools.client.api;

import org.openapitools.client.ApiException;
import org.openapitools.client.model.DefaultError;
import org.openapitools.client.model.InlineObject74;
import org.openapitools.client.model.InlineObject75;
import org.openapitools.client.model.InlineResponse20041;
import org.openapitools.client.model.SuccessError;
import org.junit.Test;
import org.junit.Ignore;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * API tests for EnvironmentsApi
 */
@Ignore
public class EnvironmentsApiTest {

    private final EnvironmentsApi api = new EnvironmentsApi();

    
    /**
     * Create a New Environment
     *
     * Create a new environment.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void addEnvironmentsTest() throws ApiException {
        InlineObject74 inlineObject74 = null;
        Object response = api.addEnvironments(inlineObject74);

        // TODO: test validations
    }
    
    /**
     * Delete a Specific Environment
     *
     * Delete an existing environment.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void deleteEnvironmentsTest() throws ApiException {
        Long id = null;
        SuccessError response = api.deleteEnvironments(id);

        // TODO: test validations
    }
    
    /**
     * Get a Specific Environment
     *
     * This endpoint will retrieve a specific environment by id.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getEnvironmentsTest() throws ApiException {
        Long id = null;
        InlineResponse20041 response = api.getEnvironments(id);

        // TODO: test validations
    }
    
    /**
     * List All Environments
     *
     * This endpoint retrieves all environments.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void listEnvironmentsTest() throws ApiException {
        Long max = null;
        Long offset = null;
        String sort = null;
        String direction = null;
        String phrase = null;
        String name = null;
        String code = null;
        Object response = api.listEnvironments(max, offset, sort, direction, phrase, name, code);

        // TODO: test validations
    }
    
    /**
     * Update Environment
     *
     * Update an existing environment.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void updateEnvironmentsTest() throws ApiException {
        Long id = null;
        InlineObject75 inlineObject75 = null;
        Object response = api.updateEnvironments(id, inlineObject75);

        // TODO: test validations
    }
    
    /**
     * Toggle Active State of Environment
     *
     * Toggle Active State of Environment. Default is to toggle the current value.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void updateEnvironmentsActiveTest() throws ApiException {
        Long id = null;
        Boolean active = null;
        Object response = api.updateEnvironmentsActive(id, active);

        // TODO: test validations
    }
    
}
