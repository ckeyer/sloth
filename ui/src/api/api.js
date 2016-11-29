import request from 'superagent'

const API_ROOT = process.env.API_ROOT

export default {
  ping: () => {
    return request.get(API_ROOT + '/_ping')
  }
}
