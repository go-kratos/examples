<template>
  <div>
    <BasicTable @register="registerTable">
      <template #toolbar>
        <a-button type="primary" @click="handleCreate"> 创建职位</a-button>
      </template>
      <template #bodyCell="{ column, record }">
        <TableAction
            v-if="column.dataIndex === 'action'"
            :actions="[
            {
              label: '编辑',
              onClick: handleEdit.bind(null, record),
            },
            {
              label: '权限分配',
              onClick: handlePermission.bind(null, record),
            },
            {
              label: '绑定用户',
              onClick: handleBind.bind(null, record),
            },
            {
              label: '删除',
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
    <PositionDrawer @register="registerDrawer" @success="handleSuccess"/>
  </div>
</template>

<script lang="ts" setup>
import {BasicTable, useTable, TableAction} from '/@/components/Table';
import {DeletePosition, ListPosition} from '/@/api/app/position';

import {useDrawer} from '/@/components/Drawer';
import PositionDrawer from './position-drawer.vue';

import {columns, searchFormSchema} from './position.data';
import {useMessage} from '/@/hooks/web/useMessage';

const {notification} = useMessage();

const [registerDrawer, {openDrawer}] = useDrawer();
const [registerTable, {reload}] = useTable({
  title: '职位列表',
  api: ListPosition,
  columns,
  formConfig: {
    labelWidth: 120,
    schemas: searchFormSchema,
  },
  useSearchForm: true,
  showTableSetting: true,
  bordered: true,
  showIndexColumn: true,
  actionColumn: {
    width: 280,
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

function handlePermission() {
}

function handleBind() {
}

function handleDelete(record: Recordable) {
  const {id = 0} = record;
  DeletePosition({id}).then(() => {
    notification.success({
      message: '删除成功',
    });
    reload();
  });
}

function handleSuccess() {
  reload();
}
</script>
