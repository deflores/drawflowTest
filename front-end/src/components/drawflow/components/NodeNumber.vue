<template>
  <div ref="el">
    <node-footer></node-footer>
    <node-header :title="'Número'" :color="color"></node-header>
    <CInputGroup>
      <CInputGroupText id="basic-addon1">#</CInputGroupText>
      <CFormInput
        placeholder="valor"
        aria-label="valor"
        aria-describedby="basic-addon1"
        v-model="value"
        @input="updateInput"
      />
    </CInputGroup>
  </div>
</template>

<script>
import { ref, onMounted, getCurrentInstance, nextTick, triggerRef } from 'vue'
import NodeHeader from './NodeHeader.vue'
export default {
  components: { NodeHeader },
  name: 'NodeNumber',
  props: ['color'],
  setup() {
    const el = ref(null)
    const color = ref('')
    const value = ref('0')
    const nodeId = ref(null)
    const dataNode = ref({})
    const df = getCurrentInstance().appContext.config.globalProperties.$df

    const updateInput = (value) => {
      let val = value.target.value
      if (isNaN(val)) {
        val = '0'
        value.target.value = '0'
      }
      dataNode.value.data.value = val
      dataNode.value.data.pythonStr = `nodo_${nodeId.value}=${val}`
      df.value.updateNodeDataFromId(nodeId.value, dataNode.value.data)
      triggerRef(df)
    }

    onMounted(async () => {
      await nextTick()
      nodeId.value = el.value.parentElement.parentElement.id.slice(5)
      dataNode.value = df.value.getNodeFromId(nodeId.value)
      console.log('initial value', dataNode.value)
      value.value = dataNode.value.data.value
      color.value = dataNode.value.data.color
    })
    return { el, value, nodeId, df, updateInput, dataNode }
  },
}
</script>
