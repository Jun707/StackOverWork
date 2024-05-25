import React from "react";


export default function SignUp() {
    const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault()

        const formData = new FormData(e.currentTarget);

        const body = {
            username: formData.get("username"),
            email: formData.get("email"),
            password: formData.get("password"),
        };

        const res = await fetch('/api/signup', {
            method:"POST",
            body: JSON.stringify(body),
            headers: {
                'Content-Type': 'application/json',
            },
        })

        await res.json()
    }

    return (
        <div>
            <form onSubmit={handleSubmit}>
                <label htmlFor="username"> Username </label>
                <input type="text" name="username"></input>
                <label htmlFor="email"> Email </label>
                <input type="text" name="email"></input>
                <label htmlFor="password"> Password </label>
                <input type="password" name="password"></input>
                <button type="submit"> Sign up </button>
            </form>

        </div>
    )
}