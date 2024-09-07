import client from "/@/api/app/index";
import {components} from "/#/openapi";

// 获取列表
export const ListUser = async (params: any) => {
    return client.GET("/admin/v1/users",
        {
            errorMessageMode: 'none',
        },
        {params: params}
    );
};

// 获取
export const GetUser = (params: any) => {
    return client.GET("/admin/v1/users/{id}",
        {
            errorMessageMode: 'none',
        },
        {params: params}
    )
};

// 创建
export const CreateUser = (params: components["schemas"]["User"]) => {
    return client.POST("/admin/v1/users",
        {
            errorMessageMode: 'none',
        },
        {data: params}
    )
};

// 更新
export const UpdateUser = (params: components["schemas"]["User"]) => {
    return client.PUT("/admin/v1/users/{id}",
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
export const DeleteUser = (params: any) => {
    return client.DELETE("/admin/v1/users/{id}",
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
