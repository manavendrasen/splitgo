import { Button } from "@/components/ui/button";
import { Plus } from "lucide-react";
const Home = () => {
  return (
    <div className="fixed bottom-0 left-0 bg-secondary flex justify-center items-center p-4 flex-1 w-full">
      <a href="/add-expense">
        <Button>
          <Plus />
        </Button>
      </a>
    </div>
  );
};

export default Home;
