import { Input } from "@/components/ui/input";
import Image from "next/image";

export default function Home() {
  return (
    <main className="w-full h-screen  flex flex-col items-center justify-center gap-y-12">
      <div className="space-y-4">
        <h1 className="text-center text-white text-5xl font-bold">Wiki Race</h1>
        <p className="text-white">Pemanfaatan Algoritma IDS dan BDS dalam Permainan WikiRace</p>
      </div>
      <Input 
        className="max-w-[500px] z-[20]"
        placeholder="Masukkan Link Wikipedia"
      />
    </main>
  );
}
