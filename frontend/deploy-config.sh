cat >./src/configs/config.ts <<EOF
interface AppConfig {
    BACKEND_BASE_URL: string;
}

const config: AppConfig = {
    BACKEND_BASE_URL: "${SERVER_URL}"
}

export default config;
EOF