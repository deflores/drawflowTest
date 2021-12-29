<template>
  <div ref="el">
    <node-footer></node-footer>
    <node-header :title="'Bloque de código'" :color="color"></node-header>
    <CInputGroup>
      <CFormTextarea
        placeholder="code"
        aria-label="code"
        aria-describedby="basic-addon1"
        v-model="value"
        @input="updateInput"
      />
    </CInputGroup>
  </div>
</template>

<script>
import {
  ref,
  onMounted,
  getCurrentInstance,
  nextTick,
  triggerRef,
  watch,
} from 'vue'
export default {
  name: 'NodeCodeBlock',
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
      dataNode.value.data.value = val
      dataNode.value.data.pythoStr = val
      df.value.updateNodeDataFromId(nodeId.value, dataNode.value.data)
      triggerRef(df)
    }
    const computePythonStr = () => {
      dataNode.value.data.pythonStr = `${dataNode.value.data.value.replace(
        '\n',
        ';',
      )}`
      df.value.updateNodeDataFromId(nodeId.value, dataNode.value.data)
    }
    onMounted(async () => {
      await nextTick()
      nodeId.value = el.value.parentElement.parentElement.id.slice(5)
      dataNode.value = df.value.getNodeFromId(nodeId.value)
      console.log('initial value', dataNode.value)
      value.value = dataNode.value.data.value
      color.value = dataNode.value.data.color
    })

    watch(
      getCurrentInstance().appContext.config.globalProperties.$df,
      //eslint-disable-next-line
      (df, preDf) => {
        dataNode.value = df.getNodeFromId(nodeId.value)
        computePythonStr()
      },
    )
    return { el, value, nodeId, df, updateInput, dataNode }
  },
}
</script>
