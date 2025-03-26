<template>
  <div class="order-list">
    <div class="header">
      <a-space>
        <a-input-search
          v-model:value="searchForm.orderNo"
          placeholder="搜索订单号"
          style="width: 200px"
          @search="onSearch"
        />
        <a-select
          v-model:value="searchForm.status"
          placeholder="订单状态"
          style="width: 200px"
          allowClear
          @change="onSearch"
        >
          <a-select-option value="pending">待付款</a-select-option>
          <a-select-option value="paid">已付款</a-select-option>
          <a-select-option value="shipped">已发货</a-select-option>
          <a-select-option value="completed">已完成</a-select-option>
          <a-select-option value="cancelled">已取消</a-select-option>
        </a-select>
        <a-range-picker
          v-model:value="searchForm.dateRange"
          :show-time="{ format: 'HH:mm' }"
          format="YYYY-MM-DD HH:mm"
          placeholder="['开始时间', '结束时间']"
          @change="onSearch"
        />
      </a-space>
    </div>

    <a-table
      :columns="columns"
      :data-source="orders"
      :loading="loading"
      :pagination="pagination"
      @change="handleTableChange"
    >
      <template #bodyCell="{ column, record }">
        <template v-if="column.key === 'status'">
          <a-tag :color="getOrderStatusColor(record.status)">
            {{ getOrderStatusText(record.status) }}
          </a-tag>
        </template>
        <template v-else-if="column.key === 'action'">
          <a-space>
            <a @click="showDetailModal(record)">详情</a>
            <a-divider type="vertical" />
            <template v-if="record.status === 'paid'">
              <a @click="handleShip(record)">发货</a>
            </template>
          </a-space>
        </template>
      </template>
    </a-table>

    <!-- 订单详情弹窗 -->
    <a-modal
      v-model:visible="detailModalVisible"
      title="订单详情"
      width="800px"
      :footer="null"
    >
      <a-descriptions :column="2" bordered>
        <a-descriptions-item label="订单号">
          {{ currentOrder.orderNo }}
        </a-descriptions-item>
        <a-descriptions-item label="订单状态">
          <a-tag :color="getOrderStatusColor(currentOrder.status)">
            {{ getOrderStatusText(currentOrder.status) }}
          </a-tag>
        </a-descriptions-item>
        <a-descriptions-item label="下单时间">
          {{ currentOrder.createdAt }}
        </a-descriptions-item>
        <a-descriptions-item label="支付时间">
          {{ currentOrder.paidAt || '-' }}
        </a-descriptions-item>
        <a-descriptions-item label="发货时间">
          {{ currentOrder.shippedAt || '-' }}
        </a-descriptions-item>
        <a-descriptions-item label="完成时间">
          {{ currentOrder.completedAt || '-' }}
        </a-descriptions-item>
        <a-descriptions-item label="收货人">
          {{ currentOrder.receiverName }}
        </a-descriptions-item>
        <a-descriptions-item label="联系电话">
          {{ currentOrder.receiverPhone }}
        </a-descriptions-item>
        <a-descriptions-item label="收货地址" :span="2">
          {{ currentOrder.receiverAddress }}
        </a-descriptions-item>
      </a-descriptions>

      <div class="order-items">
        <h3>商品信息</h3>
        <a-table
          :columns="itemColumns"
          :data-source="currentOrder.items"
          :pagination="false"
        >
          <template #bodyCell="{ column, record }">
            <template v-if="column.key === 'image'">
              <a-image
                :width="50"
                :src="record.product.images[0]"
                :fallback="defaultImage"
              />
            </template>
          </template>
        </a-table>
      </div>

      <div class="order-summary">
        <div class="summary-item">
          <span>商品总额：</span>
          <span>¥{{ currentOrder.totalAmount }}</span>
        </div>
        <div class="summary-item">
          <span>运费：</span>
          <span>¥{{ currentOrder.shippingFee }}</span>
        </div>
        <div class="summary-item total">
          <span>实付金额：</span>
          <span>¥{{ currentOrder.paidAmount }}</span>
        </div>
      </div>
    </a-modal>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { message } from 'ant-design-vue'
