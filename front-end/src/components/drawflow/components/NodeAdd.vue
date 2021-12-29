<template>
  <div ref="el">
    <node-footer></node-footer>
    <node-header :color="color" :title="'Sumatoria'" />
    <br />
    <span>{{ value }}</span>
    <node-info :inputs="inputs" />
  </div>
</template>

<script>
import { getCurrentInstance, nextTick, ref, onMounted, watch } from 'vue'
export default {
  name: 'NodeAdd',
  props: ['color'],
  setup() {
    const el = ref(null)
    const nodeId = ref(null)
    const color = ref('')
    const value = ref('0')
    const dataNode = ref({})
    const inputs = ref([])
    const df = getCurrentInstance().appContext.config.globalProperties.$df
    onMounted(async () => {
      await nextTick()
      nodeId.value = el.value.parentElement.parentElement.id.slice(5)
      dataNode.value = df.value.getNodeFromId(nodeId.value)
      value.value = dataNode.value.data.value
      color.value = dataNode.value.data.color
    })
    const computeAdd = () => {
      let sum = 0
      let arr = []
      dataNode.value.inputs.input_1.connections.forEach((n) => {
        sum += parseFloat(df.value.getNodeFromId(n.node).data.value)
        arr.push(`nodo_${n.node}`)
      })
      value.value = sum
      if (arr.length > 0) {
        dataNode.value.data.pythonStr = `nodo_${nodeId.value}=${arr.join(
          ' + ',
        )}`
      }
      dataNode.value.data.value = sum
      df.value.updateNodeDataFromId(nodeId.value, dataNode.value.data)
    }
    watch(
      getCurrentInstance().appContext.config.globalProperties.$df,
      //eslint-disable-next-line
      (df, preDf) => {
        dataNode.value = df.getNodeFromId(nodeId.value)
        inputs.value = dataNode.value.inputs.input_1.connections

        computeAdd()
      },
    )
    return { el, df, dataNode, nodeId, value, inputs }
  },
}
</script>
