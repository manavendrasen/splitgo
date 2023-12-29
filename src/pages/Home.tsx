import { Background } from "@/components/background/Background";
import { Button } from "@/components/ui/button";
import { Plus } from "lucide-react";
import Logo from "@/assets/logo.svg";
const Home = () => {
  return (
    <main>
      <Background />
      <div className="flex justify-center items-center h-[calc(100vh-100px)] mix-blend-overlay">
        <img src={Logo} alt="Logo" className="w-48" />
      </div>
      <div
        className="fixed bottom-10 right-1/2"
        style={{
          transform: "translateX(50%)",
        }}
      >
        <a href="/add-expense">
          <Button className="rounded-full p-4 font-semibold w-14 h-14 bg-gradient-to-br from-[#9A97F9] to-[#514DCF] text-foreground">
            <Plus />
          </Button>
        </a>
      </div>
    </main>
  );
};

export default Home;
