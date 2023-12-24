import "./App.css";
import AddTransaction from "./pages/AddTransaction";
import Home from "./pages/Home";
import { ThemeProvider } from "./store/ThemeProvider";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
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
      <RouterProvider router={router} />
    </ThemeProvider>
  );
}

export default App;
