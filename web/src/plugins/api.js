import Vue from 'vue'
import { GraphQLClient } from 'graphql-request'
import SSE from '../vendor/sse'
import { capitalize } from '../util/misc'

const graphQLEndpoint = `${process.env.VUE_APP_API_URL}/query`

class GraphQLSSEClient {
  constructor (url, options) {
    this.url = url
    this.options = options
  }

  request (query, variables) {
    const source = new SSE(this.url, {
      method: 'POST',
      headers: this.options.headers || undefined,
      payload: JSON.stringify({
        query,
        variables
      })
    })

    let handle = null

    function tryReconnect () {
      if (handle !== null) return

      handle = setInterval(() => {
        if ([source.OPEN, source.CONNECTING].includes(source.readyState)) {
          clearTimeout(handle)
          handle = null
          return
        }
        source.stream()
      }, 1000)
    }

    source.disconnect = () => {
      if (handle) {
        clearInterval(handle)
        source.removeAllListeners('error')
        source.removeAllListeners('readystatechange')
      }
      return source.close()
    }

    source.addEventListener('error', tryReconnect)
    source.addEventListener('readystatechange', e => {
      tryReconnect()
    })

    return source
  }

  setHeader (key, value) {
    if (!this.options.headers) {
      this.options.headers = {}
    }
    delete this.options.headers[key.toLowerCase()]
    delete this.options.headers[key.toUpperCase()]
    delete this.options.headers[capitalize(key)]
    this.options.headers[key] = value
  }
}

function authHeader () {
  // TODO: add authentication and authorization
  const token = ''
  return token ? { Authorization: `Bearer ${token}` } : {}
}

const api = {
  graphql: new GraphQLClient(graphQLEndpoint, {
    headers: authHeader()
  }),
  sse: new GraphQLSSEClient(graphQLEndpoint, {
    headers: Object.assign({
      'Content-Type': 'application/json'
    }, authHeader())
  })
}

export default {
  install () {
    Vue.$api = api
    Vue.prototype.$api = api
  }
}
