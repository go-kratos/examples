import {h} from 'vue';
import {Switch} from 'ant-design-vue';
import {BasicColumn} from '/@/components/Table';
import {FormSchema} from '/@/components/Table';
import {useMessage} from '/@/hooks/web/useMessage';
import {SwitchStatusEnum} from '/@/enums/httpEnum';
import {UpdatePosition} from '/@/api/app/position';
import {Position} from "/@/api/app";

export const columns: BasicColumn[] = [
    {
        title: '排序',
        dataIndex: 'orderNo',
        width: 50,
    },
    {
        title: '职位名称',
        dataIndex: 'name',
        width: 200,
    },
    {
        title: '备注',
        dataIndex: 'remark',
    },
    {
        title: '状态',
        dataIndex: 'status',
        width: 120,
        customRender: ({record}) => {
            const rd = record as Position;

            if (!Reflect.has(rd, 'pendingStatus')) {
                rd['pendingStatus'] = false;
            }
            return h(Switch, {
                checked: rd.status === SwitchStatusEnum.ON,
                checkedChildren: '已启用',
                unCheckedChildren: '已禁用',
                loading: rd['pendingStatus'],
                onChange(checked: boolean) {
                    rd['pendingStatus'] = true;
                    const newStatus: any = checked ? SwitchStatusEnum.ON : SwitchStatusEnum.OFF;
                    const {createMessage} = useMessage();
                    UpdatePosition({id: rd.id, position: {status: newStatus}})
                        .then(async () => {
                            rd.status = newStatus;
                            await createMessage.success(`已成功修改职位状态`);
                        })
                        .catch(async () => {
                            await createMessage.error('修改职位状态失败');
                        })
                        .finally(() => {
                            rd['pendingStatus'] = false;
                        });
                },
            });
        },
    },
];

export const searchFormSchema: FormSchema[] = [
    {
        field: 'name',
        label: '职位名称',
        component: 'Input',
        colProps: {span: 8},
    },
    {
        field: 'status',
        label: '状态',
        component: 'Select',
        componentProps: {
            options: [
                {label: '启用', value: SwitchStatusEnum.ON},
                {label: '停用', value: SwitchStatusEnum.OFF},
            ],
        },
        colProps: {span: 8},
    },
];

export const formSchema: FormSchema[] = [
    {
        field: 'name',
        label: '职位名称',
        required: true,
        component: 'Input',
    },
    {
        field: 'status',
        label: '状态',
        component: 'RadioButtonGroup',
        defaultValue: SwitchStatusEnum.OFF,
        componentProps: {
            options: [
                {label: '启用', value: SwitchStatusEnum.ON},
                {label: '停用', value: SwitchStatusEnum.OFF},
            ],
        },
    },
    {
        label: '备注',
        field: 'remark',
        component: 'InputTextArea',
    },
    {
        label: ' ',
        field: 'menu',
        slot: 'menu',
        component: 'Input',
    },
];
