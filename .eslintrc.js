module.exports = {
    parser: '@typescript-eslint/parser',
    parserOptions: {
        ecmaVersion: 'latest',
        sourceType: 'module',
        project: 'tsconfig.json',
    },
    plugins: ['@typescript-eslint', 'import'],
    extends: ['plugin:@typescript-eslint/strict-type-checked', 'prettier'],
    root: true,
    env: {
        node: true,
    },
    ignorePatterns: [
        '.eslintrc.js',
        'nyc.config.js',
        'coverage',
        'artifacts',
        'visualizer',
    ],
    rules: {
        '@typescript-eslint/restrict-template-expressions': [
            'error',
            {
                allowNumber: true,
                allowBoolean: true,
            },
        ],
        '@typescript-eslint/explicit-function-return-type': 'warn',
        '@typescript-eslint/explicit-module-boundary-types': 'warn',
        '@typescript-eslint/no-explicit-any': 'warn',
        '@typescript-eslint/no-unused-vars': 'warn',
        '@typescript-eslint/consistent-type-imports': 'warn',
        '@typescript-eslint/no-extraneous-class': 'warn',
        '@typescript-eslint/require-await': 'warn',
        '@typescript-eslint/no-floating-promises': 'warn',
        '@typescript-eslint/no-unsafe-member-access': 'warn',
        'import/order': [
            'error',
            {
                groups: ['builtin', 'external', 'parent', 'sibling', 'index'],
                pathGroups: [
                    {
                        pattern: '@custom-lib/**',
                        group: 'external',
                        position: 'after',
                    },
                ],
                pathGroupsExcludedImportTypes: ['builtin'],
                alphabetize: {
                    order: 'asc',
                },
                'newlines-between': 'always',
            },
        ],
        'sort-imports': [
            'error',
            {
                allowSeparatedGroups: true,
                ignoreDeclarationSort: true,
            },
        ],
        'no-multiple-empty-lines': [
            'error',
            {
                max: 1,
                maxEOF: 0,
                maxBOF: 0,
            },
        ],
        'import/first': 'error',
        'import/newline-after-import': 'error',
    },
}
