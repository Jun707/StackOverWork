import React from "react";


export default function SignIn() {
    const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault()

        const formData = new FormData(e.currentTarget);

        const body = {
            email: formData.get("email"),
            password: formData.get("password"),
        };

        const res = await fetch('/api/signin', {
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
                <label htmlFor="email"> Email </label>
                <input type="text" name="email"></input>
                <label htmlFor="password"> Password </label>
                <input type="password" name="password"></input>
                <button type="submit"> Sign in </button>
            </form>

        </div>
    )
}