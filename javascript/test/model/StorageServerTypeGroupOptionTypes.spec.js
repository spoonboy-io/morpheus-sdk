/**
 * Morpheus API
 * Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.
 *
 * The version of the OpenAPI document: 6.2.1
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
    instance = new MorpheusApi.StorageServerTypeGroupOptionTypes();
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

  describe('StorageServerTypeGroupOptionTypes', function() {
    it('should create an instance of StorageServerTypeGroupOptionTypes', function() {
      // uncomment below and update the code to test StorageServerTypeGroupOptionTypes
      //var instane = new MorpheusApi.StorageServerTypeGroupOptionTypes();
      //expect(instance).to.be.a(MorpheusApi.StorageServerTypeGroupOptionTypes);
    });

    it('should have the property id (base name: "id")', function() {
      // uncomment below and update the code to test the property id
      //var instance = new MorpheusApi.StorageServerTypeGroupOptionTypes();
      //expect(instance).to.be();
    });

    it('should have the property name (base name: "name")', function() {
      // uncomment below and update the code to test the property name
      //var instance = new MorpheusApi.StorageServerTypeGroupOptionTypes();
      //expect(instance).to.be();
    });

    it('should have the property description (base name: "description")', function() {
      // uncomment below and update the code to test the property description
      //var instance = new MorpheusApi.StorageServerTypeGroupOptionTypes();
      //expect(instance).to.be();
    });

    it('should have the property code (base name: "code")', function() {
      // uncomment below and update the code to test the property code
      //var instance = new MorpheusApi.StorageServerTypeGroupOptionTypes();
      //expect(instance).to.be();
    });

    it('should have the property fieldName (base name: "fieldName")', function() {
      // uncomment below and update the code to test the property fieldName
      //var instance = new MorpheusApi.StorageServerTypeGroupOptionTypes();
      //expect(instance).to.be();
    });

    it('should have the property fieldLabel (base name: "fieldLabel")', function() {
      // uncomment below and update the code to test the property fieldLabel
      //var instance = new MorpheusApi.StorageServerTypeGroupOptionTypes();
      //expect(instance).to.be();
    });

    it('should have the property fieldCode (base name: "fieldCode")', function() {
      // uncomment below and update the code to test the property fieldCode
      //var instance = new MorpheusApi.StorageServerTypeGroupOptionTypes();
      //expect(instance).to.be();
    });

    it('should have the property fieldContext (base name: "fieldContext")', function() {
      // uncomment below and update the code to test the property fieldContext
      //var instance = new MorpheusApi.StorageServerTypeGroupOptionTypes();
      //expect(instance).to.be();
    });

    it('should have the property fieldGroup (base name: "fieldGroup")', function() {
      // uncomment below and update the code to test the property fieldGroup
      //var instance = new MorpheusApi.StorageServerTypeGroupOptionTypes();
      //expect(instance).to.be();
    });

    it('should have the property fieldClass (base name: "fieldClass")', function() {
      // uncomment below and update the code to test the property fieldClass
      //var instance = new MorpheusApi.StorageServerTypeGroupOptionTypes();
      //expect(instance).to.be();
    });

    it('should have the property fieldAddOn (base name: "fieldAddOn")', function() {
      // uncomment below and update the code to test the property fieldAddOn
      //var instance = new MorpheusApi.StorageServerTypeGroupOptionTypes();
      //expect(instance).to.be();
    });

    it('should have the property fieldComponent (base name: "fieldComponent")', function() {
      // uncomment below and update the code to test the property fieldComponent
      //var instance = new MorpheusApi.StorageServerTypeGroupOptionTypes();
      //expect(instance).to.be();
    });

    it('should have the property fieldInput (base name: "fieldInput")', function() {
      // uncomment below and update the code to test the property fieldInput
      //var instance = new MorpheusApi.StorageServerTypeGroupOptionTypes();
      //expect(instance).to.be();
    });

    it('should have the property placeHolder (base name: "placeHolder")', function() {
      // uncomment below and update the code to test the property placeHolder
      //var instance = new MorpheusApi.StorageServerTypeGroupOptionTypes();
      //expect(instance).to.be();
    });

    it('should have the property verifyPattern (base name: "verifyPattern")', function() {
      // uncomment below and update the code to test the property verifyPattern
      //var instance = new MorpheusApi.StorageServerTypeGroupOptionTypes();
      //expect(instance).to.be();
    });

    it('should have the property helpBlock (base name: "helpBlock")', function() {
      // uncomment below and update the code to test the property helpBlock
      //var instance = new MorpheusApi.StorageServerTypeGroupOptionTypes();
      //expect(instance).to.be();
    });

    it('should have the property helpBlockFieldCode (base name: "helpBlockFieldCode")', function() {
      // uncomment below and update the code to test the property helpBlockFieldCode
      //var instance = new MorpheusApi.StorageServerTypeGroupOptionTypes();
      //expect(instance).to.be();
    });

    it('should have the property defaultValue (base name: "defaultValue")', function() {
      // uncomment below and update the code to test the property defaultValue
      //var instance = new MorpheusApi.StorageServerTypeGroupOptionTypes();
      //expect(instance).to.be();
    });

    it('should have the property optionSource (base name: "optionSource")', function() {
      // uncomment below and update the code to test the property optionSource
      //var instance = new MorpheusApi.StorageServerTypeGroupOptionTypes();
      //expect(instance).to.be();
    });

    it('should have the property optionSourceType (base name: "optionSourceType")', function() {
      // uncomment below and update the code to test the property optionSourceType
      //var instance = new MorpheusApi.StorageServerTypeGroupOptionTypes();
      //expect(instance).to.be();
    });

    it('should have the property optionList (base name: "optionList")', function() {
      // uncomment below and update the code to test the property optionList
      //var instance = new MorpheusApi.StorageServerTypeGroupOptionTypes();
      //expect(instance).to.be();
    });

    it('should have the property type (base name: "type")', function() {
      // uncomment below and update the code to test the property type
      //var instance = new MorpheusApi.StorageServerTypeGroupOptionTypes();
      //expect(instance).to.be();
    });

    it('should have the property advanced (base name: "advanced")', function() {
      // uncomment below and update the code to test the property advanced
      //var instance = new MorpheusApi.StorageServerTypeGroupOptionTypes();
      //expect(instance).to.be();
    });

    it('should have the property required (base name: "required")', function() {
      // uncomment below and update the code to test the property required
      //var instance = new MorpheusApi.StorageServerTypeGroupOptionTypes();
      //expect(instance).to.be();
    });

    it('should have the property exportMeta (base name: "exportMeta")', function() {
      // uncomment below and update the code to test the property exportMeta
      //var instance = new MorpheusApi.StorageServerTypeGroupOptionTypes();
      //expect(instance).to.be();
    });

    it('should have the property editable (base name: "editable")', function() {
      // uncomment below and update the code to test the property editable
      //var instance = new MorpheusApi.StorageServerTypeGroupOptionTypes();
      //expect(instance).to.be();
    });

    it('should have the property creatable (base name: "creatable")', function() {
      // uncomment below and update the code to test the property creatable
      //var instance = new MorpheusApi.StorageServerTypeGroupOptionTypes();
      //expect(instance).to.be();
    });

    it('should have the property config (base name: "config")', function() {
      // uncomment below and update the code to test the property config
      //var instance = new MorpheusApi.StorageServerTypeGroupOptionTypes();
      //expect(instance).to.be();
    });

    it('should have the property displayOrder (base name: "displayOrder")', function() {
      // uncomment below and update the code to test the property displayOrder
      //var instance = new MorpheusApi.StorageServerTypeGroupOptionTypes();
      //expect(instance).to.be();
    });

    it('should have the property wrapperClass (base name: "wrapperClass")', function() {
      // uncomment below and update the code to test the property wrapperClass
      //var instance = new MorpheusApi.StorageServerTypeGroupOptionTypes();
      //expect(instance).to.be();
    });

    it('should have the property enabled (base name: "enabled")', function() {
      // uncomment below and update the code to test the property enabled
      //var instance = new MorpheusApi.StorageServerTypeGroupOptionTypes();
      //expect(instance).to.be();
    });

    it('should have the property noBlank (base name: "noBlank")', function() {
      // uncomment below and update the code to test the property noBlank
      //var instance = new MorpheusApi.StorageServerTypeGroupOptionTypes();
      //expect(instance).to.be();
    });

    it('should have the property dependsOnCode (base name: "dependsOnCode")', function() {
      // uncomment below and update the code to test the property dependsOnCode
      //var instance = new MorpheusApi.StorageServerTypeGroupOptionTypes();
      //expect(instance).to.be();
    });

    it('should have the property visibleOnCode (base name: "visibleOnCode")', function() {
      // uncomment below and update the code to test the property visibleOnCode
      //var instance = new MorpheusApi.StorageServerTypeGroupOptionTypes();
      //expect(instance).to.be();
    });

    it('should have the property requireOnCode (base name: "requireOnCode")', function() {
      // uncomment below and update the code to test the property requireOnCode
      //var instance = new MorpheusApi.StorageServerTypeGroupOptionTypes();
      //expect(instance).to.be();
    });

    it('should have the property contextualDefault (base name: "contextualDefault")', function() {
      // uncomment below and update the code to test the property contextualDefault
      //var instance = new MorpheusApi.StorageServerTypeGroupOptionTypes();
      //expect(instance).to.be();
    });

    it('should have the property displayValueOnDetails (base name: "displayValueOnDetails")', function() {
      // uncomment below and update the code to test the property displayValueOnDetails
      //var instance = new MorpheusApi.StorageServerTypeGroupOptionTypes();
      //expect(instance).to.be();
    });

    it('should have the property showOnCreate (base name: "showOnCreate")', function() {
      // uncomment below and update the code to test the property showOnCreate
      //var instance = new MorpheusApi.StorageServerTypeGroupOptionTypes();
      //expect(instance).to.be();
    });

    it('should have the property showOnEdit (base name: "showOnEdit")', function() {
      // uncomment below and update the code to test the property showOnEdit
      //var instance = new MorpheusApi.StorageServerTypeGroupOptionTypes();
      //expect(instance).to.be();
    });

    it('should have the property localCredential (base name: "localCredential")', function() {
      // uncomment below and update the code to test the property localCredential
      //var instance = new MorpheusApi.StorageServerTypeGroupOptionTypes();
      //expect(instance).to.be();
    });

  });

}));
