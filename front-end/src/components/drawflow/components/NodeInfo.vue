<template>
  <div>
    <div v-for="input in inputs" :key="input.node">
      <CBadge :size="'sm'" :color="'info'">nodo_{{ input.node }}</CBadge>
    </div>
  </div>
</template>
<script>
import { defineComponent } from 'vue'
export default defineComponent({
  name: 'NodeInfo',
  props: ['inputs'],
})
</script>
<style scoped>
.cnode-footer {
  position: absolute;
  margin-top: -25px;
}
</style>
