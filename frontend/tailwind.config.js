/** @type {import('tailwindcss').Config} */
export default {
  content: ['./src/**/*.{html,js,svelte,ts}'],
  theme: {
    extend: {
      colors: {
        'primary': 'var(--color-primary)',
        'secondary': 'var(--color-secondary)',
        'secondary-light': 'var(--color-secondary-light)',
        'warning': 'var(--color-warning)',
        'buttons': 'var(--color-buttons)',
        'nav-icon': 'var(--color-nav-icon)',
        'vote-star': 'var(--color-vote-star)',
        'chat-bubble-me': 'var(--color-chat-bubble-me)',
        'chat-bubble-you': 'var(--color-chat-bubble-you)',
        'typography-header': 'var(--color-typography-header)',
        'typography-main': 'var(--color-typography-main)',
        'typography-button': 'var(--color-typography-button)',
        'typography-chat-me': 'var(--color-typography-chat-me)',
        'typography-chat-you': 'var(--color-typography-chat-you)',
        'typography-nav': 'var(--color-typography-nav)',
        'typography-grey': 'var(--color-typography-grey)',
        'canvas-primary': 'var(--color-canvas-primary)',
        'canvas-secondary': 'var(--color-canvas-secondary)',
      },
      fontFamily: {
        sans: ['OpenMojiFont', 'sans-serif'],
        title: ['Edo', 'sans-serif']
      }
    },
  },
  plugins: [],
  darkMode: "class"
}

