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
import org.openapitools.client.model.InlineObject198;
import org.openapitools.client.model.InlineObject199;
import org.openapitools.client.model.InlineResponse200123;
import org.openapitools.client.model.Model200Success;
import org.junit.Test;
import org.junit.Ignore;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * API tests for PriceSetsApi
 */
@Ignore
public class PriceSetsApiTest {

    private final PriceSetsApi api = new PriceSetsApi();

    
    /**
     * Creates a Price Set
     *
     * Creates a price set. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void addPriceSetsTest() throws ApiException {
        InlineObject198 inlineObject198 = null;
        Object response = api.addPriceSets(inlineObject198);

        // TODO: test validations
    }
    
    /**
     * Deactivates a Price Set
     *
     * Deactivates a price set. This does the same thing as the delete action in the UI, hiding it and making it unavailable to new resources. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void deactivatePriceSetsTest() throws ApiException {
        Long id = null;
        Model200Success response = api.deactivatePriceSets(id);

        // TODO: test validations
    }
    
    /**
     * Retrieves a Specific Price Set
     *
     * Retrieves a specific price set. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getPriceSetsTest() throws ApiException {
        Long id = null;
        InlineResponse200123 response = api.getPriceSets(id);

        // TODO: test validations
    }
    
    /**
     * Retrieves all Price Sets
     *
     * Retrieves all price sets. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void listPriceSetsTest() throws ApiException {
        Long max = null;
        Long offset = null;
        String sort = null;
        String direction = null;
        String phrase = null;
        String name = null;
        Boolean includeInactive = null;
        String type = null;
        Object response = api.listPriceSets(max, offset, sort, direction, phrase, name, includeInactive, type);

        // TODO: test validations
    }
    
    /**
     * Updates a Price Set
     *
     * Updates a price set. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void updatePriceSetsTest() throws ApiException {
        Long id = null;
        InlineObject199 inlineObject199 = null;
        Object response = api.updatePriceSets(id, inlineObject199);

        // TODO: test validations
    }
    
}
