import request from 'superagent'

const API_ROOT = process.env.API_ROOT

export default {
  ping: () => {
    return request.get(API_ROOT + '/_ping')
  },
  login: (body) => {
    return request.post(API_ROOT + '/login')
      .send(body)
      .withCredentials()
      .set('Accept', 'application/json')
  }
}
