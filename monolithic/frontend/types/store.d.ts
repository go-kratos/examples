import {ErrorTypeEnum} from '/@/enums/exceptionEnum';
import {MenuModeEnum, MenuTypeEnum} from '/@/enums/menuEnum';

// Lock screen information
export interface LockInfo {
    // Password required
    pwd?: string | undefined;
    // Is it locked?
    isLock?: boolean;
}

// Error-log information
export interface ErrorLogInfo {
    // Type of error
    type: ErrorTypeEnum;
    // Error file
    file: string;
    // Error name
    name?: string;
    // Error message
    message: string;
    // Error stack
    stack?: string;
    // Error detail
    detail: string;
    // Error url
    url: string;
    // Error time
    time?: string;
}

// 用户信息
export interface UserInfo {
    // 用户id
    id: string | number;
    // 用户名
    userName: string;
    // 昵称
    nickName: string;
    // 真实名字
    realName: string;
    // 头像
    avatar: string;
    // 电话号码
    phone: string;
    // 家庭住址
    address: string;
    // 电子邮箱
    email: string;
    // 个人描述
    description?: string;

    // 主页
    homePath?: string;
    // 角色信息
    roles: RoleInfo[];
}

export interface BeforeMiniState {
    // 收缩
    menuCollapsed?: boolean;
    // 分割
    menuSplit?: boolean;
    // 模式
    menuMode?: MenuModeEnum;
    // 类型
    menuType?: MenuTypeEnum;
}
