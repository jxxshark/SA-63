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
    EntPatientEdges,
    EntPatientEdgesFromJSON,
    EntPatientEdgesFromJSONTyped,
    EntPatientEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntPatient
 */
export interface EntPatient {
    /**
     * Age holds the value of the "age" field.
     * @type {number}
     * @memberof EntPatient
     */
    age?: number;
    /**
     * 
     * @type {EntPatientEdges}
     * @memberof EntPatient
     */
    edges?: EntPatientEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntPatient
     */
    id?: number;
    /**
     * Name holds the value of the "name" field.
     * @type {string}
     * @memberof EntPatient
     */
    name?: string;
}

export function EntPatientFromJSON(json: any): EntPatient {
    return EntPatientFromJSONTyped(json, false);
}

export function EntPatientFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntPatient {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'age': !exists(json, 'age') ? undefined : json['age'],
        'edges': !exists(json, 'edges') ? undefined : EntPatientEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
        'name': !exists(json, 'name') ? undefined : json['name'],
    };
}

export function EntPatientToJSON(value?: EntPatient | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'age': value.age,
        'edges': EntPatientEdgesToJSON(value.edges),
        'id': value.id,
        'name': value.name,
    };
}


