<template>
  <PageWrapper :title="`用户` + userName + `的资料`" content="" contentBackground @back="goBack">
    <template #extra>
      <a-button type="primary" danger> 禁用账号 </a-button>
      <a-button type="primary"> 修改密码 </a-button>
    </template>
    <template #footer>
      <a-tabs default-active-key="detail" v-model:activeKey="currentKey">
        <a-tab-pane key="detail" tab="用户资料" />
        <a-tab-pane key="logs" tab="操作日志" />
      </a-tabs>
    </template>
    <div class="pt-4 m-4 desc-wrap">
      <template v-if="currentKey === 'detail'">
        <div v-if="userInfo !== undefined">
          用户名{{ userInfo.userName }} 姓名{{ userInfo.realName }} 邮箱{{ userInfo.email }}
        </div>
      </template>
      <template v-if="currentKey === 'logs'">
        <div v-if="userInfo !== undefined">
          <div v-for="i in 10" :key="i">这是用户{{ userInfo.realName }}操作日志Tab</div>
        </div>
      </template>
    </div>
  </PageWrapper>
</template>

<script lang="ts" setup>
  import { useRoute } from 'vue-router';
  import { PageWrapper } from '/@/components/Page';
  import { useGo } from '/@/hooks/web/usePage';
  import { useTabs } from '/@/hooks/web/useTabs';
  import { Tabs } from 'ant-design-vue';
  import { GetUserByUsername } from '/@/api/app/user';
  import { onBeforeMount } from 'vue';
  import { User } from '/&/user';

  const ATabs = Tabs;
  const ATabPane = Tabs.TabPane;

  const route = useRoute();
  const go = useGo();

  // 此处可以得到用户ID
  const userName = ref(route.params?.id);
  const currentKey = ref('detail');
  const { setTitle } = useTabs();
  const userInfo = ref<User>();

  onBeforeMount(async () => {
    await nextTick();
    // 获取用户信息
    userInfo.value = await GetUserByUsername({ username: userName.value as string });
    await setTitle('详情：用户' + userInfo.value.userName);
    console.log(userInfo);
  });

  // 页面左侧点击返回链接时的操作
  function goBack() {
    go('/system/users');
  }
</script>

<style></style>
