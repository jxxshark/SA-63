/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface ControllersSpecializedappoint
 */
export interface ControllersSpecializedappoint {
    /**
     * 
     * @type {string}
     * @memberof ControllersSpecializedappoint
     */
    date?: string;
    /**
     * 
     * @type {number}
     * @memberof ControllersSpecializedappoint
     */
    patientID?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersSpecializedappoint
     */
    specializeddiagID?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersSpecializedappoint
     */
    userID?: number;
}

export function ControllersSpecializedappointFromJSON(json: any): ControllersSpecializedappoint {
    return ControllersSpecializedappointFromJSONTyped(json, false);
}

export function ControllersSpecializedappointFromJSONTyped(json: any, ignoreDiscriminator: boolean): ControllersSpecializedappoint {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'date': !exists(json, 'date') ? undefined : json['date'],
        'patientID': !exists(json, 'patientID') ? undefined : json['patientID'],
        'specializeddiagID': !exists(json, 'specializeddiagID') ? undefined : json['specializeddiagID'],
        'userID': !exists(json, 'userID') ? undefined : json['userID'],
    };
}

export function ControllersSpecializedappointToJSON(value?: ControllersSpecializedappoint | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'date': value.date,
        'patientID': value.patientID,
        'specializeddiagID': value.specializeddiagID,
        'userID': value.userID,
    };
}


