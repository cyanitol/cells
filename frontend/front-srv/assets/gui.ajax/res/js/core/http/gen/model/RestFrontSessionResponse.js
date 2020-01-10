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


import ApiClient from '../ApiClient';
import RestToken from './RestToken';





/**
* The RestFrontSessionResponse model module.
* @module model/RestFrontSessionResponse
* @version 1.0
*/
export default class RestFrontSessionResponse {
    /**
    * Constructs a new <code>RestFrontSessionResponse</code>.
    * @alias module:model/RestFrontSessionResponse
    * @class
    */

    constructor() {
        

        
        

        

        
    }

    /**
    * Constructs a <code>RestFrontSessionResponse</code> from a plain JavaScript object, optionally creating a new instance.
    * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
    * @param {Object} data The plain JavaScript object bearing properties of interest.
    * @param {module:model/RestFrontSessionResponse} obj Optional instance to populate.
    * @return {module:model/RestFrontSessionResponse} The populated <code>RestFrontSessionResponse</code> instance.
    */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new RestFrontSessionResponse();

            
            
            

            if (data.hasOwnProperty('Token')) {
                obj['Token'] = RestToken.constructFromObject(data['Token']);
            }
            if (data.hasOwnProperty('Trigger')) {
                obj['Trigger'] = ApiClient.convertToType(data['Trigger'], 'String');
            }
            if (data.hasOwnProperty('TriggerInfo')) {
                obj['TriggerInfo'] = ApiClient.convertToType(data['TriggerInfo'], {'String': 'String'});
            }
            if (data.hasOwnProperty('RedirectTo')) {
                obj['RedirectTo'] = ApiClient.convertToType(data['RedirectTo'], 'String');
            }
        }
        return obj;
    }

    /**
    * @member {module:model/RestToken} Token
    */
    Token = undefined;
    /**
    * @member {String} Trigger
    */
    Trigger = undefined;
    /**
    * @member {Object.<String, String>} TriggerInfo
    */
    TriggerInfo = undefined;
    /**
    * @member {String} RedirectTo
    */
    RedirectTo = undefined;








}


