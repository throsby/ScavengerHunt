"use client"
import { useState, useEffect } from "react";
import { ApiResponse } from "../../types/api";


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
        <div className="hunt" key={hunt.hunt_id}>
            <span className="title">{hunt.title}</span>
            <span className="description">{hunt.description}</span>
            <span className="creator">{hunt.created_by}</span>
        </div>))
        // console.log(hunt.hunt_id, Object.values(hunt)))
    
    const renderable = hunts[0] != undefined ? <div className="hunts">{jsxHunts}</div> : <div>Chill bros!</div>

    return(<>{renderable}</>)
}