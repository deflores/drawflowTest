export function computePython(data, template, nodeId) {
  let keys = Object.keys(data) // let pluginDirectives = Directives const
  keys.forEach((elem) => {
    const reg = new RegExp(`^${elem}$`, 'gi')
    template = template.replace(reg, data[elem])
  })
  const reg = new RegExp(`nodeId`, 'gi')
  template = template.replaceAll(reg, nodeId)

  console.log(template)
  return template
}
