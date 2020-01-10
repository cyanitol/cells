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

var _modelRestFrontBinaryRequest = require('../model/RestFrontBinaryRequest');

var _modelRestFrontBinaryRequest2 = _interopRequireDefault(_modelRestFrontBinaryRequest);

var _modelRestFrontBinaryResponse = require('../model/RestFrontBinaryResponse');

var _modelRestFrontBinaryResponse2 = _interopRequireDefault(_modelRestFrontBinaryResponse);

var _modelRestFrontBootConfResponse = require('../model/RestFrontBootConfResponse');

var _modelRestFrontBootConfResponse2 = _interopRequireDefault(_modelRestFrontBootConfResponse);

var _modelRestFrontEnrollAuthRequest = require('../model/RestFrontEnrollAuthRequest');

var _modelRestFrontEnrollAuthRequest2 = _interopRequireDefault(_modelRestFrontEnrollAuthRequest);

var _modelRestFrontEnrollAuthResponse = require('../model/RestFrontEnrollAuthResponse');

var _modelRestFrontEnrollAuthResponse2 = _interopRequireDefault(_modelRestFrontEnrollAuthResponse);

var _modelRestFrontMessagesResponse = require('../model/RestFrontMessagesResponse');

var _modelRestFrontMessagesResponse2 = _interopRequireDefault(_modelRestFrontMessagesResponse);

var _modelRestFrontPluginsResponse = require('../model/RestFrontPluginsResponse');

var _modelRestFrontPluginsResponse2 = _interopRequireDefault(_modelRestFrontPluginsResponse);

var _modelRestFrontSessionDelResponse = require('../model/RestFrontSessionDelResponse');

var _modelRestFrontSessionDelResponse2 = _interopRequireDefault(_modelRestFrontSessionDelResponse);

var _modelRestFrontSessionGetResponse = require('../model/RestFrontSessionGetResponse');

var _modelRestFrontSessionGetResponse2 = _interopRequireDefault(_modelRestFrontSessionGetResponse);

var _modelRestFrontSessionRequest = require('../model/RestFrontSessionRequest');

var _modelRestFrontSessionRequest2 = _interopRequireDefault(_modelRestFrontSessionRequest);

var _modelRestFrontSessionResponse = require('../model/RestFrontSessionResponse');

var _modelRestFrontSessionResponse2 = _interopRequireDefault(_modelRestFrontSessionResponse);

var _modelRestFrontStateResponse = require('../model/RestFrontStateResponse');

var _modelRestFrontStateResponse2 = _interopRequireDefault(_modelRestFrontStateResponse);

var _modelRestSettingsMenuResponse = require('../model/RestSettingsMenuResponse');

var _modelRestSettingsMenuResponse2 = _interopRequireDefault(_modelRestSettingsMenuResponse);

/**
* FrontendService service.
* @module api/FrontendServiceApi
* @version 1.0
*/

