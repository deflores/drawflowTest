<template>
  <CBadge :color="'dark'" class="cnode-footer" :shape="'rounded-pill'">
    <div class="">nodo_{{ nodeId }}</div>
  </CBadge>
</template>
<script>
import { defineComponent } from 'vue'
export default defineComponent({
  name: 'NodeFooter',
  computed: {
    nodeId: (c) => {
      return c.$parent.nodeId
    },
  },
})
</script>
<style scoped>
.cnode-footer {
  position: absolute;
  margin-top: -25px;
}
</style>
