import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'

import CoreuiVue from '@coreui/vue'
import CIcon from '@coreui/icons-vue'
import { iconsSet as icons } from '@/assets/icons'
import DocsCallout from '@/components/coreui-c/DocsCallout'
import DocsExample from '@/components/coreui-c/DocsExample'
import CustomDrawPlugin from './components/drawflow/componentsPlugin'
import axios from 'axios'
import VueAxios from 'vue-axios'

const app = createApp(App)
app.use(store)
app.use(router)
app.use(CoreuiVue)
app.use(CustomDrawPlugin)
app.use(VueAxios, axios)
app.provide('icons', icons)
app.component('CIcon', CIcon)
app.component('DocsCallout', DocsCallout)
app.component('DocsExample', DocsExample)

app.mount('#app')
