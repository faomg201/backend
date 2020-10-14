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
    EntMainingreEdges,
    EntMainingreEdgesFromJSON,
    EntMainingreEdgesFromJSONTyped,
    EntMainingreEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntMainingre
 */
export interface EntMainingre {
    /**
     * MAININGREDIENTNAME holds the value of the "MAIN_INGREDIENT_NAME" field.
     * @type {string}
     * @memberof EntMainingre
     */
    mAININGREDIENTNAME?: string;
    /**
     * 
     * @type {EntMainingreEdges}
     * @memberof EntMainingre
     */
    edges?: EntMainingreEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntMainingre
     */
    id?: number;
}

export function EntMainingreFromJSON(json: any): EntMainingre {
    return EntMainingreFromJSONTyped(json, false);
}

export function EntMainingreFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntMainingre {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'mAININGREDIENTNAME': !exists(json, 'MAIN_INGREDIENT_NAME') ? undefined : json['MAIN_INGREDIENT_NAME'],
        'edges': !exists(json, 'edges') ? undefined : EntMainingreEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntMainingreToJSON(value?: EntMainingre | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'MAIN_INGREDIENT_NAME': value.mAININGREDIENTNAME,
        'edges': EntMainingreEdgesToJSON(value.edges),
        'id': value.id,
    };
}


