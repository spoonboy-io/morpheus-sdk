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
import java.math.BigDecimal;
import org.openapitools.client.model.DefaultError;
import org.threeten.bp.OffsetDateTime;
import org.junit.Test;
import org.junit.Ignore;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * API tests for ActivityApi
 */
@Ignore
public class ActivityApiTest {

    private final ActivityApi api = new ActivityApi();

    
    /**
     * Retrieves Activity
     *
     * Retrieves activity. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void listActivityTest() throws ApiException {
        Long max = null;
        Long offset = null;
        String sort = null;
        String order = null;
        String phrase = null;
        String name = null;
        Long userId = null;
        BigDecimal tenantId = null;
        String timeframe = null;
        OffsetDateTime start = null;
        OffsetDateTime end = null;
        Object response = api.listActivity(max, offset, sort, order, phrase, name, userId, tenantId, timeframe, start, end);

        // TODO: test validations
    }
    
}
