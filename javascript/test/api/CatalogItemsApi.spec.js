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
    instance = new MorpheusApi.CatalogItemsApi();
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

  describe('CatalogItemsApi', function() {
    describe('addCatalogItemType', function() {
      it('should call addCatalogItemType successfully', function(done) {
        //uncomment below and update the code to test addCatalogItemType
        //instance.addCatalogItemType(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('getCatalogItemType', function() {
      it('should call getCatalogItemType successfully', function(done) {
        //uncomment below and update the code to test getCatalogItemType
        //instance.getCatalogItemType(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('listCatalogItemTypes', function() {
      it('should call listCatalogItemTypes successfully', function(done) {
        //uncomment below and update the code to test listCatalogItemTypes
        //instance.listCatalogItemTypes(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('removeCatalogItemType', function() {
      it('should call removeCatalogItemType successfully', function(done) {
        //uncomment below and update the code to test removeCatalogItemType
        //instance.removeCatalogItemType(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('updateCatalogItemType', function() {
      it('should call updateCatalogItemType successfully', function(done) {
        //uncomment below and update the code to test updateCatalogItemType
        //instance.updateCatalogItemType(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('updateCatalogItemTypeLogo', function() {
      it('should call updateCatalogItemTypeLogo successfully', function(done) {
        //uncomment below and update the code to test updateCatalogItemTypeLogo
        //instance.updateCatalogItemTypeLogo(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
  });

}));
