import Link from 'next/link';
import { buttonVariants } from '../ui/button';

export function Sidebar() {
  return (
    <nav className="flex flex-col w-64 h-fit overflow-hidden rounded border">
      {pageLinks.map((link) => (
        <Link
          key={link.name}
          className={buttonVariants({
            variant: 'ghost',
            className: 'w-full justify-start not-last:border-b rounded-none',
          })}
          href={link.href}
        >
          {link.name}
        </Link>
      ))}
    </nav>
  );
}

const pageLinks = [
  {
    name: 'Home',
    href: '/',
  },
  {
    name: 'Movies',
    href: '/movies',
  },
  {
    name: 'TV Shows',
    href: '/tv-shows',
  },
  {
    name: 'Watchlist',
    href: '/watchlist',
  },
];
