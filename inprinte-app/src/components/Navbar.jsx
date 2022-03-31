import { Link } from "react-router-dom";
import "../styles/Navbar.css";
import InprinteLogo from '../images/logo.svg'
import CartLogo from '../images/cart.svg'
import ConnexionLogo from '../images/connexion.svg'


export const Navbar = () => {
    return (
        <>
            <nav className="nav">
                <div className="nav-container">
                    <div className="left">
                        <Link to="/">
                            <img src={InprinteLogo} alt="Inprinte Logo" />
                        </Link>
                    </div>
                    <div className="middle">
                        <ul>
                            <li>
                                <Link to="/accueil">Accueil</Link>
                            </li>
                            <li>
                                <Link to="/boutique">Boutique</Link>
                            </li>
                            <li>
                                <Link to="/propositions">Propositions</Link>
                            </li>
                            <li>
                                <Link to="/demande">Demande</Link>
                            </li>
                            <li>
                                <Link to="/Contact">Contact</Link>
                            </li>
                        </ul>
                    </div>
                    <div className="right">
                        <ul>
                            <li>
                                <Link to='/connexion'><img className="con" src={ConnexionLogo} alt="Connexion Logo" /></Link>
                                <Link to='/panier'><img className="cart" src={CartLogo} alt="Cart Logo" /></Link>
                            </li>
                        </ul>
                    </div>
                </div>
            </nav>
        </>
    );
};