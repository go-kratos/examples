import { createApp } from 'vue';
import App from './App.vue';

import PrimeVue from 'primevue/config';
import ToastService from 'primevue/toastservice';

import 'primevue/resources/themes/saga-blue/theme.css';
import 'primevue/resources/primevue.min.css';
import 'primeicons/primeicons.css';
import 'primeflex/primeflex.css';

import Dropdown from 'primevue/dropdown';
import Tag from 'primevue/tag';
import Panel from 'primevue/panel';
import Toast from 'primevue/toast';


const app = createApp(App);

app.use(PrimeVue);
app.use(ToastService);

app.component('Dropdown', Dropdown);
app.component('Tag', Tag);
app.component('Panel', Panel);
app.component('Toast', Toast);

app.mount('#app');
