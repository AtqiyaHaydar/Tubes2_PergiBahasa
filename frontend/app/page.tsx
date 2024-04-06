"use client"

import { useEffect, useState } from "react";
import axios, { AxiosResponse } from "axios";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { Switch } from "@/components/ui/switch";
import { toast } from "@/components/ui/use-toast";

export default function Home() {
  const [query, setQuery] = useState<string>(""); /* Judul Artikel Awal */
  const [objective, setObjective] = useState<string>(""); /* Judul Artikel Tujuan */
  const [algorithm, setAlgorithm] = useState<boolean>(false); /* Default IDS */
  const [searchTerm, setSearchTerm] = useState<string[]>([]); /* Menampilkan Hasil Yang Didapat Dari Wikipedia API */
  const [result, setResult] = useState<string[]>([]); /* Hasil Pencarian */

  /* Fungsi Untuk Mengirim Request dan Menerima Response Dari Backend */
  const handleSearch = async () => {
    toast({
      title: "Searching...",
      description: "Artikel Awal: " + query + " | Artikel Tujuan: " + objective + " | Mode: " + (algorithm ? "BFS" : "IDS"),
    })

    try {
      const response = await axios.get(`http://localhost:8080/scrape?query=${query}`)
      setResult(response.data)
    } catch (error) {
      console.error('Error fetching data:', error)
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

  /* Fungsi Menampilkan Hasil Pencarian Dari Query */
  const handleQuery = async () => {
    const value = query.trim();
  
    if (!value) {
      console.error('Query parameter is empty');
      return;
    }
  
    try {
      const response = await axios.get(
        `http://localhost:8080/api/wikipedia?query=${encodeURIComponent(value)}`
      );
  
      console.log('Response:', response.data.query.search.map((item: any) => item.title));
      setSearchTerm(response.data.query.search.map((item: any) => item.title));
    } catch (error) {
      console.error('Error fetching data:', error);
    }
  };
   
  return (
    <main className="w-full h-full my-24 flex flex-col items-center justify-center gap-y-12">
      {/* Judul Website */}
      <div className="space-y-4">
        <h1 className="text-center text-white text-5xl font-bold">Wiki Race</h1>
        <p className="text-white text-center">Pemanfaatan Algoritma IDS dan BDS dalam Permainan WikiRace</p>
      </div>

      {/* Input Artikel Awal dan Tujuan */}
      <div className="flex flex-col md:flex-row justify-between gap-8 py-4 items-center">
        <div className="flex flex-col items-center gap-y-2">
          <p className="font-bold text-white ">Artikel Awal</p>
          <Input 
            type="text"
            className="w-[300px] z-[20]"
            placeholder="Judul Artikel Awal"
            value={query}
            onChange={
              (e) => {
                setQuery(e.target.value)
                handleQuery()
              }
            }
          />
        </div>
        <p className="text-white text-2xl font-bold mx-4">
          TO
        </p>
        <div className="flex flex-col items-center gap-y-2">
          <p className="font-bold text-white ">Artikel Tujuan</p>
          <Input 
            type="text"
            className="w-[300px] z-[20]"
            placeholder="Judul Artikel Tujuan"
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
      <ul className="m-12 px-4 py-8 border-2 border-white min-h-[50px] md:min-w-[750px] rounded-xl grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-x-9 gap-y-1 text-center">
        {result && result.map((link: string, index: number) => (
          <li key={index}>
            <p className="text-white">{link}</p>
          </li>
        ))}
      </ul>
    </main>
  );
}