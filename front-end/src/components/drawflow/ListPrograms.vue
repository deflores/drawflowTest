<template>
  <CInputGroup>
    <CInputGroupText id="basic-addon1">Nombre </CInputGroupText>
    <CFormInput
      aria-describedby="basic-addon1"
      v-model="programName"
      @input="(val) => (programName = val.target.value)"
    />
    <CButton color="primary" @click="getList($event)">Buscar</CButton>
  </CInputGroup>
  <CTable hover>
    <CTableHead>
      <CTableRow>
        <CTableHeaderCell scope="col">#</CTableHeaderCell>
        <CTableHeaderCell scope="col">Nombre</CTableHeaderCell>
        <CTableHeaderCell scope="col"></CTableHeaderCell>
      </CTableRow>
    </CTableHead>
    <CTableBody>
      <CTableRow v-for="p in programs" :key="p.uid">
        <CTableHeaderCell scope="row">{{ p.uid }}</CTableHeaderCell>
        <CTableDataCell>{{ p.name }}</CTableDataCell>
        <CTableDataCell
          ><CButton color="primary" @click="emitImport(p.uid)"
            >Ver</CButton
          ></CTableDataCell
        >
      </CTableRow>
    </CTableBody>
  </CTable>
</template>
<script>
import { onMounted } from '@vue/runtime-core'
export default {
  name: 'ListPrograms',
  data() {
    return { programs: [], programName: '' }
  },
  setup() {
    onMounted(async () => {})
    return {}
  },
  emits: ['shouldImport'],
  methods: {
    getList() {
      var url = 'http://localhost/drawflow'
      if (this.programName !== '')
        url = `http://localhost/drawflow?name=${this.programName}`

      this.axios.get(url).then((response) => {
        console.log(response.data)
        this.programs = response.data
      })
    },
    emitImport(id) {
      console.log('shoul import child', id)
      this.$emit('shouldImport', id)
    },
  },
}
</script>
