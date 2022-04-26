
import React, { useState } from 'react';
import "../styles/Connexion.css"
import { Navbar } from '../components/Navbar';

import { Link } from "react-router-dom";
const Connexion = () => {

  const auth = localStorage.getItem('auth');
  if (auth) {
    const authObj = JSON.parse(auth);
    if (authObj !== undefined) {
      window.location.href = '/profil';
    }
  }

  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  let handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const request = new Request(
        'http://localhost:8080/connexion',
        {
          method: 'POST',
          body: JSON.stringify({ email: email, password : password }),
          headers: new Headers({ 
            'Content-Type': 'application/json',
          }),
        }
      );
      return fetch(request)
        .then((response) => {
          if (response.status < 200 || response.status >= 300) {
            throw new Error(response.statusText);
          }
          return response.json();
        })
        .then((auth) => {
          localStorage.setItem(
            'auth',
            JSON.stringify({ ...auth })
          );
          window.location.href = '/profil';
        })
        .catch(() => {
          throw new Error('Network error');
        });
    } catch (err) {
      console.log(err);
    }
  };

    return (

        <div className="">
          <Navbar />
            <div className="connexion">
                <form className="form-connexion" onSubmit={handleSubmit}>
                    <h2>Connexion</h2>
                    <div className="form-group">
                        <input value={email} type="email" placeholder="Adresse mail" onChange={(e) => setEmail(e.target.value)}/>
                    </div>
                    <div className="form-group">
                        <input value={password} type="password" placeholder="Mot de passe" onChange={(e) => setPassword(e.target.value)}/>
                    </div>
                    <div className="form-group">
                    <input type="submit" value="Se connecter" />
                        
                    </div>
                    <div className="form-group">
                        <Link to="/inscription">s'inscrire</Link>
                    </div>
                </form>
            </div>
        </div>
    );
};

export default Connexion;