import { getOrders, shipOrder } from '@/api/order'

const columns = [
  {
    title: '订单号',
    dataIndex: 'orderNo',
    key: 'orderNo',
  },
  {
    title: '下单时间',
    dataIndex: 'createdAt',
    key: 'createdAt',
  },
  {
    title: '收货人',
    dataIndex: 'receiverName',
    key: 'receiverName',
  },
  {
    title: '订单金额',
    dataIndex: 'paidAmount',
    key: 'paidAmount',
  },
  {
    title: '订单状态',
    dataIndex: 'status',
    key: 'status',
  },
  {
    title: '操作',
    key: 'action',
    width: 150,
  },
]

const itemColumns = [
  {
    title: '商品图片',
    dataIndex: 'image',
    key: 'image',
    width: 80,
  },
  {
    title: '商品名称',
    dataIndex: ['product', 'title'],
    key: 'title',
  },
  {
    title: '单价',
    dataIndex: 'price',
    key: 'price',
  },
  {
    title: '数量',
    dataIndex: 'quantity',
    key: 'quantity',
  },
  {
    title: '小计',
    dataIndex: 'subtotal',
    key: 'subtotal',
  },
]

const loading = ref(false)
const orders = ref([])
const detailModalVisible = ref(false)
const currentOrder = ref({})
const defaultImage = 'https://via.placeholder.com/50'

const searchForm = reactive({
  orderNo: '',
  status: undefined,
  dateRange: [],
})

const pagination = reactive({
  current: 1,
  pageSize: 10,
  total: 0,
})

// 获取订单列表
const fetchOrders = async () => {
  try {
    loading.value = true
    const params = {
      page: pagination.current,
      pageSize: pagination.pageSize,
      orderNo: searchForm.orderNo,
      status: searchForm.status,
    }
    if (searchForm.dateRange.length === 2) {
      params.startTime = searchForm.dateRange[0].format('YYYY-MM-DD HH:mm:ss')
      params.endTime = searchForm.dateRange[1].format('YYYY-MM-DD HH:mm:ss')
    }
    const res = await getOrders(params)
    orders.value = res.data.orders
    pagination.total = res.data.total
  } catch (error) {
    message.error('获取订单列表失败')
  } finally {
    loading.value = false
  }
}

// 搜索
const onSearch = () => {
  pagination.current = 1
  fetchOrders()
}

// 表格变化
const handleTableChange = (pag) => {
  pagination.current = pag.current
  pagination.pageSize = pag.pageSize
  fetchOrders()
}

// 显示订单详情
const showDetailModal = (record) => {
  currentOrder.value = record
  detailModalVisible.value = true
}

// 处理发货
const handleShip = async (record) => {
  try {
    await shipOrder(record.id)
    message.success('发货成功')
    fetchOrders()
  } catch (error) {
    message.error('发货失败')
  }
}

const getOrderStatusColor = (status) => {
  const colors = {
    pending: 'orange',
    paid: 'blue',
    shipped: 'green',
    completed: 'purple',
    cancelled: 'red',
  }
  return colors[status] || 'default'
}

const getOrderStatusText = (status) => {
  const texts = {
    pending: '待付款',
    paid: '已付款',
    shipped: '已发货',
    completed: '已完成',
    cancelled: '已取消',
  }
  return texts[status] || status
}

onMounted(() => {
  fetchOrders()
})
</script>

<style scoped>
.order-list {
  padding: 24px;
}

.header {
  margin-bottom: 24px;
}

.order-items {
  margin-top: 24px;
}

.order-items h3 {
  margin-bottom: 16px;
}

.order-summary {
  margin-top: 24px;
  text-align: right;
}

.summary-item {
  margin-bottom: 8px;
}

.summary-item.total {
  font-size: 16px;
  font-weight: bold;
  color: #ff4d4f;
}
</style> 
