import request from '@/utils/request'

// 获取订单列表
export function getOrders(params) {
  return request({
    url: '/api/orders',
    method: 'get',
    params,
  })
}

// 获取订单详情
export function getOrder(id) {
  return request({
    url: `/api/orders/${id}`,
    method: 'get',
  })
}

// 发货
export function shipOrder(id) {
  return request({
    url: `/api/orders/${id}/ship`,
    method: 'post',
  })
}

// 取消订单
export function cancelOrder(id) {
  return request({
    url: `/api/orders/${id}/cancel`,
    method: 'post',
  })
}

// 获取订单统计数据
export function getOrderStats() {
  return request({
    url: '/api/orders/stats',
    method: 'get',
  })
} 
