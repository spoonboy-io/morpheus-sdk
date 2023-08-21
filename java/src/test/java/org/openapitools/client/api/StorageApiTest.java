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
import org.openapitools.client.model.InlineObject232;
import org.openapitools.client.model.InlineObject233;
import org.openapitools.client.model.InlineObject234;
import org.openapitools.client.model.InlineObject235;
import org.openapitools.client.model.InlineObject236;
import org.openapitools.client.model.InlineObject237;
import org.openapitools.client.model.InlineResponse200145;
import org.openapitools.client.model.InlineResponse200146;
import org.openapitools.client.model.InlineResponse200147;
import org.openapitools.client.model.InlineResponse200148;
import org.openapitools.client.model.InlineResponse200149;
import org.openapitools.client.model.Model200Success;
import org.junit.Test;
import org.junit.Ignore;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * API tests for StorageApi
 */
@Ignore
public class StorageApiTest {

    private final StorageApi api = new StorageApi();

    
    /**
     * Creates a Storage Bucket
     *
     * Creates a storage bucket. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void addStorageBucketsTest() throws ApiException {
        InlineObject232 inlineObject232 = null;
        Object response = api.addStorageBuckets(inlineObject232);

        // TODO: test validations
    }
    
    /**
     * Creates a Storage Server
     *
     * Creates a storage server. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void addStorageServersTest() throws ApiException {
        InlineObject234 inlineObject234 = null;
        Object response = api.addStorageServers(inlineObject234);

        // TODO: test validations
    }
    
    /**
     * Creates a Storage Volume
     *
     * This endpoint allows creating a Storage Volume. Only certain types of storage servers support creating and deleting storage volumes, such as 3Par and Isilon. Configuration options vary by &#x60;Storage Volume Type&#x60;. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void addStorageVolumesTest() throws ApiException {
        InlineObject236 inlineObject236 = null;
        Object response = api.addStorageVolumes(inlineObject236);

        // TODO: test validations
    }
    
    /**
     * Retrieves a Specific Storage Bucket
     *
     * Retrieves a specific storage bucket. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getStorageBucketsTest() throws ApiException {
        Long id = null;
        InlineResponse200145 response = api.getStorageBuckets(id);

        // TODO: test validations
    }
    
    /**
     * Retrieves a Specific Storage Server Type
     *
     * Retrieves a specific storage server type. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getStorageServerTypesTest() throws ApiException {
        Long id = null;
        InlineResponse200146 response = api.getStorageServerTypes(id);

        // TODO: test validations
    }
    
    /**
     * Retrieves a Specific Storage Server
     *
     * Retrieves a specific storage server. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getStorageServersTest() throws ApiException {
        Long id = null;
        InlineResponse200147 response = api.getStorageServers(id);

        // TODO: test validations
    }
    
    /**
     * Retrieves a Specific Storage Volume Type
     *
     * Retrieves a specific storage volume type. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getStorageVolumeTypesTest() throws ApiException {
        Long id = null;
        InlineResponse200148 response = api.getStorageVolumeTypes(id);

        // TODO: test validations
    }
    
    /**
     * Retrieves a Specific Storage Volume
     *
     * Retrieves a specific storage volume. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getStorageVolumesTest() throws ApiException {
        Long id = null;
        InlineResponse200149 response = api.getStorageVolumes(id);

        // TODO: test validations
    }
    
    /**
     * Retrieves all Storage Buckets
     *
     * Retrieves all storage buckets. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void listStorageBucketsTest() throws ApiException {
        Long max = null;
        Long offset = null;
        String sort = null;
        String direction = null;
        String phrase = null;
        String name = null;
        Object response = api.listStorageBuckets(max, offset, sort, direction, phrase, name);

        // TODO: test validations
    }
    
    /**
     * Retrieves all Storage Server Types
     *
     * Fetch a paginated list of available storage server types. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void listStorageServerTypesTest() throws ApiException {
        Long max = null;
        Long offset = null;
        String sort = null;
        String direction = null;
        String name = null;
        String code = null;
        String phrase = null;
        Object response = api.listStorageServerTypes(max, offset, sort, direction, name, code, phrase);

        // TODO: test validations
    }
    
    /**
     * Retrieves all Storage Servers
     *
     * Retrieves all storage servers. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void listStorageServersTest() throws ApiException {
        Long max = null;
        Long offset = null;
        String sort = null;
        String direction = null;
        String phrase = null;
        String name = null;
        Object response = api.listStorageServers(max, offset, sort, direction, phrase, name);

        // TODO: test validations
    }
    
    /**
     * Retrieves all Storage Volume Types
     *
     * Fetch a paginated list of available storage volume types. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void listStorageVolumeTypesTest() throws ApiException {
        Long max = null;
        Long offset = null;
        String sort = null;
        String direction = null;
        String name = null;
        String code = null;
        String phrase = null;
        Object response = api.listStorageVolumeTypes(max, offset, sort, direction, name, code, phrase);

        // TODO: test validations
    }
    
    /**
     * Retrieves all Storage Volumes
     *
     * Retrieves all storage volumes. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void listStorageVolumesTest() throws ApiException {
        Long max = null;
        Long offset = null;
        String sort = null;
        String direction = null;
        String phrase = null;
        String name = null;
        Object response = api.listStorageVolumes(max, offset, sort, direction, phrase, name);

        // TODO: test validations
    }
    
    /**
     * Deletes a Storage Bucket
     *
     * Deletes a specified storage bucket. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void removeStorageBucketsTest() throws ApiException {
        Long id = null;
        Model200Success response = api.removeStorageBuckets(id);

        // TODO: test validations
    }
    
    /**
     * Deletes a Storage Server
     *
     * Deletes a specified storage server. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void removeStorageServersTest() throws ApiException {
        Long id = null;
        Model200Success response = api.removeStorageServers(id);

        // TODO: test validations
    }
    
    /**
     * Deletes a Storage Volume
     *
     * Deletes a specified storage volume. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void removeStorageVolumesTest() throws ApiException {
        Long id = null;
        Model200Success response = api.removeStorageVolumes(id);

        // TODO: test validations
    }
    
    /**
     * Updates a Storage Bucket
     *
     * Updates a storage bucket. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void updateStorageBucketsTest() throws ApiException {
        Long id = null;
        InlineObject233 inlineObject233 = null;
        Object response = api.updateStorageBuckets(id, inlineObject233);

        // TODO: test validations
    }
    
    /**
     * Updates a Storage Server
     *
     * Updates a storage server. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void updateStorageServersTest() throws ApiException {
        Long id = null;
        InlineObject235 inlineObject235 = null;
        Object response = api.updateStorageServers(id, inlineObject235);

        // TODO: test validations
    }
    
    /**
     * Updates a Storage Volume
     *
     * Updates a storage volume. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void updateStorageVolumesTest() throws ApiException {
        Long id = null;
        InlineObject237 inlineObject237 = null;
        Object response = api.updateStorageVolumes(id, inlineObject237);

        // TODO: test validations
    }
    
}