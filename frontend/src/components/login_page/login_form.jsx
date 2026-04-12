import React, { useState } from 'react';
import {TextField} from "@mui/material";

function LoginForm() {
    const [formData, setFormData] = useState({ email: '', password: '' });

    const handleChange = (e) => {
        const { name, value } = e.target;
        setFormData({ ...formData, [name]: value });
    };

    const handleSubmit = (e) => {
        e.preventDefault(); // Prevents page reload
        // if (!formData.email || !formData.password) {
        //     alert("Please fill in all fields.");
        //     return;
        // }
        // console.log("Submitting:", formData);
    };

    return (
        <form onSubmit={handleSubmit}>
            <h2>Login</h2>
            <div>
                <TextField
                    name="email"
                    label="Email"
                    value={formData.email}
                    variant="standard"
                    onChange={handleChange}
                    required
                />
            </div>
            <div>
                <TextField
                    name="password"
                    label="password"
                    value={formData.password}
                    variant="standard"
                    onChange={handleChange}
                    type="password"
                    required
                />
            </div>
            <button type="submit">Log In</button>
        </form>
    );
}

export default LoginForm;