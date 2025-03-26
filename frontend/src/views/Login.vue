<template>
  <div class="login-container">
    <div class="login-box">
      <h2>登录</h2>
      <a-form
        :model="formState"
        name="login"
        @finish="onFinish"
        autocomplete="off"
      >
        <a-form-item
          name="username"
          :rules="[{ required: true, message: '请输入用户名!' }]"
          :validate-status="formState.username ? 'success' : ''"
        >
          <a-input v-model:value="formState.username" placeholder="用户名">
            <template #prefix>
              <UserOutlined />
            </template>
          </a-input>
        </a-form-item>

        <a-form-item
          name="password"
          :rules="[{ required: true, message: '请输入密码!' }]"
          :validate-status="formState.password ? 'success' : ''"
        >
          <a-input-password v-model:value="formState.password" placeholder="密码">
            <template #prefix>
              <LockOutlined />
            </template>
          </a-input-password>
        </a-form-item>

        <a-form-item>
          <a-button type="primary" html-type="submit" :loading="loading" block>
            登录
          </a-button>
        </a-form-item>

        <div class="register-link">
          <a-button type="link" @click="goToRegister" block>
            还没有账号？立即注册
          </a-button>
        </div>
      </a-form>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { UserOutlined, LockOutlined } from '@ant-design/icons-vue'
import { message } from 'ant-design-vue'
import { useRouter } from 'vue-router'
import { login } from '@/api/user'

const router = useRouter()
const loading = ref(false)

const formState = reactive({
  username: '',
  password: '',
})

const onFinish = async (values) => {
  loading.value = true
  try {
    const res = await login(values)
    localStorage.setItem('token', res.token)
    message.success('登录成功')
    router.push('/')
  } catch (error) {
    console.error('登录失败:', error)
    message.error(error.response?.data?.message || '登录失败')
  } finally {
    loading.value = false
  }
}

const goToRegister = () => {
  router.push('/register')
}
</script>

<style scoped>
.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  background-color: #f0f2f5;
}

.login-box {
  width: 100%;
  max-width: 400px;
  padding: 40px;
  background: white;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.15);
}

h2 {
  text-align: center;
  margin-bottom: 24px;
  color: #1890ff;
}

.register-link {
  text-align: center;
  margin-top: 16px;
}
</style> 
