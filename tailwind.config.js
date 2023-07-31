/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./views/**/*.{html,js}"],
  theme: {
    fontFamily: {
      sans: 'Avenir, -apple-system, BlinkMacSystemFont, Segoe UI, Roboto, Oxygen, Ubuntu, Cantarell, Fira Sans, Droid Sans, Helvetica Neue, sans-serif',
    },
    container: {
      center: true,
      padding: '0.5rem',
      screens: {
        xs: '414px',
        sm: '640px',
        md: '728px',
        lg: '984px',
        xl: '1240px',
        '2xl': '1440px',
      },
    },
    colors: {
      primary: "#7f56d9",
      red: '#F55B5D',
      blue: '#4166f5',
      gray: {
        lightest: "#f8f8f7",
        lighter: "#F2F2F1",
        light: "#f3f3f3",
        DEFAULT: "#BEBEC1",
        dark: "#4F4F54",
        darkest: "#3E3E40",
      },
      white: "#FFFFFF",
      light: "#F5F5F5",
      yellow: {
        calm: "#F9D783",
        DEFAULT: "#F7E67B",
      },
      sand: {
        light: "#F7F6F4",
        DEFAULT: "#FAF4EB",
        dark: "#CAA073",
      },
      black: "#1D1D1B",
    },
    fontSize: {
      sm: "0.75rem",
      base: "1rem",
      "lg": "1.125rem",
      "3xl": "1.250rem",
      "4xl": "1.500rem",
      "5xl": "2rem",
      "6xl": "2.625rem",
      "7xl": "4.500rem",
    },
    lineHeight: {
      sm: "1.500rem",
      base: "2rem",
      "5xl": "2.500rem",
      "6xl": "3.125rem",
      "7xl": "4rem",
    },
    extend: {},
  },
  plugins: [],
}
