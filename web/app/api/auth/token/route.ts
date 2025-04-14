import { getAccessToken } from '@auth0/nextjs-auth0';
import { NextResponse } from 'next/server';

export async function GET() {
  try {
    const { accessToken } = await getAccessToken();
    return NextResponse.json({ accessToken });
  } catch (error) {
    console.error('Error fetching access token:', error);
    return NextResponse.json(
      { error: '無法獲取 access token' },
      { status: 500 }
    );
  }
}