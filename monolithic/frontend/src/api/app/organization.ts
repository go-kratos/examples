import client from "/@/api/app/index";
import {components} from "/#/openapi";

// 获取列表
export const ListOrganization = async (params: any) => {
    return client.GET("/admin/v1/orgs",
        {
            errorMessageMode: 'none',
        },
        {params: params}
    );
};

// 获取
export const GetOrganization = (params: any) => {
    return client.GET("/admin/v1/orgs/{id}",
        {
            errorMessageMode: 'none',
        },
        {params: params}
    )
};

// 创建
export const CreateOrganization = (params: components["schemas"]["CreateOrganizationRequest"]) => {
    return client.POST("/admin/v1/orgs",
        {
            errorMessageMode: 'none',
        },
        {data: params}
    )
};

// 更新
export const UpdateOrganization = (params: components["schemas"]["Organization"]) => {
    return client.PUT("/admin/v1/orgs/{id}",
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
export const DeleteOrganization = (params: any) => {
    return client.DELETE("/admin/v1/orgs/{id}",
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
