-- 默认的角色
TRUNCATE TABLE kratos_monolithic.public.role;
INSERT INTO kratos_monolithic.public.role(id, parent_id, create_by, order_no, name, code, status, remark, create_time)
VALUES (1, null, 0, 1, '系统管理员', 'SYS_ADMIN', 'ON', '系统管理员拥有对整个系统的管理权限', now()),
       (3, null, 0, 3, '普通用户', 'CUSTOMER_USER', 'ON', '普通用户没有管理权限，只有设备和APP的使用权限', now()),
       (4, null, 0, 4, '游客', 'GUEST_USER', 'ON', '游客只有非常有限的数据读取权限', now());

-- 后台目录
TRUNCATE TABLE public.menu;
INSERT INTO public.menu(id, parent_id, order_no, type, name, title, path, redirect, component, icon, status, current_active_menu, keep_alive, is_ext, show, hide_tab, hide_menu, hide_breadcrumb, create_time)
VALUES (1, null, 1, 'FOLDER', 'Dashboard', 'routes.menu.dashboard.index', '/dashboard', '/dashboard/location', 'LAYOUT', 'ant-design:dashboard-outlined', 'ON', null, 'f', 'f', 't', 'f', 'f', 'f', now()),
       (2, 1, 2, 'MENU', 'AnalysisViewer', 'routes.menu.dashboard.analysis', 'analysis', null, 'app/dashboard/analysis/index', 'ant-design:bar-chart-outlined', 'ON', null, 'f', 'f', 't', 'f', 'f', 'f', now()),

       (4, null, 4, 'FOLDER', 'AccountManage', 'routes.menu.account.account', '/account', null, 'LAYOUT', 'mdi:account-box', 'ON', null, 'f', 'f', 't', 'f', 't', 'f', now()),
       (5, 4, 5, 'MENU', 'AccountSettings', 'routes.menu.dashboard.analysis', 'settings', null, 'app/account/settings/index', 'ant-design:bar-chart-outlined', 'ON', null, 'f', 'f', 't', 'f', 'f', 'f', now()),

       (6, null, 6, 'FOLDER', 'System', 'routes.menu.system.system', '/system', null, 'LAYOUT', 'ion:settings-outline', 'ON', null, 'f', 'f', 't', 'f', 'f', 'f', now()),
       (7, 6, 7, 'MENU', 'MenuManagement', 'routes.menu.system.menu', 'menu', null, 'app/system/menu/index', 'ion:menu-outline', 'ON', null, 'f', 'f', 't', 'f', 'f', 'f', now()),
       (8, 6, 8, 'MENU', 'RoleManagement', 'routes.menu.system.role', 'role', null, 'app/system/role/index', 'ant-design:team-outlined', 'ON', null, 'f', 'f', 't', 'f', 'f', 'f', now()),
       (11, 6, 11, 'MENU', 'OrganizationManagement', 'routes.menu.system.org', 'org', null, 'app/system/org/index', 'ant-design:apartment-outlined', 'ON', null, 'f', 'f', 't', 'f', 'f', 'f', now()),
       (12, 6, 12, 'MENU', 'PositionManagement', 'routes.menu.system.position', 'position', null, 'app/system/position/index', 'ion:person-circle-outline', 'ON', null, 'f', 'f', 't', 'f', 'f', 'f', now()),
       (13, 6, 13, 'MENU', 'UserManagement', 'routes.menu.system.user', 'users', null, 'app/system/users/index', 'ion:person-outline', 'ON', null, 'f', 'f', 't', 'f', 'f', 'f', now()),
       (14, 6, 14, 'MENU', 'UserDetail', 'routes.menu.system.user-detail', 'users/detail/:id', null, 'app/system/users/detail/index', null, 'ON', '/system/user', 't', 'f', 't', 'f', 't', 'f', now());
