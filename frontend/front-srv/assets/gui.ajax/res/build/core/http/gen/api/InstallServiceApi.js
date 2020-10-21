/**
 * Pydio Cells Rest API
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * OpenAPI spec version: 1.0
 * 
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 * Do not edit the class manually.
 *
 */

'use strict';

exports.__esModule = true;

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { 'default': obj }; }

function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError('Cannot call a class as a function'); } }

var _ApiClient = require("../ApiClient");

var _ApiClient2 = _interopRequireDefault(_ApiClient);

var _modelInstallGetAgreementResponse = require('../model/InstallGetAgreementResponse');

var _modelInstallGetAgreementResponse2 = _interopRequireDefault(_modelInstallGetAgreementResponse);

var _modelInstallGetDefaultsResponse = require('../model/InstallGetDefaultsResponse');

var _modelInstallGetDefaultsResponse2 = _interopRequireDefault(_modelInstallGetDefaultsResponse);

var _modelInstallInstallEventsResponse = require('../model/InstallInstallEventsResponse');

var _modelInstallInstallEventsResponse2 = _interopRequireDefault(_modelInstallInstallEventsResponse);

var _modelInstallInstallRequest = require('../model/InstallInstallRequest');

var _modelInstallInstallRequest2 = _interopRequireDefault(_modelInstallInstallRequest);

var _modelInstallInstallResponse = require('../model/InstallInstallResponse');

var _modelInstallInstallResponse2 = _interopRequireDefault(_modelInstallInstallResponse);

var _modelInstallPerformCheckRequest = require('../model/InstallPerformCheckRequest');

var _modelInstallPerformCheckRequest2 = _interopRequireDefault(_modelInstallPerformCheckRequest);

var _modelInstallPerformCheckResponse = require('../model/InstallPerformCheckResponse');

var _modelInstallPerformCheckResponse2 = _interopRequireDefault(_modelInstallPerformCheckResponse);

/**
* InstallService service.
* @module api/InstallServiceApi
* @version 1.0
*/

