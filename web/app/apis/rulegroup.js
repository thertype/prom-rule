import { createApi } from '@ajax'

const prefix = '/api/v1';
/* rulegroup */
export const getStrategy = createApi(`${prefix}/plans`, { method: 'get' }) // get list
export const addStrategy = createApi(`${prefix}/plans`) // add 
export const updateStrategy = createApi(`${prefix}/plans/:id`, { method: 'put' }) // update
export const deleteStrategy = createApi(`${prefix}/plans/:id`, { method: 'delete' }) // update
