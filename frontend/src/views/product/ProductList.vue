<template>
  <div class="product-list">
    <div class="header">
      <a-input-search
        v-model:value="searchQuery"
        placeholder="搜索商品"
        style="width: 200px"
        @search="onSearch"
      />
      <a-button type="primary" @click="showAddModal">
        <plus-outlined />添加商品
      </a-button>
    </div>

    <a-table
      :columns="columns"
      :data-source="products"
      :loading="loading"
      :pagination="pagination"
      @change="handleTableChange"
    >
      <template #bodyCell="{ column, record }">
        <template v-if="column.key === 'main_image'">
          <a-image
            :width="50"
            :src="record.main_image"
            :fallback="defaultImage"
          />
        </template>
        <template v-else-if="column.key === 'action'">
          <a-space>
            <a @click="showEditModal(record)">编辑</a>
            <a-divider type="vertical" />
            <a-popconfirm
              title="确定要删除这个商品吗？"
              @confirm="handleDelete(record.id)"
            >
              <a class="danger">删除</a>
            </a-popconfirm>
          </a-space>
        </template>
      </template>
    </a-table>
    <!-- 添加/编辑商品模态框 -->
    <a-modal
      v-model:visible="modalVisible"
      :title="modalTitle"
      @ok="handleModalOk"
      @cancel="handleModalCancel"
      :confirmLoading="modalLoading"
    >
      <a-form
        ref="formRef"
        :model="formState"
        :rules="rules"
        :label-col="{ span: 6 }"
        :wrapper-col="{ span: 16 }"
      >
        <a-form-item label="商品名称" name="title">
          <a-input v-model:value="formState.title" placeholder="请输入商品名称" />
        </a-form-item>
        <a-form-item label="商品分类" name="category_id">
          <a-select v-model:value="formState.category_id" placeholder="请选择商品分类">
            <a-select-option v-for="category in categories" :key="category.id" :value="category.id">
              {{ category.name }}
            </a-select-option>
          </a-select>
        </a-form-item>
        <a-form-item label="商品价格" name="price">
          <a-input-number
            v-model:value="formState.price"
            :min="0"
            :precision="2"
            style="width: 100%"
            placeholder="请输入商品价格"
          />
        </a-form-item>
        <a-form-item label="商品库存" name="stock">
          <a-input-number
            v-model:value="formState.stock"
            :min="0"
            :precision="0"
            style="width: 100%"
            placeholder="请输入商品库存"
          />
        </a-form-item>
        <a-form-item label="商品图片" name="main_image">
          <a-upload
            v-model:file-list="fileList"
            name="file"
            list-type="picture-card"
            class="avatar-uploader"
            :show-upload-list="false"
            :before-upload="beforeUpload"
          >
            <img v-if="formState.main_image" :src="formState.main_image" alt="avatar" />
            <div v-else>
              <plus-outlined />
              <div class="ant-upload-text">上传</div>
            </div>
          </a-upload>
        </a-form-item>
        <a-form-item label="商品描述" name="description">
          <a-textarea
            v-model:value="formState.description"
            :rows="4"
            placeholder="请输入商品描述"
          />
        </a-form-item>
      </a-form>
    </a-modal>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { message } from 'ant-design-vue'
import { PlusOutlined } from '@ant-design/icons-vue'
import { getProducts, createProduct, updateProduct, deleteProduct, uploadImage } from '@/api/product'
import { getCategories } from '@/api/category'

const columns = [
  {
    title: '商品图片',
    dataIndex: 'main_image',
    key: 'main_image',
  },
  {
    title: '商品名称',
    dataIndex: 'title',
    key: 'title',
  },
  {
    title: '分类',
    dataIndex: ['category', 'name'],
    key: 'category',
  },
  {
    title: '价格',
    dataIndex: 'price',
    key: 'price',
    render: (price) => `¥${price.toFixed(2)}`,
  },
  {
    title: '库存',
    dataIndex: 'stock',
    key: 'stock',
  },
  {
    title: '操作',
    key: 'action',
  },
]

const products = ref([])
const categories = ref([])
const loading = ref(false)
const modalVisible = ref(false)
const modalLoading = ref(false)
const modalTitle = ref('添加商品')
const formRef = ref(null)
const fileList = ref([])
const defaultImage = 'https://via.placeholder.com/50'

