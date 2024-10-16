import React, { useState, useEffect } from 'react';

const CalorieForm = ({ onSubmit, existingData }) => {
  const [formData, setFormData] = useState({ dish: '', fat: '', ingredients: '', calories: '' });

  useEffect(() => {
    if (existingData) {
      setFormData(existingData);
    }
  }, [existingData]);

  const handleSubmit = (e) => {
    e.preventDefault();
    onSubmit(formData);
    setFormData({ dish: '', fat: '', ingredients: '', calories: '' });
  };

  return (
    <form onSubmit={handleSubmit} className="mb-3">
      <div className="mb-3">
        <label className="form-label">Dish</label>
        <input
          type="text"
          className="form-control"
          value={formData.dish}
          onChange={(e) => setFormData({ ...formData, dish: e.target.value })}
          required
        />
      </div>
      <div className="mb-3">
        <label className="form-label">Fat (g)</label>
        <input
          type="number"
          className="form-control"
          value={formData.fat}
          onChange={(e) => setFormData({ ...formData, fat: e.target.value })}
          required
        />
      </div>
      <div className="mb-3">
        <label className="form-label">Ingredients</label>
        <input
          type="text"
          className="form-control"
          value={formData.ingredients}
          onChange={(e) => setFormData({ ...formData, ingredients: e.target.value })}
          required
        />
      </div>
      <div className="mb-3">
        <label className="form-label">Calories</label>
        <input
          type="number"
          className="form-control"
          value={formData.calories}
          onChange={(e) => setFormData({ ...formData, calories: e.target.value })}
          required
        />
      </div>
      <button type="submit" className="btn btn-primary">
        {existingData ? "Edit" : "Save"}
      </button>
    </form>
  );
};

export default CalorieForm;
