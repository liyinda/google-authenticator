import request from '@/utils/request'

export function loginByUsername(loginname, password) {
  const data = {
    loginname,
    password
  }
  return request({
    url: '/login',
    method: 'post',
    data
  })
}

export function logout() {
  return request({
    url: '/logout',
    method: 'post'
  })
}

export function getUserInfo(token) {
  return request({
    url: '/userinfo',
    method: 'get',
    params: { token }
  })
}

