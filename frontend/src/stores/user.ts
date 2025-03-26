import { defineStore } from 'pinia'
import axios from 'axios'

interface User {
  id: number
  username: string
  role: string
}

export const useUserStore = defineStore('user', {
  state: () => ({
    user: null as User | null,
    token: localStorage.getItem('token') || null,
    loading: false,
    error: null as string | null,
  }),

  actions: {
    async login(username: string, password: string) {
      this.loading = true
      try {
        const response = await axios.post('/api/admin/login', {
          username,
          password,
        })
        this.token = response.data.token
        this.user = response.data.user
        localStorage.setItem('token', this.token)
      } catch (error) {
        this.error = '登录失败'
      } finally {
        this.loading = false
      }
    },

    logout() {
      this.user = null
      this.token = null
      localStorage.removeItem('token')
    },

    async checkAuth() {
      if (!this.token) return false
      try {
        const response = await axios.get('/api/admin/me', {
          headers: {
            Authorization: `Bearer ${this.token}`,
          },
        })
        this.user = response.data
        return true
      } catch (error) {
        this.logout()
        return false
      }
    },
  },
}) 
