// src/context/AuthContext.tsx

import React, { createContext, useContext, useState, useEffect, ReactNode } from "react";
import axios from "axios";
import { getUrl } from "../utils/geturl";

interface User {
  id: string;
  // Add other user properties here if needed
}

interface AuthContextType {
  user: User | null;
  setUser: React.Dispatch<React.SetStateAction<User | null>>;
  isLoading: boolean;
  error: string | null; // Added error state
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export const AuthProvider = ({ children }: { children: ReactNode }) => {
  const [user, setUser] = useState<User | null>(null);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<string | null>(null); // Added error state

  const URL = getUrl();

  useEffect(() => {
    const checkAuth = async () => {
      setIsLoading(true);
      try {
        const response = await axios.get(`${URL}/v1/user/`, { withCredentials: true });
        setUser(response.data);
        setError(null); // Clear any previous errors
      } catch (error: any) {
        console.error("Auth error:", error);
        setUser(null);
        setError("Failed to authenticate. Please try again."); // Set error message
      } finally {
        setIsLoading(false);
      }
    };
    checkAuth();
  }, [URL]); // URL is now correctly added to the dependency array

  return (
    <AuthContext.Provider value={{ user, setUser, isLoading, error }}>
      {children}
    </AuthContext.Provider>
  );
};

export const useAuth = () => {
  const context = useContext(AuthContext);
  if (!context) {
    throw new Error("useAuth must be used within an AuthProvider");
  }
  return context;
};
