import "./App.css";
import AddTransaction from "./pages/AddTransaction";
import Home from "./pages/Home";
import { ParticipantProvider } from "./store/ParticipantStore";
import { ThemeProvider } from "./store/ThemeProvider";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import { Toaster } from "@/components/ui/sonner";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { ReactQueryDevtools } from "@tanstack/react-query-devtools";

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
  const queryClient = new QueryClient();
  return (
    <ThemeProvider>
      <QueryClientProvider client={queryClient}>
        <ParticipantProvider>
          <RouterProvider router={router} />
          <Toaster richColors />
        </ParticipantProvider>
        <ReactQueryDevtools initialIsOpen />
      </QueryClientProvider>
    </ThemeProvider>
  );
}

export default App;
