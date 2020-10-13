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
import {
    EntSpecializedappoint,
    EntSpecializedappointFromJSON,
    EntSpecializedappointFromJSONTyped,
    EntSpecializedappointToJSON,
    EntUser,
    EntUserFromJSON,
    EntUserFromJSONTyped,
    EntUserToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntSpecializeddiagEdges
 */
export interface EntSpecializeddiagEdges {
    /**
     * Appointment holds the value of the appointment edge.
     * @type {Array<EntSpecializedappoint>}
     * @memberof EntSpecializeddiagEdges
     */
    appointment?: Array<EntSpecializedappoint>;
    /**
     * 
     * @type {EntUser}
     * @memberof EntSpecializeddiagEdges
     */
    user?: EntUser;
}

export function EntSpecializeddiagEdgesFromJSON(json: any): EntSpecializeddiagEdges {
    return EntSpecializeddiagEdgesFromJSONTyped(json, false);
}

export function EntSpecializeddiagEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntSpecializeddiagEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'appointment': !exists(json, 'Appointment') ? undefined : ((json['Appointment'] as Array<any>).map(EntSpecializedappointFromJSON)),
        'user': !exists(json, 'User') ? undefined : EntUserFromJSON(json['User']),
    };
}

export function EntSpecializeddiagEdgesToJSON(value?: EntSpecializeddiagEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'appointment': value.appointment === undefined ? undefined : ((value.appointment as Array<any>).map(EntSpecializedappointToJSON)),
        'user': EntUserToJSON(value.user),
    };
}


