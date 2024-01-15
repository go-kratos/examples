<template>
  <PageWrapper dense contentFullHeight contentClass="flex">
    <OrgTree class="w-1/4 xl:w-1/5" @select="handleSelect"/>
    <BasicTable @register="registerTable" class="w-3/4 xl:w-4/5" :searchInfo="searchInfo">
      <template #toolbar>
        <a-button type="primary" @click="handleCreate">创建账号</a-button>
      </template>
      <template #bodyCell="{ column, record }">
        <TableAction
            v-if="column.dataIndex === 'action'"
            :actions="[
            {
              label: '编辑',
              tooltip: '编辑用户资料',
              onClick: handleEdit.bind(null, record),
            },
            {
              label: '删除',
              color: 'error',
              tooltip: '删除此账号',
              popConfirm: {
                title: '是否确认删除',
                confirm: handleDelete.bind(null, record),
              },
            },
          ]"
        />
      </template>
    </BasicTable>
    <UserModal @register="registerModal" @success="handleSuccess"/>
  </PageWrapper>
</template>

<script lang="ts" setup>
import {reactive} from "vue";
import {BasicTable, useTable, TableAction} from '/@/components/Table';
import {PageWrapper} from '/@/components/Page';
import {useModal} from '/@/components/Modal';
import {useGo} from '/@/hooks/web/usePage';
import {useMessage} from '/@/hooks/web/useMessage';

import UserModal from './user-modal.vue';
import OrgTree from './org-tree.vue';

import {columns, searchFormSchema} from './users.data';
import {DeleteUser, ListUser} from '/@/api/app/user';

const {notification} = useMessage();
const go = useGo();
const searchInfo = reactive<Recordable>({});

const [registerModal, {openModal}] = useModal();
const [registerTable, {reload, updateTableDataRecord}] = useTable({
  title: '账号列表',
  api: ListUser,
  columns,
  formConfig: {
    labelWidth: 120,
    schemas: searchFormSchema,
    autoSubmitOnEnter: true,
  },
  useSearchForm: true,
  showTableSetting: true,
  showIndexColumn: false,
  bordered: true,
  canResize: true,
  handleSearchInfoFn(info) {
    console.log('handleSearchInfoFn', info);
  },
  actionColumn: {
    width: 120,
    title: '操作',
    dataIndex: 'action',
  },
});

function handleCreate() {
  openModal(true, {
    isUpdate: false,
  });
}

function handleEdit(record: Recordable) {
  console.log(record);
  openModal(true, {
    record,
    isUpdate: true,
  });
}

function handleDelete(record: Recordable) {
  const {id = 0} = record;
  DeleteUser({id}).then(() => {
    notification.success({
      message: '删除成功',
    });
    reload();
  });
}

function handleSuccess({isUpdate, values}) {
  if (isUpdate) {
    // 注意：updateTableDataRecord要求表格的rowKey属性为string并且存在于每一行的record的keys中
    const result = updateTableDataRecord(values.id, values);
    console.log(result);
  } else {
    reload();
  }
}

function handleSelect(id = '') {
  searchInfo.Id = id;
  reload();
}

function handleView(record: Recordable) {
  console.log(record);
  go('/system/users/detail/' + record.userName);
}
</script>
