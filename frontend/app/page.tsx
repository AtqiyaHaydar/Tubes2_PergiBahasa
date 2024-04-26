"use client"

import Image from "next/image";
import { useEffect, useState, useRef } from "react";
import axios from "axios";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { Switch } from "@/components/ui/switch";
import { toast } from "@/components/ui/use-toast";
import {
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger,
} from "@/components/ui/tooltip"
import { ScrollArea } from "@/components/ui/scroll-area"
import ArrayGraph from "./ArrayGraph";

interface SearchResultI {
  ns: number;
  pageid: number;
  size: number;
  snippet: string;
  thumbnail?: { source: string };
  timestamp: string;
  title: string;
  wordcount: number;
}

export default function Home() {
  const [query, setQuery] = useState<string>(""); /* Judul Artikel Awal */
  const [objective, setObjective] = useState<string>(""); /* Judul Artikel Tujuan */
  const [algorithm, setAlgorithm] = useState<boolean>(false); /* Default IDS */
  const [searchTerm, setSearchTerm] = useState<SearchResultI[]>([]); /* Menampilkan Hasil Yang Didapat Dari Wikipedia API (Artikel Awal) */
  const [searchTermObjective, setSearchTermObjective] = useState<SearchResultI[]>([]); /* Menampilkan Hasil Yang Didapat Dari Wikipedia API (Artikel Tujuan) */
  const [isSelectOpen, setIsSelectOpen] = useState<boolean>(true); /* Menampilkan Hasil Yang Didapat Dari Wikipedia API (Artikel Awal) */
  const [isSelectOpenObjective, setIsSelectOpenObjective] = useState<boolean>(true); /* Menampilkan Hasil Yang Didapat Dari Wikipedia API (Artikel Tujuan) */
  const [result, setResult] = useState<string[]>([]); /* Hasil Pencarian */
  const [numArticleVisited, setNumArticleVisited] = useState<number>(0); /* Jumlah Artikel Dilalui */ 
  const [time, setTime] = useState<number>(0); /* Waktu Pencarian */

  /* Fungsi Untuk Mengirim Request dan Menerima Response Dari Backend */
  const handleSearch = async () => {
    toast({
      title: "Searching...",
      description: "Artikel Awal: " + query + " | Artikel Tujuan: " + objective
    })


    let theQuery = query.replace(/\s/g, "_"); /* Mengganti Spasi Dengan Underscore */
    let theObjective = objective.replace(/\s/g, "_"); /* Mengganti Spasi Dengan Underscore */

    try {
      let timeStart = new Date().getTime(); /* Waktu Mulai Pencarian */

      const mode = algorithm ? 'bfsfunc' : 'idsfunc';
      const response = await axios.get(`http://localhost:8080/api/${mode}?query=${theQuery}&query2=${theObjective}`);
      console.log(response.data);

      let timeEnd = new Date().getTime(); /* Waktu Selesai Pencarian */
      let timeDiff = timeEnd - timeStart; /* Waktu Total Pencarian */
      setTime(timeDiff); /* Set Waktu Pencarian */
      setResult(response.data.keywords); /* Set Hasil Pencarian */
      setNumArticleVisited(response.data.number); /* Set Jumlah Artikel Dilalui */

    } catch (error) {
      console.error('Error fetching data:', error);
    }
  };

  /* Use Effect : Notifikasi Toast Bahwa Mode Sudah Diganti */
  useEffect(() => {
    let mode = algorithm ? "BFS" : "IDS"
    let desc = algorithm ? "Breadth First Search" : "Iterative Deepening First Search"

    toast({
      title: "Set Mode To " + mode,
      description: desc
    })
  }, [algorithm])

  /* Fungsi Menampilkan Hasil Pencarian Dari Query Dengan Wikipedia API */
  const handleQuery = async () => {
    const value = query.trim();

    try {
      if (value === "") {
        setSearchTerm([]);
        return;
      }

      const response = await axios.get(
        `http://localhost:8080/api/wikipedia?query=${encodeURIComponent(value)}`
      );

  
      setSearchTerm(response.data.query.search);
    } catch (error) {
      console.error('Error fetching data:', error);
    }
  };

  /* Fungsi Menampilkan Hasil Pencarian Dari Objective Dengan Wikipedia API */
  const handleObjective = async () => {
    const value = objective.trim();

    try {
      if (value === "") {
        setSearchTermObjective([]);
        return;
      }

      const response = await axios.get(
        `http://localhost:8080/api/wikipedia?query=${encodeURIComponent(value)}`
      );

      setSearchTermObjective(response.data.query.search);
    } catch (error) {
      console.error('Error fetching data:', error); 
    }
  }

  /* Use Effect : Debounce Time Untuk Memperbarui Query Sekarang */
  useEffect(() => {
    if (isSelectOpen) {
      const timerId = setTimeout(() => {
        handleQuery();
      }, 500); // Debouncing Time
  
      return () => {
        clearTimeout(timerId);
      };
    }
  }, [query]);

  /* Use Effect : Debounce Time Untuk Memperbarui Objective Sekarang */
  useEffect(() => {
    if (isSelectOpenObjective) {
      const timerId = setTimeout(() => {
        handleObjective();
      }, 500); // Debouncing Time
  
      return () => {
        clearTimeout(timerId);
      };
    }
  }, [objective])
   
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
          <p className="font-bold text-white">Artikel Awal</p>
          <Input 
            type="text"
            className="w-[300px] z-[20]"
            placeholder="Judul Artikel Awal"
            value={query}
            onChange={
              (e) => {
                setQuery(e.target.value)
                handleQuery()
                setIsSelectOpen(true)
              }
            }
          />
          {/* Menampilkan Hasil Pencarian Dari Wikipedia API pada Input Artikel Awal */}
          {searchTerm.length > 0 && (
            <div className="absolute top-[325px]">
              <ScrollArea className="h-[175px] bg-white rounded-md border z-[20] w-[300px]">
                <ul className="py-2 gap-y-4">
                  {searchTerm.map((item, index) => (
                    <li 
                      key={index} 
                      className="text-black h-[45px] px-4 hover:bg-black/10 cursor-pointer transition-all py-1.5 flex items-center justify-center gap-x-2"
                      onClick={() => {
                        setQuery(item.title)
                        setSearchTerm([])
                        setIsSelectOpen(false)
                      }}
                    >
                      <p>{item.title}</p>
                    </li>
                  ))}
                </ul>
              </ScrollArea>
            </div>
          )}
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
            onChange={(e) => {
              setObjective(e.target.value)
              handleObjective()
              setIsSelectOpenObjective(true)
            }}
          />
          {/* Menampilkan Hasil Pencarian Dari Wikipedia API pada Input Artikel Tujuan */}
          {searchTermObjective.length > 0 && (
            <div className="absolute top-[325px]">
              <ScrollArea className="h-[175px] bg-white rounded-md border z-[20] w-[300px]">
                <ul className="py-2 gap-y-4">
                  {searchTermObjective.map((item, index) => (
                    <li 
                      key={index} 
                      className="text-black h-[45px] px-4 hover:bg-black/10 cursor-pointer transition-all py-1.5 flex items-center justify-center gap-x-2"
                      onClick={() => {
                        setObjective(item.title)
                        setSearchTermObjective([])
                        setIsSelectOpenObjective(false)
                      }}
                    >
                      <p>{item.title}</p>
                    </li>
                  ))}
                </ul>
              </ScrollArea>
            </div>
          )}
        </div>
      </div>

      {/* Switch : Mengganti Mode Algoritma */}
      <div className="flex gap-x-4 text-white z-[20] items-center justify-center">
        <TooltipProvider>
          <Tooltip>
            <TooltipTrigger asChild>
              <p className="hover:bg-white/10 px-3 py-2 rounded-sm cursor-pointer transition-all">IDS</p>
            </TooltipTrigger>
            <TooltipContent className="bg-white/10 text-white">
              <p>Iterative Deepening First Search</p>
            </TooltipContent>
          </Tooltip>
        </TooltipProvider>
        <Switch 
          onClick={() => setAlgorithm(prev => !prev)}
        />
        <TooltipProvider>
          <Tooltip>
            <TooltipTrigger asChild>
              <p className="hover:bg-white/10 px-3 py-2 rounded-sm cursor-pointer transition-all">BFS</p>
            </TooltipTrigger>
            <TooltipContent className="bg-white/10 text-white">
              <p>Breadth First Search</p>
            </TooltipContent>
          </Tooltip>
        </TooltipProvider>
      </div>

      {/* Button : Mencari Hasil Pencarian */}
      <Button className="z-[20] w-[125px]" onClick={handleSearch}>Search</Button>

      {/* Mapping Hasil Pencarian */}
      <div className="m-12 px-16 py-8 border-2 border-white/75 min-h-[50px] w-[90%] rounded-lg overflow-hidden">
        <div className="space-y-8 w-full">
          <div className="flex flex-col md:flex-row text-center space-y-2 w-full justify-between text-white">
            <h3 className="text-2xl font-semibold text-center">Results</h3>
            <p>Search Duration : {time} ms</p>
          </div>
          <div className="flex justify-between">
            <div className="mt-8 text-white text-xl text-center md:text-left space-y-2">
              <p>Jumlah Artikel Diperiksa : {numArticleVisited} </p>
              <p>Jumlah Artikel Dilalui : {result.length - 2 < 0 ? 0 : result.length - 2} </p>
            </div>
            <div className="flex items-center justify-center">
              {/* Menganimasikan Menggunakan Graph */}
              {result && (
                <div className="z-20 text-white">
                  <ArrayGraph data={result} width={250} height={250} />
                </div>
              )}
            </div>
          </div>
        </div>
      </div>
    </main>
  );
}