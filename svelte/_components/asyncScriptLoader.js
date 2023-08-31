function loader (urls, test, callback) {
  let remaining = urls.length
  
  function maybeCallback () {
    remaining = --remaining
    if (remaining < 1) {
      callback()
    }
  }
  
  if (test()) {
    return callback()
  }
  
  for (const { type, url, content, options = { async: true, defer: true }} of urls) {
    const isScript = type === 'script'
    const tag = document.createElement(isScript ? 'script': 'link')
    const attribute = isScript ? 'src' : 'href'
    const hasUrl = Boolean(url).valueOf()
    
    if (isScript) {
      tag.async = options.async
      tag.defer = options.defer
    } else {
      tag.rel = 'stylesheet'
    }
    
    if (hasUrl) {
      tag[attribute] = url
    } else {
      tag.appendChild(
        document.createTextNode(content)
      )
    }
    
    tag.onload = maybeCallback
    document.body.appendChild(tag)
  }
}

export default loader;