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
import org.openapitools.client.model.InlineObject267;
import org.openapitools.client.model.InlineObject268;
import org.openapitools.client.model.InlineObject269;
import org.openapitools.client.model.InlineObject270;
import org.openapitools.client.model.InlineObject271;
import org.openapitools.client.model.InlineObject272;
import org.openapitools.client.model.InlineObject273;
import org.openapitools.client.model.InlineObject274;
import org.openapitools.client.model.InlineResponse200168;
import org.openapitools.client.model.InlineResponse200169;
import org.openapitools.client.model.Model200Success;
import org.junit.Test;
import org.junit.Ignore;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * API tests for WikiApi
 */
@Ignore
public class WikiApiTest {

    private final WikiApi api = new WikiApi();

    
    /**
     * Create a Wiki Page
     *
     * Creates a Wiki Page 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void addWikiTest() throws ApiException {
        InlineObject272 inlineObject272 = null;
        Object response = api.addWiki(inlineObject272);

        // TODO: test validations
    }
    
    /**
     * Retrieves a specific Wiki page
     *
     * This endpoint retrieves a specific wiki page. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getWikiTest() throws ApiException {
        Long id = null;
        InlineResponse200168 response = api.getWiki(id);

        // TODO: test validations
    }
    
    /**
     * Retrieves an App Wiki Page
     *
     * This endpoint retrieves an app Wiki page. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getWikiAppTest() throws ApiException {
        Long id = null;
        InlineResponse200168 response = api.getWikiApp(id);

        // TODO: test validations
    }
    
    /**
     * Retrieves all Wiki categories associated with the account
     *
     * This endpoint retrieves all Wiki categories associated with the account. The results are not paginated. The categories returned are those of the found pages. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getWikiCategoriesTest() throws ApiException {
        String phrase = null;
        String pagePhrase = null;
        InlineResponse200169 response = api.getWikiCategories(phrase, pagePhrase);

        // TODO: test validations
    }
    
    /**
     * Retrieves a Cloud Wiki Page
     *
     * This endpoint retrieves a cloud Wiki page. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getWikiCloudTest() throws ApiException {
        Long id = null;
        InlineResponse200168 response = api.getWikiCloud(id);

        // TODO: test validations
    }
    
    /**
     * Retrieves a Cluster Wiki Page
     *
     * This endpoint retrieves a cluster Wiki page. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getWikiClusterTest() throws ApiException {
        Integer clusterId = null;
        InlineResponse200168 response = api.getWikiCluster(clusterId);

        // TODO: test validations
    }
    
    /**
     * Retrieves a Group Wiki Page
     *
     * This endpoint retrieves a group Wiki page. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getWikiGroupTest() throws ApiException {
        Long id = null;
        InlineResponse200168 response = api.getWikiGroup(id);

        // TODO: test validations
    }
    
    /**
     * Retrieves an Instance Wiki Page
     *
     * This endpoint retrieves an instance Wiki page. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getWikiInstanceTest() throws ApiException {
        Long id = null;
        InlineResponse200168 response = api.getWikiInstance(id);

        // TODO: test validations
    }
    
    /**
     * Retrieves a Server Wiki Page
     *
     * This endpoint retrieves a server Wiki page. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getWikiServerTest() throws ApiException {
        Long id = null;
        InlineResponse200168 response = api.getWikiServer(id);

        // TODO: test validations
    }
    
    /**
     * Retrieves wiki pages associated with the account.
     *
     * This endpoint retrieves wiki pages associated with the account. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void listWikiTest() throws ApiException {
        String name = null;
        String phrase = null;
        InlineResponse200168 response = api.listWiki(name, phrase);

        // TODO: test validations
    }
    
    /**
     * Deletes a Wiki Page
     *
     * Deletes the specified Wiki page. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void removeWikiTest() throws ApiException {
        Long id = null;
        Model200Success response = api.removeWiki(id);

        // TODO: test validations
    }
    
    /**
     * Updates a Wiki Page
     *
     * Updates a Wiki Page 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void updateWikiTest() throws ApiException {
        Long id = null;
        InlineObject273 inlineObject273 = null;
        Object response = api.updateWiki(id, inlineObject273);

        // TODO: test validations
    }
    
    /**
     * Update an App Wiki Page
     *
     * Updates an app Wiki page. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void updateWikiAppTest() throws ApiException {
        Long id = null;
        InlineObject267 inlineObject267 = null;
        Object response = api.updateWikiApp(id, inlineObject267);

        // TODO: test validations
    }
    
    /**
     * Update a Cloud Wiki Page
     *
     * Updates a cloud Wiki page. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void updateWikiCloudTest() throws ApiException {
        Long id = null;
        InlineObject274 inlineObject274 = null;
        Object response = api.updateWikiCloud(id, inlineObject274);

        // TODO: test validations
    }
    
    /**
     * Update a Cluster Wiki Page
     *
     * Updates a cluster Wiki page. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void updateWikiClusterTest() throws ApiException {
        Integer clusterId = null;
        InlineObject268 inlineObject268 = null;
        Object response = api.updateWikiCluster(clusterId, inlineObject268);

        // TODO: test validations
    }
    
    /**
     * Update a Group Wiki Page
     *
     * Updates a group Wiki page. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void updateWikiGroupTest() throws ApiException {
        Long id = null;
        InlineObject269 inlineObject269 = null;
        Object response = api.updateWikiGroup(id, inlineObject269);

        // TODO: test validations
    }
    
    /**
     * Update an Instance Wiki Page
     *
     * Updates an instance Wiki page. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void updateWikiInstanceTest() throws ApiException {
        Long id = null;
        InlineObject270 inlineObject270 = null;
        Object response = api.updateWikiInstance(id, inlineObject270);

        // TODO: test validations
    }
    
    /**
     * Update a Server Wiki Page
     *
     * Updates a server Wiki page. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void updateWikiServerTest() throws ApiException {
        Long id = null;
        InlineObject271 inlineObject271 = null;
        Object response = api.updateWikiServer(id, inlineObject271);

        // TODO: test validations
    }
    
}