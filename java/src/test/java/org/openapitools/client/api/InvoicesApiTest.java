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
import org.openapitools.client.model.InlineObject102;
import org.junit.Test;
import org.junit.Ignore;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * API tests for InvoicesApi
 */
@Ignore
public class InvoicesApiTest {

    private final InvoicesApi api = new InvoicesApi();

    
    /**
     * Get a Specific Invoice Line Item
     *
     * Get details about a specific invoice line item.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getInvoiceLineItemsTest() throws ApiException {
        Long id = null;
        Object response = api.getInvoiceLineItems(id);

        // TODO: test validations
    }
    
    /**
     * Get a Specific Invoice
     *
     * Get details about a specific invoice.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getInvoicesTest() throws ApiException {
        Long id = null;
        Object response = api.getInvoices(id);

        // TODO: test validations
    }
    
    /**
     * List All Invoice Line Items
     *
     * This endpoint retrieves a list of all invoice line items for the specified parameters.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void listInvoiceLineItemsTest() throws ApiException {
        Long max = null;
        Long offset = null;
        String sort = null;
        String direction = null;
        String phrase = null;
        String name = null;
        String startDate = null;
        String endDate = null;
        String period = null;
        String refType = null;
        Long refId = null;
        Long zoneId = null;
        Long siteId = null;
        Long instanceId = null;
        Long containerId = null;
        Long serverId = null;
        Long userId = null;
        Long projectId = null;
        Boolean active = null;
        Long accountId = null;
        Boolean includeTotals = null;
        Object response = api.listInvoiceLineItems(max, offset, sort, direction, phrase, name, startDate, endDate, period, refType, refId, zoneId, siteId, instanceId, containerId, serverId, userId, projectId, active, accountId, includeTotals);

        // TODO: test validations
    }
    
    /**
     * List All Invoices
     *
     * This endpoint retrieves a list of invoices for the specified parameters.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void listInvoicesTest() throws ApiException {
        Long max = null;
        Long offset = null;
        String sort = null;
        String direction = null;
        String phrase = null;
        String name = null;
        String startDate = null;
        String endDate = null;
        String period = null;
        String refType = null;
        Long refId = null;
        String refStatus = null;
        Long zoneId = null;
        Long siteId = null;
        Long instanceId = null;
        Long containerId = null;
        Long serverId = null;
        Long userId = null;
        Long projectId = null;
        Boolean active = null;
        Long accountId = null;
        Boolean includeLineItems = null;
        Boolean includeTotals = null;
        String tags = null;
        Object response = api.listInvoices(max, offset, sort, direction, phrase, name, startDate, endDate, period, refType, refId, refStatus, zoneId, siteId, instanceId, containerId, serverId, userId, projectId, active, accountId, includeLineItems, includeTotals, tags);

        // TODO: test validations
    }
    
    /**
     * Update Invoice Tags
     *
     * Update, Add, or Remove invoice tag(s).
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void updateInvoicesTest() throws ApiException {
        Long id = null;
        InlineObject102 inlineObject102 = null;
        Object response = api.updateInvoices(id, inlineObject102);

        // TODO: test validations
    }
    
}
