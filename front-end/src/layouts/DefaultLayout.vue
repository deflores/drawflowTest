<template>
  <div>
    <AppSidebar />
    <div class="wrapper d-flex flex-column min-vh-100 bg-light">
      <AppHeader />
      <div class="body flex-grow-1 px-3">
        <router-view />
      </div>
      <AppFooter />
    </div>
  </div>
</template>
<script>
import AppFooter from '@/components/coreui-c/AppFooter.vue'
import AppHeader from '@/components/coreui-c/AppHeader.vue'
import AppSidebar from '@/components/coreui-c/AppSidebar.vue'

export default {
  name: 'DefaultLayout',
  components: {
    AppFooter,
    AppHeader,
    AppSidebar,
  },
  mounted() {
    let recaptchaScript = document.createElement('script')
    recaptchaScript.setAttribute(
      'src',
      'https://unpkg.com/drawflow@0.0.45/dist/drawflow.min.js',
    )
    document.head.appendChild(recaptchaScript)
  },
}
</script>
