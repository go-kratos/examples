import createClient from "/@/utils/openapi-axios";
import type {components, paths} from "/#/openapi";

const client = createClient<paths>();
export default client;

export type User = components['schemas']['User'];
export type Menu = components['schemas']['Menu'];
export type Role = components['schemas']['Role'];
export type Position = components['schemas']['Position'];
export type Organization = components['schemas']['Organization'];
export type Dict = components['schemas']['Dict'];
export type DictDetail = components['schemas']['DictDetail'];
export type KratosStatus = components['schemas']['KratosStatus'];
