<script setup lang="ts">
import { reactive, ref } from 'vue';
import { AuthSignUpDtoI } from '@uniscope/shared/vite';
import { ElNotification, FormInstance, FormRules } from 'element-plus';

import { useRouter } from 'vue-router';

import { useAuthStore } from '@/auth/model/auth.store';

const authStore = useAuthStore();
const router = useRouter();

const signupFormRef = ref<FormInstance>();
const signupForm = reactive<AuthSignUpDtoI>({
    nickname: '',
    login: '',
    password: '',
    passwordRepeated: '',
});
const signupFormRules = reactive<FormRules<AuthSignUpDtoI>>({
    nickname: [
        {
            required: true,
            message: 'Введите имя пользователя',
            trigger: 'blur',
        },
        {
            min: 3,
            message: 'Имя пользователя должно быть не менее 3 символов',
            trigger: 'blur',
        },
        {
            max: 20,
            message: 'Имя пользователя должно быть не более 20 символов',
            trigger: 'blur',
        },
    ],
    login: [
        { required: true, message: 'Введите логин', trigger: 'blur' },
        {
            min: 6,
            message: 'Логин должен быть не менее 6 символов',
            trigger: 'blur',
        },
        {
            max: 20,
            message: 'Логин должен быть не более 20 символов',
            trigger: 'blur',
        },
    ],
    password: [
        { required: true, message: 'Введите пароль', trigger: 'blur' },
        {
            min: 6,
            message: 'Пароль должен быть не менее 6 символов',
            trigger: 'blur',
        },
        {
            max: 20,
            message: 'Пароль должен быть не более 20 символов',
            trigger: 'blur',
        },
        {
            validator: (rule, value, callback) => {
                new RegExp(
                    /((?=.*\d)|(?=.*\W+))(?![.\n])(?=.*[A-Z])(?=.*[a-z]).*$/
                ).test(value)
                    ? callback()
                    : callback(new Error('Пароль слишком простой'));
            },
        },
    ],
    passwordRepeated: [
        { required: true, message: 'Повторите пароль', trigger: 'blur' },
        {
            min: 6,
            message: 'Пароль должен быть не менее 6 символов',
            trigger: 'blur',
        },
        {
            max: 20,
            message: 'Пароль должен быть не более 20 символов',
            trigger: 'blur',
        },
        {
            validator: (rule, value, callback) => {
                if (value !== signupForm.password) {
                    callback(new Error('Пароли не совпадают'));
                } else {
                    callback();
                }
            },
            trigger: 'blur',
        },
    ],
});

async function submitForm(formEl: FormInstance | undefined) {
    if (!formEl) return;
    await formEl.validate(async (valid) => {
        if (valid) {
            try {
                await authStore.signUp(signupForm);
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
        ref="signupFormRef"
        :model="signupForm"
        :rules="signupFormRules"
        class="signup-form"
        label-suffix=":"
    >
        <ElFormItem>
            <h2 class="signup-form__header">Форма регистрации</h2>
        </ElFormItem>
        <ElFormItem
            required
            class="signup-form__item"
            label="Имя пользователя"
            prop="nickname"
        >
            <ElInput v-model="signupForm.nickname" />
        </ElFormItem>
        <ElFormItem
            required
            class="signup-form__item"
            label="Логин"
            prop="login"
        >
            <ElInput v-model="signupForm.login" />
        </ElFormItem>
        <ElFormItem
            required
            class="signup-form__item"
            label="Пароль"
            prop="password"
        >
            <ElInput v-model="signupForm.password" />
        </ElFormItem>
        <ElFormItem
            required
            class="signup-form__item"
            label="Повторите пароль"
            prop="passwordRepeated"
        >
            <ElInput v-model="signupForm.passwordRepeated" />
        </ElFormItem>
        <div class="signup-form__item signup-form__button">
            <ElButton type="primary" @click.prevent="submitForm(signupFormRef)">
                Зарегистрироваться
            </ElButton>
        </div>
        <router-link
            :to="{ name: 'signin' }"
            class="signup-form__item signup-form__link el-link el-link--info"
        >
            Уже есть аккаунт? Войти
        </router-link>
    </ElForm>
</template>

<style scoped lang="scss">
.signup-form {
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
