import "./App.css";
import AddTransaction from "./pages/AddTransaction";
import Home from "./pages/Home";
import { ParticipantProvider } from "./store/ParticipantStore";
import { ThemeProvider } from "./store/ThemeProvider";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import { Toaster } from "@/components/ui/sonner";
const router = createBrowserRouter([
  {
    path: "/",
    element: <Home />,
  },
  {
    path: "/add-expense",
    element: <AddTransaction />,
  },
]);

function App() {
  return (
    <ThemeProvider>
      <ParticipantProvider>
        <RouterProvider router={router} />
        <Toaster richColors />
      </ParticipantProvider>
    </ThemeProvider>
  );
}

export default App;
