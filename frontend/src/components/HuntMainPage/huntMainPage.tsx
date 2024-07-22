"use client"
import { useState, useEffect, SyntheticEvent } from "react";
import { ApiResponse } from "../../../types/api";
import styles from './huntMainPage.module.css';
import { Hunt } from "../../../types/models";


export default function HuntMainPage() {
    const [hunts, setHunts] = useState<ApiResponse>([])
    const [selectedHuntId, setSelectedHuntId] = useState<number | null>(null)

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

    function handleClick(e: SyntheticEvent, respective_hunt_id: number) {
        console.log(e.currentTarget)
        console.log(respective_hunt_id)

        setSelectedHuntId(respective_hunt_id)
    }
    
    const jsxHunts = hunts.map((hunt) => (
        <div style={{ "backgroundImage": `linear-gradient(lightslategrey, rgba(0, 0, 0, 0)), url(/output-${hunt.hunt_id}.jpg)` } as React.CSSProperties} onClick={(event) => { handleClick(event, hunt.hunt_id) }} className={`${styles.hunt} ${hunt.hunt_id === selectedHuntId ? styles.chosen : ""}`} key={hunt.hunt_id}>
            {hunt.hunt_id !== selectedHuntId && (
            <>
                <span className={styles.title}>{hunt.title}</span>
                <span className={styles.description}>{hunt.description}</span>
                <span className={styles.creator}>{hunt.created_by}</span>
            </>
            )}
        </div>))
    
    const renderable = hunts.length > 0 ? <div className={styles.hunts}>{jsxHunts}</div> : <div>Chill bros!</div>

    return(<>{renderable}</>)
}