/// <reference types="vite/client" />
interface ImportMetaEnv {
  readonly VITE_BASE_URL: string;
}

interface importMeta {
  readonly env: ImportMetaEnv;
}
