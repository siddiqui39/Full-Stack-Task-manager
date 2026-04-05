import React from "react";
import ReactDOM from "react-dom/client";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import App from "./App"; // Task/dashboard page
import Login from "./pages/Login";
import Register from "./pages/Register";
import "./index.css";

ReactDOM.createRoot(document.getElementById("root")).render(
  <BrowserRouter>
    <Routes>
      <Route path="/" element={<App />} />          {/* Home or tasks page */}
      <Route path="/login" element={<Login />} />   {/* Login page */}
      <Route path="/register" element={<Register />} /> {/* Register page */}
      <Route path="/tasks" element={<App />} />     {/* Tasks page */}
    </Routes>
  </BrowserRouter>
);