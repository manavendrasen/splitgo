import { Button } from "@/components/ui/button";
import { Plus } from "lucide-react";
const Home = () => {
  return (
    <main>
      <div className="flex justify-center items-center mt-10">
        <p className="text-lg text-center font-semibold">SplitGo</p>
      </div>
      <div className="fixed bottom-5 right-5 bg-secondary rounded-full">
        <a href="/add-expense">
          <Button className="rounded-full">
            <Plus />
          </Button>
        </a>
      </div>
    </main>
  );
};

export default Home;
