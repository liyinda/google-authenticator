import request from '@/utils/request'

export function getUserList(query) {
  return request({
    url: '/admin/list',
    method: 'get',
    params: query
  })
}

