import request from '@/utils/request'

export function getUserList(query) {
  return request({
    url: '/home/list',
    method: 'get',
    params: query
  })
}

export function updateUser(data) {
  return request({
    url: '/home/edit',
    method: 'post',
    data
  })
}

export function addUser(data) {
  return request({
    url: '/home/add',
    method: 'post',
    data
  })
}

export function getGoogleCode(data) {
  return request({
    url: '/home/getgooglecode',
    method: 'get',
    params: { data }
  })
}
