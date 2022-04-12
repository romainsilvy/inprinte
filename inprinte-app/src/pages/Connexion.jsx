import React from 'react';
import "../styles/Connexion.css"
import { Link } from "react-router-dom";
const Connexion = () => {
    return (
        <div className="">
            <div className="connexion">
                <form className="form">
                    <h2>Connexion</h2>
                    <div className="form-group">
                        <input type="text" placeholder="Pseudo" />
                    </div>
                    <div className="form-group">
                        <input type="text" placeholder="Mot de passe" />
                    </div>
                    <div className="form-group">
                        <input type="submit" value="Se connecter" />
                    </div>
                    <div className="form-group">
                        <Link to="/inscription"><input type="submit" value="S'inscrire" /></Link>
                    </div>
                </form>
            </div>
        </div>
    );
};

export default Connexion;