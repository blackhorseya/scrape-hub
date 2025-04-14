import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  env: {
    AUTH0_AUDIENCE: process.env.AUTH0_AUDIENCE,
    NEXT_PUBLIC_API_URL: process.env.NEXT_PUBLIC_API_URL,
  },
  auth: {
    authorizationParams: {
      audience: process.env.AUTH0_AUDIENCE,
      scope: 'openid profile email',
    },
  },
};

export default nextConfig;