var FrontendServiceApi = (function () {

  /**
  * Constructs a new FrontendServiceApi. 
  * @alias module:api/FrontendServiceApi
  * @class
  * @param {module:ApiClient} [apiClient] Optional API client implementation to use,
  * default to {@link module:ApiClient#instance} if unspecified.
  */

  function FrontendServiceApi(apiClient) {
    _classCallCheck(this, FrontendServiceApi);

    this.apiClient = apiClient || _ApiClient2['default'].instance;
  }

  /**
   * Remove Session
   * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with an object containing data of type {@link module:model/RestFrontSessionDelResponse} and HTTP response
   */

  FrontendServiceApi.prototype.fronSessionDelWithHttpInfo = function fronSessionDelWithHttpInfo() {
    var postBody = null;

    var pathParams = {};
    var queryParams = {};
    var headerParams = {};
    var formParams = {};

    var authNames = [];
    var contentTypes = ['application/json'];
    var accepts = ['application/json'];
    var returnType = _modelRestFrontSessionDelResponse2['default'];

    return this.apiClient.callApi('/frontend/session', 'DELETE', pathParams, queryParams, headerParams, formParams, postBody, authNames, contentTypes, accepts, returnType);
  };

  /**
   * Remove Session
   * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with data of type {@link module:model/RestFrontSessionDelResponse}
   */

  FrontendServiceApi.prototype.fronSessionDel = function fronSessionDel() {
    return this.fronSessionDelWithHttpInfo().then(function (response_and_data) {
      return response_and_data.data;
    });
  };

  /**
   * Add some data to the initial set of parameters loaded by the frontend
   * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with an object containing data of type {@link module:model/RestFrontBootConfResponse} and HTTP response
   */

  FrontendServiceApi.prototype.frontBootConfWithHttpInfo = function frontBootConfWithHttpInfo() {
    var postBody = null;

    var pathParams = {};
    var queryParams = {};
    var headerParams = {};
    var formParams = {};

    var authNames = [];
    var contentTypes = ['application/json'];
    var accepts = ['application/json'];
    var returnType = _modelRestFrontBootConfResponse2['default'];

    return this.apiClient.callApi('/frontend/bootconf', 'GET', pathParams, queryParams, headerParams, formParams, postBody, authNames, contentTypes, accepts, returnType);
  };

  /**
   * Add some data to the initial set of parameters loaded by the frontend
   * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with data of type {@link module:model/RestFrontBootConfResponse}
   */

  FrontendServiceApi.prototype.frontBootConf = function frontBootConf() {
    return this.frontBootConfWithHttpInfo().then(function (response_and_data) {
      return response_and_data.data;
    });
  };

  /**
   * Generic endpoint that can be implemented by 2FA systems for enrollment
   * @param {module:model/RestFrontEnrollAuthRequest} body 
   * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with an object containing data of type {@link module:model/RestFrontEnrollAuthResponse} and HTTP response
   */

  FrontendServiceApi.prototype.frontEnrollAuthWithHttpInfo = function frontEnrollAuthWithHttpInfo(body) {
    var postBody = body;

    // verify the required parameter 'body' is set
    if (body === undefined || body === null) {
      throw new Error("Missing the required parameter 'body' when calling frontEnrollAuth");
    }

    var pathParams = {};
    var queryParams = {};
    var headerParams = {};
    var formParams = {};

    var authNames = [];
    var contentTypes = ['application/json'];
    var accepts = ['application/json'];
    var returnType = _modelRestFrontEnrollAuthResponse2['default'];

    return this.apiClient.callApi('/frontend/enroll', 'POST', pathParams, queryParams, headerParams, formParams, postBody, authNames, contentTypes, accepts, returnType);
  };

  /**
   * Generic endpoint that can be implemented by 2FA systems for enrollment
   * @param {module:model/RestFrontEnrollAuthRequest} body 
   * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with data of type {@link module:model/RestFrontEnrollAuthResponse}
   */

  FrontendServiceApi.prototype.frontEnrollAuth = function frontEnrollAuth(body) {
    return this.frontEnrollAuthWithHttpInfo(body).then(function (response_and_data) {
      return response_and_data.data;
    });
  };

  /**
   * Serve list of I18n messages
   * @param {String} lang 
   * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with an object containing data of type {@link module:model/RestFrontMessagesResponse} and HTTP response
   */

  FrontendServiceApi.prototype.frontMessagesWithHttpInfo = function frontMessagesWithHttpInfo(lang) {
    var postBody = null;

    // verify the required parameter 'lang' is set
    if (lang === undefined || lang === null) {
      throw new Error("Missing the required parameter 'lang' when calling frontMessages");
    }

    var pathParams = {
      'Lang': lang
    };
    var queryParams = {};
    var headerParams = {};
    var formParams = {};

    var authNames = [];
    var contentTypes = ['application/json'];
    var accepts = ['application/json'];
    var returnType = _modelRestFrontMessagesResponse2['default'];

    return this.apiClient.callApi('/frontend/messages/{Lang}', 'GET', pathParams, queryParams, headerParams, formParams, postBody, authNames, contentTypes, accepts, returnType);
  };

  /**
   * Serve list of I18n messages
   * @param {String} lang 
   * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with data of type {@link module:model/RestFrontMessagesResponse}
   */

  FrontendServiceApi.prototype.frontMessages = function frontMessages(lang) {
    return this.frontMessagesWithHttpInfo(lang).then(function (response_and_data) {
      return response_and_data.data;
    });
  };

  /**
   * Serve list of I18n messages
   * @param {String} lang 
   * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with an object containing data of type {@link module:model/RestFrontPluginsResponse} and HTTP response
   */

  FrontendServiceApi.prototype.frontPluginsWithHttpInfo = function frontPluginsWithHttpInfo(lang) {
    var postBody = null;

    // verify the required parameter 'lang' is set
    if (lang === undefined || lang === null) {
      throw new Error("Missing the required parameter 'lang' when calling frontPlugins");
    }

    var pathParams = {
      'Lang': lang
    };
    var queryParams = {};
    var headerParams = {};
    var formParams = {};

    var authNames = [];
    var contentTypes = ['application/json'];
    var accepts = ['application/json'];
    var returnType = _modelRestFrontPluginsResponse2['default'];

    return this.apiClient.callApi('/frontend/plugins/{Lang}', 'GET', pathParams, queryParams, headerParams, formParams, postBody, authNames, contentTypes, accepts, returnType);
  };

  /**
   * Serve list of I18n messages
   * @param {String} lang 
   * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with data of type {@link module:model/RestFrontPluginsResponse}
   */

  FrontendServiceApi.prototype.frontPlugins = function frontPlugins(lang) {
    return this.frontPluginsWithHttpInfo(lang).then(function (response_and_data) {
      return response_and_data.data;
    });
  };

  /**
   * Upload frontend binaries (avatars / logos / bg images)
   * @param {String} binaryType 
   * @param {String} uuid 
   * @param {module:model/RestFrontBinaryRequest} body 
   * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with an object containing data of type {@link module:model/RestFrontBinaryResponse} and HTTP response
   */

  FrontendServiceApi.prototype.frontPutBinaryWithHttpInfo = function frontPutBinaryWithHttpInfo(binaryType, uuid, body) {
    var postBody = body;

    // verify the required parameter 'binaryType' is set
    if (binaryType === undefined || binaryType === null) {
      throw new Error("Missing the required parameter 'binaryType' when calling frontPutBinary");
    }

    // verify the required parameter 'uuid' is set
    if (uuid === undefined || uuid === null) {
      throw new Error("Missing the required parameter 'uuid' when calling frontPutBinary");
    }

    // verify the required parameter 'body' is set
    if (body === undefined || body === null) {
      throw new Error("Missing the required parameter 'body' when calling frontPutBinary");
    }

    var pathParams = {
      'BinaryType': binaryType,
      'Uuid': uuid
    };
    var queryParams = {};
    var headerParams = {};
    var formParams = {};

    var authNames = [];
    var contentTypes = ['application/json'];
    var accepts = ['application/json'];
    var returnType = _modelRestFrontBinaryResponse2['default'];

    return this.apiClient.callApi('/frontend/binaries/{BinaryType}/{Uuid}', 'POST', pathParams, queryParams, headerParams, formParams, postBody, authNames, contentTypes, accepts, returnType);
  };

  /**
   * Upload frontend binaries (avatars / logos / bg images)
   * @param {String} binaryType 
   * @param {String} uuid 
   * @param {module:model/RestFrontBinaryRequest} body 
   * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with data of type {@link module:model/RestFrontBinaryResponse}
   */

  FrontendServiceApi.prototype.frontPutBinary = function frontPutBinary(binaryType, uuid, body) {
    return this.frontPutBinaryWithHttpInfo(binaryType, uuid, body).then(function (response_and_data) {
      return response_and_data.data;
    });
  };

  /**
   * Serve frontend binaries directly (avatars / logos / bg images)
   * @param {String} binaryType 
   * @param {String} uuid 
   * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with an object containing data of type {@link module:model/RestFrontBinaryResponse} and HTTP response
   */

  FrontendServiceApi.prototype.frontServeBinaryWithHttpInfo = function frontServeBinaryWithHttpInfo(binaryType, uuid) {
    var postBody = null;

    // verify the required parameter 'binaryType' is set
    if (binaryType === undefined || binaryType === null) {
      throw new Error("Missing the required parameter 'binaryType' when calling frontServeBinary");
    }

    // verify the required parameter 'uuid' is set
    if (uuid === undefined || uuid === null) {
      throw new Error("Missing the required parameter 'uuid' when calling frontServeBinary");
    }

    var pathParams = {
      'BinaryType': binaryType,
      'Uuid': uuid
    };
    var queryParams = {};
    var headerParams = {};
    var formParams = {};

    var authNames = [];
    var contentTypes = ['application/json'];
    var accepts = ['application/json'];
    var returnType = _modelRestFrontBinaryResponse2['default'];

    return this.apiClient.callApi('/frontend/binaries/{BinaryType}/{Uuid}', 'GET', pathParams, queryParams, headerParams, formParams, postBody, authNames, contentTypes, accepts, returnType);
  };

  /**
   * Serve frontend binaries directly (avatars / logos / bg images)
   * @param {String} binaryType 
   * @param {String} uuid 
   * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with data of type {@link module:model/RestFrontBinaryResponse}
   */

  FrontendServiceApi.prototype.frontServeBinary = function frontServeBinary(binaryType, uuid) {
    return this.frontServeBinaryWithHttpInfo(binaryType, uuid).then(function (response_and_data) {
      return response_and_data.data;
    });
  };

  /**
   * Handle Session POST
   * @param {module:model/RestFrontSessionRequest} body 
   * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with an object containing data of type {@link module:model/RestFrontSessionResponse} and HTTP response
   */

  FrontendServiceApi.prototype.frontSessionWithHttpInfo = function frontSessionWithHttpInfo(body) {
    var postBody = body;

    // verify the required parameter 'body' is set
    if (body === undefined || body === null) {
      throw new Error("Missing the required parameter 'body' when calling frontSession");
    }

    var pathParams = {};
    var queryParams = {};
    var headerParams = {};
    var formParams = {};

    var authNames = [];
    var contentTypes = ['application/json'];
    var accepts = ['application/json'];
    var returnType = _modelRestFrontSessionResponse2['default'];

    return this.apiClient.callApi('/frontend/session', 'POST', pathParams, queryParams, headerParams, formParams, postBody, authNames, contentTypes, accepts, returnType);
  };

  /**
   * Handle Session POST
   * @param {module:model/RestFrontSessionRequest} body 
   * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with data of type {@link module:model/RestFrontSessionResponse}
   */

  FrontendServiceApi.prototype.frontSession = function frontSession(body) {
    return this.frontSessionWithHttpInfo(body).then(function (response_and_data) {
      return response_and_data.data;
    });
  };

  /**
   * Retrieve session info
   * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with an object containing data of type {@link module:model/RestFrontSessionGetResponse} and HTTP response
   */

  FrontendServiceApi.prototype.frontSessionGetWithHttpInfo = function frontSessionGetWithHttpInfo() {
    var postBody = null;

    var pathParams = {};
    var queryParams = {};
    var headerParams = {};
    var formParams = {};

    var authNames = [];
    var contentTypes = ['application/json'];
    var accepts = ['application/json'];
    var returnType = _modelRestFrontSessionGetResponse2['default'];

    return this.apiClient.callApi('/frontend/session', 'GET', pathParams, queryParams, headerParams, formParams, postBody, authNames, contentTypes, accepts, returnType);
  };

  /**
   * Retrieve session info
   * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with data of type {@link module:model/RestFrontSessionGetResponse}
   */

  FrontendServiceApi.prototype.frontSessionGet = function frontSessionGet() {
    return this.frontSessionGetWithHttpInfo().then(function (response_and_data) {
      return response_and_data.data;
    });
  };

  /**
   * Send XML state registry
   * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with an object containing data of type {@link module:model/RestFrontStateResponse} and HTTP response
   */

  FrontendServiceApi.prototype.frontStateWithHttpInfo = function frontStateWithHttpInfo() {
    var postBody = null;

    var pathParams = {};
    var queryParams = {};
    var headerParams = {};
    var formParams = {};

    var authNames = [];
    var contentTypes = ['application/json'];
    var accepts = ['application/json'];
    var returnType = _modelRestFrontStateResponse2['default'];

    return this.apiClient.callApi('/frontend/state', 'GET', pathParams, queryParams, headerParams, formParams, postBody, authNames, contentTypes, accepts, returnType);
  };

  /**
   * Send XML state registry
   * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with data of type {@link module:model/RestFrontStateResponse}
   */

  FrontendServiceApi.prototype.frontState = function frontState() {
    return this.frontStateWithHttpInfo().then(function (response_and_data) {
      return response_and_data.data;
    });
  };

  /**
   * Sends a tree of nodes to be used a menu in the Settings panel
   * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with an object containing data of type {@link module:model/RestSettingsMenuResponse} and HTTP response
   */

  FrontendServiceApi.prototype.settingsMenuWithHttpInfo = function settingsMenuWithHttpInfo() {
    var postBody = null;

    var pathParams = {};
    var queryParams = {};
    var headerParams = {};
    var formParams = {};

    var authNames = [];
    var contentTypes = ['application/json'];
    var accepts = ['application/json'];
    var returnType = _modelRestSettingsMenuResponse2['default'];

    return this.apiClient.callApi('/frontend/settings-menu', 'GET', pathParams, queryParams, headerParams, formParams, postBody, authNames, contentTypes, accepts, returnType);
  };

  /**
   * Sends a tree of nodes to be used a menu in the Settings panel
   * @return {Promise} a {@link https://www.promisejs.org/|Promise}, with data of type {@link module:model/RestSettingsMenuResponse}
   */

  FrontendServiceApi.prototype.settingsMenu = function settingsMenu() {
    return this.settingsMenuWithHttpInfo().then(function (response_and_data) {
      return response_and_data.data;
    });
  };

  return FrontendServiceApi;
})();

exports['default'] = FrontendServiceApi;
module.exports = exports['default'];
