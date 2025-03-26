import request from '@/utils/request'

// 基础 URL 配置
const BASE_URL = 'http://localhost:8080'

// 转换后端数据为前端格式
function transformResponse(data) {
  if (Array.isArray(data)) {
    return data.map(item => ({
      id: item.id,
      title: item.title,
      price: item.price,
      original_price: item.original_price,
      discount: item.discount,
      main_image: item.main_image ? `${BASE_URL}${item.main_image}` : '',
      images: (item.images || []).map(img => `${BASE_URL}${img}`),
      description: item.description || '',
      category_id: item.category_id,
      sales: item.sales,
      created_at: item.created_at,
      updated_at: item.updated_at,
      stock: item.stock,
      category: item.category
    }))
  }
  return data
}

export function getProducts(params) {
  return request({
    url: '/api/products',
    method: 'get',
    params
  }).then(response => {
    if (response.code === 0) {
      return {
        ...response,
        data: transformResponse(response.data)
      }
    }
    return response
  })
}

export function getProduct(id) {
  return request({
    url: `/api/products/${id}`,
    method: 'get'
  }).then(response => {
    if (response.code === 0) {
      return {
        ...response,
        data: transformResponse(response.data)
      }
    }
    return response
  })
}

export function createProduct(data) {
  // 验证必填字段
  if (!data.title || !data.category_id || !data.price || !data.stock) {
    return Promise.reject(new Error('请填写所有必填字段'));
  }

  // 处理图片URL，确保是相对路径
  const productData = { ...data };
  if (productData.main_image) {
    productData.main_image = productData.main_image.replace(BASE_URL, '');
  }
  if (productData.images) {
    productData.images = productData.images.map(img => img.replace(BASE_URL, ''));
  }

  return request({
    url: '/api/products/create',
    method: 'post',
    data: productData
  }).then(response => {
    if (response.code === 0) {
      return {
        ...response,
        data: transformResponse(response.data)
      }
    }
    return response
  });
}

export function updateProduct(id, data) {
  // 验证必填字段
  if (!data.title || !data.category_id || !data.price || !data.stock) {
    return Promise.reject(new Error('请填写所有必填字段'));
  }

  // 处理图片URL，确保是相对路径
  const productData = { ...data };
  if (productData.main_image) {
    productData.main_image = productData.main_image.replace(BASE_URL, '');
  }
  if (productData.images) {
    productData.images = productData.images.map(img => img.replace(BASE_URL, ''));
  }

  return request({
    url: `/api/products/${id}`,
    method: 'put',
    data: productData
  }).then(response => {
    if (response.code === 0) {
      return {
        ...response,
        data: transformResponse(response.data)
      }
    }
    return response
  })
}

export function deleteProduct(id) {
  return request({
    url: `/api/products/${id}`,
    method: 'delete'
  })
}

export function uploadImage(data) {
  return request({
    url: '/api/products/upload',
    method: 'post',
    data,
    headers: {
      'Content-Type': 'multipart/form-data'
    }
  }).then(response => {
    if (response.code === 0) {
      return {
        ...response,
        data: {
          ...response.data,
          url: `${BASE_URL}${response.data.url}`  // 添加基础 URL
        }
      }
    }
    return response
  })
} 
