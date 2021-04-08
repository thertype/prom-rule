import { createApi } from '@ajax'

const prefix = '/api/v1';
/* rulegroup */
export const getRuleGroup = createApi(`${prefix}/regroup`, { method: 'get' }) // get list
export const addRuleGroup = createApi(`${prefix}/regroup`) // add
export const updateRuleGroup = createApi(`${prefix}/regroup/:id`, { method: 'put' }) // update
export const deleteRuleGroup = createApi(`${prefix}/regroup/:id`, { method: 'delete' }) // update

/* reunion */

export const getRuleUnion = createApi(`${prefix}/regroup/:id/reunion`, { method: 'get' }) // get list
export const addRuleUnion = createApi(`${prefix}/regroup/:id/reunion`) // add
export const updateRuleUnion = createApi(`${prefix}/reunion/:id`, { method: 'put' }) // update
export const deleteRuleUnion = createApi(`${prefix}/reunion/:id`, { method: 'delete' }) // update

