import Base from './interceptor'

const API = '/apis/'

const Apis = {
  getAll(data) {
    return Base.request(`${API}`, { method: 'GET' })
  },
}

export default Apis