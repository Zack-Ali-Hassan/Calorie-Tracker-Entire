import React, { useState } from "react";
import axios from "axios";
const API_URL =
  import.meta.env.MODE == "development" ? "http://localhost:6767/api" : "/api";
const Login = () => {
  const [formData, setFormData] = useState({
    email: "",
    password: "",
  });

  const handleChange = (e) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    // try {
    //   const response = await axios.post(API_URL + "/login/", formData);
    //   console.log(response.data);
    //   toast.success(response.data.msg);
    //   // Store the token in local storage for future requests
    //   localStorage.setItem("token", response.data.token);
    // } catch (error) {
    //   console.error("Error logging in", error);
    //   toast.error("Login error");
    // }
  };

  return (
    <div className="container">
      <h2>Login</h2>
      <form onSubmit={handleSubmit}>
        <div className="mb-3">
          <label className="form-label">Email address</label>
          <input
            type="email"
            name="email"
            className="form-control"
            value={formData.email}
            onChange={handleChange}
            required
          />
        </div>
        <div className="mb-3">
          <label className="form-label">Password</label>
          <input
            type="password"
            name="password"
            className="form-control"
            value={formData.password}
            onChange={handleChange}
            required
          />
        </div>
        <button type="submit" className="btn btn-primary">
          Login
        </button>
      </form>
    </div>
  );
};

export default Login;
