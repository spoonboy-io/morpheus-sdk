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
import java.io.File;
import org.openapitools.client.model.InlineObject107;
import org.openapitools.client.model.InlineObject108;
import org.openapitools.client.model.InlineObject109;
import org.openapitools.client.model.InlineObject110;
import org.openapitools.client.model.InlineObject111;
import org.openapitools.client.model.InlineObject112;
import org.openapitools.client.model.InlineObject113;
import org.openapitools.client.model.InlineObject114;
import org.openapitools.client.model.InlineObject115;
import org.openapitools.client.model.InlineObject117;
import org.openapitools.client.model.InlineObject118;
import org.openapitools.client.model.InlineObject119;
import org.openapitools.client.model.InlineObject120;
import org.openapitools.client.model.InlineObject121;
import org.openapitools.client.model.InlineObject122;
import org.openapitools.client.model.InlineObject123;
import org.openapitools.client.model.InlineObject124;
import org.openapitools.client.model.InlineObject263;
import org.openapitools.client.model.InlineObject264;
import org.openapitools.client.model.InlineResponse200136;
import org.openapitools.client.model.InlineResponse200165;
import org.openapitools.client.model.InlineResponse20068;
import org.openapitools.client.model.InlineResponse20069;
import org.openapitools.client.model.InlineResponse20070;
import org.openapitools.client.model.InlineResponse20071;
import org.openapitools.client.model.InlineResponse20072;
import org.openapitools.client.model.InlineResponse20073;
import org.openapitools.client.model.InlineResponse20074;
import org.openapitools.client.model.InlineResponse20075;
import org.openapitools.client.model.InlineResponse20076;
import org.openapitools.client.model.Model200Success;
import org.threeten.bp.OffsetDateTime;
import org.openapitools.client.model.ScriptSuccessId;
import org.openapitools.client.model.SuccessId;
import org.junit.Test;
import org.junit.Ignore;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * API tests for LibraryApi
 */
@Ignore
public class LibraryApiTest {

