'use client';

import { buttonVariants } from '@/components/ui/button';
import Link from 'next/link';

export default function GlobalError({ error }: { error: Error }) {
  return (
    <div className="flex flex-col items-center justify-center h-full gap-4">
      <h2>{error.message}</h2>
      <Link href="/" className={buttonVariants({})}>
        Return Home
      </Link>
    </div>
  );
}
