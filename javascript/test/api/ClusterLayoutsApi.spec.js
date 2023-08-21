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
    instance = new MorpheusApi.ClusterLayoutsApi();
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

  describe('ClusterLayoutsApi', function() {
    describe('addClusterLayoutClone', function() {
      it('should call addClusterLayoutClone successfully', function(done) {
        //uncomment below and update the code to test addClusterLayoutClone
        //instance.addClusterLayoutClone(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('addClusterLayouts', function() {
      it('should call addClusterLayouts successfully', function(done) {
        //uncomment below and update the code to test addClusterLayouts
        //instance.addClusterLayouts(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('deleteClusterLayout', function() {
      it('should call deleteClusterLayout successfully', function(done) {
        //uncomment below and update the code to test deleteClusterLayout
        //instance.deleteClusterLayout(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('getClusterLayout', function() {
      it('should call getClusterLayout successfully', function(done) {
        //uncomment below and update the code to test getClusterLayout
        //instance.getClusterLayout(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('listClusterLayouts', function() {
      it('should call listClusterLayouts successfully', function(done) {
        //uncomment below and update the code to test listClusterLayouts
        //instance.listClusterLayouts(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('updateClusterLayout', function() {
      it('should call updateClusterLayout successfully', function(done) {
        //uncomment below and update the code to test updateClusterLayout
        //instance.updateClusterLayout(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
  });

}));
