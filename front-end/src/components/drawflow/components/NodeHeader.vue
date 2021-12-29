<template>
  <CBadge :color="color">
    {{ title }}
  </CBadge>
</template>
<script>
import { defineComponent } from 'vue'
export default defineComponent({
  name: 'NodeHeader',
  props: ['title', 'color'],
})
</script>
