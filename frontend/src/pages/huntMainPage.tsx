"use client"
import { useState, useEffect } from "react";
import { ApiResponse } from "../../types/api";
import styles from './huntMainPage.module.css';


export default function HuntMainPage() {
    const [hunts, setHunts] = useState<ApiResponse>([])

    useEffect(()=> {
        async function fetchHunts(): Promise<void>{
            try{
                let req = await fetch("http://localhost:8080/scavengerhunts")
                let res: ApiResponse = await req.json()
                setHunts(res)
            }
            catch(err) {
                console.log(err)
            }
        }
        fetchHunts()
    },[])

    const jsxHunts = hunts.map(hunt => (
        <div className={styles.hunt} key={hunt.hunt_id}>
            <span className={styles.title}>{hunt.title}</span>
            <span className={styles.description}>{hunt.description}</span>
            <span className={styles.creator}>{hunt.created_by}</span>
        </div>))
        // console.log(hunt.hunt_id, Object.values(hunt)))
    
    const renderable = hunts[0] != undefined ? <div className={styles.hunts}>{jsxHunts}</div> : <div>Chill bros!</div>

    return(<>{renderable}</>)
}