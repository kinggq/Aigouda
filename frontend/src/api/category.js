import request from '@/utils/request'

export function getCategories() {
  return request({
    url: '/api/categories/list',
    method: 'post'
  })
}

export function getCategory(id) {
  return request({
    url: '/api/categories/detail',
    method: 'post',
    data: { id }
  })
}

export function createCategory(data) {
  return request({
    url: '/api/categories/create',
    method: 'post',
    data
  })
}

export function updateCategory(id, data) {
  data = { id, ...data }
  return request({
    url: `/api/categories/update/${id}`,
    method: 'post',
    data
  })
}

export function deleteCategory(id) {
  return request({
    url: `/api/categories/delete/${id}`,
    method: 'delete',
    data: { id }
  })
}

export function getCategoryProducts(id, params) {
  return request({
    url: '/api/categories/products',
    method: 'post',
    data: { id, ...params }
  })
} 
