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
import org.openapitools.client.model.InlineObject228;
import org.openapitools.client.model.InlineObject229;
import org.openapitools.client.model.InlineResponse200142;
import org.openapitools.client.model.Model200Success;
import org.junit.Test;
import org.junit.Ignore;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * API tests for ServicePlansApi
 */
@Ignore
public class ServicePlansApiTest {

    private final ServicePlansApi api = new ServicePlansApi();

    
    /**
     * Activates a Service Plan
     *
     * Activates a service plan. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void activateServicePlansTest() throws ApiException {
        Long id = null;
        Model200Success response = api.activateServicePlans(id);

        // TODO: test validations
    }
    
    /**
     * Creates a Service Plan
     *
     * Creates a service plan. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void addServicePlansTest() throws ApiException {
        InlineObject228 inlineObject228 = null;
        Object response = api.addServicePlans(inlineObject228);

        // TODO: test validations
    }
    
    /**
     * Deactivates a Service Plan
     *
     * Deactivates a service plan. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void deactivateServicePlansTest() throws ApiException {
        Long id = null;
        Model200Success response = api.deactivateServicePlans(id);

        // TODO: test validations
    }
    
    /**
     * Retrieves a Specific Service Plan
     *
     * Retrieves a specific service plan. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getServicePlansTest() throws ApiException {
        Long id = null;
        InlineResponse200142 response = api.getServicePlans(id);

        // TODO: test validations
    }
    
    /**
     * Retrieves all Service Plans
     *
     * Retrieves all service plans. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void listServicePlansTest() throws ApiException {
        Long max = null;
        Long offset = null;
        String sort = null;
        String direction = null;
        String phrase = null;
        String name = null;
        Boolean includeZones = null;
        Long provisionTypeId = null;
        Boolean includeInactive = null;
        Object response = api.listServicePlans(max, offset, sort, direction, phrase, name, includeZones, provisionTypeId, includeInactive);

        // TODO: test validations
    }
    
    /**
     * Deletes a Service Plan
     *
     * Deletes a specified service plan. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void removeServicePlansTest() throws ApiException {
        Long id = null;
        Model200Success response = api.removeServicePlans(id);

        // TODO: test validations
    }
    
    /**
     * Updates a Service Plan
     *
     * Updates a service plan. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void updateServicePlansTest() throws ApiException {
        Long id = null;
        InlineObject229 inlineObject229 = null;
        Object response = api.updateServicePlans(id, inlineObject229);

        // TODO: test validations
    }
    
}
