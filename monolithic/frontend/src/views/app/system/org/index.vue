<template>
  <div>
    <BasicTable @register="registerTable">
      <template #toolbar>
        <a-button type="primary" @click="handleCreate"> 创建部门</a-button>
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
    <OrgModal @register="registerModal" @success="handleSuccess"/>
  </div>
</template>

<script lang="ts" setup>
import {BasicTable, useTable, TableAction} from '/@/components/Table';
import {useModal} from '/@/components/Modal';
import {useMessage} from '/@/hooks/web/useMessage';

import {DeleteOrganization, ListOrganization} from '/@/api/app/organization';

import OrgModal from './org-modal.vue';
import {columns, searchFormSchema} from './org.data';

const {notification} = useMessage();

const [registerModal, {openModal}] = useModal();
const [registerTable, {reload}] = useTable({
  title: '部门列表',
  api: ListOrganization,
  columns,
  formConfig: {
    labelWidth: 120,
    schemas: searchFormSchema,
  },
  pagination: false,
  striped: false,
  useSearchForm: true,
  showTableSetting: true,
  bordered: true,
  showIndexColumn: false,
  canResize: true,
  actionColumn: {
    width: 80,
    title: '操作',
    dataIndex: 'action',
    fixed: undefined,
  },
});

function handleCreate() {
  openModal(true, {
    isUpdate: false,
  });
}

function handleEdit(record: Recordable) {
  openModal(true, {
    record,
    isUpdate: true,
  });
}

function handleDelete(record: Recordable) {
  const {id = 0} = record;
  DeleteOrganization({id}).then(() => {
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
