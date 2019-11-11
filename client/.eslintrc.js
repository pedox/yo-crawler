module.exports = {
  'env': {
    'browser': true,
    'es6': true
  },
  'extends': [
    'eslint:recommended',
    'plugin:react/recommended'
  ],
  'globals': {
    'Atomics': 'readonly',
    'SharedArrayBuffer': 'readonly'
  },
  'parserOptions': {
    'ecmaFeatures': {
    'jsx': true
    },
    'ecmaVersion': 2018,
    'sourceType': 'module'
  },
  'plugins': [
    'react',
    'react-hooks'
  ],
  'rules': {
    'indent': [
      'error',
      2
    ],
    'object-curly-spacing': [
      'error',
      'always'
    ],
    'linebreak-style': [
      'error',
      'unix'
    ],
    'quotes': [
      'error',
      'single'
    ],
    'semi': [
      'error',
      'never'
    ],
    'prefer-arrow-callback': [
      'error'
    ],
    'no-duplicate-imports': [
      'error'
    ],
    'react/prop-types': [
      'warn'
    ],
    'react/display-name': [
      'off'
    ],
    'no-unused-vars': [
      'off'
    ],
    'react-hooks/exhaustive-deps': 'off'
  }
}
