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
    EntSpecializedappointEdges,
    EntSpecializedappointEdgesFromJSON,
    EntSpecializedappointEdgesFromJSONTyped,
    EntSpecializedappointEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntSpecializedappoint
 */
export interface EntSpecializedappoint {
    /**
     * Date holds the value of the "Date" field.
     * @type {string}
     * @memberof EntSpecializedappoint
     */
    date?: string;
    /**
     * 
     * @type {EntSpecializedappointEdges}
     * @memberof EntSpecializedappoint
     */
    edges?: EntSpecializedappointEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntSpecializedappoint
     */
    id?: number;
}

export function EntSpecializedappointFromJSON(json: any): EntSpecializedappoint {
    return EntSpecializedappointFromJSONTyped(json, false);
}

export function EntSpecializedappointFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntSpecializedappoint {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'date': !exists(json, 'Date') ? undefined : json['Date'],
        'edges': !exists(json, 'edges') ? undefined : EntSpecializedappointEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntSpecializedappointToJSON(value?: EntSpecializedappoint | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'Date': value.date,
        'edges': EntSpecializedappointEdgesToJSON(value.edges),
        'id': value.id,
    };
}


