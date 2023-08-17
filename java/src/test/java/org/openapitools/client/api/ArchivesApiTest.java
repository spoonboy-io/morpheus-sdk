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
import org.openapitools.client.model.AddArchiveBucket200Response;
import org.openapitools.client.model.AddArchiveBucketRequest;
import org.openapitools.client.model.AddArchiveFile200Response;
import org.openapitools.client.model.AddArchiveFileLink200Response;
import org.openapitools.client.model.DefaultError;
import java.io.File;
import org.openapitools.client.model.GetArchiveBucket200Response;
import org.openapitools.client.model.GetArchiveFileDetail200Response;
import org.openapitools.client.model.GetArchiveFileLinks200Response;
import org.openapitools.client.model.ListArchiveBuckets200Response;
import org.openapitools.client.model.ListArchiveFiles200Response;
import org.openapitools.client.model.Model200Success;
import org.openapitools.client.model.UpdateArchiveBucketRequest;
import org.junit.jupiter.api.Disabled;
import org.junit.jupiter.api.Test;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * API tests for ArchivesApi
 */
@Disabled
public class ArchivesApiTest {

    private final ArchivesApi api = new ArchivesApi();

    /**
     * Create an Archive Bucket
     *
     * Create an Archive Bucket
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void addArchiveBucketTest() throws ApiException {
        AddArchiveBucketRequest addArchiveBucketRequest = null;
        AddArchiveBucket200Response response = api.addArchiveBucket(addArchiveBucketRequest);
        // TODO: test validations
    }

    /**
     * Upload Archive File
     *
     * Upload a file to the specified archive bucket and file path.  This will overwrite the file if it already exists. 
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void addArchiveFileTest() throws ApiException {
        String bucket = null;
        String filepath = null;
        String filename = null;
        File _file = null;
        AddArchiveFile200Response response = api.addArchiveFile(bucket, filepath, filename, _file);
        // TODO: test validations
    }

    /**
     * Create an Archive File Link
     *
     * This returns a secret token that can be used to download the file via a public url, without any other authentication or authorization. File links can be set to expire after a certain amount of time.
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void addArchiveFileLinkTest() throws ApiException {
        Long id = null;
        Long expireSeconds = null;
        AddArchiveFileLink200Response response = api.addArchiveFileLink(id, expireSeconds);
        // TODO: test validations
    }

    /**
     * Delete an Archive Bucket
     *
     * Will delete an archive bucket from the system and make it no longer usable.
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void deleteArchiveBucketTest() throws ApiException {
        Long id = null;
        Model200Success response = api.deleteArchiveBucket(id);
        // TODO: test validations
    }

    /**
     * Delete Archive File
     *
     * Permanently delete a file or directory.  Deleting a directory will also delete all the files under it. 
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void deleteArchiveFileTest() throws ApiException {
        Long id = null;
        Model200Success response = api.deleteArchiveFile(id);
        // TODO: test validations
    }

    /**
     * Delete an Archive File Link
     *
     * This will delete the link from the system, so it can no longer be used.
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void deleteArchiveFileLinkTest() throws ApiException {
        Long id = null;
        Long linkId = null;
        Model200Success response = api.deleteArchiveFileLink(id, linkId);
        // TODO: test validations
    }

    /**
     * Get a Specific Archive Bucket
     *
     * This endpoint retrieves a specific archive bucket.
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void getArchiveBucketTest() throws ApiException {
        Long id = null;
        GetArchiveBucket200Response response = api.getArchiveBucket(id);
        // TODO: test validations
    }

    /**
     * Download an Archive File
     *
     * Download the file as an authorized user with access to the bucket.  Downloading a directory will return a .zip file containing all files under it. 
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void getArchiveFileTest() throws ApiException {
        String bucket = null;
        String filepath = null;
        api.getArchiveFile(bucket, filepath);
        // TODO: test validations
    }

    /**
     * Get Archive File Details
     *
     * Get details about a specific archive file.
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void getArchiveFileDetailTest() throws ApiException {
        Long id = null;
        GetArchiveFileDetail200Response response = api.getArchiveFileDetail(id);
        // TODO: test validations
    }

    /**
     * Get Archive File Links
     *
     * This endpoint retrieves the links that have been created for the specified file.
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void getArchiveFileLinksTest() throws ApiException {
        Long id = null;
        GetArchiveFileLinks200Response response = api.getArchiveFileLinks(id);
        // TODO: test validations
    }

    /**
     * Get All Archive Buckets
     *
     * This endpoint retrieves all archive buckets associated with the account.
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void listArchiveBucketsTest() throws ApiException {
        String name = null;
        String phrase = null;
        ListArchiveBuckets200Response response = api.listArchiveBuckets(name, phrase);
        // TODO: test validations
    }

    /**
     * Get All Archive Files
     *
     * This endpoint retrieves all files in an archive bucket under the specified &#x60;filePath&#x60;.
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void listArchiveFilesTest() throws ApiException {
        String bucket = null;
        String filepath = null;
        String name = null;
        String phrase = null;
        Boolean fullTree = null;
        ListArchiveFiles200Response response = api.listArchiveFiles(bucket, filepath, name, phrase, fullTree);
        // TODO: test validations
    }

    /**
     * Update an Archive Bucket
     *
     * Update an Archive Bucket
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void updateArchiveBucketTest() throws ApiException {
        Long id = null;
        UpdateArchiveBucketRequest updateArchiveBucketRequest = null;
        AddArchiveBucket200Response response = api.updateArchiveBucket(id, updateArchiveBucketRequest);
        // TODO: test validations
    }

}