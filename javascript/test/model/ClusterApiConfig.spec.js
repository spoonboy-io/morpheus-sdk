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
    instance = new MorpheusApi.ClusterApiConfig();
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

  describe('ClusterApiConfig', function() {
    it('should create an instance of ClusterApiConfig', function() {
      // uncomment below and update the code to test ClusterApiConfig
      //var instane = new MorpheusApi.ClusterApiConfig();
      //expect(instance).to.be.a(MorpheusApi.ClusterApiConfig);
    });

    it('should have the property serviceUrl (base name: "serviceUrl")', function() {
      // uncomment below and update the code to test the property serviceUrl
      //var instance = new MorpheusApi.ClusterApiConfig();
      //expect(instance).to.be();
    });

    it('should have the property serviceHost (base name: "serviceHost")', function() {
      // uncomment below and update the code to test the property serviceHost
      //var instance = new MorpheusApi.ClusterApiConfig();
      //expect(instance).to.be();
    });

    it('should have the property servicePath (base name: "servicePath")', function() {
      // uncomment below and update the code to test the property servicePath
      //var instance = new MorpheusApi.ClusterApiConfig();
      //expect(instance).to.be();
    });

    it('should have the property serviceHostname (base name: "serviceHostname")', function() {
      // uncomment below and update the code to test the property serviceHostname
      //var instance = new MorpheusApi.ClusterApiConfig();
      //expect(instance).to.be();
    });

    it('should have the property servicePort (base name: "servicePort")', function() {
      // uncomment below and update the code to test the property servicePort
      //var instance = new MorpheusApi.ClusterApiConfig();
      //expect(instance).to.be();
    });

    it('should have the property serviceUsername (base name: "serviceUsername")', function() {
      // uncomment below and update the code to test the property serviceUsername
      //var instance = new MorpheusApi.ClusterApiConfig();
      //expect(instance).to.be();
    });

    it('should have the property servicePassword (base name: "servicePassword")', function() {
      // uncomment below and update the code to test the property servicePassword
      //var instance = new MorpheusApi.ClusterApiConfig();
      //expect(instance).to.be();
    });

    it('should have the property servicePasswordHash (base name: "servicePasswordHash")', function() {
      // uncomment below and update the code to test the property servicePasswordHash
      //var instance = new MorpheusApi.ClusterApiConfig();
      //expect(instance).to.be();
    });

    it('should have the property serviceToken (base name: "serviceToken")', function() {
      // uncomment below and update the code to test the property serviceToken
      //var instance = new MorpheusApi.ClusterApiConfig();
      //expect(instance).to.be();
    });

    it('should have the property serviceAccess (base name: "serviceAccess")', function() {
      // uncomment below and update the code to test the property serviceAccess
      //var instance = new MorpheusApi.ClusterApiConfig();
      //expect(instance).to.be();
    });

    it('should have the property serviceCert (base name: "serviceCert")', function() {
      // uncomment below and update the code to test the property serviceCert
      //var instance = new MorpheusApi.ClusterApiConfig();
      //expect(instance).to.be();
    });

    it('should have the property serviceVersion (base name: "serviceVersion")', function() {
      // uncomment below and update the code to test the property serviceVersion
      //var instance = new MorpheusApi.ClusterApiConfig();
      //expect(instance).to.be();
    });

  });

}));
