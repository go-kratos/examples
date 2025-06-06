<template>
  <div>
    <BasicTable @register="registerTable" @fetch-success="onFetchSuccess">
      <template #toolbar>
        <Button type="primary" @click="handleCreate"> 创建菜单</Button>
      </template>
      <template #bodyCell="{ column, record }">
        <TableAction
            v-if="column.dataIndex === 'action'"
            :actions="[
            {
              icon: 'clarity:note-edit-line',
              onClick: handleEdit.bind(null, record),
            },
            {
              icon: 'ant-design:delete-outlined',
              color: 'error',
              popConfirm: {
                title: '是否确认删除',
                confirm: handleDelete.bind(null, record),
              },
            },
          ]"
        />
      </template>
    </BasicTable>
    <MenuDrawer @register="registerDrawer" @success="handleSuccess"/>
  </div>
</template>

<script lang="ts" setup>
import {Button} from 'ant-design-vue';
import {BasicTable, useTable, TableAction} from '/@/components/Table';
import {useDrawer} from '/@/components/Drawer';
import MenuDrawer from './menu-drawer.vue';

import {ListMenu, DeleteMenu} from '/@/api/app/menu';

import {columns, searchFormSchema} from './menu.data';
import {useMessage} from '/@/hooks/web/useMessage';
import {nextTick} from "vue";

const {notification} = useMessage();
const enableExpandAll = false;

const [registerDrawer, {openDrawer}] = useDrawer();
const [registerTable, {reload, expandAll}] = useTable({
  title: '菜单列表',
  api: ListMenu,
  columns,
  formConfig: {
    labelWidth: 120,
    schemas: searchFormSchema,
  },
  isTreeTable: true,
  pagination: false,
  striped: true,
  useSearchForm: true,
  showTableSetting: true,
  bordered: true,
  showIndexColumn: false,
  canResize: true,
  defaultExpandAllRows: false,
  actionColumn: {
    width: 80,
    title: '操作',
    dataIndex: 'action',
    fixed: undefined,
  },
});

function handleCreate() {
  openDrawer(true, {
    isUpdate: false,
  });
}

function handleEdit(record: Recordable) {
  openDrawer(true, {
    record,
    isUpdate: true,
  });
}

function handleDelete(record: Recordable) {
  const {id = 0} = record;
  DeleteMenu({id}).then(() => {
    notification.success({
      message: '删除成功',
    });
    reload();
  });
}

function handleSuccess() {
  // console.log('handleSuccess', values);
  reload();
}

function onFetchSuccess() {
  // 演示默认展开所有表项
  if (enableExpandAll) {
    nextTick(expandAll);
  }
}
</script>
