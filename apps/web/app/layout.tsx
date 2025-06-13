'use client';

import './globals.css';

import { Sidebar, TopBar } from '@/components/navigation';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';

const queryClient = new QueryClient();

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className="max-w-7xl mx-auto p-8 flex flex-col gap-4 h-screen">
        <QueryClientProvider client={queryClient}>
          <TopBar />
          <div className="flex w-full h-full gap-4">
            <Sidebar />
            <main className="border flex-1 rounded-md py-2 px-4">
              {children}
            </main>
          </div>
        </QueryClientProvider>
      </body>
    </html>
  );
}
