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
    signed: false,
  },
  {
    name: 'Movies',
    href: '/movies',
    signed: false,
  },
  {
    name: 'Genres',
    href: '/genres',
    signed: false,
  },
  {
    name: 'Add Movie',
    href: '/add-movie',
    signed: true,
  },
  {
    name: 'Manage Catalogs',
    href: '/manage-catalogs',
    signed: true,
  },
  {
    name: 'GraphQL',
    href: '/graphql',
    signed: true,
  },
];
