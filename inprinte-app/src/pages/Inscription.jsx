import React from "react";
import "../styles/Inscription.css"
import { Link } from "react-router-dom";

export default function Inscription() {
    return (
        <>
            <div className="inscription">
                <form className="form">
                    <h2>Inscription</h2>
                    <div className="form-group">
                        <input type="text" placeholder="Prénom" require/>
                    </div>
                    <div className="form-group">
                        <input type="text" placeholder="Nom" require/>
                    </div>
                    <div className="form-group">
                        <input type="email" placeholder="Adresse mail" require/>
                    </div>
                    <div className="form-group">
                        <input type="password" placeholder="Mot de passe" require/>
                    </div>
                    <div className="form-group">
                        <input type="tel" placeholder="Téléphone" require/>
                    </div>
                    <div className="form-group">
                        <Link to="/connexion"><input type="submit" value="S'inscrire" /></Link>
                    </div>
                </form>
            </div>
        </>
    );
}
