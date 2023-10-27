<script setup lang="ts">
import { reactive, ref } from 'vue';
import { AuthSignInDtoI } from '@uniscope/shared/vite';
import { ElNotification, FormInstance, FormRules } from 'element-plus';

import { useRouter } from 'vue-router';

import { useAuthStore } from '@/auth/model/auth.store';

const authStore = useAuthStore();
const router = useRouter();

const signinFormRef = ref<FormInstance>();
const signinForm = reactive<AuthSignInDtoI>({
    login: '',
    password: '',
});
const signinFormRules = reactive<FormRules<AuthSignInDtoI>>({
    login: [{ required: true, message: 'Введите логин', trigger: 'blur' }],
    password: [{ required: true, message: 'Введите пароль', trigger: 'blur' }],
});

async function submitForm(formEl: FormInstance | undefined) {
    if (!formEl) return;
    await formEl.validate(async (valid) => {
        if (valid) {
            try {
                await authStore.signIn(signinForm);
                formEl.resetFields();
                await router.push({ name: 'home' });
            } catch (error: any) {
                for (const errorMessage of error.response.data.message) {
                    ElNotification.error({
                        title: 'Ошибка',
                        message: errorMessage,
                        duration: 3000,
                    });
                }
            }
        }
    });
}
</script>

<template>
    <ElForm
        ref="signinFormRef"
        :model="signinForm"
        :rules="signinFormRules"
        class="signin-form"
        label-suffix=":"
    >
        <ElFormItem>
            <h2 class="signin-form__header">Форма входа</h2>
        </ElFormItem>
        <ElFormItem
            required
            class="signin-form__item"
            label="Логин"
            prop="login"
        >
            <ElInput v-model="signinForm.login" />
        </ElFormItem>
        <ElFormItem
            required
            class="signin-form__item"
            label="Пароль"
            prop="password"
        >
            <ElInput v-model="signinForm.password" />
        </ElFormItem>
        <div class="signin-form__item signin-form__button">
            <ElButton type="primary" @click.prevent="submitForm(signinFormRef)">
                Войти
            </ElButton>
        </div>
        <router-link
            :to="{ name: 'signup' }"
            class="signup-form__item signup-form__link el-link el-link--info"
        >
            Нет аккаунта? Зарегистрироваться
        </router-link>
    </ElForm>
</template>

<style scoped lang="scss">
.signin-form {
    display: flex;
    flex-direction: column;
    width: 400px;
    padding: 20px;
    border-radius: 7px;
    border: 1px solid var(--el-border-color);

    &__header {
        width: 100%;
        text-align: center;
        margin-bottom: 5px;
    }

    &__item {
        margin-bottom: 35px;
    }

    &__item:last-child {
        margin: 0 auto;
    }

    &__button {
        display: flex;
        justify-content: center;
        margin-bottom: 15px;
    }

    &__link {
        font-size: 12px;
    }
}
</style>
