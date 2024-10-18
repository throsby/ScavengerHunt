"use client"
import { useQuery } from "@tanstack/react-query"


export default function ScavengerHunt(){

    const { isPending, error, data } = useQuery({
        queryKey: ['repoData'],
        queryFn: () => (
            // if (window.location.hostname === 'localhost') {
            //     fetch("http://localhost:8080/scavengerhunts/2").then((res) => res.json())
            //   } else {
            //     setApiBaseUrl('http://192.168.x.x:5000'); // Replace with your computer's local IP
            //   }
            fetch("http://localhost:8080/scavengerhunts/2").then((res) => res.json())      
        ),
    })
    
    console.log(isPending)
    console.log(data)





    return(<></>)
}