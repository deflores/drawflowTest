<template>
  <div ref="el">
    <node-footer></node-footer>
    <node-header :title="'Bloque For'" :color="color"></node-header>
    <CInputGroup>
      <CInputGroupText id="basic-addon1">for</CInputGroupText>
      <CFormInput
        placeholder="desde"
        aria-label="desde"
        aria-describedby="basic-addon1"
        v-model="valueDesde"
        @input="updateInput('valueDesde', '0', $event)"
      />
      <CInputGroupText id="basic-addon1">to</CInputGroupText>
      <CFormInput
        placeholder="hasta"
        aria-label="hasta"
        aria-describedby="basic-addon1"
        v-model="valueHasta"
        @input="updateInput('valueHasta', '1', $event)"
      />
    </CInputGroup>
    <CFormTextarea
      placeholder="code"
      aria-label="code"
      aria-describedby="basic-addon1"
      v-model="value"
      @input="updateInputCode"
    />
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
  name: 'NodeFor',
  props: ['color'],
  setup() {
    const el = ref(null)
    const color = ref('')
    const value = ref('0')
    const valueHasta = ref('0')
    const valueDesde = ref('0')
    const nodeId = ref(null)
    const dataNode = ref({})
    const df = getCurrentInstance().appContext.config.globalProperties.$df

    const computePythonStr = () => {
      dataNode.value.data.pythonStr = `for nodo_${nodeId.value} in range(${
        dataNode.value.data.valueDesde
      }, ${dataNode.value.data.valueHasta}):  
      ${dataNode.value.data.value.replace('\n', ';')}`
      df.value.updateNodeDataFromId(nodeId.value, dataNode.value.data)
    }
    const updateInputCode = (value) => {
      let val = value.target.value
      dataNode.value.data.value = val
      df.value.updateNodeDataFromId(nodeId.value, dataNode.value.data)
      triggerRef(df)
    }
    const updateInput = (field, defval, value) => {
      let val = value.target.value
      if (isNaN(val)) {
        val = defval
      }
      dataNode.value.data[field] = val
      value.target.value = val
      df.value.updateNodeDataFromId(nodeId.value, dataNode.value.data)
      triggerRef(df)
    }

    onMounted(async () => {
      await nextTick()
      nodeId.value = el.value.parentElement.parentElement.id.slice(5)
      dataNode.value = df.value.getNodeFromId(nodeId.value)
      value.value = dataNode.value.data.value
      color.value = dataNode.value.data.color
      valueHasta.value = dataNode.value.data.valueHasta
      valueDesde.value = dataNode.value.data.valueDesde
    })

    watch(
      getCurrentInstance().appContext.config.globalProperties.$df,
      //eslint-disable-next-line
      (df, preDf) => {
        dataNode.value = df.getNodeFromId(nodeId.value)
        computePythonStr()
      },
    )
    return {
      el,
      value,
      nodeId,
      df,
      updateInput,
      dataNode,
      updateInputCode,
      valueHasta,
      valueDesde,
    }
  },
}
</script>
