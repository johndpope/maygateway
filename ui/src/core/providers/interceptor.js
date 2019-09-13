import axios from 'axios'
import { defaultsDeep } from 'lodash'

const getBaseConfig = () => ({
  headers: {
    Accept: 'application/json',
    'content-type': 'application/json;charset=utf-8',
  },
})

const getAxiosOptions = baseURL => axios.create({ baseURL })

const mergeOptions = (axiosConfig, path, options) => {
  const mergedOptions = defaultsDeep(options, getBaseConfig())

  return axiosConfig(path, mergedOptions).then(({ data }) => data)
}

const onResponseError = error => Promise.reject(error)


const onRequestValidate = config => config

const getAxiosConfiguration = (HOST) => {
  const axiosConfig = getAxiosOptions(HOST)
  axiosConfig.request = (path, options) => mergeOptions(axiosConfig, path, options)
  axiosConfig.interceptors.request.use(onRequestValidate)
  axiosConfig.interceptors.response.use(null, onResponseError)

  return axiosConfig
}

const baseApi = {
  async request(path, options) {
    const axiosConfig = getAxiosConfiguration('http://localhost:8080/v1')

    return axiosConfig.request(path, options)
  },
}

export const ApmApi = {
  async request(path, options) {
    const axiosConfig = getAxiosConfiguration(global.ENVIRONMENT.APM)

    return axiosConfig.request(path, options)
  },
}

export default baseApi