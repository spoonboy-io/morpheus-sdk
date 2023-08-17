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
    instance = new MorpheusApi.BlueprintKubernetesCreateKubernetesGit();
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

  describe('BlueprintKubernetesCreateKubernetesGit', function() {
    it('should create an instance of BlueprintKubernetesCreateKubernetesGit', function() {
      // uncomment below and update the code to test BlueprintKubernetesCreateKubernetesGit
      //var instance = new MorpheusApi.BlueprintKubernetesCreateKubernetesGit();
      //expect(instance).to.be.a(MorpheusApi.BlueprintKubernetesCreateKubernetesGit);
    });

    it('should have the property repoId (base name: "repoId")', function() {
      // uncomment below and update the code to test the property repoId
      //var instance = new MorpheusApi.BlueprintKubernetesCreateKubernetesGit();
      //expect(instance).to.be();
    });

    it('should have the property path (base name: "path")', function() {
      // uncomment below and update the code to test the property path
      //var instance = new MorpheusApi.BlueprintKubernetesCreateKubernetesGit();
      //expect(instance).to.be();
    });

    it('should have the property integrationId (base name: "integrationId")', function() {
      // uncomment below and update the code to test the property integrationId
      //var instance = new MorpheusApi.BlueprintKubernetesCreateKubernetesGit();
      //expect(instance).to.be();
    });

    it('should have the property branch (base name: "branch")', function() {
      // uncomment below and update the code to test the property branch
      //var instance = new MorpheusApi.BlueprintKubernetesCreateKubernetesGit();
      //expect(instance).to.be();
    });

  });

}));
