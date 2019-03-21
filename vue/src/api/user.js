import request from '@/utils/request'

export function getUserList(query) {
  return request({
    url: '/home/userlist',
    method: 'get',
    params: query
  })
}

export function updateUser(data) {
  return request({
    url: '/home/useredit',
    method: 'put',
    data
  })
}

export function addUser(data) {
  return request({
    url: '/home/useradd',
    method: 'post',
    data
  })
}

export function getGoogleCode(data) {
  return request({
    url: '/home/userlist',
    method: 'get'
  })
}
