import { Link } from "react-router-dom";
import "../styles/Navbar.css";
import InprinteLogo from '../images/logo.svg'
import CartLogo from '../images/cart.svg'
import ConnexionLogo from '../images/connexion.svg'


export const Navbar = () => {
    return (
        <>
            <nav className="nav">
                <Link to="/" className="left">
                    <img src={InprinteLogo} alt="Inprinte Logo" />
                </Link>
                <ul className="middle">
                    <li>
                        <Link to="/">Accueil</Link>
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
                <div className="right">
                    <Link to='/connexion'><img className="con" src={ConnexionLogo} alt="Connexion Logo" /></Link>
                    <Link to='/panier'><img className="cart" src={CartLogo} alt="Cart Logo" /></Link>
                </div>
            </nav>
        </>
    );
};