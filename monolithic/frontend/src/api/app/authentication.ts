import client from "/@/api/app/index";
import {ErrorMessageMode} from "/#/axios";


// 用户登陆
export const Logon = (params: any, mode?: ErrorMessageMode) => {
    return client.POST("/admin/v1/login",
        {errorMessageMode: mode},
        {data: params}
    )
};

// 用户登出
export const Logout = (params: any) => {
    return client.POST("/admin/v1/logout",
        {
            errorMessageMode: 'none',
        },
        {data: params}
    )
};

// 获取用户信息
export const GetMe = () => {
    return client.GET("/admin/v1/me",
        {
            errorMessageMode: 'none',
        }
    )
};
