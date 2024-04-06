"use client"

import { useState } from "react";
import axios from "axios";

import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";

export default function Home() {
  const [query, setQuery] = useState<string>("");
  const [result, setResult] = useState<string[]>([]);

  const handleSearch = async () => {
    console.log("CLIKED!")
    console.log("Query:", query)

    try {
      const response = await axios.get(`http://localhost:8080/scrape?query=${query}`)
      console.log("Response:", response.data)
      setResult(response.data)
    } catch (error) {
      console.log('Error fetching data:', error)
    }
  }
  
  return (
    <main className="w-full h-screen  flex flex-col items-center justify-center gap-y-12">
      <div className="space-y-4">
        <h1 className="text-center text-white text-5xl font-bold">Wiki Race</h1>
        <p className="text-white">Pemanfaatan Algoritma IDS dan BDS dalam Permainan WikiRace</p>
      </div>
      <Input 
        type="text"
        className="max-w-[500px] z-[20]"
        placeholder="Masukkan Link Wikipedia"
        value={query}
        onChange={(e) => setQuery(e.target.value)}
      />
      <Button className="z-[20]" onClick={handleSearch}>Search</Button>
      <ul className="grid grid-cols-3 gap-x-9 gap-y-1 text-center">
        {result.slice(0, 50).map((link: string, index: number) => (
          <li key={index}>
            <p className="text-white">{link}</p>
          </li>
        ))}
      </ul>
    </main>
  );
}
