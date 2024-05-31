/**
 * @description: Request result set
 */
export enum ResultEnum {
  SUCCESS = 0,
  ERROR = -1,
  TIMEOUT = 401,
  TYPE = 'success',
}

/**
 * @description: request method
 */
export enum RequestEnum {
  GET = 'GET',
  POST = 'POST',
  PUT = 'PUT',
  DELETE = 'DELETE',
}

/**
 * @description:  contentType
 */
export enum ContentTypeEnum {
  // json
  JSON = 'application/json;charset=UTF-8',
  // form-data qs
  FORM_URLENCODED = 'application/x-www-form-urlencoded;charset=UTF-8',
  // form-data  upload
  FORM_DATA = 'multipart/form-data;charset=UTF-8',
}

/**
 * @description:  开关状态
 */
export const SwitchStatusEnum = {
  // 开启
  ON: 'ON',

  // 关闭
  OFF: 'OFF',
};
export const isOn = (status?: string) => status === SwitchStatusEnum.ON;

export const MenuTypeEnum = {
  // 目录
  FOLDER: 'FOLDER',

  // 菜单
  MENU: 'MENU',

  // 按钮
  BUTTON: 'BUTTON',
};
export const isDir = (menuType?: string) => menuType === MenuTypeEnum.FOLDER;
export const isMenu = (menuType?: string) => menuType === MenuTypeEnum.MENU;
export const isButton = (menuType?: string) => menuType === MenuTypeEnum.BUTTON;
