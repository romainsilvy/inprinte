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
                        <input type="text" placeholder="Pseudo" require/>
                    </div>
                    <div className="form-group">
                        <input type="text" placeholder="Adresse mail" require/>
                    </div>
                    <div className="form-group">
                        <input type="text" placeholder="Mot de passe" require/>
                    </div>
                    <div className="form-group">
                        <Link to="/connexion"><input type="submit" value="S'inscrire" /></Link>
                    </div>
                </form>
            </div>
        </>
    );
}