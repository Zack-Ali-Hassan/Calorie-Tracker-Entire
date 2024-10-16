import React from 'react';

const CalorieTable = ({ calories, onEdit, onDelete }) => {
  return (
    <table className="table">
      <thead>
        <tr>
          <th>Dish</th>
          <th>Fat</th>
          <th>Ingredients</th>
          <th>Calories</th>
          <th>Actions</th>
        </tr>
      </thead>
      <tbody>
        {calories.map((calorie) => (
          <tr key={calorie._id}>
            <td>{calorie.dish}</td>
            <td>{calorie.fat}</td>
            <td>{calorie.ingredients}</td>
            <td>{calorie.calories}</td>
            <td>
              <button className="btn btn-warning me-2" onClick={() => onEdit(calorie)}>Edit</button>
              <button className="btn btn-danger" onClick={() => onDelete(calorie._id)}>Delete</button>
            </td>
          </tr>
        ))}
      </tbody>
    </table>
  );
};

export default CalorieTable;
