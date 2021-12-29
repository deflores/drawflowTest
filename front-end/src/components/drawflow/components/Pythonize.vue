<template>
  <div class="drawflow-code">
    <div class="justify-content-center m-3">
      <CButton color="info" @click="fetchData">Procesar</CButton><br />
      <span class="text-info"
        >presione el botón procesar para refrescar los resultados</span
      >
    </div>
    <CContainer class="text-white">
      <h5 class="bg-dark">Código</h5>
      <pre>{{ formattedCode }}</pre>
    </CContainer>
    <CContainer class="text-white">
      <h5 class="bg-dark">Resultado</h5>
      <p style="white-space: pre-line">{{ resultado }}</p>
    </CContainer>
  </div>
</template>
<script>
import axios from 'axios'

export default {
  name: 'DrawCustom',
  props: ['nodes'],
  data() {
    return { resultado: '', formattedCode: '' }
  },
  emit: ['error'],
  computed: {
    data: (vm) => {
      const noditos = Object.keys(vm.nodes) // let pluginDirectives = Directives const
      let arr = []
      noditos.forEach((node) => {
        arr.push(vm.nodes[node])
      })
      console.log(arr)
      const inputs = arr.filter((n) => n.data.type === 'input')
      const operations = arr.filter(
        (n) => n.data.type === 'operation' || n.data.type === 'bloque',
      )
      let str = ''
      let nodesProc = []
      inputs.forEach((n) => {
        str = str + n.data.pythonStr + '\n'
        nodesProc.push(n.id)
      })

      operations.sort(function (a, b) {
        if (a.id > b.id) {
          return 1
        }
        if (a.id < b.id) {
          return -1
        }
        // a must be equal to b
        return 0
      })
      console.log('sorted', operations)
      operations.forEach((n) => {
        if (n.inputs.input_1 && n.inputs.input_1.connections.length > 0) {
          n.inputs.input_1.connections.forEach((nn) => {
            const elnode = arr.find((nnn) => nnn.id === parseInt(nn.node))
            console.log('print elnode', elnode)
            if (!nodesProc.find((x) => x === elnode.id)) {
              console.log('print imputs')
              str = str + elnode.data.pythonStr + '\n'
              nodesProc.push(elnode.id)
            }
          })
        }
        if (!nodesProc.find((x) => x === n.id)) {
          str = str + n.data.pythonStr + '\n'
          nodesProc.push(n.id)
        }
      })

      return str
    },
  },
  methods: {
    async fetchData() {
      console.log('las daas', this.data)
      const payload = {
        code: this.data,
      }
      const response = await axios.post(
        'http://localhost/drawflow/save/code',
        payload,
      )
      if (response.status === 200) {
        const formatResponse = await axios.get('http://localhost:5000/format')
        if (formatResponse.data.code === 'success') {
          this.formattedCode = formatResponse.data.message
          const readExec = await axios.get('http://localhost:5000/execute')
          if (readExec.data.code === 'success') {
            const readR = await axios.get('http://localhost:5000/read')
            if (readR.status === 200) {
              this.resultado = readR.data.result
            } else {
              this.resultado = readR.data.message

              this.$emit('error', 'Error', readR.data.message, 'danger')
            }
          } else {
            this.resultado = readExec.data.message
          }
        } else {
          this.$emit('error', 'Error', formatResponse.data.message, 'danger')
        }
      }
    },
  },
}
</script>
