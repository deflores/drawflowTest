<template>
  <div ref="el">
    <node-footer></node-footer>
    <node-header :title="'Bloque If'" :color="color"></node-header>
    <CInputGroup>
      <CInputGroupText id="basic-addon">If</CInputGroupText>
      <CFormSelect
        placeholder="nodo"
        aria-label="nodo"
        aria-describedby="basic-addon"
        @input="updateInput('nodeFrom', $event)"
      >
        <option value="0" />
        <option v-for="op in inputs" :key="op.node" :value="op.node">
          nodo_{{ op.node }}
        </option>
      </CFormSelect>

      <CInputGroupText id="basic-addon2">igual</CInputGroupText>
      <CFormInput
        placeholder="valor"
        aria-label="valor"
        aria-describedby="basic-addon2"
        v-model="valueCompare"
        @input="updateInput('valueCompare', $event)"
      />
    </CInputGroup>
    <CInputGroup>
      <CInputGroupText id="basic-addon1"
        >salida si cumple<span>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</span>
      </CInputGroupText>
      <CFormInput
        placeholder="valor"
        aria-label="valor"
        aria-describedby="basic-addon1"
        v-model="valueIf"
        @input="updateInput('valueIf', $event)"
      />
    </CInputGroup>
    <CInputGroup>
      <CInputGroupText id="basic-addon3">salida si no cumple</CInputGroupText>
      <CFormInput
        placeholder="valor"
        aria-label="valor"
        aria-describedby="basic-addon3"
        v-model="valueElse"
        @input="updateInput('valueElse', $event)"
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
  name: 'NodeIfElse',
  props: ['color'],
  setup() {
    const el = ref(null)
    const color = ref('')
    const value = ref('0')
    const nodeId = ref(null)
    const dataNode = ref({})
    const inputs = ref([])
    const nodeFrom = ref('')
    const valueElse = ref('')
    const valueIf = ref('')
    const valueCompare = ref('')
    const df = getCurrentInstance().appContext.config.globalProperties.$df

    const computeOperation = () => {
      if (
        dataNode.value.data.nodeFrom &&
        dataNode.value.data.nodeFrom !== '0' &&
        dataNode.value.data.valueCompare &&
        dataNode.value.data.valueIf &&
        dataNode.value.data.valueElse
      ) {
        const nodeFromVal = df.value.getNodeFromId(dataNode.value.data.nodeFrom)
          .data.value
        let val = ''
        if (nodeFromVal == dataNode.value.data.valueCompare) {
          val = dataNode.value.data.valueIf
        } else {
          val = dataNode.value.data.valueElse
        }
        dataNode.value.data.pythonStr = `if nodo_${dataNode.value.data.nodeFrom} == ${dataNode.value.data.valueCompare}: nodo_${nodeId.value}=${dataNode.value.data.valueIf};
else: nodo_${nodeId.value}=${dataNode.value.data.valueElse};`
        dataNode.value.data.value = val
        df.value.updateNodeDataFromId(nodeId.value, dataNode.value.data)
      }
    }

    const updateInput = (obj, value) => {
      let val = value.target.value
      dataNode.value.data[obj] = val
      df.value.updateNodeDataFromId(nodeId.value, dataNode.value.data)
      triggerRef(df)
    }

    onMounted(async () => {
      await nextTick()
      nodeId.value = el.value.parentElement.parentElement.id.slice(5)
      dataNode.value = df.value.getNodeFromId(nodeId.value)
      value.value = dataNode.value.data.value
      color.value = dataNode.value.data.color
      valueIf.value = dataNode.value.data.valueIf
      valueElse.value = dataNode.value.data.valueElse
      valueCompare.value = dataNode.value.data.valueCompare
    })

    watch(
      getCurrentInstance().appContext.config.globalProperties.$df,
      //eslint-disable-next-line
      (df, preDf) => {
        dataNode.value = df.getNodeFromId(nodeId.value)
        inputs.value = dataNode.value.inputs.input_1.connections
        computeOperation()
      },
    )

    return {
      el,
      value,
      nodeId,
      df,
      dataNode,

      computeOperation,
      inputs,
      nodeFrom,
      updateInput,
      valueIf,
      valueElse,
      valueCompare,
    }
  },
}
</script>
