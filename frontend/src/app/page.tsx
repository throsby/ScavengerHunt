"use client"
import Image from "next/image";
import styles from "./page.module.css";
import Header from "@/components/Header/Header"
import ScavengerHunt from "@/components/ScavengerHunt/ScavengerHunt";
import HuntMainPage from "@/components/HuntMainPage/huntMainPage";

import {
  QueryClient,
  QueryClientProvider,
} from '@tanstack/react-query'

const queryClient = new QueryClient()

export default function Home() {
  
    return (
      <>
        <QueryClientProvider client={queryClient}>
          <Header/>
          <HuntMainPage  />
          <ScavengerHunt />
        </QueryClientProvider>
      </>
  );
}
