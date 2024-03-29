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
    instance = new MorpheusApi.CheckWebConfig();
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

  describe('CheckWebConfig', function() {
    it('should create an instance of CheckWebConfig', function() {
      // uncomment below and update the code to test CheckWebConfig
      //var instane = new MorpheusApi.CheckWebConfig();
      //expect(instance).to.be.a(MorpheusApi.CheckWebConfig);
    });

    it('should have the property webMethod (base name: "webMethod")', function() {
      // uncomment below and update the code to test the property webMethod
      //var instance = new MorpheusApi.CheckWebConfig();
      //expect(instance).to.be();
    });

    it('should have the property webUrl (base name: "webUrl")', function() {
      // uncomment below and update the code to test the property webUrl
      //var instance = new MorpheusApi.CheckWebConfig();
      //expect(instance).to.be();
    });

    it('should have the property ignoreSSL (base name: "ignoreSSL")', function() {
      // uncomment below and update the code to test the property ignoreSSL
      //var instance = new MorpheusApi.CheckWebConfig();
      //expect(instance).to.be();
    });

    it('should have the property checkUser (base name: "checkUser")', function() {
      // uncomment below and update the code to test the property checkUser
      //var instance = new MorpheusApi.CheckWebConfig();
      //expect(instance).to.be();
    });

    it('should have the property checkPassword (base name: "checkPassword")', function() {
      // uncomment below and update the code to test the property checkPassword
      //var instance = new MorpheusApi.CheckWebConfig();
      //expect(instance).to.be();
    });

    it('should have the property textCheckOn (base name: "textCheckOn")', function() {
      // uncomment below and update the code to test the property textCheckOn
      //var instance = new MorpheusApi.CheckWebConfig();
      //expect(instance).to.be();
    });

    it('should have the property webTextMatch (base name: "webTextMatch")', function() {
      // uncomment below and update the code to test the property webTextMatch
      //var instance = new MorpheusApi.CheckWebConfig();
      //expect(instance).to.be();
    });

    it('should have the property tunnelOn (base name: "tunnelOn")', function() {
      // uncomment below and update the code to test the property tunnelOn
      //var instance = new MorpheusApi.CheckWebConfig();
      //expect(instance).to.be();
    });

    it('should have the property sshHost (base name: "sshHost")', function() {
      // uncomment below and update the code to test the property sshHost
      //var instance = new MorpheusApi.CheckWebConfig();
      //expect(instance).to.be();
    });

    it('should have the property sshPort (base name: "sshPort")', function() {
      // uncomment below and update the code to test the property sshPort
      //var instance = new MorpheusApi.CheckWebConfig();
      //expect(instance).to.be();
    });

    it('should have the property sshUser (base name: "sshUser")', function() {
      // uncomment below and update the code to test the property sshUser
      //var instance = new MorpheusApi.CheckWebConfig();
      //expect(instance).to.be();
    });

    it('should have the property sshPassword (base name: "sshPassword")', function() {
      // uncomment below and update the code to test the property sshPassword
      //var instance = new MorpheusApi.CheckWebConfig();
      //expect(instance).to.be();
    });

  });

}));
