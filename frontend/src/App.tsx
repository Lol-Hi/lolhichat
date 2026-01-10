import React from 'react';
import './App.css';
import { BrowserRouter, Route, Routes } from "react-router-dom";

import Home from "./pages/Home";
import Login from "./pages/Login";
import SignUp from "./pages/SignUp";
import NewThread from "./pages/NewThread";

import NavBar from "./components/NavBar";

function App() {
  return (
		<div className="App">
			<BrowserRouter>
				<NavBar />
				<Routes>
					<Route path="/" element={<Home />} />
					<Route path="/login" element={<Login />} />
					<Route path="/signUp" element={<SignUp />} />
					<Route path="/newThread" element={<NewThread />} />
				</Routes>
			</BrowserRouter>
		</div>
  );
}

export default App;
