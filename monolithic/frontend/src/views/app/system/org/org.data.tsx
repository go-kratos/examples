import {h} from 'vue';
import {Tag} from 'ant-design-vue';
import {BasicColumn} from '/@/components/Table';
import {FormSchema} from '/@/components/Table';
import {isOn, SwitchStatusEnum} from '/@/enums/httpEnum';
import {Organization} from "/@/api/app";

export const columns: BasicColumn[] = [
    {
        title: '部门名称',
        dataIndex: 'name',
        width: 160,
        align: 'left',
    },
    {
        title: '排序',
        dataIndex: 'orderNo',
        width: 50,
    },
    {
        title: '状态',
        dataIndex: 'status',
        width: 80,
        customRender: ({record}) => {
            const {status} = record as Organization;
            const enable = isOn(status);
            const color = enable ? '#108ee9' : '#FF0000';
            const text = enable ? '启用' : '停用';
            return h(Tag, {color: color}, () => text);
        },
    },
    {
        title: '创建时间',
        dataIndex: 'createTime',
        width: 180,
    },
    {
        title: '备注',
        dataIndex: 'remark',
    },
];

export const searchFormSchema: FormSchema[] = [
    {
        field: 'name',
        label: '部门名称',
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
        label: '部门名称',
        component: 'Input',
        required: true,
    },
    {
        field: 'parentId',
        label: '上级部门',
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
        field: 'orderNo',
        label: '排序',
        component: 'InputNumber',
        required: true,
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
        required: true,
    },
    {
        label: '备注',
        field: 'remark',
        component: 'InputTextArea',
    },
];
