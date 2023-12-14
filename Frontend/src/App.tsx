import { Route, Routes } from "react-router-dom";
import { ToastContainer } from "react-toastify";
import "./App.css";
import HomeApp from "./pages/HomeApp";
import Dashboard from "./pages/Dashboard";
import MyUrls from "./pages/MyUrls";

function App() {
  return (
    <>
      <Routes>
        <Route path="/" element={<HomeApp />} />
        <Route path="/dashboard" element={<Dashboard />} />
        <Route path="/myUrls" element={<MyUrls />} />
      </Routes>
      <ToastContainer />
    </>
  );
}

export default App;