const searchQuery = ref('')
const pagination = reactive({
  current: 1,
  pageSize: 10,
  total: 0,
})

const formState = ref({
  title: '',
  price: 0,
  original_price: 0,
  discount: 0,
  main_image: '',
  description: '',
  category_id: undefined,
  stock: 0
})

const rules = {
  title: [{ required: true, message: '请输入商品名称' }],
  category_id: [{ required: true, message: '请选择商品分类' }],
  price: [{ required: true, message: '请输入商品价格' }],
  stock: [{ required: true, message: '请输入商品库存' }],
}

// 获取商品列表
const fetchProducts = async () => {
  loading.value = true
  try {
    const params = {
      page: pagination.current,
      page_size: pagination.pageSize,
      search: searchQuery.value,
    }
    const res = await getProducts(params)
    products.value = res.data
    pagination.total = res.total
  } catch (error) {
    console.error('获取商品列表失败:', error)
  } finally {
    loading.value = false
  }
}

// 获取分类列表
const fetchCategories = async () => {
  try {
    const res = await getCategories()
    categories.value = res.data
  } catch (error) {
    console.error('获取分类列表失败:', error)
  }
}

// 表格变化处理
const handleTableChange = (pag) => {
  pagination.current = pag.current
  pagination.pageSize = pag.pageSize
  fetchProducts()
}

// 搜索处理
const onSearch = () => {
  pagination.current = 1
  fetchProducts()
}

// 显示添加模态框
const showAddModal = () => {
  modalTitle.value = '添加商品'
  resetForm()
  modalVisible.value = true
}

// 显示编辑模态框
const showEditModal = (record) => {
  modalTitle.value = '编辑商品'
  handleEdit(record)
}

// 重置表单
const resetForm = () => {
  formState.value = {
    title: '',
    price: 0,
    original_price: 0,
    discount: 0,
    main_image: '',
    description: '',
    category_id: undefined,
    stock: 0
  }
  fileList.value = []
  modalVisible.value = true
}

// 处理图片上传成功
const handleUploadSuccess = (response) => {
  if (response.code === 0) {
    formState.value.main_image = response.data.url
  } else {
    message.error(response.message || '图片上传失败')
  }
}

// 处理编辑
const handleEdit = (record) => {
  formState.value = {
    ...record,
    title: record.title,
    price: record.price,
    original_price: record.original_price,
    discount: record.discount,
    main_image: record.main_image,
    description: record.description,
    category_id: record.category_id,
    stock: record.stock
  }
  fileList.value = record.main_image ? [{ url: record.main_image }] : []
  modalVisible.value = true
}

// 上传前处理
const beforeUpload = async (file) => {
  const isImage = file.type.startsWith('image/')
  if (!isImage) {
    message.error('只能上传图片文件!')
    return false
  }
  const isLt2M = file.size / 1024 / 1024 < 2
  if (!isLt2M) {
    message.error('图片大小不能超过 2MB!')
    return false
  }
  try {
    const formData = new FormData()
    formData.append('file', file)
    const res = await uploadImage(formData)
    handleUploadSuccess(res)
    return false
  } catch (error) {
    message.error('上传失败')
    return false
  }
}

// 模态框确认
const handleModalOk = async () => {
  try {
    await formRef.value.validate()
    modalLoading.value = true
    if (formState.value.id) {
      await updateProduct(formState.value.id, formState.value)
      message.success('更新成功')
    } else {
      await createProduct(formState.value)
      message.success('添加成功')
    }
    modalVisible.value = false
    fetchProducts()
  } catch (error) {
    console.error('保存失败:', error)
  } finally {
    modalLoading.value = false
  }
}

// 模态框取消
const handleModalCancel = () => {
  modalVisible.value = false
  formRef.value?.resetFields()
}

// 删除商品
const handleDelete = async (id) => {
  try {
    await deleteProduct(id)
    message.success('删除成功')
    fetchProducts()
  } catch (error) {
    console.error('删除失败:', error)
  }
}

onMounted(() => {
  fetchProducts()
  fetchCategories()
})
</script>

<style scoped>
.product-list {
  padding: 24px;
}

.header {
  display: flex;
  justify-content: space-between;
  margin-bottom: 16px;
}

.avatar-uploader {
  text-align: center;
}

.avatar-uploader img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.ant-upload-text {
  margin-top: 8px;
  color: #666;
}

.danger {
  color: #ff4d4f;
}
</style> 

