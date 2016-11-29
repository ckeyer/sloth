import createLogger from 'vuex/src/plugins/logger'

export default
  process.env.NODE_ENV !== 'production' ? [createLogger()] : [ ]
