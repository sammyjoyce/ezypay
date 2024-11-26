import { reactRouter } from "@react-router/dev/vite";
import tailwindcss from '@tailwindcss/vite';
import { defineConfig } from "vite";
import tsconfigPaths from "vite-tsconfig-paths";

export default defineConfig({
  plugins: [tailwindcss(), reactRouter(), tsconfigPaths()],
  resolve: {
    extensions: ['.ts', '.js', '.tsx', '.jsx'],
  },
  server: {
    port: 3000,
  },
  define: {
    APP_NAME: JSON.stringify("{{project_name}}"),
    APP_DESCRIPTION: JSON.stringify("{{project_description}}"),
  },
});
