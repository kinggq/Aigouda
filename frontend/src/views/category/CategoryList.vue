<template>
  <div class="category-list">
    <div class="header">
      <a-input-search
        v-model:value="searchQuery"
        placeholder="搜索分类"
        style="width: 200px"
        @search="onSearch"
      />
      <a-button type="primary" @click="showAddModal">
        <plus-outlined />添加分类
      </a-button>
    </div>

    <a-table
      :columns="columns"
      :data-source="categories"
      :loading="loading"
      :pagination="pagination"
      @change="handleTableChange"
    >
      <template #bodyCell="{ column, record }">
        <template v-if="column.key === 'action'">
          <a-space>
            <a @click="showEditModal(record)">编辑</a>
            <a-divider type="vertical" />
            <a-popconfirm
              title="确定要删除这个分类吗？"
              @confirm="handleDelete(record.id)"
            >
              <a class="danger">删除</a>
            </a-popconfirm>
          </a-space>
        </template>
      </template>
    </a-table>

    <!-- 添加/编辑分类模态框 -->
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
        <a-form-item label="分类名称" name="name">
          <a-input v-model:value="formState.name" placeholder="请输入分类名称" />
        </a-form-item>
        <a-form-item label="父级分类" name="parent_id">
          <a-select
            v-model:value="formState.parent_id"
            placeholder="请选择父级分类"
            allowClear
          >
            <a-select-option v-for="category in categories" :key="category.id" :value="category.id">
              {{ category.name }}
            </a-select-option>
          </a-select>
        </a-form-item>
        <a-form-item label="描述" name="description">
          <a-textarea
            v-model:value="formState.description"
            :rows="4"
            placeholder="请输入分类描述"
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
import { getCategories, createCategory, updateCategory, deleteCategory } from '@/api/category'

const columns = [
  {
    title: '分类名称',
    dataIndex: 'name',
    key: 'name',
  },
  {
    title: '父级分类',
    dataIndex: ['parent', 'name'],
    key: 'parent',
  },
  {
    title: '描述',
    dataIndex: 'description',
    key: 'description',
  },
  {
    title: '操作',
    key: 'action',
  },
]

const categories = ref([])
const loading = ref(false)
const modalVisible = ref(false)
const modalLoading = ref(false)
const modalTitle = ref('添加分类')
const formRef = ref(null)

const searchQuery = ref('')
const pagination = reactive({
  current: 1,
  pageSize: 10,
  total: 0,
})

const formState = reactive({
  id: null,
  name: '',
  parent_id: undefined,
  description: '',
})

const rules = {
  name: [{ required: true, message: '请输入分类名称' }],
}

// 获取分类列表
const fetchCategories = async () => {
  loading.value = true
  try {
    const params = {
      page: pagination.current,
      page_size: pagination.pageSize,
      search: searchQuery.value,
    }
    const res = await getCategories(params)
    categories.value = res.data
    pagination.total = res.total
  } catch (error) {
    console.error('获取分类列表失败:', error)
  } finally {
    loading.value = false
  }
}

// 表格变化处理
const handleTableChange = (pag) => {
  pagination.current = pag.current
  pagination.pageSize = pag.pageSize
  fetchCategories()
}

// 搜索处理
const onSearch = () => {
  pagination.current = 1
  fetchCategories()
}

// 显示添加模态框
const showAddModal = () => {
  modalTitle.value = '添加分类'
  Object.assign(formState, {
    id: null,
    name: '',
    parent_id: undefined,
    description: '',
  })
  modalVisible.value = true
}

// 显示编辑模态框
const showEditModal = (record) => {
  modalTitle.value = '编辑分类'
  Object.assign(formState, record)
  modalVisible.value = true
}

// 模态框确认
const handleModalOk = async () => {
  try {
    await formRef.value.validate()
    modalLoading.value = true
    if (formState.id) {
      await updateCategory(formState.id, formState)
      message.success('更新成功')
    } else {
      await createCategory(formState)
      message.success('添加成功')
    }
    modalVisible.value = false
    fetchCategories()
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

// 删除分类
const handleDelete = async (id) => {
  try {
    await deleteCategory(id)
    message.success('删除成功')
    fetchCategories()
  } catch (error) {
    console.error('删除失败:', error)
  }
}

onMounted(() => {
  fetchCategories()
})
</script>

<style scoped>
.category-list {
  padding: 24px;
}

.header {
  display: flex;
  justify-content: space-between;
  margin-bottom: 16px;
}

.danger {
  color: #ff4d4f;
}
</style> 
