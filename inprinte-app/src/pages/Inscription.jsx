import React, { useState } from 'react';
import "../styles/Inscription.css"
import { Link } from "react-router-dom";


export default function Inscription() {
    const auth = localStorage.getItem('auth');
    if (auth) {
      const authObj = JSON.parse(auth);
      if (authObj !== undefined) {
        window.location.href = '/profil';
      }
    }



    const [firstName, setFirstName] = useState("");
    const [lastName, setLastName] = useState("");
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [phone, setPhone] = useState("");


  let handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const request = new Request(
        'http://localhost:8080/inscription',
        {
          method: 'POST',
          body: JSON.stringify({firstName: firstName, lastName: lastName, email: email, password : password, phone: phone }),
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
        })
        .then(() => {
          window.location.href = '/connexion';
        })
        .catch(() => {
          throw new Error('Network error');
        });
    } catch (err) {
      console.log(err);
    }
  };




    return (
        <>
            <div className="inscription">
                <form className="form" onSubmit={handleSubmit}>
                    <h2>Inscription</h2>
                    <div className="form-group">
                        <input value={firstName} type="text" placeholder="Prénom" require onChange={(e) => setFirstName(e.target.value)}/>
                    </div>
                    <div className="form-group">
                        <input value={lastName} type="text" placeholder="Nom" require onChange={(e) => setLastName(e.target.value)}/>
                    </div>
                    <div className="form-group">
                        <input value={email} type="email" placeholder="Adresse mail" require onChange={(e) => setEmail(e.target.value)}/>
                    </div>
                    <div className="form-group">
                        <input value={password} type="password" placeholder="Mot de passe" require onChange={(e) => setPassword(e.target.value)}/>
                    </div>
                    <div className="form-group">
                        <input value={phone} type="tel" placeholder="Téléphone" require onChange={(e) => setPhone(e.target.value)}/>
                    </div>
                    <div className="form-group">
                        <input type="submit" value="S'inscrire" />
                    </div>
                </form>
            </div>
        </>
    );
}
