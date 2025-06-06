import type {
    ErrorResponse,
    FilterKeys,
    HasRequiredKeys,
    HttpMethod,
    MediaType,
    OperationRequestBodyContent,
    PathsWithMethod,
    ResponseObjectMap,
    SuccessResponse,
} from "openapi-typescript-helpers";
import {defHttp} from "/@/utils/http/axios";
import type {RequestOptions as AxiosRequestOptions} from '/#/axios';


export interface DefaultParamsOption {
    params?: {
        query?: Record<string, unknown>
    };
}

export type ParamsOption<T> = T extends {
        parameters: any
    }
    ? HasRequiredKeys<T["parameters"]> extends never
        ? {
            params?: T["parameters"]
        }
        : {
            params: T["parameters"]
        }
    : DefaultParamsOption;
// v7 breaking change: TODO uncomment for openapi-typescript@7 support
// : never;

export type RequestBodyOption<T> = OperationRequestBodyContent<T> extends never
    ? {
        data?: never
    }
    : undefined extends OperationRequestBodyContent<T>
        ? {
            data?: OperationRequestBodyContent<T>
        }
        : {
            data: OperationRequestBodyContent<T>
        };

export type FetchSuccessResponse<T> = FilterKeys<SuccessResponse<ResponseObjectMap<T>>, MediaType>;
export type FetchErrorResponse<T> = FilterKeys<ErrorResponse<ResponseObjectMap<T>>, MediaType>;

export type FetchResponse<T> =
    | {
    data: FilterKeys<SuccessResponse<ResponseObjectMap<T>>, MediaType>;
    error?: never;
    response: Response;
}
    | {
    data?: never;
    error: FilterKeys<ErrorResponse<ResponseObjectMap<T>>, MediaType>;
    response: Response;
};

export type RequestOptions<T> = ParamsOption<T> & RequestBodyOption<T>;

export type FetchOptions<T> = RequestOptions<T> & Omit<RequestInit, "body">;


