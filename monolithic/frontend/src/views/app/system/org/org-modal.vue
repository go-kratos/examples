<template>
  <BasicModal v-bind="$attrs" @register="registerModal" :title="getTitle" @ok="handleSubmit">
    <BasicForm @register="registerForm"/>
  </BasicModal>
</template>

<script lang="ts" setup>
import {computed, ref, unref} from "vue";

import {BasicModal, useModalInner} from '/@/components/Modal';
import {BasicForm, useForm} from '/@/components/Form/index';

import {ListOrganization, CreateOrganization, UpdateOrganization} from '/@/api/app/organization';

import {formSchema} from './org.data';


const emit = defineEmits(['success', 'register']);

const isUpdate = ref(true);
const rowId = ref('');

const [registerForm, {resetFields, setFieldsValue, updateSchema, validate}] = useForm({
  labelWidth: 100,
  schemas: formSchema,
  showActionButtonGroup: false,
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
  await updateSchema({
    field: 'parentId',
    componentProps: {treeData},
  });
});

const getTitle = computed(() => (!unref(isUpdate) ? '创建部门' : '编辑部门'));

async function handleSubmit() {
  try {
    const values = await validate();
    const _rowId = unref(rowId);
    const _isUpdate = unref(isUpdate);
    setModalProps({confirmLoading: true});
    console.log(!_isUpdate ? '创建部门' : '编辑部门', _rowId, values);

    // API提交更改
    if (_isUpdate) {
      await UpdateOrganization({id: _rowId, org: values});
    } else {
      await CreateOrganization({org: values});
    }

    closeModal();
    emit('success');
  } finally {
    setModalProps({confirmLoading: false});
  }
}
</script>
