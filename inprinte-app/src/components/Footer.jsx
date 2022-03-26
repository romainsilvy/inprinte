import React from 'react';
import {
    Container,
    Row,
    Column,
    FooterLink,
    Heading,
} from "../styles/FooterStyles";

const Footer = () => {
    return (
        <div className="footer">
            <Container>
                <Row>
                    <Column>
                        <Heading>A Propos De Nous</Heading>
                        <FooterLink href="/F.A.Q">F.A.Q</FooterLink>
                        <FooterLink href="/mentions-légales">Mentions légales</FooterLink>
                    </Column>
                    <Column>
                        <Heading>Nos Services</Heading>
                        <FooterLink href="/accueil">Accueil</FooterLink>
                        <FooterLink href="/boutique">Boutique</FooterLink>
                        <FooterLink href="/propositions">Propositions</FooterLink>
                        <FooterLink href="/profil">Profil</FooterLink>
                        <FooterLink href="/panier">Panier</FooterLink>
                    </Column>
                    <Column>
                        <Heading>Contact</Heading>
                        <FooterLink href="/tel">04 XX XX XX XX </FooterLink>
                        <FooterLink href="/mobile">08 XX XX XX XX</FooterLink>
                        <FooterLink href="/email">email@email.com</FooterLink>
                        <FooterLink href="/adresse">27 rue raoul servant, 69007 Lyon FRANCE</FooterLink>
                    </Column>
                    <Column>
                        <Heading>Communauté</Heading>
                        <FooterLink href="https://facebok.com">
                            <i className="fab fa-facebook-f">
                                <span style={{ marginLeft: "10px" }}>
                                    Facebook
                                </span>
                            </i>
                        </FooterLink>
                        <FooterLink href="https://instagram.com">
                            <i className="fab fa-instagram">
                                <span style={{ marginLeft: "10px" }}>
                                    Instagram
                                </span>
                            </i>
                        </FooterLink>
                        <FooterLink href="https://twitter.com">
                            <i className="fab fa-twitter">
                                <span style={{ marginLeft: "10px" }}>
                                    Twitter
                                </span>
                            </i>
                        </FooterLink>
                    </Column>
                </Row>
            </Container>
        </div>
    );
};

export default Footer;