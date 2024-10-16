import React, { useState, useEffect } from 'react';
import axios from 'axios';
import CalorieForm from './components/forms/CalorieForm';
import CalorieTable from './components/tables/CalorieTable';

const API_URL = import.meta.env.MODE == "development" ? "http://localhost:6767/api" : "/api";

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
    if (editData) {
      await axios.patch(`${API_URL}/${editData._id}`, data);
    } else {
      await axios.post(API_URL, data);
    }
    setEditData(null);
    fetchCalories();  // Refresh data
  };

  const handleEdit = (calorie) => {
    setEditData(calorie);
  };

  const handleDelete = async (id) => {
    await axios.delete(`${API_URL}/${id}`);
    fetchCalories();
  };

  return (
    <div className="container mt-5">
      <h1>Calorie Tracker Entire</h1>

      {/* Form for Create or Update */}
      <CalorieForm onSubmit={handleSubmit} existingData={editData} />

      {/* Display Table of Items */}
      <CalorieTable calories={calories} onEdit={handleEdit} onDelete={handleDelete} />
    </div>
  );
};

export default App;
