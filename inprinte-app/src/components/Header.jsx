import React from 'react';
import { Link } from 'react-router-dom';

const Header = () => {
    return (
        <div>
            <ul className="header-ul">
                <li><Link to="/">Logo</Link></li>
                <li><Link to="/accueil">Accueil</Link></li>
                <li><Link to="/boutique">Boutique</Link></li>
                <li><Link to="/propositions">Propositions</Link></li>
                <li><Link to="/contact">Contact</Link></li>
            </ul>
        </div>
    );
};

export default Header;