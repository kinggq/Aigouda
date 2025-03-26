<template>
  <div class="register-container">
    <div class="register-box">
      <h2>注册</h2>
      <a-form
        :model="formState"
        name="register"
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
          :rules="[
            { required: true, message: '请输入密码!' },
            { min: 6, message: '密码长度不能小于6位!' }
          ]"
          :validate-status="formState.password ? 'success' : ''"
        >
          <a-input-password v-model:value="formState.password" placeholder="密码">
            <template #prefix>
              <LockOutlined />
            </template>
          </a-input-password>
        </a-form-item>

        <a-form-item
          name="confirmPassword"
          :rules="[
            { required: true, message: '请确认密码!' },
            { validator: validateConfirmPassword }
          ]"
          :validate-status="formState.confirmPassword ? 'success' : ''"
        >
          <a-input-password v-model:value="formState.confirmPassword" placeholder="确认密码">
            <template #prefix>
              <LockOutlined />
            </template>
          </a-input-password>
        </a-form-item>

        <a-form-item>
          <a-button type="primary" html-type="submit" :loading="loading" block>
            注册
          </a-button>
        </a-form-item>

        <div class="login-link">
          <a-button type="link" @click="goToLogin" block>
            已有账号？立即登录
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
import { register } from '@/api/user'

const router = useRouter()
const loading = ref(false)

const formState = reactive({
  username: '',
  password: '',
  confirmPassword: '',
})

const validateConfirmPassword = async (rule, value) => {
  if (value !== formState.password) {
    throw new Error('两次输入的密码不一致!')
  }
  return Promise.resolve()
}

const onFinish = async (values) => {
  loading.value = true
  try {
    await register({
      username: values.username,
      password: values.password,
    })
    message.success('注册成功，请登录')
    router.push('/login')
  } catch (error) {
    console.error('注册失败:', error)
  } finally {
    loading.value = false
  }
}

const goToLogin = () => {
  router.push('/login')
}
</script>

<style scoped>
.register-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  background-color: #f0f2f5;
}

.register-box {
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

.login-link {
  text-align: center;
  margin-top: 16px;
}
</style> 
