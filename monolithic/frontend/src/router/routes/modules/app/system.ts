import type { AppRouteModule } from '/@/router/types';

import { LAYOUT } from '/@/router/constant';
import { t } from '/@/hooks/web/useI18n';

const system: AppRouteModule = {
  path: '/system',
  name: 'System',
  component: LAYOUT,
  meta: {
    orderNo: 2000,
    icon: 'ant-design:setting-outline',
    title: t('routes.menu.system.moduleName'),
  },
  children: [
    {
      path: 'account',
      name: 'AccountPage',
      component: () => import('/@/views/app/system/account/index.vue'),
      meta: {
        title: t('routes.menu.system.account'),
        hideMenu: true,
      },
    },

    {
      path: 'users',
      name: 'UserManagement',
      meta: {
        icon: 'ion:person-outline',
        title: t('routes.menu.system.user'),
        ignoreKeepAlive: false,
      },
      component: () => import('/@/views/app/system/users/index.vue'),
    },
    {
      path: 'users/detail/:id',
      name: 'UserDetail',
      meta: {
        hideMenu: true,
        title: t('routes.menu.system.user-detail'),
        ignoreKeepAlive: true,
        currentActiveMenu: '/system/user',
      },
      component: () => import('/@/views/app/system/users/detail/index.vue'),
    },

    {
      path: 'menu',
      name: 'MenuManagement',
      meta: {
        icon: 'ion:menu-outline',
        title: t('routes.menu.system.menu'),
        ignoreKeepAlive: true,
      },
      component: () => import('/@/views/app/system/menu/index.vue'),
    },
    {
      path: 'org',
      name: 'OrganizationManagement',
      meta: {
        icon: 'ant-design:apartment-outlined',
        title: t('routes.menu.system.org'),
        ignoreKeepAlive: true,
      },
      component: () => import('/@/views/app/system/org/index.vue'),
    },

    {
      path: 'role',
      name: 'RoleManagement',
      meta: {
        icon: 'ant-design:team-outlined',
        title: t('routes.menu.system.role'),
        ignoreKeepAlive: true,
        hideMenu: false,
      },
      component: () => import('/@/views/app/system/role/index.vue'),
    },
    {
      path: 'position',
      name: 'PositionManagement',
      meta: {
        icon: 'ion:person-circle-outline',
        title: t('routes.menu.system.position'),
        ignoreKeepAlive: true,
        hideMenu: false,
      },
      component: () => import('/@/views/app/system/position/index.vue'),
    },
  ],
};

export default system;
