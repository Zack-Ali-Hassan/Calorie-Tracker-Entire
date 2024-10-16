import React, { useState, useEffect } from "react";
import axios from "axios";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Header from "./components/UI/Header";
import CalorieForm from "./components/forms/CalorieForm";
import CalorieTable from "./components/tables/CalorieTable";
import Profile from "./components/UI/Profile";
import Logout from "./components/UI/Logout";
import toast from "react-hot-toast";
import Signup from "./components/auth/Signup";
import Login from "./components/auth/Login";

const API_URL =
  import.meta.env.MODE == "development" ? "http://localhost:6767/api" : "/api";

const App = () => {
  const [calories, setCalories] = useState([]);
  const [editData, setEditData] = useState(null);

  useEffect(() => {
    fetchCalories();
  }, []);

  const fetchCalories = async () => {
    try {
      const response = await axios.get(API_URL + "/entire/");
      setCalories(response.data);
    } catch (error) {
      console.error("Error fetching data", error);
    }
  };

  const handleSubmit = async (data) => {
    try {
      if (editData) {
        const response = await axios.patch(`${API_URL}/entire/${editData._id}`, {
        dish: data.dish,
        fat: parseFloat(data.fat),  
        ingredients: data.ingredients,
        calories: parseFloat(data.calories)  
      });
      console.log("Success updated data : ", response.data.msg);
      toast.success(response.data.msg);
       } else {
         const response = await axios.post(API_URL + "/entire/", {
          dish: data.dish,
          fat: parseFloat(data.fat),  
          ingredients: data.ingredients,
          calories: parseInt(data.calories)  
        });
        console.log("Success inserted data : ", response.data.msg);
        toast.success(response.data.msg);
       }
       setEditData(null);
       fetchCalories(); // Refresh data
    } catch (error) {
      console.error("Error inserting/updating data", error);
      toast.error("Error inserting/updating data")
    }
   
  };

  const handleEdit = (calorie) => {
    setEditData(calorie);
  };

  const handleDelete = async (id) => {
    try {
      if (confirm(`Are you sure you want to delete ${id} ?`)) {
        const response = await axios.delete(`${API_URL}/entire/${id}`);
        fetchCalories();
        toast.success(response.data.msg);
      }
    } catch (error) {
      console.error("Error deleing data", error);
    }
  };

  return (
    <Router>
      {/* Header Component */}
      <Header />

      {/* Routes for different pages */}
      <div className="container mt-5">
        <Routes>
          <Route
            path="/"
            element={
              <div>
                <h1>Calorie Tracker Entire</h1>
                <CalorieForm onSubmit={handleSubmit} existingData={editData} />
                <CalorieTable
                  calories={calories}
                  onEdit={handleEdit}
                  onDelete={handleDelete}
                />
              </div>
            }
          />
          <Route path="/profile" element={<Profile />} />
          <Route path="/logout" element={<Logout />} />
          <Route path="/signup" element={<Signup />} /> {/* Signup route */}
          <Route path="/login" element={<Login />} />    {/* Login route */}
        </Routes>
      </div>
    </Router>
  );
};

export default App;
