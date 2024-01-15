<template>
  <div>
    <BasicTable @register="registerTable">
      <template #toolbar>
        <a-button type="primary" @click="handleCreate"> 创建角色</a-button>
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
    <RoleDrawer @register="registerDrawer" @success="handleSuccess"/>
  </div>
</template>

<script lang="ts" setup>
import {BasicTable, useTable, TableAction} from '/@/components/Table';
import {DeleteRole, ListRole} from '/@/api/app/role';

import {useDrawer} from '/@/components/Drawer';
import RoleDrawer from './role-drawer.vue';

import {columns, searchFormSchema} from './role.data';
import {useMessage} from '/@/hooks/web/useMessage';

const {notification} = useMessage();

const [registerDrawer, {openDrawer}] = useDrawer();
const [registerTable, {reload}] = useTable({
  title: '角色列表',
  api: ListRole,
  columns,
  formConfig: {
    labelWidth: 120,
    schemas: searchFormSchema,
  },
  useSearchForm: true,
  showTableSetting: true,
  bordered: true,
  showIndexColumn: false,
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
  DeleteRole({id}).then(() => {
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
