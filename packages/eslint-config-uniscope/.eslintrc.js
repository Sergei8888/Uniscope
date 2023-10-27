module.exports = {
    env: {
        node: true,
    },
    root: true,
    plugins: [
        'turbo',
        'prettier',
        '@typescript-eslint/eslint-plugin',
        'import',
    ],
    extends: [
        // JS base
        'eslint:recommended',
        // TS
        'plugin:@typescript-eslint/recommended',
        // Prettier
        'prettier',
    ],
    rules: {
        'import/order': [
            'error',
            {
                pathGroups: [
                    {
                        pattern: '@uniscope/shared',
                        group: 'sibling',
                        position: 'before',
                    },
                ],

                'newlines-between': 'always-and-inside-groups',

                groups: [
                    'builtin',
                    'external',
                    ['parent', 'internal', 'sibling'],
                    'index',
                    'type',
                    'object',
                ],

                warnOnUnassignedImports: true,
            },
        ],
        'import/no-default-export': 'error',
        // Typescript
        '@typescript-eslint/explicit-member-accessibility': [
            'error',
            {
                accessibility: 'explicit',
                overrides: {
                    accessors: 'explicit',
                    constructors: 'no-public',
                    methods: 'explicit',
                    properties: 'off',
                    parameterProperties: 'explicit',
                },
            },
        ],
        '@typescript-eslint/no-explicit-any': 'off',
        // == and != restrictions
        eqeqeq: 1,
        // Production rules
        'no-debugger': process.env.NODE_ENV === 'production' ? 'error' : 'off',
        // Comment formatting
        'spaced-comment': ['error', 'always', { exceptions: ['-', '+'] }],
        // Fixing false positive end of line error after git
        'prettier/prettier': [
            'error',
            {
                endOfLine: 'auto',
            },
        ],
    },
};
