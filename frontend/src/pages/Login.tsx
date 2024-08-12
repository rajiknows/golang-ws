// src/pages/Login.tsx

import React, { useState } from "react";
import { useNavigate, Link } from "react-router-dom";
import axios from "axios";
import { getUrl } from "../utils/geturl";

const Login: React.FC = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  // const { setUser } = useAuth();
  const navigate = useNavigate();
  const url: string | undefined = getUrl();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      await axios.post(
        `${url}/v1/user/login`,
        {
          email,
          password,
        },
        {
          withCredentials: true,
          headers: {
            "Content-Type": "application/json",
          },
        },
      );
      // Adjust according to your API response structure
      // setUser({ id: response.data.id });
      navigate("/");
    } catch (error) {
      console.error("Login failed:", error);
    }
  };

  return (
    <div className="flex items-center justify-center min-h-screen bg-gray-100">
      <div className="bg-white p-6 rounded shadow-md w-80">
        <form onSubmit={handleSubmit}>
          <h2 className="text-2xl font-bold mb-4">Login</h2>
          <input
            type="email"
            placeholder="Email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            className="border p-2 mb-4 w-full"
          />
          <input
            type="password"
            placeholder="Password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            className="border p-2 mb-4 w-full"
          />
          <button
            type="submit"
            className="bg-blue-500 text-white py-2 px-4 rounded w-full mb-4"
          >
            Login
          </button>
        </form>
        <Link to="/register">
          <button className="bg-green-500 text-white py-2 px-4 rounded w-full">
            Register
          </button>
        </Link>
      </div>
    </div>
  );
};

export default Login;
