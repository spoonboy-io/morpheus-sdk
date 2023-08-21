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
import org.openapitools.client.model.InlineObject16;
import org.openapitools.client.model.InlineObject17;
import org.openapitools.client.model.InlineObject18;
import org.openapitools.client.model.InlineObject19;
import org.openapitools.client.model.InlineResponse20010;
import org.openapitools.client.model.InlineResponse20011;
import org.openapitools.client.model.InlineResponse20012;
import org.openapitools.client.model.InlineResponse20013;
import org.openapitools.client.model.Model200Success;
import org.junit.Test;
import org.junit.Ignore;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * API tests for BackupsApi
 */
@Ignore
public class BackupsApiTest {

    private final BackupsApi api = new BackupsApi();

    
    /**
     * Creates a Backup Job
     *
     * This endpoint allows creating a Backup Job. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void addBackupJobsTest() throws ApiException {
        InlineObject18 inlineObject18 = null;
        Object response = api.addBackupJobs(inlineObject18);

        // TODO: test validations
    }
    
    /**
     * Creates a Backup
     *
     * This endpoint allows creating a Backup. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void addBackupsTest() throws ApiException {
        InlineObject16 inlineObject16 = null;
        Object response = api.addBackups(inlineObject16);

        // TODO: test validations
    }
    
    /**
     * Executes a Backup Job
     *
     * Executes a backup job. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void executeBackupJobsTest() throws ApiException {
        Long id = null;
        Object body = null;
        Object response = api.executeBackupJobs(id, body);

        // TODO: test validations
    }
    
    /**
     * Executes a Backup
     *
     * Executes a backup. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void executeBackupsTest() throws ApiException {
        Long id = null;
        Object body = null;
        Object response = api.executeBackups(id, body);

        // TODO: test validations
    }
    
    /**
     * Retrieves a Specific Backup Job
     *
     * Retrieves a specific backup job. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getBackupJobsTest() throws ApiException {
        Long id = null;
        InlineResponse20011 response = api.getBackupJobs(id);

        // TODO: test validations
    }
    
    /**
     * Retrieves a Specific Backup Restore
     *
     * Retrieves a specific backup restore. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getBackupRestoresTest() throws ApiException {
        Long id = null;
        InlineResponse20013 response = api.getBackupRestores(id);

        // TODO: test validations
    }
    
    /**
     * Retrieves a Specific Backup Result
     *
     * Retrieves a specific backup result. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getBackupResultsTest() throws ApiException {
        Long id = null;
        InlineResponse20012 response = api.getBackupResults(id);

        // TODO: test validations
    }
    
    /**
     * Retrieves a Specific Backup
     *
     * Retrieves a specific backup. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getBackupsTest() throws ApiException {
        Long id = null;
        InlineResponse20010 response = api.getBackups(id);

        // TODO: test validations
    }
    
    /**
     * Retrieves all Backup Jobs
     *
     * Retrieves all backup jobs. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void listBackupJobsTest() throws ApiException {
        Long max = null;
        Long offset = null;
        String sort = null;
        String direction = null;
        String phrase = null;
        String name = null;
        String code = null;
        String externalId = null;
        Object response = api.listBackupJobs(max, offset, sort, direction, phrase, name, code, externalId);

        // TODO: test validations
    }
    
    /**
     * Retrieves all Backup Restores
     *
     * Retrieves all backup restores. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void listBackupRestoresTest() throws ApiException {
        Long max = null;
        Long offset = null;
        String sort = null;
        String direction = null;
        String name = null;
        String phrase = null;
        Long backupId = null;
        String backupResultId = null;
        Long containerId = null;
        Object response = api.listBackupRestores(max, offset, sort, direction, name, phrase, backupId, backupResultId, containerId);

        // TODO: test validations
    }
    
    /**
     * Retrieves all Backup Results
     *
     * Retrieves all backup results. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void listBackupResultsTest() throws ApiException {
        Long max = null;
        Long offset = null;
        String sort = null;
        String direction = null;
        String name = null;
        String phrase = null;
        Long backupId = null;
        String backupSetId = null;
        Long instanceId = null;
        Long containerId = null;
        Long serverId = null;
        Object response = api.listBackupResults(max, offset, sort, direction, name, phrase, backupId, backupSetId, instanceId, containerId, serverId);

        // TODO: test validations
    }
    
    /**
     * Retrieves all Backups
     *
     * Retrieves all backups. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void listBackupsTest() throws ApiException {
        Long max = null;
        Long offset = null;
        String sort = null;
        String direction = null;
        String phrase = null;
        String name = null;
        Object response = api.listBackups(max, offset, sort, direction, phrase, name);

        // TODO: test validations
    }
    
    /**
     * Deletes a Backup Job
     *
     * Deletes a specified backup job. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void removeBackupJobsTest() throws ApiException {
        Long id = null;
        Model200Success response = api.removeBackupJobs(id);

        // TODO: test validations
    }
    
    /**
     * Deletes a Backup Restore
     *
     * Deletes a specified backup restore. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void removeBackupRestoresTest() throws ApiException {
        Long id = null;
        Model200Success response = api.removeBackupRestores(id);

        // TODO: test validations
    }
    
    /**
     * Deletes a Backup Result
     *
     * Deletes a specified backup result. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void removeBackupResultsTest() throws ApiException {
        Long id = null;
        Model200Success response = api.removeBackupResults(id);

        // TODO: test validations
    }
    
    /**
     * Deletes a Backup
     *
     * Deletes a specified backup. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void removeBackupsTest() throws ApiException {
        Long id = null;
        Model200Success response = api.removeBackups(id);

        // TODO: test validations
    }
    
    /**
     * Updates a Backup Job
     *
     * Updates a backup job. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void updateBackupJobsTest() throws ApiException {
        Long id = null;
        InlineObject19 inlineObject19 = null;
        Object response = api.updateBackupJobs(id, inlineObject19);

        // TODO: test validations
    }
    
    /**
     * Updates a Backup
     *
     * Updates a backup. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void updateBackupsTest() throws ApiException {
        Long id = null;
        InlineObject17 inlineObject17 = null;
        Object response = api.updateBackups(id, inlineObject17);

        // TODO: test validations
    }
    
}
