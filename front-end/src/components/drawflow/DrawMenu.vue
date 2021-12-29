<template>
  <CListGroup style="overflow-x: auto" class="drawflow-menu drawflow-menu-left">
    <CListGroupItem
      v-for="operation in operations"
      :key="operation.key"
      :class="'bg-' + operation.data.color"
      @click="clicked(operation, $event)"
    >
      {{ operation.key }}<br />
      <CBadge color="dark" size="sm" text-color="high-emphasis-inverse">{{
        operation.shortcut
      }}</CBadge>
    </CListGroupItem>
  </CListGroup>
</template>
<script>
import operations from './operations'
export default {
  name: 'DrawMenu',
  emits: ['addNode'],
  setup() {
    return { operations }
  },
  methods: {
    clicked(node) {
      this.$emit('addNode', node)
    },
  },
}
</script>
