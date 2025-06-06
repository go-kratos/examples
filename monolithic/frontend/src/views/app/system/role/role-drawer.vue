<template>
  <BasicDrawer
      v-bind="$attrs"
      @register="registerDrawer"
      showFooter
      :title="getTitle"
      width="500px"
      @ok="handleSubmit"
  >
    <BasicForm @register="registerForm">
      <template #menu="{ model, field }">
        <BasicTree
            v-model:value="model[field]"
            :treeData="treeData"
            :fieldNames="{ title: 'name', key: 'id' }"
            checkable
            toolbar
            title="菜单分配"
        />
      </template>
    </BasicForm>
  </BasicDrawer>
</template>

<script lang="ts" setup>
import {computed, ref, unref} from "vue";
import {BasicForm, useForm} from '/@/components/Form/index';
import {BasicDrawer, useDrawerInner} from '/@/components/Drawer';
import {BasicTree, TreeItem} from '/@/components/Tree';

import {CreateRole, UpdateRole} from '/@/api/app/role';
import {ListMenu} from '/@/api/app/menu';

import {formSchema} from './role.data';


const emit = defineEmits(['success', 'register']);

const isUpdate = ref(true);
const treeData = ref<TreeItem[]>([]);
const rowId = ref('');

const [registerForm, {resetFields, setFieldsValue, validate}] = useForm({
  labelWidth: 90,
  schemas: formSchema,
  showActionButtonGroup: false,
});

const [registerDrawer, {setDrawerProps, closeDrawer}] = useDrawerInner(async (data) => {
  await resetFields();
  setDrawerProps({confirmLoading: false});
  // 需要在setFieldsValue之前先填充treeData，否则Tree组件可能会报key not exist警告
  if (unref(treeData).length === 0) {
    const menuData = await ListMenu({});
    treeData.value = menuData.items as any as TreeItem[];
  }
  isUpdate.value = !!data?.isUpdate;

  if (unref(isUpdate)) {
    rowId.value = data.record.id;
    await setFieldsValue({
      ...data.record,
    });
  }
});

const getTitle = computed(() => (!unref(isUpdate) ? '创建角色' : '编辑角色'));

async function handleSubmit() {
  try {
    const values = await validate();
    const _rowId = unref(rowId);
    const _isUpdate = unref(isUpdate);
    setDrawerProps({confirmLoading: true});
    console.log(!_isUpdate ? '创建角色' : '编辑角色', _rowId, values);

    // API提交更改
    if (_isUpdate) {
      await UpdateRole({id: _rowId, role: values});
    } else {
      await CreateRole({role: values});
    }

    closeDrawer();
    emit('success');
  } finally {
    setDrawerProps({confirmLoading: false});
  }
}
</script>
