module.exports = {
    env: {
        browser: true,
        node: true,
        'vue/setup-compiler-macros': true,
    },
    parser: 'vue-eslint-parser',
    parserOptions: {
        parser: {
            js: 'espree',
            ts: '@typescript-eslint/parser',
            '<template>': 'espree',
        },
        requireConfigFile: false,
    },
    extends: [
        // Import resolver
        'plugin:import/typescript',
        // Vue 3
        'plugin:vue/vue3-recommended',
        // General config
        'uniscope',
    ],
    settings: {
        'import/parsers': {
            '@typescript-eslint/parser': ['.ts', '.tsx'],
        },

        'import/resolver': {
            alias: [
                ['@', './src'],
                ['@/*', './src/*'],
            ],
            node: {
                extensions: ['.js', '.jsx', '.ts', '.tsx', '.vue', '.d.ts'],
            },
            typescript: {
                alwaysTryTypes: true, // always try to resolve types under `<root>@types` directory even it doesn't contain any source code, like `@types/unist`
                // use <root>/path/to/folder/tsconfig.json
                project: __dirname,
            },
        },
    },
    // add your custom rules here
    rules: {
        // General
        'vue/component-tags-order': [
            'error',
            {
                order: ['script', 'template', 'style'],
            },
        ],
        // Vue template
        'vue/html-indent': [
            'error',
            4,
            {
                attribute: 1,
                baseIndent: 1,
                closeBracket: 0,
                alignAttributesVertically: true,
                ignores: [],
            },
        ],
        // Hack to deal with prettier styling
        'vue/max-attributes-per-line': [
            'error',
            {
                singleline: {
                    max: 20,
                },
                multiline: {
                    max: 1,
                },
            },
        ],
        'vue/singleline-html-element-content-newline': 'off',
        'vue/html-self-closing': [
            'error',
            {
                html: {
                    void: 'always',
                    normal: 'never',
                    component: 'always',
                },
            },
        ],
        'vue/block-lang': [
            'error',
            {
                script: { lang: 'ts' },
                style: { lang: 'scss' },
                template: { lang: 'html' },
            },
        ],
        // Vue component naming
        'vue/component-name-in-template-casing': ['error', 'PascalCase'],
        'vue/component-definition-name-casing': ['error', 'PascalCase'],
        // Fix script setup macros errors
        'vue/script-setup-uses-vars': 'error',
    },
};
