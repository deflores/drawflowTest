<template>
  <CContainer ref="el" fluid class="drawi p-0">
    <CAlert color="info"
      >Tip: Para utilizar los <strong>shortcuts</strong> y no entrar en
      conflicto con los de tu navegador, presiona de manera separada las teclas
      definidas
    </CAlert>
    <div class="d-flex justify-content-end bg-light p-1">
      <CInputGroup>
        <CInputGroupText id="basic-addon1">Nombre </CInputGroupText>
        <CFormInput
          class="form-control"
          aria-label="name"
          aria-describedby="basic-addon1"
          v-model="drname"
        />
      </CInputGroup>
    </div>
    <CRow>
      >
      <CCol class="p-0 drawi-space bg-cece">
        <div class="sticky"></div>
        <DrawMenu @add-node="recibeNode" />
        <CListGroup class="p-0 drawflow-menu drawflow-menu-right">
          <CListGroupItem class="bg-primary" @click="save"
            >Guardar</CListGroupItem
          >
          <CListGroupItem class="bg-secondary" @click="openModal"
            >Importar</CListGroupItem
          >
          <CListGroupItem class="bg-success" @click="nuevo"
            >Nuevo</CListGroupItem
          >
          <CListGroupItem class="bg-dark" @click="toggleShowPythonCode">{{
            showCode ? 'Editor' : 'Python'
          }}</CListGroupItem>
        </CListGroup>
        <div id="denitos" :class="!showCode ? 'visible' : 'no-visible'"></div>
        <pythonize
          :class="showCode ? 'visible' : 'no-visible'"
          :nodes="code"
          @error="createToast"
        ></pythonize>
      </CCol>
    </CRow>
    <CModal
      class="show d-block"
      :backdrop="false"
      :visible="showPrograms"
      @close="closeModal"
    >
      <CModalHeader class="bg-light">
        <CModalTitle>Programas</CModalTitle>
      </CModalHeader>
      <CModalBody
        ><list-programs @should-import="shouldImport"></list-programs
      ></CModalBody>
      <CModalFooter>
        <CButton color="secondary" @click="closeModal">Cerrar</CButton>
      </CModalFooter>
    </CModal>
    <CToaster placement="top-end">
      <CToast v-for="(toast, index) in toasts" :key="index">
        <CToastHeader closeButton :class="'bg-' + toast.color">
          <span class="me-auto fw-bold text-white">{{ toast.title }}</span>
        </CToastHeader>
        <CToastBody>
          {{ toast.content }}
        </CToastBody>
      </CToast>
    </CToaster>
    <CModal
      backdrop="static"
      :visible="confirmacion"
      @close="
        () => {
          confirmacion = false
        }
      "
    >
      <CModalHeader>
        <CModalTitle>Programa Nuevo</CModalTitle>
      </CModalHeader>
      <CModalBody>
        <span>
          {{
            presaved
              ? 'Guardaste un programa si desseas crear uno nuevo a partir del recientemente guardado presiona "Mantener" sino presiona "Nuevo"'
              : '¿Deseas crear un programa nuevo?'
          }}
        </span>
      </CModalBody>
      <CModalFooter>
        <CButton
          color="secondary"
          @click="
            () => {
              confirmacion = false
            }
          "
        >
          {{ presaved ? 'Mantener' : 'No' }}
        </CButton>
        <CButton color="primary" @click="confirmaNuevo">{{
          presaved ? 'Nuevo' : 'Sí'
        }}</CButton>
      </CModalFooter>
    </CModal>
  </CContainer>
</template>

<script>
import Drawflow from 'drawflow'
import {
  shallowRef,
  h,
  getCurrentInstance,
  render,
  resolveDynamicComponent,
  onMounted,
  triggerRef,
  ref,
  nextTick,
} from 'vue'
import DrawMenu from './DrawMenu.vue'
import nodes from './operations'
import axios from 'axios'
import Pythonize from './components/Pythonize.vue'

