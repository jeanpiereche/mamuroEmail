/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./index.html", "./src/**/*.{vue,js,ts,jsx,tsx}"],
  theme: {
    extend: {
      colors: {
        "primary-color" : "#00668A",
        "secondary-color" : "#004E71",
        "color-table": "#86b3f0",
      }
    },
    fontFamily:{
      Roboto : ["Roboto, sans-serif"],
    },
    container: {
      padding: "2rem",
      center: true,
    },
    screens:{
      sm: "640px",
      md: "768px",
    },
  },
  plugins: [],
}
