import Image from "next/image";
import styles from "./page.module.css";
import Header from "@/components/Header/Header"

import HuntMainPage from "@/components/HuntMainPage/huntMainPage";




export default function Home() {
  
    return (
      <>
        <Header/>
        <HuntMainPage  />
      </>
  );
}