var InstallServiceApi = (function () {

  /**
  * Constructs a new InstallServiceApi. 
  * @alias module:api/InstallServiceApi
  * @class
  * @param {module:ApiClient} [apiClient] Optional API client implementation to use,
  * default to {@link module:ApiClient#instance} if unspecified.
  */

  function InstallServiceApi(apiClient) {
    _classCallCheck(this, InstallServiceApi);

    this.apiClient = apiClient || _ApiClient2['default'].instance;
  }

  /**
   * Load a textual agreement for using the software
   * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with an object containing data of type {@link module:model/InstallGetAgreementResponse} and HTTP response
   */

  InstallServiceApi.prototype.getAgreementWithHttpInfo = function getAgreementWithHttpInfo() {
    var postBody = null;

    var pathParams = {};
    var queryParams = {};
    var headerParams = {};
    var formParams = {};

    var authNames = [];
    var contentTypes = ['application/json'];
    var accepts = ['application/json'];
    var returnType = _modelInstallGetAgreementResponse2['default'];

    return this.apiClient.callApi('/install/agreement', 'GET', pathParams, queryParams, headerParams, formParams, postBody, authNames, contentTypes, accepts, returnType);
  };

  /**
   * Load a textual agreement for using the software
   * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with data of type {@link module:model/InstallGetAgreementResponse}
   */

  InstallServiceApi.prototype.getAgreement = function getAgreement() {
    return this.getAgreementWithHttpInfo().then(function (response_and_data) {
      return response_and_data.data;
    });
  };

  /**
   * Loads default values for install form
   * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with an object containing data of type {@link module:model/InstallGetDefaultsResponse} and HTTP response
   */

  InstallServiceApi.prototype.getInstallWithHttpInfo = function getInstallWithHttpInfo() {
    var postBody = null;

    var pathParams = {};
    var queryParams = {};
    var headerParams = {};
    var formParams = {};

    var authNames = [];
    var contentTypes = ['application/json'];
    var accepts = ['application/json'];
    var returnType = _modelInstallGetDefaultsResponse2['default'];

    return this.apiClient.callApi('/install', 'GET', pathParams, queryParams, headerParams, formParams, postBody, authNames, contentTypes, accepts, returnType);
  };

  /**
   * Loads default values for install form
   * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with data of type {@link module:model/InstallGetDefaultsResponse}
   */

  InstallServiceApi.prototype.getInstall = function getInstall() {
    return this.getInstallWithHttpInfo().then(function (response_and_data) {
      return response_and_data.data;
    });
  };

  /**
   * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with an object containing data of type {@link module:model/InstallInstallEventsResponse} and HTTP response
   */

  InstallServiceApi.prototype.installEventsWithHttpInfo = function installEventsWithHttpInfo() {
    var postBody = null;

    var pathParams = {};
    var queryParams = {};
    var headerParams = {};
    var formParams = {};

    var authNames = [];
    var contentTypes = ['application/json'];
    var accepts = ['application/json'];
    var returnType = _modelInstallInstallEventsResponse2['default'];

    return this.apiClient.callApi('/install/events', 'GET', pathParams, queryParams, headerParams, formParams, postBody, authNames, contentTypes, accepts, returnType);
  };

  /**
   * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with data of type {@link module:model/InstallInstallEventsResponse}
   */

  InstallServiceApi.prototype.installEvents = function installEvents() {
    return this.installEventsWithHttpInfo().then(function (response_and_data) {
      return response_and_data.data;
    });
  };

  /**
   * Perform a check during install (like a valid DB connection)
   * @param {module:model/InstallPerformCheckRequest} body 
   * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with an object containing data of type {@link module:model/InstallPerformCheckResponse} and HTTP response
   */

  InstallServiceApi.prototype.performInstallCheckWithHttpInfo = function performInstallCheckWithHttpInfo(body) {
    var postBody = body;

    // verify the required parameter 'body' is set
    if (body === undefined || body === null) {
      throw new Error("Missing the required parameter 'body' when calling performInstallCheck");
    }

    var pathParams = {};
    var queryParams = {};
    var headerParams = {};
    var formParams = {};

    var authNames = [];
    var contentTypes = ['application/json'];
    var accepts = ['application/json'];
    var returnType = _modelInstallPerformCheckResponse2['default'];

    return this.apiClient.callApi('/install/check', 'POST', pathParams, queryParams, headerParams, formParams, postBody, authNames, contentTypes, accepts, returnType);
  };

  /**
   * Perform a check during install (like a valid DB connection)
   * @param {module:model/InstallPerformCheckRequest} body 
   * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with data of type {@link module:model/InstallPerformCheckResponse}
   */

  InstallServiceApi.prototype.performInstallCheck = function performInstallCheck(body) {
    return this.performInstallCheckWithHttpInfo(body).then(function (response_and_data) {
      return response_and_data.data;
    });
  };

  /**
   * Post values to be saved for install
   * @param {module:model/InstallInstallRequest} body 
   * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with an object containing data of type {@link module:model/InstallInstallResponse} and HTTP response
   */

  InstallServiceApi.prototype.postInstallWithHttpInfo = function postInstallWithHttpInfo(body) {
    var postBody = body;

    // verify the required parameter 'body' is set
    if (body === undefined || body === null) {
      throw new Error("Missing the required parameter 'body' when calling postInstall");
    }

    var pathParams = {};
    var queryParams = {};
    var headerParams = {};
    var formParams = {};

    var authNames = [];
    var contentTypes = ['application/json'];
    var accepts = ['application/json'];
    var returnType = _modelInstallInstallResponse2['default'];

    return this.apiClient.callApi('/install', 'POST', pathParams, queryParams, headerParams, formParams, postBody, authNames, contentTypes, accepts, returnType);
  };

  /**
   * Post values to be saved for install
   * @param {module:model/InstallInstallRequest} body 
   * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with data of type {@link module:model/InstallInstallResponse}
   */

  InstallServiceApi.prototype.postInstall = function postInstall(body) {
    return this.postInstallWithHttpInfo(body).then(function (response_and_data) {
      return response_and_data.data;
    });
  };

  return InstallServiceApi;
})();

exports['default'] = InstallServiceApi;
module.exports = exports['default'];
