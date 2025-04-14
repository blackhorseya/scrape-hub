import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  env: {
    AUTH0_AUDIENCE: process.env.AUTH0_AUDIENCE,
  },
  auth: {
    authorizationParams: {
      audience: process.env.AUTH0_AUDIENCE,
      scope: 'openid profile email',
    },
  },
};

export default nextConfig;
