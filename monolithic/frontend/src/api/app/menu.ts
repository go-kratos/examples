import client from "/@/api/app/index";
import {components} from "/#/openapi";

// 获取列表
export const ListMenu = async (params: any) => {
    return client.GET("/admin/v1/menus",
        {
            errorMessageMode: 'none',
        },
        {params: params}
    );
};

// 获取
export const GetMenu = (params: any) => {
    return client.GET("/admin/v1/menus/{id}",
        {
            errorMessageMode: 'none',
        },
        {params: params}
    )
};

// 创建
export const CreateMenu = (params: components["schemas"]["CreateMenuRequest"]) => {
    return client.POST("/admin/v1/menus",
        {
            errorMessageMode: 'none',
        },
        {data: params}
    )
};

// 更新
export const UpdateMenu = (params: components["schemas"]["Menu"]) => {
    return client.PUT("/admin/v1/menus/{id}",
        {
            errorMessageMode: 'none',
        },
        {
            params: {
                path: {
                    id: params.id!,
                }
            },
            data: params
        }
    )
};

// 删除
export const DeleteMenu = (params: any) => {
    return client.DELETE("/admin/v1/menus/{id}",
        {
            errorMessageMode: 'none',
        },
        {
            params: {
                path: {
                    id: params.id,
                }
            },
        }
    )
};
