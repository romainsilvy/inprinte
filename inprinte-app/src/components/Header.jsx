import React from 'react';
import { Nav, Bars, NavLink, NavMenu, NavBtn, NavBtnLink } from '../styles/HeaderStyles';

const Header = () => {
    return (
        <>
        <Nav>
            <NavLink to="/">
                <h1>Logo</h1>
            </NavLink>
            <Bars />
            <NavMenu> 
                <NavLink to="/accueil">
                    Accueil
                </NavLink>
                <NavLink to="/boutique">
                    Boutique
                </NavLink>
                <NavLink to="/propositions">
                    Propositions
                </NavLink>
                <NavLink to="/demande">
                    Demande
                </NavLink>
                <NavLink to="/contact">
                    Contact
                </NavLink>
            </NavMenu>
            <NavBtn> 
                <NavBtnLink to="/signup">Sign Up</NavBtnLink>
            </NavBtn>
        </Nav>
        </>
    );
};

export default Header;