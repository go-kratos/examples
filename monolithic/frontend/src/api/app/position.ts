import client from "/@/api/app/index";
import {components} from "/#/openapi";

// 获取列表
export const ListPosition = async (params: any) => {
    return client.GET("/admin/v1/positions",
        {
            errorMessageMode: 'none',
        },
        {params: params}
    );
};

// 获取
export const GetPosition = (params: any) => {
    return client.GET("/admin/v1/positions/{id}",
        {
            errorMessageMode: 'none',
        },
        {params: params}
    )
};

// 创建
export const CreatePosition = (params: components["schemas"]["CreatePositionRequest"]) => {
    return client.POST("/admin/v1/positions",
        {
            errorMessageMode: 'none',
        },
        {data: params}
    )
};

// 更新
export const UpdatePosition = (params: components["schemas"]["Position"]) => {
    return client.PUT("/admin/v1/positions/{id}",
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
export const DeletePosition = (params: any) => {
    return client.DELETE("/admin/v1/positions/{id}",
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
