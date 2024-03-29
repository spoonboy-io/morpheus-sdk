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
import org.openapitools.client.model.InlineObject200;
import org.openapitools.client.model.InlineObject201;
import org.openapitools.client.model.InlineResponse200124;
import org.openapitools.client.model.Model200Success;
import org.junit.Test;
import org.junit.Ignore;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * API tests for PricesApi
 */
@Ignore
public class PricesApiTest {

    private final PricesApi api = new PricesApi();

    
    /**
     * Creates a Price
     *
     * Creates a price. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void addPricesTest() throws ApiException {
        InlineObject200 inlineObject200 = null;
        Object response = api.addPrices(inlineObject200);

        // TODO: test validations
    }
    
    /**
     * Deactivates a Price
     *
     * Deactivates a price. This does the same thing as the delete action in the UI, hiding it and making it unavailable to new resources. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void deactivatePricesTest() throws ApiException {
        Long id = null;
        Model200Success response = api.deactivatePrices(id);

        // TODO: test validations
    }
    
    /**
     * Retrieves a Specific Price
     *
     * Retrieves a specific price. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getPricesTest() throws ApiException {
        Long id = null;
        InlineResponse200124 response = api.getPrices(id);

        // TODO: test validations
    }
    
    /**
     * Retrieves all Prices
     *
     * Retrieves all prices. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void listPricesTest() throws ApiException {
        Long max = null;
        Long offset = null;
        String sort = null;
        String direction = null;
        String phrase = null;
        String name = null;
        Boolean includeInactive = null;
        String priceType = null;
        String platform = null;
        String currency = null;
        String type = null;
        Object response = api.listPrices(max, offset, sort, direction, phrase, name, includeInactive, priceType, platform, currency, type);

        // TODO: test validations
    }
    
    /**
     * Updates a Price
     *
     * Updates a price. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void updatePricesTest() throws ApiException {
        Long id = null;
        InlineObject201 inlineObject201 = null;
        Object response = api.updatePrices(id, inlineObject201);

        // TODO: test validations
    }
    
}
