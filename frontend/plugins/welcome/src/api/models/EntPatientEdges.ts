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
} from './';

/**
 * 
 * @export
 * @interface EntPatientEdges
 */
export interface EntPatientEdges {
    /**
     * Appointment holds the value of the appointment edge.
     * @type {Array<EntSpecializedappoint>}
     * @memberof EntPatientEdges
     */
    appointment?: Array<EntSpecializedappoint>;
}

export function EntPatientEdgesFromJSON(json: any): EntPatientEdges {
    return EntPatientEdgesFromJSONTyped(json, false);
}

export function EntPatientEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntPatientEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'appointment': !exists(json, 'Appointment') ? undefined : ((json['appointment'] as Array<any>).map(EntSpecializedappointFromJSON)),
    };
}

export function EntPatientEdgesToJSON(value?: EntPatientEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'appointment': value.appointment === undefined ? undefined : ((value.appointment as Array<any>).map(EntSpecializedappointToJSON)),
    };
}


