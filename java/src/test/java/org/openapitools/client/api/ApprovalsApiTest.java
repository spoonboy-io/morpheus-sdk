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
import org.openapitools.client.model.ApprovalItem;
import org.openapitools.client.model.DefaultError;
import org.openapitools.client.model.GetApprovals200Response;
import org.openapitools.client.model.ListApprovals200Response;
import org.openapitools.client.model.Model200Success;
import org.junit.jupiter.api.Disabled;
import org.junit.jupiter.api.Test;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * API tests for ApprovalsApi
 */
@Disabled
public class ApprovalsApiTest {

    private final ApprovalsApi api = new ApprovalsApi();

    /**
     * Retrieves a Specific Approval
     *
     * Retrieves a specific approval. 
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void getApprovalsTest() throws ApiException {
        Long id = null;
        GetApprovals200Response response = api.getApprovals(id);
        // TODO: test validations
    }

    /**
     * Retrieves a Specific Approval Item
     *
     * Retrieves a specific approval item. 
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void getApprovalsItemTest() throws ApiException {
        Long id = null;
        ApprovalItem response = api.getApprovalsItem(id);
        // TODO: test validations
    }

    /**
     * Retrieves all Approvals
     *
     * Retrieves all approvals. 
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void listApprovalsTest() throws ApiException {
        Long max = null;
        Long offset = null;
        String sort = null;
        String direction = null;
        String phrase = null;
        String name = null;
        ListApprovals200Response response = api.listApprovals(max, offset, sort, direction, phrase, name);
        // TODO: test validations
    }

    /**
     * Updates a Specific Approval Item
     *
     * Updates a specific approval item. 
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void updateApprovalsItemTest() throws ApiException {
        Long id = null;
        String action = null;
        Model200Success response = api.updateApprovalsItem(id, action);
        // TODO: test validations
    }

}