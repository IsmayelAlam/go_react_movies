import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Label } from '@radix-ui/react-label';
import React from 'react';

export default function Auth() {
  return (
    <form className="space-y-4 max-w-md mx-auto mt-10 p-6">
      <div className="flex flex-col gap-1">
        <Label htmlFor="email">Email</Label>
        <Input id="email" />
      </div>
      <div className="flex flex-col gap-1">
        <Label htmlFor="password">Password</Label>
        <Input id="password" />
      </div>
      <Button>Sign In</Button>
    </form>
  );
}