export default function createClient<Paths extends {}>() {

    async function coreFetch<P extends keyof Paths, M extends HttpMethod>(
        url: P,
        fetchOptions: FetchOptions<M extends keyof Paths[P] ? Paths[P][M] : never>,
        options?: AxiosRequestOptions
    )
        : Promise<FetchSuccessResponse<M extends keyof Paths[P] ? Paths[P][M] : unknown>> {

        const {
            method,
            data: requestBody,
            params = {},
        } = fetchOptions || {};

        // console.log('mode', JSON.stringify(params), JSON.stringify(requestBody), JSON.stringify(options))

        return await defHttp.request(
            {url: url as string, data: requestBody, params: params, method: method},
            options
        )
    }

    type GetPaths = PathsWithMethod<Paths, "get">;
    type PutPaths = PathsWithMethod<Paths, "put">;
    type PostPaths = PathsWithMethod<Paths, "post">;
    type DeletePaths = PathsWithMethod<Paths, "delete">;
    type OptionsPaths = PathsWithMethod<Paths, "options">;
    type HeadPaths = PathsWithMethod<Paths, "head">;
    type PatchPaths = PathsWithMethod<Paths, "patch">;
    type TracePaths = PathsWithMethod<Paths, "trace">;

    type GetFetchOptions<P extends GetPaths> = FetchOptions<
        FilterKeys<Paths[P], "get">
    >;
    type PutFetchOptions<P extends PutPaths> = FetchOptions<
        FilterKeys<Paths[P], "put">
    >;
    type PostFetchOptions<P extends PostPaths> = FetchOptions<
        FilterKeys<Paths[P], "post">
    >;
    type DeleteFetchOptions<P extends DeletePaths> = FetchOptions<
        FilterKeys<Paths[P], "delete">
    >;
    type OptionsFetchOptions<P extends OptionsPaths> = FetchOptions<
        FilterKeys<Paths[P], "options">
    >;
    type HeadFetchOptions<P extends HeadPaths> = FetchOptions<
        FilterKeys<Paths[P], "head">
    >;
    type PatchFetchOptions<P extends PatchPaths> = FetchOptions<
        FilterKeys<Paths[P], "patch">
    >;
    type TraceFetchOptions<P extends TracePaths> = FetchOptions<
        FilterKeys<Paths[P], "trace">
    >;

    return {
        /** Call a GET endpoint */
        async GET<P extends GetPaths>(
            url: P,
            options?: AxiosRequestOptions,
            ...init: HasRequiredKeys<GetFetchOptions<P>> extends never
                ? [GetFetchOptions<P>?]
                : [GetFetchOptions<P>]
        ) {
            return coreFetch<P, "get">(
                url,
                {...init[0], method: "GET"} as any,
                options
            );
        },
        /** Call a PUT endpoint */
        async PUT<P extends PutPaths>(
            url: P,
            options?: AxiosRequestOptions,
            ...init: HasRequiredKeys<PutFetchOptions<P>> extends never
                ? [PutFetchOptions<P>?]
                : [PutFetchOptions<P>]
        ) {
            return coreFetch<P, "put">(
                url,
                {...init[0], method: "PUT"} as any,
                options
            );
        },
        /** Call a POST endpoint */
        async POST<P extends PostPaths>(
            url: P,
            options?: AxiosRequestOptions,
            ...init: HasRequiredKeys<PostFetchOptions<P>> extends never
                ? [PostFetchOptions<P>?]
                : [PostFetchOptions<P>]
        ) {
            return coreFetch<P, "post">(
                url,
                {...init[0], method: "POST"} as any,
                options
            );
        },
        /** Call a DELETE endpoint */
        async DELETE<P extends DeletePaths>(
            url: P,
            options?: AxiosRequestOptions,
            ...init: HasRequiredKeys<DeleteFetchOptions<P>> extends never
                ? [DeleteFetchOptions<P>?]
                : [DeleteFetchOptions<P>]
        ) {
            return coreFetch<P, "delete">(
                url,
                {
                    ...init[0],
                    method: "DELETE",
                } as any,
                options
            );
        },
        /** Call a OPTIONS endpoint */
        async OPTIONS<P extends OptionsPaths>(
            url: P,
            options?: AxiosRequestOptions,
            ...init: HasRequiredKeys<OptionsFetchOptions<P>> extends never
                ? [OptionsFetchOptions<P>?]
                : [OptionsFetchOptions<P>]
        ) {
            return coreFetch<P, "options">(
                url,
                {
                    ...init[0],
                    method: "OPTIONS",
                } as any,
                options
            );
        },
        /** Call a HEAD endpoint */
        async HEAD<P extends HeadPaths>(
            url: P,
            options?: AxiosRequestOptions,
            ...init: HasRequiredKeys<HeadFetchOptions<P>> extends never
                ? [HeadFetchOptions<P>?]
                : [HeadFetchOptions<P>]
        ) {
            return coreFetch<P, "head">(
                url,
                {...init[0], method: "HEAD"} as any,
                options
            );
        },
        /** Call a PATCH endpoint */
        async PATCH<P extends PatchPaths>(
            url: P,
            options?: AxiosRequestOptions,
            ...init: HasRequiredKeys<PatchFetchOptions<P>> extends never
                ? [PatchFetchOptions<P>?]
                : [PatchFetchOptions<P>]
        ) {
            return coreFetch<P, "patch">(
                url,
                {...init[0], method: "PATCH"} as any,
                options
            );
        },
        /** Call a TRACE endpoint */
        async TRACE<P extends TracePaths>(
            url: P,
            options?: AxiosRequestOptions,
            ...init: HasRequiredKeys<TraceFetchOptions<P>> extends never
                ? [TraceFetchOptions<P>?]
                : [TraceFetchOptions<P>]
        ) {
            return coreFetch<P, "trace">(
                url,
                {...init[0], method: "TRACE"} as any,
                options
            );
        },
    };
}
