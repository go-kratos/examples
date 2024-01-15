import type { ErrorMessageMode } from '/#/axios';
import { useMessage } from '/@/hooks/web/useMessage';
import { useI18n } from '/@/hooks/web/useI18n';
// import router from '/@/router';
// import { PageEnum } from '/@/enums/pageEnum';
import { useUserStoreWithOut } from '/@/store/modules/user';
import projectSetting from '/@/settings/projectSetting';
import { SessionTimeoutProcessingEnum } from '/@/enums/appEnum';

const { createMessage, createErrorModal } = useMessage();
const error = createMessage.error!;
const stp = projectSetting.sessionTimeoutProcessing;

export function checkStatus(
  status: number,
  reason: string,
  msg: string,
  mode: ErrorMessageMode = 'message',
): void {
  const { t } = useI18n();
  const userStore = useUserStoreWithOut();

  let jumpToLogin = false;
  let errMessage = reason == '' ? msg : reason;
  if (errMessage === '') {
    switch (status) {
      case 400:
        errMessage = `${msg}`;
        break;

      case 401:
        // 401: Not logged in
        // 跳转到登陆页面
        jumpToLogin = true;
        errMessage = msg || t('sys.api.UNAUTHORIZED');
        break;

      case 403:
        errMessage = t('sys.api.ACCESS_FORBIDDEN');
        break;

      // 请求不存在
      case 404:
        errMessage = t('sys.api.RESOURCE_NOT_FOUND');
        break;

      case 405:
        errMessage = t('sys.api.METHOD_NOT_ALLOWED');
        break;

      case 408:
        errMessage = t('sys.api.REQUEST_TIMEOUT');
        break;

      case 500:
        errMessage = t('sys.api.INTERNAL_SERVER_ERROR');
        break;

      case 501:
        errMessage = t('sys.api.NOT_IMPLEMENTED');
        break;

      case 502:
        errMessage = t('sys.api.NETWORK_ERROR');
        break;

      case 503:
        errMessage = t('sys.api.SERVICE_UNAVAILABLE');
        break;

      case 504:
        errMessage = t('sys.api.NETWORK_TIMEOUT');
        break;

      case 505:
        errMessage = t('sys.api.REQUEST_NOT_SUPPORT');
        break;
      default:
    }
  }

  if (jumpToLogin) {
    userStore.setToken(undefined);
    if (stp === SessionTimeoutProcessingEnum.PAGE_COVERAGE) {
      userStore.setSessionTimeout(true);
    } else {
      userStore.logout().then();
    }
  }

  // console.log('checkStatus', status, errMessage, mode);

  if (errMessage) {
    if (mode === 'modal') {
      createErrorModal({ title: t('sys.api.errorTip'), content: errMessage });
    } else if (mode === 'message') {
      error(errMessage);
    }
  }
}
