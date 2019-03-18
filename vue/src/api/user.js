import request from '@/utils/request'

export function getUserList(query) {
  return request({
    url: '/admin/list',
    method: 'get',
    params: query
  })
}

export function updateUser(data) {
  return request({
    url: '/admin/update',
    method: 'post',
    data
  })
}
