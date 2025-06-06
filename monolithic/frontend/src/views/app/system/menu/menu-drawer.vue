<template>
  <BasicDrawer
      v-bind="$attrs"
      @register="registerDrawer"
      showFooter
      :title="getTitle"
      width="50%"
      @ok="handleSubmit"
  >
    <BasicForm @register="registerForm"/>
  </BasicDrawer>
</template>

<script lang="ts" setup>
import {BasicForm, useForm} from '/@/components/Form/index';
import {formSchema} from './menu.data';
import {BasicDrawer, useDrawerInner} from '/@/components/Drawer';

import {CreateMenu, ListMenu, UpdateMenu} from '/@/api/app/menu';
import {computed, ref, unref} from "vue";

const emit = defineEmits(['success', 'register']);
const isUpdate = ref(true);
const menuId = ref('');

const [registerForm, {resetFields, setFieldsValue, updateSchema, validate}] = useForm({
  labelWidth: 100,
  schemas: formSchema,
  showActionButtonGroup: false,
  baseColProps: {lg: 12, md: 24},
});

const [registerDrawer, {setDrawerProps, closeDrawer}] = useDrawerInner(async (data) => {
  await resetFields();
  setDrawerProps({confirmLoading: false});
  isUpdate.value = !!data?.isUpdate;

  if (unref(isUpdate)) {
    menuId.value = data.record.id;
    await setFieldsValue({
      ...data.record,
    });
  }
  const menuData = await ListMenu({});
  const treeData = menuData.items;
  await updateSchema({
    field: 'parentId',
    componentProps: {treeData},
  });
});

const getTitle = computed(() => (!unref(isUpdate) ? '创建菜单' : '编辑菜单'));

async function handleSubmit() {
  try {
    const values = await validate();
    const _isUpdate = unref(isUpdate);
    const _menuId = unref(menuId);

    setDrawerProps({confirmLoading: true});
    console.log(!_isUpdate ? '创建菜单' : '编辑菜单', _menuId, values);

    // API提交更改
    if (_isUpdate) {
      await UpdateMenu({id: Number(_menuId), ...values});
    } else {
      await CreateMenu({menu: values});
    }

    closeDrawer();
    emit('success', values);
  } finally {
    setDrawerProps({confirmLoading: false});
  }
}
</script>
