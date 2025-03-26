<template>
  <div class="dashboard">
    <a-row :gutter="16">
      <a-col :span="6">
        <a-card>
          <template #title>
            <span>
              <ShoppingOutlined />
              商品总数
            </span>
          </template>
          <div class="card-content">
            <h2>{{ statistics.productCount }}</h2>
          </div>
        </a-card>
      </a-col>
      <a-col :span="6">
        <a-card>
          <template #title>
            <span>
              <UserOutlined />
              用户总数
            </span>
          </template>
          <div class="card-content">
            <h2>{{ statistics.userCount }}</h2>
          </div>
        </a-card>
      </a-col>
      <a-col :span="6">
        <a-card>
          <template #title>
            <span>
              <OrderedListOutlined />
              订单总数
            </span>
          </template>
          <div class="card-content">
            <h2>{{ statistics.orderCount }}</h2>
          </div>
        </a-card>
      </a-col>
      <a-col :span="6">
        <a-card>
          <template #title>
            <span>
              <MoneyCollectOutlined />
              总销售额
            </span>
          </template>
          <div class="card-content">
            <h2>¥{{ statistics.totalSales }}</h2>
          </div>
        </a-card>
      </a-col>
    </a-row>

    <a-row :gutter="16" class="mt-16">
      <a-col :span="12">
        <a-card title="最近订单">
          <a-table
            :columns="orderColumns"
            :data-source="recentOrders"
            :pagination="false"
          >
            <template #bodyCell="{ column, record }">
              <template v-if="column.key === 'status'">
                <a-tag :color="getOrderStatusColor(record.status)">
                  {{ getOrderStatusText(record.status) }}
                </a-tag>
              </template>
            </template>
          </a-table>
        </a-card>
      </a-col>
      <a-col :span="12">
        <a-card title="热门商品">
          <a-table
            :columns="productColumns"
            :data-source="hotProducts"
            :pagination="false"
          >
            <template #bodyCell="{ column, record }">
              <template v-if="column.key === 'image'">
                <a-image
                  :width="50"
                  :src="record.images[0]"
                  :fallback="defaultImage"
                />
              </template>
            </template>
          </a-table>
        </a-card>
      </a-col>
    </a-row>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import {
  ShoppingOutlined,
  UserOutlined,
  OrderedListOutlined,
  MoneyCollectOutlined,
} from '@ant-design/icons-vue'

const statistics = ref({
  productCount: 0,
  userCount: 0,
  orderCount: 0,
  totalSales: 0,
})

const orderColumns = [
  {
    title: '订单号',
    dataIndex: 'orderNo',
    key: 'orderNo',
  },
  {
    title: '用户',
    dataIndex: 'username',
    key: 'username',
  },
  {
    title: '金额',
    dataIndex: 'amount',
    key: 'amount',
  },
  {
    title: '状态',
    dataIndex: 'status',
    key: 'status',
  },
  {
    title: '时间',
    dataIndex: 'createdAt',
    key: 'createdAt',
  },
]

const productColumns = [
  {
    title: '商品图片',
    dataIndex: 'image',
    key: 'image',
    width: 80,
  },
  {
    title: '商品名称',
    dataIndex: 'title',
    key: 'title',
  },
  {
    title: '价格',
    dataIndex: 'price',
    key: 'price',
  },
  {
    title: '销量',
    dataIndex: 'sales',
    key: 'sales',
  },
]

const recentOrders = ref([])
const hotProducts = ref([])
const defaultImage = 'https://via.placeholder.com/50'

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

onMounted(async () => {
  // TODO: 获取统计数据
  statistics.value = {
    productCount: 100,
    userCount: 50,
    orderCount: 200,
    totalSales: 50000,
  }

  // TODO: 获取最近订单
  recentOrders.value = [
    {
      orderNo: 'ORDER001',
      username: '张三',
      amount: 100,
      status: 'pending',
      createdAt: '2024-03-20 10:00:00',
    },
    // ... 更多订单数据
  ]

  // TODO: 获取热门商品
  hotProducts.value = [
    {
      id: 1,
      title: '商品1',
      price: 100,
      sales: 50,
      images: ['https://via.placeholder.com/50'],
    },
    // ... 更多商品数据
  ]
})
</script>

<style scoped>
.dashboard {
  padding: 24px;
}

.mt-16 {
  margin-top: 16px;
}

.card-content {
  text-align: center;
}

.card-content h2 {
  margin: 0;
  font-size: 24px;
  color: #1890ff;
}
</style> 
