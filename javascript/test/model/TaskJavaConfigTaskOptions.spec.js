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
    instance = new MorpheusApi.TaskJavaConfigTaskOptions();
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

  describe('TaskJavaConfigTaskOptions', function() {
    it('should create an instance of TaskJavaConfigTaskOptions', function() {
      // uncomment below and update the code to test TaskJavaConfigTaskOptions
      //var instane = new MorpheusApi.TaskJavaConfigTaskOptions();
      //expect(instance).to.be.a(MorpheusApi.TaskJavaConfigTaskOptions);
    });

    it('should have the property username (base name: "username")', function() {
      // uncomment below and update the code to test the property username
      //var instance = new MorpheusApi.TaskJavaConfigTaskOptions();
      //expect(instance).to.be();
    });

    it('should have the property port (base name: "port")', function() {
      // uncomment below and update the code to test the property port
      //var instance = new MorpheusApi.TaskJavaConfigTaskOptions();
      //expect(instance).to.be();
    });

    it('should have the property jsScript (base name: "jsScript")', function() {
      // uncomment below and update the code to test the property jsScript
      //var instance = new MorpheusApi.TaskJavaConfigTaskOptions();
      //expect(instance).to.be();
    });

    it('should have the property host (base name: "host")', function() {
      // uncomment below and update the code to test the property host
      //var instance = new MorpheusApi.TaskJavaConfigTaskOptions();
      //expect(instance).to.be();
    });

    it('should have the property localScriptGitRef (base name: "localScriptGitRef")', function() {
      // uncomment below and update the code to test the property localScriptGitRef
      //var instance = new MorpheusApi.TaskJavaConfigTaskOptions();
      //expect(instance).to.be();
    });

    it('should have the property password (base name: "password")', function() {
      // uncomment below and update the code to test the property password
      //var instance = new MorpheusApi.TaskJavaConfigTaskOptions();
      //expect(instance).to.be();
    });

    it('should have the property passwordHash (base name: "passwordHash")', function() {
      // uncomment below and update the code to test the property passwordHash
      //var instance = new MorpheusApi.TaskJavaConfigTaskOptions();
      //expect(instance).to.be();
    });

    it('should have the property sshKey (base name: "sshKey")', function() {
      // uncomment below and update the code to test the property sshKey
      //var instance = new MorpheusApi.TaskJavaConfigTaskOptions();
      //expect(instance).to.be();
    });

    it('should have the property localScriptGitId (base name: "localScriptGitId")', function() {
      // uncomment below and update the code to test the property localScriptGitId
      //var instance = new MorpheusApi.TaskJavaConfigTaskOptions();
      //expect(instance).to.be();
    });

  });

}));
