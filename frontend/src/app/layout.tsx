import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";
import { Children } from "react";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "Scavenger Hunt",
  description: "United to support the medium of scavenger",
};

export default function RootLayout({ children, }: Readonly<{children: React.ReactNode;}>) {
  return (
    <html lang="en">
      <body className={inter.className}>{children}</body>
    </html>
  );
}
