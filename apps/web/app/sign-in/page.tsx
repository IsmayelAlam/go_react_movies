import { signinAction } from '@/actions/signin';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Label } from '@radix-ui/react-label';
import React from 'react';

export default function Auth() {
  return (
    <form
      className="space-y-4 max-w-md mx-auto mt-10 p-6"
      action={signinAction}
    >
      <div className="flex flex-col gap-1">
        <Label htmlFor="email">Email</Label>
        <Input id="email" name="email" />
      </div>
      <div className="flex flex-col gap-1">
        <Label htmlFor="password">Password</Label>
        <Input id="password" name="password" />
      </div>
      <Button type="submit">Sign In</Button>
    </form>
  );
}
