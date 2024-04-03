import type { Metadata } from "next";
import { Montserrat } from "next/font/google";
import "./globals.css";
import StarsCanvas from "@/components/StarBackground";

const montserrat = Montserrat({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "Wiki Race",
  description: "Tubes 2 Strategi Algoritma",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className={`${montserrat.className} bg-[#030014]`}>
        <StarsCanvas />
        {children}
      </body>
    </html>
  );
}
