"use client"
import { useQuery } from "@tanstack/react-query"


export default function ScavengerHunt(){

    const { isPending, error, data } = useQuery({
        queryKey: ['repoData'],
        queryFn: () => (
            fetch("http://localhost:8080/scavengerhunts/2").then((res) => res.json())      
            ),
    })
    
    console.log(isPending)
    console.log(data)





    return(<></>)
}