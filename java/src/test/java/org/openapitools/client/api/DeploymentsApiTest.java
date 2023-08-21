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
import java.io.File;
import org.openapitools.client.model.InlineObject67;
import org.openapitools.client.model.InlineObject68;
import org.openapitools.client.model.InlineObject69;
import org.openapitools.client.model.InlineObject70;
import org.openapitools.client.model.InlineResponse20038;
import org.openapitools.client.model.InlineResponse20039;
import org.openapitools.client.model.Model200Success;
import org.threeten.bp.OffsetDateTime;
import org.junit.Test;
import org.junit.Ignore;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * API tests for DeploymentsApi
 */
@Ignore
public class DeploymentsApiTest {

    private final DeploymentsApi api = new DeploymentsApi();

    
    /**
     * Upload a Deployment File
     *
     * This endpoint will upload a file for a specific deployment version. This will overwrite the file if one with the same name exists already.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void addDeploymentFileTest() throws ApiException {
        Long deploymentId = null;
        Long id = null;
        String filepath = null;
        File file = null;
        Model200Success response = api.addDeploymentFile(deploymentId, id, filepath, file);

        // TODO: test validations
    }
    
    /**
     * Create a new Deployment Version
     *
     * This endpoint will create a new deployment version that is ready to have files uploaded to it. The default type is file, which has files directly uploaded via Morpheus. Alternatively, the type git or fetch can be used to just point to a repository or remote url.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void addDeploymentVersionTest() throws ApiException {
        Long deploymentId = null;
        InlineObject69 inlineObject69 = null;
        Object response = api.addDeploymentVersion(deploymentId, inlineObject69);

        // TODO: test validations
    }
    
    /**
     * Create a new Deployment
     *
     * This endpoint will create a new deployment that is ready to have versions added to it.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void addDeploymentsTest() throws ApiException {
        InlineObject67 inlineObject67 = null;
        Object response = api.addDeployments(inlineObject67);

        // TODO: test validations
    }
    
    /**
     * Delete a Deployment
     *
     * This endpoint will delete an existing deployment.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void deleteDeploymentTest() throws ApiException {
        Long deploymentId = null;
        Model200Success response = api.deleteDeployment(deploymentId);

        // TODO: test validations
    }
    
    /**
     * Delete a Deployment File
     *
     * This endpoint will delete an existing deployment file. To recursively delete a directory and all of its contents, the force parameter must be specified.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void deleteDeploymentFileTest() throws ApiException {
        Long deploymentId = null;
        Long id = null;
        String filepath = null;
        String force = null;
        Model200Success response = api.deleteDeploymentFile(deploymentId, id, filepath, force);

        // TODO: test validations
    }
    
    /**
     * Delete a Deployment Version
     *
     * This endpoint will delete an existing deployment version.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void deleteDeploymentVersionTest() throws ApiException {
        Long deploymentId = null;
        Long id = null;
        Model200Success response = api.deleteDeploymentVersion(deploymentId, id);

        // TODO: test validations
    }
    
    /**
     * Get a Specific Deployment
     *
     * This endpoint retrieves a specific deployment. By default the 5 most recent versions are returned, more can be returned by specifying the maxVersions parameter.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getDeploymentTest() throws ApiException {
        Long deploymentId = null;
        Long maxVersions = null;
        InlineResponse20038 response = api.getDeployment(deploymentId, maxVersions);

        // TODO: test validations
    }
    
    /**
     * Get a Specific Deployment Version
     *
     * This endpoint retrieves a specific deployment version.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getDeploymentVersionTest() throws ApiException {
        Long deploymentId = null;
        Long id = null;
        InlineResponse20039 response = api.getDeploymentVersion(deploymentId, id);

        // TODO: test validations
    }
    
    /**
     * List Deployment Files
     *
     * This endpoint returns a list of files for a specific deployment version. This only applies to deploy type file. Files are sorted alphabetically, with directories appearing at the beginning of the list.  The filepath parameter can be specified to search for specific files or directories. To list files under a directory, use a trailing / in the filepath parameter.  To list a specific file, provide it&#39;s full path. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void listDeploymentFilesTest() throws ApiException {
        Long deploymentId = null;
        Long id = null;
        String filepath = null;
        Long max = null;
        Long offset = null;
        String phrase = null;
        Long version = null;
        Object response = api.listDeploymentFiles(deploymentId, id, filepath, max, offset, phrase, version);

        // TODO: test validations
    }
    
    /**
     * Get All Versions For a Deployment
     *
     * This endpoint returns a paginated list of versions for a specific deployment.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void listDeploymentVersionsTest() throws ApiException {
        Long deploymentId = null;
        Long max = null;
        Long offset = null;
        String phrase = null;
        Long version = null;
        String type = null;
        String dateCreated = null;
        OffsetDateTime lastUpdated = null;
        Object response = api.listDeploymentVersions(deploymentId, max, offset, phrase, version, type, dateCreated, lastUpdated);

        // TODO: test validations
    }
    
    /**
     * Get All Deployments
     *
     * This endpoint returns a paginated list of deployments.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void listDeploymentsTest() throws ApiException {
        Long max = null;
        Long offset = null;
        String phrase = null;
        String name = null;
        String description = null;
        String dateCreated = null;
        OffsetDateTime lastUpdated = null;
        Object response = api.listDeployments(max, offset, phrase, name, description, dateCreated, lastUpdated);

        // TODO: test validations
    }
    
    /**
     * Updating a Deployment
     *
     * This endpoint will update an existing deployment.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void updateDeploymentTest() throws ApiException {
        Long deploymentId = null;
        InlineObject68 inlineObject68 = null;
        Object response = api.updateDeployment(deploymentId, inlineObject68);

        // TODO: test validations
    }
    
    /**
     * Updating a Deployment Version
     *
     * This endpoint will update an existing deployment version.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void updateDeploymentVersionTest() throws ApiException {
        Long deploymentId = null;
        Long id = null;
        InlineObject70 inlineObject70 = null;
        Object response = api.updateDeploymentVersion(deploymentId, id, inlineObject70);

        // TODO: test validations
    }
    
}