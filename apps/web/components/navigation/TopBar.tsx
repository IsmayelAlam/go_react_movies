import Link from 'next/link';
import { Button } from '../ui/button';

export function TopBar() {
  return (
    <header className="flex items-center justify-between">
      <Link className="text-2xl font-bold" href={'/'}>
        Movies
      </Link>
      <Button variant="outline" className="mr-2">
        Sign In
      </Button>
    </header>
  );
}
