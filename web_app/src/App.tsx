import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import StudentsList from "./modules/students/screens/List";

function App() {
  const queryClient = new QueryClient();

  return (
    <QueryClientProvider client={queryClient}>
      <div data-testid="container" >
        <StudentsList />
      </div>
    </QueryClientProvider>
  );
}

export default App;
