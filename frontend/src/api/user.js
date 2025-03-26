import request from '@/utils/request'

export function login(data) {
  return request({
    url: '/api/login',
    method: 'post',
    data
  })
}

export function register(data) {
  return request({
    url: '/api/register',
    method: 'post',
    data
  })
}

export function getProfile() {
  return request({
    url: '/api/user/profile',
    method: 'post'
  })
}

export function updateProfile(data) {
  return request({
    url: '/api/user/profile/update',
    method: 'post',
    data
  })
}

export function updatePassword(data) {
  return request({
    url: '/api/user/password/update',
    method: 'post',
    data
  })
} 
