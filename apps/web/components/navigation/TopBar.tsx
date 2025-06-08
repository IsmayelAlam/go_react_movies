import Link from 'next/link';
import { buttonVariants } from '../ui/button';

export function TopBar() {
  return (
    <header className="flex items-center justify-between">
      <Link className="text-2xl font-bold" href={'/'}>
        Movies
      </Link>
      <Link
        className={buttonVariants({ variant: 'outline' })}
        href={'/sign-in'}
      >
        Sign In
      </Link>
    </header>
  );
}
