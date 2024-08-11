// src/pages/Home.tsx

import React from "react";
import { useAuth } from "../context/AuthContext";

const Home: React.FC = () => {
  const { user } = useAuth();

  if (!user) {
    return <div>Loading...</div>; // or redirect to login if user is not available
  }

  return (
    <div className="flex items-center justify-center min-h-screen bg-gray-100">
      <h1 className="text-3xl font-bold">Welcome, {user.id}</h1>
      {/* Display more user info here if available */}
    </div>
  );
};

export default Home;
