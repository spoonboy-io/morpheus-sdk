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
import org.openapitools.client.model.InlineObject50;
import org.openapitools.client.model.InlineObject51;
import org.openapitools.client.model.InlineResponse20025;
import org.openapitools.client.model.Model200Success;
import org.openapitools.client.model.SuccessId;
import org.junit.Test;
import org.junit.Ignore;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * API tests for ClusterLayoutsApi
 */
@Ignore
public class ClusterLayoutsApiTest {

    private final ClusterLayoutsApi api = new ClusterLayoutsApi();

    
    /**
     * Clone a Cluster Layout
     *
     * Use this command to clone a cluster layout.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void addClusterLayoutCloneTest() throws ApiException {
        Long id = null;
        String name = null;
        String description = null;
        String computeVersion = null;
        SuccessId response = api.addClusterLayoutClone(id, name, description, computeVersion);

        // TODO: test validations
    }
    
    /**
     * Create a Cluster Layout
     *
     * Use this command to create a cluster layout.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void addClusterLayoutsTest() throws ApiException {
        InlineObject50 inlineObject50 = null;
        SuccessId response = api.addClusterLayouts(inlineObject50);

        // TODO: test validations
    }
    
    /**
     * Delete a Cluster Layout
     *
     * Will delete a cluster layout
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void deleteClusterLayoutTest() throws ApiException {
        Long id = null;
        Model200Success response = api.deleteClusterLayout(id);

        // TODO: test validations
    }
    
    /**
     * Get a Specific Cluster Layout
     *
     * This endpoint retrieves a specific cluster layout.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getClusterLayoutTest() throws ApiException {
        Long id = null;
        InlineResponse20025 response = api.getClusterLayout(id);

        // TODO: test validations
    }
    
    /**
     * Get All Cluster Layouts
     *
     * This endpoint retrieves all cluster layouts.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void listClusterLayoutsTest() throws ApiException {
        Long max = null;
        Long offset = null;
        String sort = null;
        String direction = null;
        String phrase = null;
        String provisionType = null;
        String labels = null;
        String allLabels = null;
        Object response = api.listClusterLayouts(max, offset, sort, direction, phrase, provisionType, labels, allLabels);

        // TODO: test validations
    }
    
    /**
     * Update a Cluster Layout
     *
     * Use this command to update an existing cluster layout.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void updateClusterLayoutTest() throws ApiException {
        Long id = null;
        InlineObject51 inlineObject51 = null;
        SuccessId response = api.updateClusterLayout(id, inlineObject51);

        // TODO: test validations
    }
    
}
