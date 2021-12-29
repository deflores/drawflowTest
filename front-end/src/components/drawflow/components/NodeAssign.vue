<template>
  <div ref="el">
    <node-footer></node-footer>
    <node-header :color="color" :title="'Asignación'" />
    <br />
    <span>{{ value }}</span>
    <node-info :inputs="inputs" />
  </div>
</template>

<script>
import { getCurrentInstance, nextTick, ref, onMounted, watch } from 'vue'

export default {
  name: 'NodeAssign',
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

    const computeValue = () => {
      let sum = ''
      dataNode.value.inputs.input_1.connections.forEach((n) => {
        const nodi = df.value.getNodeFromId(n.node)
        if (nodi.name !== 'NodeCodeBlock' && nodi.nodeName !== 'NodeFor') {
          sum = nodi.data.value
          dataNode.value.data.pythonStr = `nodo_${nodeId.value} = nodo_${n.node}`
        } else {
          df.value.removeSingleConnection(
            n.node,
            nodeId.value,
            'output_1',
            'input_1',
          )
        }
      })
      value.value = sum

      dataNode.value.data.value = sum

      df.value.updateNodeDataFromId(nodeId.value, dataNode.value.data)
    }
    watch(
      getCurrentInstance().appContext.config.globalProperties.$df,
      //eslint-disable-next-line
      (df, preDf) => {
        dataNode.value = df.getNodeFromId(nodeId.value)
        inputs.value = dataNode.value.inputs.input_1.connections

        computeValue()
      },
    )
    return { el, df, dataNode, nodeId, value, inputs }
  },
}
</script>
