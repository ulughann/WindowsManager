/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      "corporate": {
        "100": "#f5f5f5",
        "200": "#ebebeb",
        "300": "#e0e0e0",
      }
    },
  },
  plugins: [require("daisyui")],
  daisyui: {
    themes: [
      "light",
      "dark",
      "cupcake",
      "bumblebee",
      "emerald",
      "corporate",
      "synthwave",
      "retro",
      "cyberpunk",
      "valentine",
      "halloween",
      "garden",
      "forest",
      "aqua",
      "lofi",
      "pastel",
      "fantasy",
      "wireframe",
      "black",
      "luxury",
      "dracula",
      "cmyk",
      "autumn",
      "business",
      "acid",
      "lemonade",
      "night",
      "coffee",
      "winter",
      {
        nord: {
          "primary": "#5E81AC",
          "secondary": "#88C0D0",
          "accent": "#5E81AC",
          "neutral": "#3B4256",
          "base-100": "#2E3440",
          "info": "#B48EAD",
          "success": "#A3BE8C",
          "warning": "#EBCB8B",
          "error": "#BF616A",
        },
      },

    ],
  },
}