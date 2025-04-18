import { BrowserRouter, Routes, Route } from "react-router-dom";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import StudentListScreen from "./modules/students/screens/List";
import StudentFormScreen from "./modules/students/screens/Form";

function App() {
  const queryClient = new QueryClient();

  return (
    <QueryClientProvider client={queryClient}>
      <BrowserRouter>
      <Routes>
        <Route path="/" element={<StudentListScreen />} />
        <Route path="/form" element={<StudentFormScreen />} />
      </Routes>
    </BrowserRouter>
    </QueryClientProvider>
  );
}

export default App;
