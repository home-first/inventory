import { createApp } from 'vue'
import App from './App.vue'
import { router } from './router'
import { store, key } from './store'
import PrimeVue from 'primevue/config';
import 'primevue/resources/primevue.min.css';
import 'primevue/resources/themes/saga-blue/theme.css';
import 'primeicons/primeicons.css';



createApp(App)
    .use(router)
    .use(store, key)
    .use(PrimeVue)
    .mount('#app')
