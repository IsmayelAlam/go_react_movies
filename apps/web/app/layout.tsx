import './globals.css';

import { Sidebar, TopBar } from '@/components/navigation';
import type { Metadata } from 'next';

export const metadata: Metadata = {
  title: 'Watch Movie App',
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className="max-w-7xl mx-auto p-8 flex flex-col gap-4 h-screen">
        <TopBar />
        <div className="flex w-full h-full gap-4">
          <Sidebar />
          <main className="border flex-1 rounded-md py-2 px-4">{children}</main>
        </div>
      </body>
    </html>
  );
}
