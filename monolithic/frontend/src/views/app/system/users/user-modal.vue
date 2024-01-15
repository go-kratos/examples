<template>
  <BasicModal v-bind="$attrs" @register="registerModal" :title="getTitle" @ok="handleSubmit">
    <BasicForm @register="registerForm"/>
  </BasicModal>
</template>

<script lang="ts" setup>
import {BasicModal, useModalInner} from '/@/components/Modal';
import {BasicForm, useForm} from '/@/components/Form/index';
import {createUserFormSchema} from './users.data';

import {CreateUser, UpdateUser} from '/@/api/app/user';
import {ListOrganization} from '/@/api/app/organization';
import {computed, ref, unref} from "vue";

const emit = defineEmits(['success', 'register']);

const isUpdate = ref(true);
const rowId = ref('');

const [registerForm, {setFieldsValue, updateSchema, resetFields, validate}] = useForm({
  labelWidth: 100,
  schemas: createUserFormSchema,
  showActionButtonGroup: false,
  actionColOptions: {
    span: 23,
  },
});

const [registerModal, {setModalProps, closeModal}] = useModalInner(async (data) => {
  await resetFields();
  setModalProps({confirmLoading: false});
  isUpdate.value = !!data?.isUpdate;

  if (unref(isUpdate)) {
    rowId.value = data.record.id;
    await setFieldsValue({
      ...data.record,
    });
  }

  const orgData = (await ListOrganization({})) || [];
  const treeData = orgData.items || [];
  await updateSchema([
    {
      field: 'password',
      show: !unref(isUpdate),
    },
    {
      field: 'org',
      componentProps: {treeData},
    },
  ]);
});

const getTitle = computed(() => (!unref(isUpdate) ? '创建账号' : '编辑账号'));

async function handleSubmit() {
  try {
    const values = await validate();
    setModalProps({confirmLoading: true});
    const _isUpdate = unref(isUpdate);
    const _rowId = unref(rowId);
    console.log(!_isUpdate ? '创建账号' : '编辑账号', _rowId, values);

    // API提交更改
    if (_isUpdate) {
      await UpdateUser({id: parseInt(_rowId), ...values});
    } else {
      await CreateUser(values);
    }

    closeModal();
    emit('success', {isUpdate: _isUpdate, values: {...values, id: rowId.value}});
  } finally {
    setModalProps({confirmLoading: false});
  }
}
</script>
