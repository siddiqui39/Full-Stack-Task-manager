import { useState } from "react";
import { useNavigate } from "react-router-dom"; // <--- Add this

function Register() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const navigate = useNavigate(); // <--- Add this

  const handleSubmit = async (e) => {
    e.preventDefault();

    try {
      const res = await fetch("http://localhost:3000/api/register", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ email, password }),
      });

      const data = await res.json();
      console.log("Response:", data);

      if (data.success) {
        alert("Registration successful! Please login.");
        navigate("/login"); // <--- Redirect after registration
      } else {
        alert("Registration failed!");
      }
    } catch (err) {
      console.error("Error:", err);
      alert("An error occurred");
    }
  };

  return (
    <div>
      <h2>Register</h2>
      <form onSubmit={handleSubmit}>
        <input
          type="email"
          placeholder="Enter email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
        />
        <br /><br />
        <input
          type="password"
          placeholder="Enter password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
        />
        <br /><br />
        <button type="submit">Register</button>
      </form>
    </div>
  );
}

export default Register;