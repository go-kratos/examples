import {h} from 'vue';
import {Tag} from 'ant-design-vue';
import {BasicColumn} from '/@/components/Table';
import {FormSchema} from '/@/components/Table';
import {Icon} from '/@/components/Icon';
import {SwitchStatusEnum, isOn, isDir, isMenu, isButton, MenuTypeEnum} from '/@/enums/httpEnum';
import {useI18n} from '/@/hooks/web/useI18n';
import {Menu} from "/@/api/app";

const {t} = useI18n();

export const columns: BasicColumn[] = [
    {
        title: '菜单名称',
        dataIndex: 'title',
        width: 200,
        align: 'left',
        customRender: ({record, text}) => {
            const {icon} = record as Menu;
            const title = t(text);
            return icon === '' || icon === undefined
                ? h('div', null, h('span', {style: {marginRight: '15px'}}, title))
                : h('div', null, [
                    h(Icon, {icon: icon}),
                    h('span', {style: {marginRight: '15px'}}, title),
                ]);
        },
    },
    {
        title: '排序',
        dataIndex: 'orderNo',
        width: 50,
    },
    {
        title: '权限标识',
        dataIndex: 'permissionCode',
        width: 180,
    },
    {
        title: '路由地址',
        dataIndex: 'path',
    },
    {
        title: '组件路径',
        dataIndex: 'component',
    },

    {
        title: '状态',
        dataIndex: 'status',
        width: 80,
        customRender: ({record}) => {
            const {status} = record as Menu;
            const enable = isOn(status);
            const color = enable ? '#108ee9' : '#FF0000';
            const text = enable ? '启用' : '停用';
            return h(Tag, {color: color}, () => text);
        },
    },
    {
        title: '更新时间',
        dataIndex: 'updateTime',
        width: 180,
    },
];

export const searchFormSchema: FormSchema[] = [
    {
        field: 'name',
        label: '菜单名称',
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
        field: 'type',
        label: '菜单类型',
        component: 'RadioButtonGroup',
        defaultValue: MenuTypeEnum.FOLDER,
        componentProps: {
            options: [
                {label: '目录', value: MenuTypeEnum.FOLDER},
                {label: '菜单', value: MenuTypeEnum.MENU},
                {label: '按钮', value: MenuTypeEnum.BUTTON},
            ],
        },
        colProps: {lg: 24, md: 24},
    },
    {
        field: 'name',
        label: '菜单名称',
        component: 'Input',
        required: true,
    },

    {
        field: 'parentId',
        label: '上级菜单',
        component: 'TreeSelect',
        componentProps: {
            fieldNames: {
                label: 'name',
                key: 'id',
                value: 'id',
            },
            getPopupContainer: () => document.body,
        },
    },

    {
        field: 'orderNo',
        label: '排序',
        component: 'InputNumber',
        required: true,
    },
    {
        field: 'icon',
        label: '图标',
        component: 'IconPicker',
        required: true,
        ifShow: ({values}) => !isButton(values?.type),
    },

    {
        field: 'path',
        label: '路由地址',
        component: 'Input',
        required: true,
        ifShow: ({values}) => !isButton(values?.type),
    },
    {
        field: 'component',
        label: '组件路径',
        component: 'Input',
        ifShow: ({values}) => isMenu(values?.type),
    },
    {
        field: 'permissionCode',
        label: '权限标识',
        component: 'Input',
        ifShow: ({values}) => !isDir(values?.type),
    },
    {
        field: 'status',
        label: '状态',
        component: 'RadioButtonGroup',
        defaultValue: SwitchStatusEnum.OFF,
        componentProps: {
            options: [
                {label: '启用', value: SwitchStatusEnum.ON},
                {label: '禁用', value: SwitchStatusEnum.OFF},
            ],
        },
    },
    {
        field: 'isExt',
        label: '是否外链',
        component: 'RadioButtonGroup',
        defaultValue: false,
        componentProps: {
            options: [
                {label: '是', value: true},
                {label: '否', value: false},
            ],
        },
        ifShow: ({values}) => !isButton(values?.type),
    },

    {
        field: 'keepAlive',
        label: '是否缓存',
        component: 'RadioButtonGroup',
        defaultValue: false,
        componentProps: {
            options: [
                {label: '是', value: true},
                {label: '否', value: false},
            ],
        },
        ifShow: ({values}) => isMenu(values?.type),
    },

    {
        field: 'show',
        label: '是否显示',
        component: 'RadioButtonGroup',
        defaultValue: false,
        componentProps: {
            options: [
                {label: '是', value: true},
                {label: '否', value: false},
            ],
        },
        ifShow: ({values}) => !isButton(values?.type),
    },
];
