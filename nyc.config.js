module.exports = {
    extends: '@istanbuljs/nyc-config-typescript',
    'check-coverage': true,
    all: true,
    include: [
        'deployment/**/!(*.test.*).[tj]s?(x)',
        'client/**/!(*.test.*).[tj]s?(x)',
    ],
    exclude: [],
    reporter: ['html', 'lcov', 'text', 'text-summary'],
    'report-dir': 'coverage',
    lines: 40,
}
