import NodeNumber from './components/NodeNumber'
import NodeAdd from './components/NodeAdd'
import NodeLess from './components/NodeLess'
import NodeFooter from './components/NodeFooter'
import NodeInfo from './components/NodeInfo'
import NodeAssign from './components/NodeAssign'
import NodeHeader from './components/NodeHeader'
import NodeCodeBlock from './components/NodeCodeBlock'
import NodeFor from './components/NodeFor'
import NodeIfElse from './components/NodeIfElse'
import ListPrograms from './ListPrograms'

const components = {
  NodeNumber,
  NodeAdd,
  NodeLess,
  NodeFooter,
  NodeInfo,
  NodeAssign,
  NodeHeader,
  NodeCodeBlock,
  NodeFor,
  NodeIfElse,
  ListPrograms,
}
export default {
  install: (app) => {
    let pluginComponents = Object.keys(components) // let pluginDirectives = Directives const
    pluginComponents.forEach((plugin) => {
      app.component(plugin, components[plugin])
    })
  },
}
