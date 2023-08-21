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
import org.openapitools.client.model.Log;
import org.threeten.bp.OffsetDateTime;
import org.junit.Test;
import org.junit.Ignore;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * API tests for LogsApi
 */
@Ignore
public class LogsApiTest {

    private final LogsApi api = new LogsApi();

    
    /**
     * Retrieves Logs
     *
     * Retrieves logs based on filters provided. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void listLogsTest() throws ApiException {
        Long max = null;
        Long offset = null;
        String sort = null;
        String order = null;
        String query = null;
        String message = null;
        String sourceType = null;
        String typeCode = null;
        Long objectId = null;
        String token = null;
        String level = null;
        Long startMs = null;
        Long endMs = null;
        OffsetDateTime startDateTime = null;
        OffsetDateTime endDateTime = null;
        Long containers = null;
        Long servers = null;
        Long clusterId = null;
        Log response = api.listLogs(max, offset, sort, order, query, message, sourceType, typeCode, objectId, token, level, startMs, endMs, startDateTime, endDateTime, containers, servers, clusterId);

        // TODO: test validations
    }
    
}
