"use client"

import { useEffect, useState } from "react";
import axios from "axios";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { Switch } from "@/components/ui/switch";
import { toast } from "@/components/ui/use-toast";

export default function Home() {
  const [query, setQuery] = useState<string>(""); /* Judul Artikel Awal */
  const [objective, setObjective] = useState<string>(""); /* Judul Artikel Tujuan */
  const [algorithm, setAlgorithm] = useState<boolean>(false); /* Default IDS */
  const [result, setResult] = useState<string[]>([]); /* Hasil Pencarian */


  /* Fungsi Untuk Mengirim Request dan Menerima Response Dari Backend */
  const handleSearch = async () => {
    toast({
      title: "Searching...",
      description: "Artikel Awal: " + query + " | Artikel Tujuan: " + objective + " | Mode: " + (algorithm ? "BFS" : "IDS"),
    })

    try {
      const response = await axios.get(`http://localhost:8080/scrape?query=${query}`)
      console.log("Response:", response.data)
      setResult(response.data)
    } catch (error) {
      console.log('Error fetching data:', error)
    }
  }

  /* Use Effect : Notifikasi Bahwa Mode Sudah Diganti */
  useEffect(() => {
    let mode = algorithm ? "BFS" : "IDS"
    let desc = algorithm ? "Breadth First Search" : "Iterative Deepening First Search"

    toast({
      title: "Set Mode To " + mode,
      description: desc
    })
  }, [algorithm])
   
  return (
    <main className="w-full h-full my-24 flex flex-col items-center justify-center gap-y-12">
      {/* Judul Website */}
      <div className="space-y-4">
        <h1 className="text-center text-white text-5xl font-bold">Wiki Race</h1>
        <p className="text-white text-center">Pemanfaatan Algoritma IDS dan BDS dalam Permainan WikiRace</p>
      </div>

      {/* Input Artikel Awal dan Tujuan */}
      <div className="flex flex-col md:flex-row justify-between gap-8">
        <div className="flex flex-col items-center gap-y-2">
          <p className="font-bold text-white ">Artikel Awal</p>
          <Input 
            type="text"
            className="w-[250px] z-[20]"
            placeholder="Masukkan Artikel Awal"
            value={query}
            onChange={(e) => setQuery(e.target.value)}
          />
        </div>
        <div className="flex flex-col items-center gap-y-2">
          <p className="font-bold text-white ">Artikel Tujuan</p>
          <Input 
            type="text"
            className="w-[250px] z-[20]"
            placeholder="Masukkan Artikel Tujuan"
            value={objective}
            onChange={(e) => setObjective(e.target.value)}
          />
        </div>
      </div>

      {/* Switch : Mengganti Mode Algoritma */}
      <div className="flex gap-x-4 text-white z-[20]">
        <p>IDF</p>
        <Switch 
          onClick={() => setAlgorithm(prev => !prev)}
        />
        <p>BFS</p>
      </div>

      {/* Button : Mencari Hasil Pencarian */}
      <Button className="z-[20] w-[125px]" onClick={handleSearch}>Search</Button>

      {/* Mapping Hasil Pencarian */}
      <ul className="m-12 p-4 border-2 border-white min-h-[50px] md:min-w-[500px] rounded-xl grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-x-9 gap-y-1 text-center">
        {result && result.map((link: string, index: number) => (
          <li key={index}>
            <p className="text-white">{link}</p>
          </li>
        ))}
      </ul>
    </main>
  );
}