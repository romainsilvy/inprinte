import React from 'react';
import InprinteLogo from '../images/logo.svg'
import "../styles/Footer.css";
import { Link } from "react-router-dom";
import bar from '../images/bar.svg'
import facebook from '../images/facebook.svg'
import twitter from '../images/twitter.svg'
import instagram from '../images/instagram.svg'

export const Footer = () => {
    return (
        <>
            <footer>
                <div className="container">
                    <Link className="logo" to="/"><img src={InprinteLogo} alt="Inprinte Logo" /></Link>
                    <div className="propos">
                        <h2>A propos de nous</h2>
                        <ul>
                            <li><Link to="/faq">FAQ</Link></li>
                            <li><Link to="/mentions-legales">Mentions Légales</Link></li>
                        </ul>
                    </div>
                    <div className="services">
                        <h2>Nos Services</h2>
                        <ul>
                            <li><Link to="/accueil">Accueil</Link></li>
                            <li><Link to="/boutique">Boutique</Link></li>
                            <li><Link to="/propositions">Propositions</Link></li>
                            <li><Link to="/profil">Profil</Link></li>
                            <li><Link to="/panier">Panier</Link></li>
                        </ul>
                    </div>
                    <div className="contact">
                        <h2>Contact</h2>
                        <ul>
                            <li><Link to="/">Telephone</Link></li>
                            <li><Link to="/">email</Link></li>
                            <li><Link to="/">addresse</Link></li>
                        </ul>
                    </div>
                    <div className="communauté">
                        <h2>Communauté</h2>
                        <ul>
                            <li><Link to="/"><img className="bar" src={instagram} alt="instagram" /></Link></li>
                            <li><Link to="/"><img className="bar" src={facebook} alt="facebook" /></Link></li>
                            <li><Link to="/"><img className="bar" src={twitter} alt="twitter" /></Link></li>
                        </ul>
                    </div>
                </div>
            </footer>
            <img className="bar" src={bar} alt="bar" />
            <div className="copyrightText">
                <p>Copyright &copy; 2022 Tous droits réservés</p>
            </div>
        </>
    );
};

export default Footer;