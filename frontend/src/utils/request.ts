import axios from 'axios'
import { message } from 'ant-design-vue'
import router from '@/router'

interface RequestConfig {
  url: string
  method: string
  data?: any
  params?: any
  headers?: Record<string, string>
}

interface Response<T = any> {
  code: number
  data: T
  message: string
}

const request = axios.create({
  baseURL: 'http://localhost:8080',
  timeout: 5000
})

// 请求拦截器
request.interceptors.request.use(
  config => {
    const token = localStorage.getItem('token')
    if (token) {
      config.headers.Authorization = `Bearer ${token}`
    }
    return config
  },
  error => {
    return Promise.reject(error)
  }
)

// 响应拦截器
request.interceptors.response.use(
  response => {
    const res = response.data
    if (res.code === 0) {
      return res
    }
    message.error(res.message || '请求失败')
    return Promise.reject(new Error(res.message || '请求失败'))
  },
  error => {
    if (error.response) {
      switch (error.response.status) {
        case 401:
          message.error('请先登录')
          router.push('/login')
          break
        case 403:
          message.error('没有权限')
          break
        case 404:
          message.error('请求的资源不存在')
          break
        case 500:
          message.error('服务器错误')
          break
        default:
          message.error(error.response.data.message || '请求失败')
      }
    } else {
      message.error('网络错误，请检查网络连接')
    }
    return Promise.reject(error)
  }
)

export default function<T = any>(config: RequestConfig): Promise<Response<T>> {
  return request(config)
} 
