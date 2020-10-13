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
    EntSpecializeddiag,
    EntSpecializeddiagFromJSON,
    EntSpecializeddiagFromJSONTyped,
    EntSpecializeddiagToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntUserEdges
 */
export interface EntUserEdges {
    /**
     * Appointment holds the value of the appointment edge.
     * @type {Array<EntSpecializedappoint>}
     * @memberof EntUserEdges
     */
    appointment?: Array<EntSpecializedappoint>;
    /**
     * Specializeddiag holds the value of the specializeddiag edge.
     * @type {Array<EntSpecializeddiag>}
     * @memberof EntUserEdges
     */
    specializeddiag?: Array<EntSpecializeddiag>;
}

export function EntUserEdgesFromJSON(json: any): EntUserEdges {
    return EntUserEdgesFromJSONTyped(json, false);
}

export function EntUserEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntUserEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'appointment': !exists(json, 'Appointment') ? undefined : ((json['Appointment'] as Array<any>).map(EntSpecializedappointFromJSON)),
        'specializeddiag': !exists(json, 'Specializeddiag') ? undefined : ((json['Specializeddiag'] as Array<any>).map(EntSpecializeddiagFromJSON)),
    };
}

export function EntUserEdgesToJSON(value?: EntUserEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'appointment': value.appointment === undefined ? undefined : ((value.appointment as Array<any>).map(EntSpecializedappointToJSON)),
        'specializeddiag': value.specializeddiag === undefined ? undefined : ((value.specializeddiag as Array<any>).map(EntSpecializeddiagToJSON)),
    };
}


