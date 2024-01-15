import client from "/@/api/app/index";
import {components} from "/#/openapi";

// 获取列表
export const ListRole = async (params: any) => {
    return client.GET("/admin/v1/roles",
        {
            errorMessageMode: 'none',
        },
        {params: params}
    );
};

// 获取
export const GetRole = (params: any) => {
    return client.GET("/admin/v1/roles/{id}",
        {
            errorMessageMode: 'none',
        },
        {params: params}
    )
};

// 创建
export const CreateRole = (params: components["schemas"]["CreateRoleRequest"]) => {
    return client.POST("/admin/v1/roles",
        {
            errorMessageMode: 'none',
        },
        {data: params}
    )
};

// 更新
export const UpdateRole = (params: components["schemas"]["Role"]) => {
    return client.PUT("/admin/v1/roles/{id}",
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
export const DeleteRole = (params: any) => {
    return client.DELETE("/admin/v1/roles/{id}",
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
