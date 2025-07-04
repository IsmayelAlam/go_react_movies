'use server';

import { backendUrl } from '@/lib/constance';

export async function signinAction(formData: FormData) {
  const email = formData.get('email') as string;
  const password = formData.get('password') as string;

  if (!email || !password) {
    throw new Error('Email and password are required');
  }

  try {
    const response = await fetch(`${backendUrl}/auth/login`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      credentials: 'include',
      body: JSON.stringify({ email, password }),
    });
    if (!response.ok) throw new Error('Failed to sign in');

    const data = await response.json();
    console.log('Sign-in successful:', data);
    return data;
  } catch (error) {
    throw new Error(
      `Sign-in error: ${error instanceof Error ? error.message : 'Unknown error'}`
    );
  }
}
// admin@example.com
// secret
