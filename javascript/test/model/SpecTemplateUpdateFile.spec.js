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
    instance = new MorpheusApi.SpecTemplateUpdateFile();
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

  describe('SpecTemplateUpdateFile', function() {
    it('should create an instance of SpecTemplateUpdateFile', function() {
      // uncomment below and update the code to test SpecTemplateUpdateFile
      //var instane = new MorpheusApi.SpecTemplateUpdateFile();
      //expect(instance).to.be.a(MorpheusApi.SpecTemplateUpdateFile);
    });

    it('should have the property sourceType (base name: "sourceType")', function() {
      // uncomment below and update the code to test the property sourceType
      //var instance = new MorpheusApi.SpecTemplateUpdateFile();
      //expect(instance).to.be();
    });

    it('should have the property content (base name: "content")', function() {
      // uncomment below and update the code to test the property content
      //var instance = new MorpheusApi.SpecTemplateUpdateFile();
      //expect(instance).to.be();
    });

    it('should have the property contentPath (base name: "contentPath")', function() {
      // uncomment below and update the code to test the property contentPath
      //var instance = new MorpheusApi.SpecTemplateUpdateFile();
      //expect(instance).to.be();
    });

    it('should have the property contentRef (base name: "contentRef")', function() {
      // uncomment below and update the code to test the property contentRef
      //var instance = new MorpheusApi.SpecTemplateUpdateFile();
      //expect(instance).to.be();
    });

    it('should have the property repository (base name: "repository")', function() {
      // uncomment below and update the code to test the property repository
      //var instance = new MorpheusApi.SpecTemplateUpdateFile();
      //expect(instance).to.be();
    });

  });

}));
