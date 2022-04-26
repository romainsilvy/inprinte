import React, { useState } from 'react'
import "../styles/Accordion.css"
import { Navbar } from '../components/Navbar';
import { Footer } from '../components/Footer';
import FAQ from '../pages/Faq';
import bar2 from '../images/bar2.svg'
function Accordion() {

    const [faqs, setfaqs] = useState([
        {
            question: "J'ai perdu mon mot de passe, comment faire pour le récuperer ?",
            answer: "Dans la partie CONNEXION, cliquez sur 'Mot de passe oublié ?' puis indiquez nous votre adresse email (qui correspond à votre identifiant) pour recevoir un nouveau mot de passe dans votre boîte mail.",
            open: false
        },
        {
            question: "Je souhaite annuler ma commande, comment faire ?",
            answer: "Pour annuler votre commande, contactez rapidement notre service client au 06 57 59 38 85 ou par email support@inprinte.com. Attention ! Une fois votre commande préparée, emballée ou expédiée, il n'est malheureusement plus possible de procéder à son annulation.",
            open: false
        },
        {
            question: "Que se passe t-il si je suis absent le jour de la livraison de mon colis ?",
            answer: "Dans le cas d'une livraison à domicile ou sur le lieu de travail - Sur l'avis de passage, il sera indiqué si votre colis sera livré le lendemain à l'adresse indiqué, ou si il sera remis en bureau de poste ou au point relais le plus proche.",
            open: false
        },
        {
            question: "Je souhaite retourner un article reçu, comment faire ?",
            answer: "Pour tout retour d'article, merci de contacter notre service client au 06 57 59 38 85 ou par e-mail à support@inprinte.com. Après accord de notre part, vous bénéficiez de 14 jours pour nous renvoyer l’article. Notez que les produits devront être retournés dans leur emballage d’origine, non déballés et ne devront pas avoir été utilisés. ",
            open: false
        },
        {
            question: "Le produit reçu est cassé ou défectueux, que faire ?",
            answer: "Contactez notre service client par téléphone au 06 57 59 38 85 ou par email support@inprinte.com dans les 14 jours suivant le réception de votre colis. Passé ce délai, aucune réclamation ne sera possible.",
            open: false
        }
    ]);

    const toggleFAQ = index => {
        setfaqs(faqs.map((faq, i) => {
            if (i === index) {
                faq.open = !faq.open
            } else {
                faq.open = false;
            }

            return faq;
        }))
    }
    return (
        <>
            <Navbar />
            <div className="accordion">
                <div className="title">
                    <div className="titleCat2">
                        <h1>FAQ</h1>
                        <img className="bar2" src={bar2} alt="bar" />
                    </div>
                </div>
                <div className="faqs">
                    {faqs.map((faq, i) => (
                        <FAQ faq={faq} index={i} toggleFAQ={toggleFAQ} />
                    ))}
                </div>
            </div>
            <Footer />
        </>
    );
}

export default Accordion;