import { Route, Routes } from "react-router-dom";
import "./App.css";
import HomeApp from "./pages/HomeApp";
import Dashboard from "./pages/Dashboard";

function App() {
  return (
    <Routes>
      <Route path="/" element={<HomeApp />} />
      <Route path="/dashboard" element={<Dashboard />} />
    </Routes>
  );
}

export default App;
