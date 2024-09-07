import {h} from 'vue';
import {Switch} from 'ant-design-vue';
import {BasicColumn, FormSchema} from '/@/components/Table';
import {SwitchStatusEnum} from '/@/enums/httpEnum';
import {useMessage} from '/@/hooks/web/useMessage';
import {UpdateUser} from '/@/api/app/user';
import {User} from "/@/api/app";

export const columns: BasicColumn[] = [
    {
        title: '用户名',
        dataIndex: 'userName',
        width: 180,
    },
    {
        title: '姓名',
        dataIndex: 'realName',
        width: 180,
    },
    {
        title: '邮箱',
        dataIndex: 'email',
        width: 200,
    },
    {
        title: '手机',
        dataIndex: 'phone',
        width: 120,
    },
    {
        title: '状态',
        dataIndex: 'status',
        width: 80,
        customRender: ({record}) => {
            const rd = record as User;

            if (!Reflect.has(rd, 'pendingStatus')) {
                rd['pendingStatus'] = false;
            }
            return h(Switch, {
                checked: rd.status === SwitchStatusEnum.ON,
                loading: rd['pendingStatus'],
                onChange(checked: boolean) {
                    rd['pendingStatus'] = true;
                    const newStatus: any = checked ? SwitchStatusEnum.ON : SwitchStatusEnum.OFF;
                    const {createMessage} = useMessage();
                    UpdateUser({id: rd.id, status: newStatus})
                        .then(async () => {
                            rd.status = newStatus;
                            await createMessage.success(`已成功修改用户状态`);
                        })
                        .catch(async () => {
                            await createMessage.error('修改用户状态失败');
                        })
                        .finally(() => {
                            rd['pendingStatus'] = false;
                        });
                },
            });
        },
    },
    {
        title: '创建时间',
        dataIndex: 'createTime',
        width: 180,
    },
];

export const searchFormSchema: FormSchema[] = [
    {
        field: 'realName',
        label: '姓名',
        component: 'Input',
        colProps: {
            span: 8,
        },
    },
    {
        field: 'phone',
        label: '手机',
        component: 'Input',
        colProps: {
            span: 8,
        },
    },
];

export const createUserFormSchema: FormSchema[] = [
    {
        field: 'userName',
        label: '用户名',
        component: 'Input',
        rules: [
            {
                required: true,
                message: '请输入用户名',
            },
        ],
    },
    {
        field: 'password',
        label: '密码',
        component: 'InputPassword',
        required: true,
        ifShow: true,
    },
    {
        field: 'realName',
        label: '名字',
        component: 'Input',
        required: true,
    },
    {
        label: '邮箱',
        field: 'email',
        component: 'Input',
        required: true,
    },
    {
        label: '手机',
        field: 'phone',
        component: 'Input',
        required: true,
    },
    {
        field: 'orgId',
        label: '所属部门',
        component: 'TreeSelect',
        componentProps: {
            fieldNames: {
                label: 'name',
                key: 'id',
                value: 'id',
            },
            getPopupContainer: () => document.body,
        },
        required: false,
    },
    {
        field: 'positionId',
        label: '所在岗位',
        component: 'TreeSelect',
        componentProps: {
            fieldNames: {
                label: 'name',
                key: 'id',
                value: 'id',
            },
            getPopupContainer: () => document.body,
        },
        required: false,
    },
    {
        label: '备注',
        field: 'remark',
        component: 'InputTextArea',
    },
];
