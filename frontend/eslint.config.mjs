// @ts-check
import sqConfig from '@akasection/eslint-config-sq/dist/nuxt.mjs'
import withNuxt from './.nuxt/eslint.config.mjs'

export default withNuxt(
  // Your custom configs here
  sqConfig,
  {
    ignores: [
      '**/dist/**',
      '**/node_modules/**',
      '.data/**',
      '.nuxt/**',
      '**/coverage/',
      'content.config.ts',
    ],
    rules: {
      'vue/max-attributes-per-line': ['warn', { singleline: 8, multiline: 1 }],
    },
  }
)
