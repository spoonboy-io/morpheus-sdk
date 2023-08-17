/**
 * Morpheus API
 * Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.
 *
 * The version of the OpenAPI document: 6.1.1
 * Contact: dev@morpheusdata.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD.
    define(['expect.js', process.cwd()+'/src/index'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    factory(require('expect.js'), require(process.cwd()+'/src/index'));
  } else {
    // Browser globals (root is window)
    factory(root.expect, root.MorpheusApi);
  }
}(this, function(expect, MorpheusApi) {
  'use strict';

  var instance;

  beforeEach(function() {
    instance = new MorpheusApi.CatalogItemTypeWorkflowCreate();
  });

  var getProperty = function(object, getter, property) {
    // Use getter method if present; otherwise, get the property directly.
    if (typeof object[getter] === 'function')
      return object[getter]();
    else
      return object[property];
  }

  var setProperty = function(object, setter, property, value) {
    // Use setter method if present; otherwise, set the property directly.
    if (typeof object[setter] === 'function')
      object[setter](value);
    else
      object[property] = value;
  }

  describe('CatalogItemTypeWorkflowCreate', function() {
    it('should create an instance of CatalogItemTypeWorkflowCreate', function() {
      // uncomment below and update the code to test CatalogItemTypeWorkflowCreate
      //var instance = new MorpheusApi.CatalogItemTypeWorkflowCreate();
      //expect(instance).to.be.a(MorpheusApi.CatalogItemTypeWorkflowCreate);
    });

    it('should have the property name (base name: "name")', function() {
      // uncomment below and update the code to test the property name
      //var instance = new MorpheusApi.CatalogItemTypeWorkflowCreate();
      //expect(instance).to.be();
    });

    it('should have the property code (base name: "code")', function() {
      // uncomment below and update the code to test the property code
      //var instance = new MorpheusApi.CatalogItemTypeWorkflowCreate();
      //expect(instance).to.be();
    });

    it('should have the property category (base name: "category")', function() {
      // uncomment below and update the code to test the property category
      //var instance = new MorpheusApi.CatalogItemTypeWorkflowCreate();
      //expect(instance).to.be();
    });

    it('should have the property description (base name: "description")', function() {
      // uncomment below and update the code to test the property description
      //var instance = new MorpheusApi.CatalogItemTypeWorkflowCreate();
      //expect(instance).to.be();
    });

    it('should have the property labels (base name: "labels")', function() {
      // uncomment below and update the code to test the property labels
      //var instance = new MorpheusApi.CatalogItemTypeWorkflowCreate();
      //expect(instance).to.be();
    });

    it('should have the property type (base name: "type")', function() {
      // uncomment below and update the code to test the property type
      //var instance = new MorpheusApi.CatalogItemTypeWorkflowCreate();
      //expect(instance).to.be();
    });

    it('should have the property visibility (base name: "visibility")', function() {
      // uncomment below and update the code to test the property visibility
      //var instance = new MorpheusApi.CatalogItemTypeWorkflowCreate();
      //expect(instance).to.be();
    });

    it('should have the property layoutCode (base name: "layoutCode")', function() {
      // uncomment below and update the code to test the property layoutCode
      //var instance = new MorpheusApi.CatalogItemTypeWorkflowCreate();
      //expect(instance).to.be();
    });

    it('should have the property iconPath (base name: "iconPath")', function() {
      // uncomment below and update the code to test the property iconPath
      //var instance = new MorpheusApi.CatalogItemTypeWorkflowCreate();
      //expect(instance).to.be();
    });

    it('should have the property enabled (base name: "enabled")', function() {
      // uncomment below and update the code to test the property enabled
      //var instance = new MorpheusApi.CatalogItemTypeWorkflowCreate();
      //expect(instance).to.be();
    });

    it('should have the property featured (base name: "featured")', function() {
      // uncomment below and update the code to test the property featured
      //var instance = new MorpheusApi.CatalogItemTypeWorkflowCreate();
      //expect(instance).to.be();
    });

    it('should have the property allowQuantity (base name: "allowQuantity")', function() {
      // uncomment below and update the code to test the property allowQuantity
      //var instance = new MorpheusApi.CatalogItemTypeWorkflowCreate();
      //expect(instance).to.be();
    });

    it('should have the property workflow (base name: "workflow")', function() {
      // uncomment below and update the code to test the property workflow
      //var instance = new MorpheusApi.CatalogItemTypeWorkflowCreate();
      //expect(instance).to.be();
    });

    it('should have the property context (base name: "context")', function() {
      // uncomment below and update the code to test the property context
      //var instance = new MorpheusApi.CatalogItemTypeWorkflowCreate();
      //expect(instance).to.be();
    });

    it('should have the property workflowConfig (base name: "workflowConfig")', function() {
      // uncomment below and update the code to test the property workflowConfig
      //var instance = new MorpheusApi.CatalogItemTypeWorkflowCreate();
      //expect(instance).to.be();
    });

  });

}));
