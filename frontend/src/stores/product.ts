import { defineStore } from 'pinia'
import axios from 'axios'

interface Product {
  id: number
  title: string
  price: number
  originalPrice: number
  discount: number
  mainImage: string
  images: string[]
  tags: string[]
  sales: number
  comments: Comment[]
}

interface Comment {
  id: number
  userId: number
  content: string
  rating: number
}

export const useProductStore = defineStore('product', {
  state: () => ({
    products: [] as Product[],
    currentProduct: null as Product | null,
    loading: false,
    error: null as string | null,
  }),

  actions: {
    async fetchProducts() {
      this.loading = true
      try {
        const response = await axios.get('/api/products')
        this.products = response.data
      } catch (error) {
        this.error = '获取商品列表失败'
      } finally {
        this.loading = false
      }
    },

    async fetchProductDetail(id: number) {
      this.loading = true
      try {
        const response = await axios.get(`/api/products/${id}`)
        this.currentProduct = response.data
      } catch (error) {
        this.error = '获取商品详情失败'
      } finally {
        this.loading = false
      }
    },
  },
}) 
