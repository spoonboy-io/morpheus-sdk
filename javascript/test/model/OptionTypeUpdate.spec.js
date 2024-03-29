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
    instance = new MorpheusApi.OptionTypeUpdate();
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

  describe('OptionTypeUpdate', function() {
    it('should create an instance of OptionTypeUpdate', function() {
      // uncomment below and update the code to test OptionTypeUpdate
      //var instane = new MorpheusApi.OptionTypeUpdate();
      //expect(instance).to.be.a(MorpheusApi.OptionTypeUpdate);
    });

    it('should have the property name (base name: "name")', function() {
      // uncomment below and update the code to test the property name
      //var instance = new MorpheusApi.OptionTypeUpdate();
      //expect(instance).to.be();
    });

    it('should have the property description (base name: "description")', function() {
      // uncomment below and update the code to test the property description
      //var instance = new MorpheusApi.OptionTypeUpdate();
      //expect(instance).to.be();
    });

    it('should have the property labels (base name: "labels")', function() {
      // uncomment below and update the code to test the property labels
      //var instance = new MorpheusApi.OptionTypeUpdate();
      //expect(instance).to.be();
    });

    it('should have the property fieldName (base name: "fieldName")', function() {
      // uncomment below and update the code to test the property fieldName
      //var instance = new MorpheusApi.OptionTypeUpdate();
      //expect(instance).to.be();
    });

    it('should have the property type (base name: "type")', function() {
      // uncomment below and update the code to test the property type
      //var instance = new MorpheusApi.OptionTypeUpdate();
      //expect(instance).to.be();
    });

    it('should have the property fieldLabel (base name: "fieldLabel")', function() {
      // uncomment below and update the code to test the property fieldLabel
      //var instance = new MorpheusApi.OptionTypeUpdate();
      //expect(instance).to.be();
    });

    it('should have the property placeholder (base name: "placeholder")', function() {
      // uncomment below and update the code to test the property placeholder
      //var instance = new MorpheusApi.OptionTypeUpdate();
      //expect(instance).to.be();
    });

    it('should have the property verifyPattern (base name: "verifyPattern")', function() {
      // uncomment below and update the code to test the property verifyPattern
      //var instance = new MorpheusApi.OptionTypeUpdate();
      //expect(instance).to.be();
    });

    it('should have the property defaultValue (base name: "defaultValue")', function() {
      // uncomment below and update the code to test the property defaultValue
      //var instance = new MorpheusApi.OptionTypeUpdate();
      //expect(instance).to.be();
    });

    it('should have the property required (base name: "required")', function() {
      // uncomment below and update the code to test the property required
      //var instance = new MorpheusApi.OptionTypeUpdate();
      //expect(instance).to.be();
    });

    it('should have the property exportMeta (base name: "exportMeta")', function() {
      // uncomment below and update the code to test the property exportMeta
      //var instance = new MorpheusApi.OptionTypeUpdate();
      //expect(instance).to.be();
    });

    it('should have the property editable (base name: "editable")', function() {
      // uncomment below and update the code to test the property editable
      //var instance = new MorpheusApi.OptionTypeUpdate();
      //expect(instance).to.be();
    });

    it('should have the property optionList (base name: "optionList")', function() {
      // uncomment below and update the code to test the property optionList
      //var instance = new MorpheusApi.OptionTypeUpdate();
      //expect(instance).to.be();
    });

  });

}));