export default {
  name: 'DrawCustom',
  components: { DrawMenu, Pythonize },
  data() {
    return {
      showPrograms: false,
      uid: null,
      drname: '',
      toasts: [],
      confirmacion: false,
      showCode: false,
      activeView: 1,
      code: [],
    }
  },
  setup() {
    const editor = shallowRef({})
    const el = ref(null)
    const controlKey = ref(false)
    const presaved = ref(false)

    const internalInstance = getCurrentInstance()
    internalInstance.appContext.app._context.config.globalProperties.$df =
      editor

    const recibeNode = (node) => {
      editor.value.addNode(
        node.nodeName,
        node.inputs,
        node.outputs,
        150,
        300,
        node.class || 'node',
        node.data,
        node.nodeName,
        'vue',
      )
    }

    /*eslint-disable-next-line*/
    const computeOperation = (args) => {
      console.log('args', args)
      triggerRef(editor)
    }

    onMounted(async () => {
      const Vue = { version: 3, h, render }
      const id = document.getElementById('denitos')

      // Pass render Vue 3 Instance
      editor.value = new Drawflow(
        id,
        Vue,
        internalInstance.appContext.app._context,
      )

      nodes.forEach((element) => {
        const theComp = resolveDynamicComponent(element.nodeName)
        editor.value.registerNode(
          element.nodeName,
          theComp,
          element.data,
          element.options,
        )
      })

      const controlKeys = (event) => {
        if (event.key.toLowerCase() === 'control') {
          controlKey.value = true
        } else if (controlKey.value) {
          const elnode = nodes.find(
            (n) => n.shortcut === `Control+${event.key}`,
          )
          if (elnode) recibeNode(elnode)
          controlKey.value = false
        }
      }

      editor.value.start()
      editor.value.on('connectionCreated', computeOperation)
      editor.value.on('connectionRemoved', computeOperation)
      editor.value.on('nodeDataChanged', computeOperation)
      editor.value.on('keydown', controlKeys)
    })

    return { editor, recibeNode, el, controlKey, presaved }
  },
  methods: {
    closeModal() {
      this.showPrograms = false
    },
    openModal() {
      this.showPrograms = true
    },
    async shouldImport(id) {
      const response = await axios.get(`http://localhost/drawflow/${id}`)
      this.editor.import(JSON.parse(response.data.data))
      this.uid = response.data.uid
      this.drname = response.data.name
      this.closeModal()
    },
    async save() {
      if (this.drname) {
        const data = this.editor.export()
        const datapayload = {
          data: JSON.stringify(data),
          name: this.drname,
          uid: this.uid,
        }

        axios
          .post(`http://localhost/drawflow/save/update`, datapayload)
          // eslint-disable-next-line
          .then((response) => {
            //this.editor.clear()
            this.uid = ''
            this.drname = ''
            this.presaved = true
            this.createToast(
              'Guardado',
              'Programa guardado exitósamente',
              'success',
            )
            this.nuevo()
          })
      } else {
        this.createToast(
          'Información',
          'Ingresa un nombre para identificar el programa',
          'warning',
        )
      }
    },
    createToast(title, content, color) {
      this.toasts.push({
        title,
        content,
        color,
      })
    },
    updateName(val) {
      this.drname = val.target.value
    },
    nuevo() {
      this.confirmacion = true
    },
    confirmaNuevo() {
      //la siguiente linea tiene un bug
      //this.editor.clear()
      this.presaved = false
      this.uid = ''
      this.drname = ''
      this.confirmacion = false
      this.$router.go()
    },
    async toggleShowPythonCode() {
      if (this.showCode) {
        this.showCode = false
      } else {
        const data = this.editor.export()
        this.code = data.drawflow.Home.data
        await nextTick()
        this.showCode = true
      }
    },
  },
}
</script>