    private final LibraryApi api = new LibraryApi();

    
    /**
     * Create a File Template
     *
     * Use this command to create a file template.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void addFileTemplateTest() throws ApiException {
        InlineObject111 inlineObject111 = null;
        SuccessId response = api.addFileTemplate(inlineObject111);

        // TODO: test validations
    }
    
    /**
     * Create an Instance Type
     *
     * Use this command to create an instance type.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void addInstanceTypeTest() throws ApiException {
        InlineObject113 inlineObject113 = null;
        Model200Success response = api.addInstanceType(inlineObject113);

        // TODO: test validations
    }
    
    /**
     * Create a Layout
     *
     * Use this command to create a layout.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void addLayoutTest() throws ApiException {
        BigDecimal instanceTypeId = null;
        InlineObject115 inlineObject115 = null;
        SuccessId response = api.addLayout(instanceTypeId, inlineObject115);

        // TODO: test validations
    }
    
    /**
     * Create a Node Type
     *
     * Use this command to create a node type.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void addNodeTypeTest() throws ApiException {
        InlineObject109 inlineObject109 = null;
        Object response = api.addNodeType(inlineObject109);

        // TODO: test validations
    }
    
    /**
     * Create an Option List
     *
     * Use this command to create an option list.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void addOptionListTest() throws ApiException {
        InlineObject119 inlineObject119 = null;
        Model200Success response = api.addOptionList(inlineObject119);

        // TODO: test validations
    }
    
    /**
     * Create an Input
     *
     * Use this command to create an option type.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void addOptionTypeTest() throws ApiException {
        InlineObject121 inlineObject121 = null;
        SuccessId response = api.addOptionType(inlineObject121);

        // TODO: test validations
    }
    
    /**
     * Create a Script
     *
     * Use this command to create a script.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void addScriptTest() throws ApiException {
        InlineObject107 inlineObject107 = null;
        ScriptSuccessId response = api.addScript(inlineObject107);

        // TODO: test validations
    }
    
    /**
     * Create a Spec Template
     *
     * Use this command to create a spec template.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void addSpecTemplateTest() throws ApiException {
        InlineObject123 inlineObject123 = null;
        SuccessId response = api.addSpecTemplate(inlineObject123);

        // TODO: test validations
    }
    
    /**
     * Create a Virtual Image
     *
     * This endpoint creates a new virtual image, without any files yet.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void addVirtualImageTest() throws ApiException {
        InlineObject263 inlineObject263 = null;
        Object response = api.addVirtualImage(inlineObject263);

        // TODO: test validations
    }
    
    /**
     * Upload Virtual Image File
     *
     * This will upload the file and associate it to the Virtual Image.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void addVirtualImageFileTest() throws ApiException {
        BigDecimal virtualImageId = null;
        String filename = null;
        String url = null;
        File body = null;
        Model200Success response = api.addVirtualImageFile(virtualImageId, filename, url, body);

        // TODO: test validations
    }
    
    /**
     * Delete a File Template
     *
     * Will delete a file template
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void deleteFileTemplateTest() throws ApiException {
        Long id = null;
        Model200Success response = api.deleteFileTemplate(id);

        // TODO: test validations
    }
    
    /**
     * Delete an Instance Type
     *
     * Will delete an instance type
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void deleteInstanceTypeTest() throws ApiException {
        BigDecimal instanceTypeId = null;
        Model200Success response = api.deleteInstanceType(instanceTypeId);

        // TODO: test validations
    }
    
    /**
     * Delete a Layout
     *
     * Will delete a layout
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void deleteLayoutTest() throws ApiException {
        Long id = null;
        Model200Success response = api.deleteLayout(id);

        // TODO: test validations
    }
    
    /**
     * Delete a Node Type
     *
     * Will delete a node type
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void deleteNodeTypeTest() throws ApiException {
        Long id = null;
        Model200Success response = api.deleteNodeType(id);

        // TODO: test validations
    }
    
    /**
     * Delete an Option List
     *
     * Will delete an option list.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void deleteOptionListTest() throws ApiException {
        Long id = null;
        Model200Success response = api.deleteOptionList(id);

        // TODO: test validations
    }
    
    /**
     * Delete an Input
     *
     * Will delete an option type.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void deleteOptionTypeTest() throws ApiException {
        Long id = null;
        Model200Success response = api.deleteOptionType(id);

        // TODO: test validations
    }
    
    /**
     * Delete a Script
     *
     * Will delete a script
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void deleteScriptTest() throws ApiException {
        Long id = null;
        Model200Success response = api.deleteScript(id);

        // TODO: test validations
    }
    
    /**
     * Delete a Spec Template
     *
     * Will delete a spec template
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void deleteSpecTemplateTest() throws ApiException {
        Long id = null;
        Model200Success response = api.deleteSpecTemplate(id);

        // TODO: test validations
    }
    
    /**
     * Get a Specific File Template
     *
     * This endpoint retrieves a specific file template.  The value of template will be masked as ************ for system owned file templates. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getFileTemplateTest() throws ApiException {
        Long id = null;
        InlineResponse20070 response = api.getFileTemplate(id);

        // TODO: test validations
    }
    
    /**
     * Get A Specific Input
     *
     * This endpoint retrieves a specific option type. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getInputTest() throws ApiException {
        Long id = null;
        InlineResponse20075 response = api.getInput(id);

        // TODO: test validations
    }
    
    /**
     * Get a Specific Instance Type
     *
     * This endpoint retrieves a specific instance type.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getInstanceTypeTest() throws ApiException {
        BigDecimal instanceTypeId = null;
        InlineResponse20071 response = api.getInstanceType(instanceTypeId);

        // TODO: test validations
    }
    
    /**
     * Get a Specific Layout
     *
     * This endpoint retrieves a specific layout.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getLayoutTest() throws ApiException {
        Long id = null;
        InlineResponse20072 response = api.getLayout(id);

        // TODO: test validations
    }
    
    /**
     * Get a Specific Node Type
     *
     * This endpoint retrieves a specific node type.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getNodeTypeTest() throws ApiException {
        Long id = null;
        InlineResponse20069 response = api.getNodeType(id);

        // TODO: test validations
    }
    
    /**
     * Get a Specific Option List
     *
     * This endpoint retrieves a specific option list.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getOptionListTest() throws ApiException {
        Long id = null;
        InlineResponse20073 response = api.getOptionList(id);

        // TODO: test validations
    }
    
    /**
     * List Items for a Specific Option List
     *
     * This endpoint retrieves the items for a specific option list.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getOptionListItemsTest() throws ApiException {
        Long id = null;
        InlineResponse20074 response = api.getOptionListItems(id);

        // TODO: test validations
    }
    
    /**
     * Get a Specific Script
     *
     * This endpoint retrieves a specific script.  The value of script will be masked as ************ for system owned scripts. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getScriptTest() throws ApiException {
        Long id = null;
        InlineResponse20068 response = api.getScript(id);

        // TODO: test validations
    }
    
    /**
     * Retrieves a Specific Security Package Type
     *
     * Retrieves a specific security package type. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getSecurityPackageTypeTest() throws ApiException {
        Long id = null;
        InlineResponse200136 response = api.getSecurityPackageType(id);

        // TODO: test validations
    }
    
    /**
     * Get a Specific Spec Template
     *
     * This endpoint retrieves a specific spec template.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getSpecTemplateTest() throws ApiException {
        Long id = null;
        InlineResponse20076 response = api.getSpecTemplate(id);

        // TODO: test validations
    }
    
    /**
     * Get a Specific Virtual Image
     *
     * This endpoint retrieves a specific virtual image and its files.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getVirtualImageTest() throws ApiException {
        BigDecimal virtualImageId = null;
        InlineResponse200165 response = api.getVirtualImage(virtualImageId);

        // TODO: test validations
    }
    
    /**
     * Get All File Templates
     *
     * This endpoint retrieves all file templates.  The value of template will be masked as ************ for system owned file templates. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void listFileTemplatesTest() throws ApiException {
        Long max = null;
        Long offset = null;
        String sort = null;
        String direction = null;
        String phrase = null;
        String name = null;
        String labels = null;
        String allLabels = null;
        String fileName = null;
        Object response = api.listFileTemplates(max, offset, sort, direction, phrase, name, labels, allLabels, fileName);

        // TODO: test validations
    }
    
    /**
     * Get All Inputs
     *
     * This endpoint retrieves all option types. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void listInputsTest() throws ApiException {
        Long max = null;
        Long offset = null;
        String sort = null;
        String direction = null;
        String phrase = null;
        String name = null;
        String code = null;
        String labels = null;
        String allLabels = null;
        String fieldName = null;
        String fieldContext = null;
        String fieldLabel = null;
        Object response = api.listInputs(max, offset, sort, direction, phrase, name, code, labels, allLabels, fieldName, fieldContext, fieldLabel);

        // TODO: test validations
    }
    
    /**
     * Get All Instance Types
     *
     * This endpoint retrieves all instance types. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void listInstanceTypesTest() throws ApiException {
        Long max = null;
        Long offset = null;
        String sort = null;
        String direction = null;
        String phrase = null;
        String name = null;
        String code = null;
        Boolean featured = null;
        String labels = null;
        String allLabels = null;
        Boolean details = null;
        Object response = api.listInstanceTypes(max, offset, sort, direction, phrase, name, code, featured, labels, allLabels, details);

        // TODO: test validations
    }
    
    /**
     * Get All Layouts
     *
     * This endpoint retrieves all layouts. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void listLayoutsTest() throws ApiException {
        Long max = null;
        Long offset = null;
        String sort = null;
        String direction = null;
        String phrase = null;
        String name = null;
        String code = null;
        String provisionType = null;
        String labels = null;
        String allLabels = null;
        Object response = api.listLayouts(max, offset, sort, direction, phrase, name, code, provisionType, labels, allLabels);

        // TODO: test validations
    }
    
    /**
     * Get All Layouts For an Instance Type
     *
     * This endpoint retrieves all layouts for a specific instance type.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void listLayoutsForInstanceTypeTest() throws ApiException {
        BigDecimal instanceTypeId = null;
        Long max = null;
        Long offset = null;
        String sort = null;
        String direction = null;
        String phrase = null;
        String name = null;
        String code = null;
        String provisionType = null;
        String labels = null;
        String allLabels = null;
        Object response = api.listLayoutsForInstanceType(instanceTypeId, max, offset, sort, direction, phrase, name, code, provisionType, labels, allLabels);

        // TODO: test validations
    }
    
    /**
     * Get All Node Types
     *
     * This endpoint retrieves all node types.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void listNodeTypesTest() throws ApiException {
        Long max = null;
        Long offset = null;
        String sort = null;
        String direction = null;
        String phrase = null;
        String name = null;
        String code = null;
        String provisionType = null;
        String labels = null;
        String allLabels = null;
        Object response = api.listNodeTypes(max, offset, sort, direction, phrase, name, code, provisionType, labels, allLabels);

        // TODO: test validations
    }
    
    /**
     * Get All Option Lists
     *
     * This endpoint retrieves all option lists.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void listOptionListsTest() throws ApiException {
        Long max = null;
        Long offset = null;
        String sort = null;
        String direction = null;
        String phrase = null;
        String name = null;
        String labels = null;
        String allLabels = null;
        Object response = api.listOptionLists(max, offset, sort, direction, phrase, name, labels, allLabels);

        // TODO: test validations
    }
    
    /**
     * Get All Scripts
     *
     * This endpoint retrieves all scripts.  The value of script will be masked as ************ for system owned scripts. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void listScriptsTest() throws ApiException {
        Long max = null;
        Long offset = null;
        String sort = null;
        String direction = null;
        String phrase = null;
        String labels = null;
        String allLabels = null;
        String scriptType = null;
        String scriptPhase = null;
        Object response = api.listScripts(max, offset, sort, direction, phrase, labels, allLabels, scriptType, scriptPhase);

        // TODO: test validations
    }
    
    /**
     * Retrieves all Security Package Types
     *
     * Retrieves all Security Package Types 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void listSecurityPackageTypesTest() throws ApiException {
        Object response = api.listSecurityPackageTypes();

        // TODO: test validations
    }
    
    /**
     * Get All Spec Templates
     *
     * This endpoint retrieves all spec templates.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void listSpecTemplatesTest() throws ApiException {
        Long max = null;
        Long offset = null;
        String sort = null;
        String direction = null;
        String phrase = null;
        String name = null;
        String labels = null;
        String allLabels = null;
        Object response = api.listSpecTemplates(max, offset, sort, direction, phrase, name, labels, allLabels);

        // TODO: test validations
    }
    
    /**
     * Get a List of Virtual Image Locations
     *
     * This endpoint retrieves a specific virtual image and its files.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void listVirtualImageLocationsTest() throws ApiException {
        BigDecimal virtualImageId = null;
        Object response = api.listVirtualImageLocations(virtualImageId);

        // TODO: test validations
    }
    
    /**
     * Get List of Virtual Images
     *
     * This endpoint retrieves a list of virtual images for the specified filter.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void listVirtualImagesTest() throws ApiException {
        Long max = null;
        Long offset = null;
        String name = null;
        String phrase = null;
        OffsetDateTime lastUpdated = null;
        String filterType = null;
        String imageType = null;
        String tags = null;
        String labels = null;
        String allLabels = null;
        Object response = api.listVirtualImages(max, offset, name, phrase, lastUpdated, filterType, imageType, tags, labels, allLabels);

        // TODO: test validations
    }
    
    /**
     * Deletes a Security Scan
     *
     * Deletes a specified security scan. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void removeSecurityScansTest() throws ApiException {
        Long id = null;
        Model200Success response = api.removeSecurityScans(id);

        // TODO: test validations
    }
    
    /**
     * Delete a Virtual Image
     *
     * Will delete a virtual image and any associated files, use removeFromCloud&#x3D;true to also delete image locations from all clouds.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void removeVirtualImageTest() throws ApiException {
        BigDecimal virtualImageId = null;
        Model200Success response = api.removeVirtualImage(virtualImageId);

        // TODO: test validations
    }
    
    /**
     * Remove Virtual Image File
     *
     * Remove Virtual Image File
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void removeVirtualImageFileTest() throws ApiException {
        BigDecimal virtualImageId = null;
        String filename = null;
        Model200Success response = api.removeVirtualImageFile(virtualImageId, filename);

        // TODO: test validations
    }
    
    /**
     * Delete a Virtual Image Location
     *
     * Will delete a virtual image location, use removeFromCloud&#x3D;true to all also delete image locations from all clouds as well.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void removeVirtualImageLocationTest() throws ApiException {
        BigDecimal virtualImageId = null;
        Long id = null;
        Model200Success response = api.removeVirtualImageLocation(virtualImageId, id);

        // TODO: test validations
    }
    
    /**
     * Toggle Featured For Instance Type
     *
     * Use this command to toggle the featured flag for an existing instance type. This will change the value from false to true, or from true to false. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void setInstanceTypeFeaturedTest() throws ApiException {
        BigDecimal instanceTypeId = null;
        Model200Success response = api.setInstanceTypeFeatured(instanceTypeId);

        // TODO: test validations
    }
    
    /**
     * Update a File Template
     *
     * Use this command to update an existing file template.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void updateFileTemplateTest() throws ApiException {
        Long id = null;
        InlineObject112 inlineObject112 = null;
        Model200Success response = api.updateFileTemplate(id, inlineObject112);

        // TODO: test validations
    }
    
    /**
     * Update an Instance Type
     *
     * Use this command to update an existing instance type.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void updateInstanceTypeTest() throws ApiException {
        BigDecimal instanceTypeId = null;
        InlineObject114 inlineObject114 = null;
        Model200Success response = api.updateInstanceType(instanceTypeId, inlineObject114);

        // TODO: test validations
    }
    
    /**
     * Update Logo For Instance Type
     *
     * Use this command to update the logo and dark logo images for an existing instance type. This endpoint expects multipart form data as the request format, not JSON. 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void updateInstanceTypeLogoTest() throws ApiException {
        BigDecimal instanceTypeId = null;
        File logo = null;
        File darkLogo = null;
        Model200Success response = api.updateInstanceTypeLogo(instanceTypeId, logo, darkLogo);

        // TODO: test validations
    }
    
    /**
     * Update a Layout
     *
     * Use this command to update an existing layout.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void updateLayoutTest() throws ApiException {
        Long id = null;
        InlineObject117 inlineObject117 = null;
        Model200Success response = api.updateLayout(id, inlineObject117);

        // TODO: test validations
    }
    
    /**
     * Update Layout Permissions
     *
     * Use this command to update permissions for an existing layout.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void updateLayoutPermissionsTest() throws ApiException {
        Long id = null;
        InlineObject118 inlineObject118 = null;
        Model200Success response = api.updateLayoutPermissions(id, inlineObject118);

        // TODO: test validations
    }
    
    /**
     * Update a Node Type
     *
     * Use this command to update an existing node type.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void updateNodeTypeTest() throws ApiException {
        Long id = null;
        InlineObject110 inlineObject110 = null;
        Model200Success response = api.updateNodeType(id, inlineObject110);

        // TODO: test validations
    }
    
    /**
     * Update an Option List
     *
     * Use this command to update an existing option list.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void updateOptionListTest() throws ApiException {
        Long id = null;
        InlineObject120 inlineObject120 = null;
        Model200Success response = api.updateOptionList(id, inlineObject120);

        // TODO: test validations
    }
    
    /**
     * Update an Input
     *
     * Use this command to update an existing option type.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void updateOptionTypeTest() throws ApiException {
        Long id = null;
        InlineObject122 inlineObject122 = null;
        Model200Success response = api.updateOptionType(id, inlineObject122);

        // TODO: test validations
    }
    
    /**
     * Update a Script
     *
     * Use this command to update an existing script.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void updateScriptTest() throws ApiException {
        Long id = null;
        InlineObject108 inlineObject108 = null;
        ScriptSuccessId response = api.updateScript(id, inlineObject108);

        // TODO: test validations
    }
    
    /**
     * Update a Spec Template
     *
     * Use this command to update an existing spec template.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void updateSpecTemplateTest() throws ApiException {
        Long id = null;
        InlineObject124 inlineObject124 = null;
        Model200Success response = api.updateSpecTemplate(id, inlineObject124);

        // TODO: test validations
    }
    
    /**
     * Update a Virtual Image
     *
     * This endpoint updates an existing virtual image.
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void updateVirtualImageTest() throws ApiException {
        BigDecimal virtualImageId = null;
        InlineObject264 inlineObject264 = null;
        Object response = api.updateVirtualImage(virtualImageId, inlineObject264);

        // TODO: test validations
    }
    
}