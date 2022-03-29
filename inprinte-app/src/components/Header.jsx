import React from 'react';
import { Nav, Bars, NavLink, NavMenu, NavBtn, NavBtnLink } from '../styles/HeaderStyles';
import InprinteLogo from '../images/logo.svg'
import UserLogo from '../images/user.svg'
const Header = () => {
    return (
        <>
        <Nav>
            <NavLink to="/">
                <img src={InprinteLogo} alt="Inprinte Logo" />
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
                <NavBtnLink to="/login">
                    <img src={UserLogo} alt="User Logo" />
                </NavBtnLink>
            </NavBtn>
        </Nav>
        </>
    );
};

export default Header;